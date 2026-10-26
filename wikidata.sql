-- wikidata.sql — 위키데이터 JSON 덤프 적재 스키마 (geuldev, PostgreSQL 16)
--
-- 2026-10-26. 기존 geuldev.public 스키마(geul-entity artifacts/wikidata/postgres.go,
-- 커밋 4225b3b)를 대신한다. 기존 테이블은 건드리지 않도록 별도 스키마 `wikidata`에 둔다.
--
-- 원칙
--   1. 원본을 먼저 보존한다. 개체마다 덤프의 JSON 한 줄을 그대로(raw) 저장하고,
--      나머지 테이블은 그 원본에서 파생한 질의용 색인이다. 해석 규칙에 버그가 있어도
--      재적재 없이 raw에서 다시 파생한다.
--   2. 적재 단계에서 아무것도 버리지 않는다. 값 없음(novalue)·값 미상(somevalue)도,
--      순서도, 출처도 남긴다. 실패한 줄은 적재 오류 테이블에 남긴다.
--   3. 판본을 추정하지 않는다. 덤프 파일·해시·적재 도구 커밋을 dump 테이블에 기록하고
--      모든 행이 그 dump를 가리킨다.
--
-- 기존 스키마의 부족한 점과 개선
--   | 기존 (public)                                   | 개선 (wikidata)                                        |
--   |-------------------------------------------------|--------------------------------------------------------|
--   | 덤프 판본·해시 기록 없음 (README "추정")          | dump: 파일, URL, 크기, MD5·SHA-1·SHA-256, 적재 도구 커밋 |
--   | 원본 JSON 보존 없음                               | entity_raw.raw: 덤프 줄 그대로 (lz4 압축)               |
--   | lastrevid·modified 없음                          | entity.lastrevid, entity.modified                       |
--   | quantity 본 값이 amount만 (단위·범위 소실)         | statement.value: datavalue.value 전체 (jsonb)           |
--   | 본 값 개체는 맨 ID, qualifier 개체는 JSON (불일치)  | 모든 snak이 같은 구조: value jsonb + value_entity       |
--   | novalue·somevalue: 본 값은 빈 값, qualifier는 버림 | snaktype 열 (value / novalue / somevalue)              |
--   | qualifier 속성 사이 순서 무작위 (Go map 순회)      | qualifier.ordinal: qualifiers-order와 속성 안 순서      |
--   | statement GUID 없음, 같은 속성 안 순서는 id 순     | statement.guid, statement.ordinal                       |
--   | 출처(references) 미적재                          | reference, statement_reference, reference_snak          |
--   | 속성 데이터 타입이 별도 집계 테이블에만             | entity.datatype (속성), 각 snak의 datatype               |
--   | sitelink 없음                                    | sitelink                                                |
--   | 삽입 오류 무시 (stmt.Exec 결과 미확인)             | load_error                                              |
--   | object_type에 datavalue 형식과 datatype이 섞임     | value_type(datavalue.type)과 datatype(snak.datatype) 분리 |
--
-- 적재 순서: 이 파일의 1~3절(타입·테이블)을 먼저 실행하고, COPY로 적재한 뒤
-- 4절(제약·색인)을 실행한다. 대량 적재 중 색인을 유지하면 수 배 느리다.
-- 4절을 실행하기 전에는 질의가 전부 순차 스캔이다. 적재 중 시험 질의는 하지 않는다.
--
-- 판(dump_id): 개체 단위 테이블은 모두 dump_id를 가진다. 판을 둘 이상 넣어 비교할 수
-- 있고, 키는 (개체, dump_id, …) 순이라 한 판 안의 개체 조회가 색인 하나로 끝난다.
--
-- 덤프 한 줄 = 개체 하나의 JSON 객체(끝의 쉼표 제외). 이 스키마는 all.json(항목 Q,
-- 속성 P)을 대상으로 한다. Lexeme(L)은 별도 덤프이며 entity.kind로 받을 수 있게만 둔다.

BEGIN;

CREATE SCHEMA IF NOT EXISTS wikidata;
SET search_path = wikidata;

-- 1. 타입 -------------------------------------------------------------------

-- snak 종류. novalue = 값이 없음이 확인됨, somevalue = 값은 있으나 미상.
CREATE TYPE snak_type AS ENUM ('value', 'novalue', 'somevalue');

-- statement rank.
CREATE TYPE statement_rank AS ENUM ('preferred', 'normal', 'deprecated');

-- 2. 판본과 원본 ------------------------------------------------------------

-- 덤프 한 판의 출처와 적재 기록. 모든 파생 행은 dump_id로 판을 가리킨다.
CREATE TABLE dump (
    dump_id        smallint    PRIMARY KEY,
    dump_date      date        NOT NULL,                 -- 덤프 폴더 날짜 (20261019)
    file_name      text        NOT NULL,                 -- wikidata-20261019-all.json.bz2
    url            text        NOT NULL,
    size_bytes     bigint      NOT NULL,
    md5            text,                                 -- 위키미디어 공식 체크섬
    sha1           text,
    sha256         text        NOT NULL,                 -- 받은 뒤 직접 계산
    published_at   timestamptz,                          -- 서버 Last-Modified
    downloaded_at  timestamptz NOT NULL,
    loader         text        NOT NULL,                 -- 적재 도구 저장소·경로
    loader_commit  text        NOT NULL,
    load_started   timestamptz,
    load_finished  timestamptz,                          -- NULL이면 적재 중이거나 중단
    checkpoint_line bigint      NOT NULL DEFAULT 0,      -- 이 줄까지는 빠짐없이 커밋됨 (재개 지점)
    entity_count   bigint,                               -- 적재를 마친 개체 수
    error_count    bigint,
    note           text
);

-- 개체 원본. 덤프 줄을 바이트 그대로 보존한다(끝의 쉼표와 줄바꿈만 뺀다).
CREATE TABLE entity_raw (
    entity_id  text     NOT NULL,                        -- Q42, P31
    dump_id    smallint NOT NULL,
    line_no    bigint   NOT NULL,                        -- 덤프 파일의 물리 줄 번호 (1부터. 1번 줄은 '[')
    raw        text     COMPRESSION lz4 NOT NULL
);

-- 적재 중 실패한 줄. 파싱 오류, 예상 밖 구조, 삽입 오류를 숨기지 않는다.
CREATE TABLE load_error (
    dump_id    smallint NOT NULL,
    line_no    bigint   NOT NULL,
    entity_id  text,                                     -- 알 수 있으면
    stage      text     NOT NULL,                        -- parse, entity, statement, qualifier, reference, insert
    message    text     NOT NULL,
    excerpt    text                                      -- 줄 앞부분 (최대 1KB)
);

-- 3. 파생 테이블 ------------------------------------------------------------

-- 개체 머리 정보.
CREATE TABLE entity (
    entity_id  text        NOT NULL,                     -- Q42
    line_no    bigint      NOT NULL,                     -- 덤프 안 줄 번호 (entity_raw와 같다)
    kind       char(1)     NOT NULL,                     -- Q, P, L
    num        bigint      NOT NULL,                     -- 42
    type       text        NOT NULL,                     -- item, property, lexeme
    datatype   text,                                     -- 속성만: wikibase-item, quantity, url, external-id …
    lastrevid  bigint,                                   -- 이 판의 revision ID
    modified   timestamptz,
    dump_id    smallint    NOT NULL
);

CREATE TABLE label (
    entity_id  text     NOT NULL,
    dump_id    smallint NOT NULL,
    language   text     NOT NULL,
    value      text     NOT NULL
);

CREATE TABLE description (
    entity_id  text     NOT NULL,
    dump_id    smallint NOT NULL,
    language   text     NOT NULL,
    value      text     NOT NULL
);

-- 별칭. ordinal은 덤프의 언어별 배열 순서(0부터).
CREATE TABLE alias (
    entity_id  text     NOT NULL,
    dump_id    smallint NOT NULL,
    language   text     NOT NULL,
    ordinal    smallint NOT NULL,
    value      text     NOT NULL
);

CREATE TABLE sitelink (
    entity_id  text     NOT NULL,
    dump_id    smallint NOT NULL,
    site       text   NOT NULL,                          -- enwiki, kowiki …
    title      text   NOT NULL,
    badges     text[] NOT NULL DEFAULT '{}'               -- 배지 항목 Q ID
);

-- statement(위키데이터 claim). snak 값 열(snaktype … value_entity)은 qualifier·reference_snak과 같다.
--   value        datavalue.value 그대로. 개체 {"entity-type":…,"id":"Q5"}, 시간 {"time":…,
--                "precision":…,"calendarmodel":…}, 수량 {"amount":…,"unit":…,"upperBound":…},
--                문자열은 JSON 문자열. novalue·somevalue는 NULL.
--   value_entity 값이 개체일 때 그 ID(Q5). 역방향 질의용.
CREATE TABLE statement (
    statement_id  bigint         NOT NULL,               -- dump_id << 55 | 줄 번호 << 20 | 개체 안 순번 (결정적)
    guid          text           NOT NULL,               -- 위키데이터 statement ID (Q42$…)
    entity_id     text           NOT NULL,
    property      text           NOT NULL,
    ordinal       smallint       NOT NULL,               -- 같은 개체·속성 안 순서 (0부터)
    rank          statement_rank NOT NULL,
    snaktype      snak_type      NOT NULL,
    datatype      text,                                  -- mainsnak.datatype
    value_type    text,                                  -- datavalue.type (wikibase-entityid, time, quantity …)
    value         jsonb,
    value_entity  text,
    dump_id       smallint       NOT NULL
);

-- qualifier. ordinal은 qualifiers-order의 속성 순서를 따라 펼친 전체 순서(0부터)이고,
-- value_ordinal은 같은 속성 안 순서다. 둘 다 덤프 그대로다.
CREATE TABLE qualifier (
    statement_id   bigint    NOT NULL,
    ordinal        smallint  NOT NULL,
    property       text      NOT NULL,
    value_ordinal  smallint  NOT NULL,
    snaktype       snak_type NOT NULL,
    datatype       text,
    value_type     text,
    value          jsonb,
    value_entity   text,
    hash           text                                  -- snak hash
);

-- 출처 묶음. 위키데이터는 같은 출처를 여러 statement가 hash로 공유한다.
CREATE TABLE reference (
    hash     text     NOT NULL,                          -- reference hash
    dump_id  smallint NOT NULL
);

-- statement와 출처 묶음의 연결. ordinal은 statement의 references 배열 순서.
CREATE TABLE statement_reference (
    statement_id  bigint   NOT NULL,
    ordinal       smallint NOT NULL,
    hash          text     NOT NULL
);

-- 출처 묶음의 항목(snak). ordinal은 snaks-order를 따라 펼친 전체 순서.
CREATE TABLE reference_snak (
    hash           text      NOT NULL,
    dump_id        smallint  NOT NULL,
    ordinal        smallint  NOT NULL,
    property       text      NOT NULL,
    value_ordinal  smallint  NOT NULL,
    snaktype       snak_type NOT NULL,
    datatype       text,
    value_type     text,
    value          jsonb,
    value_entity   text
);

COMMIT;

-- 4. 적재를 마친 뒤 ----------------------------------------------------------
--
-- 4.1 중단 뒤 재개했다면: checkpoint_line 이후 줄의 행은 중단 전에 일부 커밋됐을 수 있다.
--     적재 도구는 재개할 때 그 줄들의 행을 먼저 지운다(-from-line 없이 재개하면 자동).
--
-- 4.2 출처 묶음 중복 정리. 같은 출처(hash)를 여러 개체가 쓰므로, 적재 도구는 최근에 본
--     hash만 기억해 거른다(메모리 한도). 남은 중복을 PRIMARY KEY 전에 지운다.
--
-- BEGIN;
-- SET search_path = wikidata;
-- CREATE TABLE reference_dedup (LIKE reference);
-- INSERT INTO reference_dedup SELECT DISTINCT ON (hash, dump_id) * FROM reference;
-- DROP TABLE reference; ALTER TABLE reference_dedup RENAME TO reference;
-- CREATE TABLE reference_snak_dedup (LIKE reference_snak);
-- INSERT INTO reference_snak_dedup SELECT DISTINCT ON (hash, dump_id, ordinal) * FROM reference_snak;
-- DROP TABLE reference_snak; ALTER TABLE reference_snak_dedup RENAME TO reference_snak;
-- COMMIT;
--
-- 4.3 제약과 색인 --------------------------------------------------------------
--
-- 질의 패턴별 색인. PRIMARY KEY와 UNIQUE는 그 자체로 색인이다(같은 열 순서의 색인을 따로 두지 않는다).
--
--   | 질의                                              | 색인                                   |
--   |---------------------------------------------------|----------------------------------------|
--   | 개체 하나의 원본·머리·레이블·별칭·설명·sitelink     | 각 테이블 PK (entity_id, dump_id, …)    |
--   | Q1~Q1000 같은 번호 범위                            | entity_dump_kind_num                   |
--   | 덤프 줄 번호로 원본 찾기 (재개, 오류 추적)           | entity_raw_dump_line                   |
--   | 개체의 statement를 속성 순서대로                    | statement_entity                       |
--   | statement GUID로 찾기                              | statement_guid                         |
--   | 속성·값으로 찾기 (P31 = Q5), 속성별 사용량           | statement_property_value               |
--   | 어떤 개체를 값으로 가리키는 statement (역방향)        | statement_value_entity                 |
--   | statement의 qualifier                              | qualifier PK (statement_id, ordinal)   |
--   | qualifier 속성·값으로 찾기 (P580 시점 등)            | qualifier_property_value               |
--   | qualifier가 가리키는 개체 (역방향)                   | qualifier_value_entity                 |
--   | statement의 출처                                   | statement_reference PK                 |
--   | 출처 묶음을 쓰는 statement                          | statement_reference_hash               |
--   | 출처 묶음의 항목                                    | reference_snak PK (hash, dump_id, …)   |
--   | 출처가 가리키는 문헌·개체 (stated in Q…, 역방향)      | reference_snak_value_entity            |
--   | 위키백과 제목 → Q ID                                | sitelink_site_title                    |
--   | 레이블·별칭 검색 (언어별, 대소문자 무시, 접두어 포함)  | label_language_value, alias_language_value |
--   | 적재 오류를 줄 순서로                                | load_error_dump_line                   |
--
-- 레이블·별칭 검색 색인은 lower(value)에 text_pattern_ops를 쓴다. 기본 정렬 규칙(C가 아닌
-- 로캘)의 btree는 LIKE 'abc%' 접두어 검색에 쓰이지 않는다. 같음 비교에도 쓰인다.
-- jsonb value 안쪽(quantity 단위, 시간 정밀도 등) 질의는 색인을 두지 않았다. 필요해지면
-- 표현식 색인(예: (value->>'unit') WHERE value_type = 'quantity')을 추가한다.
--
-- BEGIN;
-- SET search_path = wikidata;
--
-- ALTER TABLE entity_raw          ADD PRIMARY KEY (entity_id, dump_id);
-- ALTER TABLE entity              ADD PRIMARY KEY (entity_id, dump_id);
-- ALTER TABLE label               ADD PRIMARY KEY (entity_id, dump_id, language);
-- ALTER TABLE description         ADD PRIMARY KEY (entity_id, dump_id, language);
-- ALTER TABLE alias               ADD PRIMARY KEY (entity_id, dump_id, language, ordinal);
-- ALTER TABLE sitelink            ADD PRIMARY KEY (entity_id, dump_id, site);
-- ALTER TABLE statement           ADD PRIMARY KEY (statement_id);
-- ALTER TABLE qualifier           ADD PRIMARY KEY (statement_id, ordinal);
-- ALTER TABLE reference           ADD PRIMARY KEY (hash, dump_id);
-- ALTER TABLE statement_reference ADD PRIMARY KEY (statement_id, ordinal);
-- ALTER TABLE reference_snak      ADD PRIMARY KEY (hash, dump_id, ordinal);
--
-- -- 참조 무결성. 적재 도구의 버그(고아 qualifier, 없는 출처 hash)를 여기서 잡는다.
-- ALTER TABLE entity_raw          ADD FOREIGN KEY (dump_id) REFERENCES dump;
-- ALTER TABLE entity              ADD FOREIGN KEY (dump_id) REFERENCES dump;
-- ALTER TABLE statement           ADD FOREIGN KEY (dump_id) REFERENCES dump;
-- ALTER TABLE reference           ADD FOREIGN KEY (dump_id) REFERENCES dump;
-- ALTER TABLE load_error          ADD FOREIGN KEY (dump_id) REFERENCES dump;
-- ALTER TABLE statement           ADD FOREIGN KEY (entity_id, dump_id) REFERENCES entity;
-- ALTER TABLE qualifier           ADD FOREIGN KEY (statement_id) REFERENCES statement;
-- ALTER TABLE statement_reference ADD FOREIGN KEY (statement_id) REFERENCES statement;
--
-- CREATE UNIQUE INDEX entity_raw_dump_line        ON entity_raw (dump_id, line_no);
-- CREATE INDEX entity_dump_kind_num               ON entity (dump_id, kind, num);
-- CREATE UNIQUE INDEX statement_guid              ON statement (guid, dump_id);
-- CREATE INDEX statement_entity                   ON statement (entity_id, dump_id, property, ordinal);
-- CREATE INDEX statement_property_value           ON statement (property, value_entity);
-- CREATE INDEX statement_value_entity             ON statement (value_entity) WHERE value_entity IS NOT NULL;
-- CREATE INDEX qualifier_property_value           ON qualifier (property, value_entity);
-- CREATE INDEX qualifier_value_entity             ON qualifier (value_entity) WHERE value_entity IS NOT NULL;
-- CREATE INDEX statement_reference_hash           ON statement_reference (hash);
-- CREATE INDEX reference_snak_value_entity        ON reference_snak (value_entity) WHERE value_entity IS NOT NULL;
-- CREATE UNIQUE INDEX sitelink_site_title         ON sitelink (site, title, dump_id);
-- CREATE INDEX label_language_value               ON label (language, lower(value) text_pattern_ops);
-- CREATE INDEX alias_language_value               ON alias (language, lower(value) text_pattern_ops);
-- CREATE INDEX load_error_dump_line               ON load_error (dump_id, line_no);
--
-- ANALYZE;
-- COMMIT;

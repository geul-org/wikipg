# 위키데이터 → GEUL 포팅

위키데이터 항목을 GEUL 문서로 옮겨 `.well-known/geul/` 아래에 공표한다. 출처·근거 기록(files/geul/sign.md, legacy.md)의 실전 시험대이기도 하다.

## 원본 판본

### 이전 판본 (20250908, 삭제)

지금 공표된 문서(2026-10-26 생성)의 원본이다. 이 DB(`geuldev.public`)는 2026-10-26 삭제했다(적재 입력이 소실되어 다시 만들 수 없다). 다음 생성부터는 아래 새 덤프를 쓴다.

| 항목 | 값 |
|---|---|
| 덤프 | Wikidata JSON 전체 덤프 **20250908** (`wikidata-20250908-all.json`) |
| 판단 근거 | 아래 "판본 판단 근거" |
| 보관 파일 | `~/latest-all.json.bz2`, 98,022,279,070 byte, 받은 시각 2025-09-17 20:49 ~ 2025-09-18 00:32 (KST) |
| DB 적재 시각 | 2025-09-22 ~ 2025-09-23 (`entities.created_at`) |
| 적재 도구 | `github.com/geul-org/geul-entity` `artifacts/wikidata/postgres.go`, 커밋 `4225b3b4297796a79e95c4cda5aa54df3acc4069` |
| 적재 입력 | `/mnt/d/latest-all.json`(압축을 푼 파일). 지금은 없다(소실). 보관 파일과 같은 덤프로 본다 |

## 새 덤프 (20261019, 적재 전)

wikiload로 geuldev의 스키마 `wikidata`에 적재할 판본이다. 위 20250908 DB와 달리 판본·해시를 추정하지 않는다. 출처 기록은 적재 뒤 `dump` 행에서 읽는다.

| 항목 | 값 |
|---|---|
| 덤프 | Wikidata JSON 전체 덤프 **20261019** (`wikidata-20261019-all.json.bz2`) |
| 보관 파일 | `/mnt/d/geul/wikidata/wikidata-20261019-all.json.bz2`, 103,488,697,311 byte, 받기 완료 2026-10-26 15:34 (KST), your.org 미러 |
| MD5 | `cfec3f8ab5fd8f6f0b70f58c590a02da` (공식 `md5sums`와 일치) |
| SHA-1 | `a05b8b809ab80576da6d838c67138ddf401055e6` (공식 `sha1sums`와 일치) |
| SHA-256 | `a249a19d12b1b0e8443248da3583090682be3f20c6bdd6ed1f7e1ae05b6b0774` (직접 계산, 같은 폴더 `.sha256`) |

### 0908 판본 판단 근거

- 파일 이름이 `latest-all.json.bz2`라 이름만으로는 판본을 알 수 없다. 덤프 사이트는 최근 판만 남겨 2025년 판의 크기·해시를 대조할 수 없었다.
- 덤프 앞부분 6,354개 개체의 `modified` 최댓값이 `2025-09-08T03:01:14Z`, `lastrevid` 최댓값이 `2402035052`다. 자주 편집되는 앞쪽 개체가 이 시각에서 멈춘다.
- DB의 가장 큰 QID는 `Q136198562`이고, 위키데이터 API 기준 첫 판이 `2025-09-10T01:34:49Z`(revid `2403031064`)다. 덤프는 여러 날에 걸쳐 만들어지므로 9월 8일(월)에 시작한 판과 맞는다.
- 9월 17일에 받았지만 그때 `latest`는 아직 0908 판을 가리키고 있었던 것으로 본다(0915 판은 완성 전).

### 0908 판의 빈칸 (legacy.md의 미상·소실·비공개)

새 스키마(`wikidata.sql`)는 아래 빈칸을 대부분 메운다: 판본·해시(`dump`), 원본 줄(`entity_raw`), qualifier 순서(`ordinal`), novalue·somevalue(`snaktype`), 출처(`reference*`).

| 항목 | 상태 |
|---|---|
| 항목별 판본(`lastrevid`) | **미적재.** DB에 없다. 덤프 판본으로 대신한다 |
| 위키데이터 참조(references) | **미적재.** `triple_references`가 비어 있다 |
| 적재 입력 파일 | **소실.** `/mnt/d/latest-all.json`이 없다 |
| qualifier 순서 | **소실(일부).** 적재 도구가 qualifier를 Go map으로 순회해 속성 사이 순서(위키데이터 `qualifiers-order`)가 무작위로 저장됐다. 같은 속성 안의 값 순서는 보존된다 |
| novalue·somevalue qualifier | **미적재.** 적재 도구가 `datavalue`가 없는 qualifier를 버렸다(예: 종료 시점 novalue = "아직 끝나지 않음") |
| 덤프 판본 | **추정.** 위 근거로 0908 판으로 판단했다 |
| 원본 해시 | **생략.** DB는 소실된 적재 입력에서 읽었으므로 보관 파일 해시는 DB 내용을 증명하지 못하고, 2025년 판은 덤프 사이트에 없어 대조할 수도 없다 |

## DB 테이블 (`geuldev.wikidata`)

정의는 `wikidata.sql`이다. wikiport가 읽는 것:

| 테이블 | wikiport가 쓰는 열 |
|---|---|
| `dump` | `url`, `sha256`, `downloaded_at`, `loader`, `loader_commit` (출처 기록) |
| `entity` | `entity_id` (있는지) |
| `label`, `description`, `alias` | `language`, `value` (별칭은 `ordinal` 순) |
| `statement` | `statement_id`(덤프 순), `property`, `rank`, `value_type`, `value`(jsonb), `value_entity` |
| `qualifier` | `statement_id`, `ordinal`(qualifiers-order), `property`, `value_type`, `value`, `value_entity` |

- 모든 질의는 `dump_id`로 한 판만 읽는다(`-dump-id`).
- 값은 개체면 맨 ID(`value_entity`), 문자열이면 따옴표 없는 문자열, 그 밖은 datavalue JSON(jsonb의 공백·키 순서)이다. novalue·somevalue는 빈 값이라 건너뛴다(보고서 `값 없음`).
- 본 값과 qualifier 값이 같은 형태다(이전 DB는 qualifier 개체만 JSON이었다).

## 작업 범위 (2026-10-25 결정)

- **도구:** Go만 쓴다. DB에서 읽어 라이브러리 `core`로 인코딩하고 `resolve.Publish`로 공표한다. Python은 쓰지 않는다.
- **Issuer:** `wikidata.geul.org` = `0x00010000` 하나(2026-10-26 결정). 위키데이터의 모든 개체 종류를 한 Issuer 아래 둔다. Issuer는 공표 주체이고, Q·P·L은 같은 출처를 같은 주체가 옮긴 것이라 나누지 않는다(geul.org도 Issuer `0x0000` 하나 안에서 문서 ID 영역으로 나눈다).
- **문서 ID:** 64비트 중 맨 위 바이트가 개체 종류의 ASCII 글자, 아래 56비트가 번호다(`port.DocNumber`).

  | 종류 | 맨 위 바이트 | 예 |
  |---|---|---|
  | 항목 Q | `0x51` | Q42 → `0x510000000000002a` |
  | 속성 P | `0x50` | P31 → `0x500000000000001f` |
  | Lexeme L | `0x4C` | L42 → `0x4c0000000000002a`. Form·Sense는 그 Lexeme 문서의 LOCAL. 문서 생성·변환은 Lexeme 덤프를 다룰 때 |
  | Entity Schema E | `0x45` | E10 → `0x450000000000000a` (같은 시점) |

  - hex만 보고 종류를 읽을 수 있다. 번호는 2^56(약 7경)까지이고, 넘는 번호는 오류로 막는다. 맨 위 바이트에 0을 두는 낭비 없이 종류 한 바이트만 쓴다(2026-10-26).
  - 이름 색인 `documents.geul`은 하나이고 이름은 소문자 종류 글자 + 번호다(`q42`, `p31`).
  - Q·P를 가리키는 Reference(WIKIDATA-QID·PID)는 문서 위치와 상관없다. 바뀐 것은 문서가 놓이는 주소뿐이다.
- **문서 단위:** QID 하나당 `.geul` 문서 하나.
- **대상:** 우선 Q1 ~ Q1000. PID는 필요한 만큼 추가한다.
- **레코드:** LABEL, DESCRIPTION, ALIAS, (속성이면) 데이터 타입, TRIPLE과 그 QUALIFIER를 작성한다(이 순서).
- **값:** 항목(QID), 문자열, 외부 ID, URL·커먼즈(URI), 날짜, monolingualtext, 좌표, 수량(아래 값 변환 표). 처음 범위는 날짜까지였고 2026-10-26에 monolingualtext·좌표·URI·수량을 더했다. 그 밖의 값(Lexeme 등)은 변환하지 않았다고 기록한다.
- **언어:** ko, en.
- **명세:** 아직 없는 명세는 설계하면서 제안하고 검토를 받는다(ALIAS, DESCRIPTION, TRIPLE, QUALIFIER는 docs/edge에 반영).

## 출처 기록 (초기 최소 세트, files/geul/legacy.md)

각 문서에 다음을 남긴다.

- 출처: SOURCE 둘. 덤프 URI + 덤프 파일 SHA-256(`dump.sha256`, SHA-256-DIGEST), 그 항목의 QID + 없음 표지 `0x0000`(항목 단위 다이제스트는 없다, docs/edge/CODEC 선택 필드). 덤프 SHA-256을 모르면 없음 표지
- 받은 시각: `dump.downloaded_at`
- 변환 도구: TRANSFORMED-BY 셋(docs/meta/lineage/TRANSFORMED-BY). 도구 Reference는 geul.org URL 정의 Reference 문서(GEUL-ORG), 버전은 전체 커밋 ID(SHA1-DIGEST)다.
  - 적재 도구 wikiload(wikiport 저장소, `0x0001000000000001`) `dump.loader_commit`. 다른 적재 도구로 채운 dump는 거부한다
  - (이전 0908 판의 적재 도구 geul-entity `0x0001000000000003` `4225b3b4297796a79e95c4cda5aa54df3acc4069`)
  - 포팅 도구 wikiport(`0x0001000000000001`) `-commit`
  - GEUL 라이브러리 geul-lib(`0x0001000000000002`) `-lib-commit`. wikiport는 `go.mod`의 `replace`로 로컬 라이브러리를 쓰므로 wikiport 커밋만으로는 라이브러리 코드가 정해지지 않는다

## 포팅 도구 (`wiki/go`)

```sh
cd draft/t002/wiki/go
# DB 접속은 lib/pq의 PG* 환경변수(PGHOST, PGUSER, PGPASSWORD, PGDATABASE=geuldev, PGSSLMODE)로 한다.
go run ./cmd/wikiport -root ../.. -from 1 -to 1000 -commit "$(git rev-parse HEAD)" -lib-commit "$(git -C ../../geul rev-parse HEAD)" -report report.tsv
```

| 플래그 | 뜻 | 기본값 |
|---|---|---|
| `-root` | `.well-known/geul`을 담은 t002 루트 | `.` |
| `-from`, `-to` | QID 범위 | 1, 1000 |
| `-langs` | 레이블·설명·별칭 언어 | `ko,en` |
| `-commit` | TRANSFORMED-BY에 남길 포팅 도구 커밋, 전체 ID (필수) | |
| `-lib-commit` | TRANSFORMED-BY에 남길 GEUL 라이브러리 커밋, 전체 ID (필수) | |
| `-report` | 개체별 건수·건너뛴 사유·원문 보존 날짜의 사유 TSV (열: id, labels, descriptions, aliases, triples, qualifiers, skipped, raw. qualifier 사유는 `qualifier ` 접두) | 없음 |
| `-dsn` | PostgreSQL DSN | 비우면 PG* 환경변수 |
| `-schema` | 적재된 덤프의 스키마(wikiload `-schema`) | `wikidata` |
| `-dump-id` | 읽을 판(wikiload `-dump-id`) | 1 |

- Q 범위를 먼저 옮기고, 그 문서들이 쓴 속성(Predicate, 값이 속성 개체인 것)을 이어서 옮긴 뒤 두 Issuer의 `documents.geul`을 쓴다. 색인은 이번 실행분과 그 Issuer 디렉터리에 이미 공표된 문서를 모두 담는다(범위를 나눠 실행해도 앞 실행의 이름이 지워지지 않는다).
- 공표는 개발 기간 동안 `resolve.Dir.Overwrite`다(2026-10-26 결정): 실행할 때마다 v1을 덮어쓰고 버전 목록도 v1 하나만 둔다. 버전을 쌓으면 용량을 감당할 수 없기 때문이다. 내용이 같으면 아무 파일도 바꾸지 않는다. 형식이 안정되면 `Publish`(버전 누적)로 돌아간다.

## 적재 도구 (`wiki/go`, wikiload)

덤프를 geuldev의 스키마 `wikidata`(`wiki/wikidata.sql`)에 적재한다. 0908 판에서 뽑은 `ko_term`·`lexeme`는 스키마 `wikidata0908`로 옮겼다(2026-10-26). 같은 DB라 stdict 등과 조인할 수 있다. 기존 `geuldev.public`(geul-entity `postgres.go`, 2026-10-26 삭제)을 대신한다. 원칙과 기존 스키마 대비 개선은 `wikidata.sql` 머리말에 있다.

```sh
cd draft/t002/wiki/go
go run ./cmd/wikiload -in /mnt/d/geul/wikidata/wikidata-20261019-all.json.bz2 -commit "$(git rev-parse HEAD)"
```

- 스키마: `psql -d geuldev -f wikidata.sql`은 1~3절(타입·테이블)만 만든다. 4절(재개 정리·출처 중복 정리·제약·색인)은 주석으로 되어 있고, 적재를 마친 뒤 주석을 풀어 실행한다.
- 테이블: `dump`(파일·URL·크기·MD5·SHA-1·SHA-256·적재 도구 커밋), `entity_raw`(덤프 줄 원문), `load_error`, `entity`, `label`, `description`, `alias`, `sitelink`, `statement`, `qualifier`, `reference`, `statement_reference`, `reference_snak`, Lexeme용 `lexeme`, `lemma`, `form`, `form_representation`, `sense`, `sense_gloss`. 개체 단위 테이블은 모두 `dump_id`를 가진다.
- 기본값은 기계를 30% 남긴다: CPU 70% 중 1/5 압축 해제(lbzip2), 2/5 JSON 파싱, 작성자(DB 연결) 4, 배치 8 MiB(메모리 약 1 GB). 2026-10-26 측정에서 작성자 4·8 MiB가 가장 빨랐다.
- Lexeme은 별도 덤프(`wikidata-YYYYMMDD-lexemes.json.bz2`)라 다른 판 번호로 적재한다: `-in …lexemes.json.bz2 -dump-id 2`. Form·Sense도 `entity`에 행(type `form`·`sense`, kind L)을 두어 그 statement(`entity_id` = `L4-F1`)가 가리킬 개체가 있다. 표제어·어형 표기는 낱말 자체라 `-langs`로 거르지 않고, 뜻풀이(gloss)는 거른다. Lexeme·Form·Sense의 statement 키는 한 줄 안에서 이어진다.
- 체크포인트에서 이어 읽는다(`-from-line 0`). 언어는 `-langs`(기본 `ko,en`, `all`은 전부).
- 플래그 전체: `go run ./cmd/wikiload -h`.
- wikiport는 이 스키마를 읽는다(2026-10-26 전환). 테스트는 `testdata/fixture.sql`을 임시 스키마에 넣어 돌린다.

## 값 변환

| 위키데이터 | GEUL |
|---|---|
| 항목 `Q…`, 속성 `P…` | WIKIDATA-QID, WIKIDATA-PID Reference |
| Lexeme `L…`, EntitySchema `E…` | 건너뜀 (Reference Codec 없음) |
| string, external-id, math, musical-notation | UTF8-TEXT (external-id는 식별자 텍스트 그대로, 2026-10-26 결정) |
| url | URI. 비ASCII·URI에 못 쓰는 바이트는 퍼센트 인코딩(RFC 3987 3.1, 되돌릴 수 있다). 스킴이 없으면 원문 UTF8-TEXT |
| commonsMedia | URI `https://commons.wikimedia.org/wiki/File:` + 파일 이름(공백 → `_`, 한 경로 조각으로 퍼센트 인코딩). 주소 앞부분은 wikiport가 붙인다 |
| geo-shape, tabular-data | URI `https://commons.wikimedia.org/wiki/` + 페이지 이름(`Data:…`), 위와 같은 인코딩 |
| time, 그레고리력 일 정밀도(기원전 포함) | GREGORIAN-DATE |
| time, 그레고리력 월·연·10년 정밀도 | EDTF |
| time, 율리우스력 일·월·연 정밀도 | GEDCOM-DATE (`JULIAN 14 OCT 1066`) |
| time, 그 밖(세기 이상 정밀도, 시간대·오차 범위 등) | datavalue JSON 원문을 UTF8-TEXT로. 날짜는 건너뛰지 않는다 |
| novalue, somevalue | 건너뜀 (DB에 값이 비어 있다) |
| monolingualtext | UTF8-LANGUAGE-TEXT (언어는 LANGUAGE-BCP47-CODE, 아래 언어 코드 대응). 그래도 BCP 47 형식이 아니면(빈 언어) datavalue 원문을 UTF8-TEXT로 |
| globecoordinate | 지구(Q2)는 COORDINATE-EARTH, 다른 천체는 COORDINATE-GEO(천체 = WIKIDATA-QID). 지구 좌표의 불확도는 precision에서 Georeferencing Best Practices 3.4.3 방법으로(정밀도 기인 불확도의 하한), 다른 천체와 precision 없음은 미상. 천체 없음·고도 있음·범위 밖·이상한 precision은 datavalue 원문을 UTF8-TEXT로 |
| quantity | QUANTITY(amount·단위·하한·상한, docs/value/numeric/QUANTITY). 수와 경계는 자릿수 그대로, 단위 `1`은 무단위(없음 표지), 위키데이터 항목 단위는 WIKIDATA-QID. 그 밖의 단위·형식은 원문 UTF8-TEXT |
| 그 밖(Lexeme 값 등) | 건너뜀 (이번 범위 밖) |

- 레이블·설명·별칭은 LABEL·DESCRIPTION·ALIAS(Edge Id, Subject, Object, `entity_descriptions`의 빈 설명은 건너뜀)이고 Object는 UTF8-LANGUAGE-TEXT, 언어는 LANGUAGE-BCP47-CODE(GEUL 언어 레지스트리 전). 언어 코드는 BCP 47 형식만 확인한다.

### 언어 코드 대응

위키미디어 언어 코드 중 BCP 47이 아닌 14개는 BCP 47 tag로 옮긴다(2026-10-26 기존 DB 전수조사에서 레이블·별칭·설명·monolingualtext의 무효 코드 전부, 대소문자 무시). 근거는 IANA Language Subtag Registry 항목이고(File-Date 2026-09-17, 2026-10-26 확보. IANA 프로토콜 레지스트리라 CC0 1.0, [IANA·IETF 공동 성명 2021-11-10](https://www.iana.org/help/licensing-terms)) MediaWiki `LanguageCode.php`와는 대조만 했다.

| 위키미디어 | BCP 47 | | 위키미디어 | BCP 47 |
|---|---|---|---|---|
| `sr-el` | `sr-Latn` | | `map-bms` | `jv-x-bms` |
| `bat-smg` | `sgs` | | `roa-tara` | `nap-x-tara` |
| `fiu-vro` | `vro` | | `de-formal` | `de-x-formal` |
| `cbk-zam` | `cbk` | | `es-formal` | `es-x-formal` |
| `simple` | `en-simple` | | `hu-formal` | `hu-x-formal` |
| `eml` | `egl` | | `nl-informal` | `nl-x-informal` |
| `roa-rup` | `rup` | | | |
| `zh-classical` | `lzh` | | | |

- 사설 확장(`x-`)의 뜻: `x-formal` 존댓말(격식체), `x-informal` 반말(비격식체), `x-bms` 자바어 반유마스 방언, `x-tara` 나폴리어 타란토 방언.
- 표에 없고 BCP 47 형식도 아닌 코드(빈 언어 등)는 레이블·설명·별칭이면 건너뛰고(보고서 `skipped`), monolingualtext면 원문 UTF8-TEXT로 보존한다(보고서 `raw`).
- 속성(P) 문서는 그 속성의 데이터 타입(`entity.datatype`)을 문서 안에 적는다. 다른 문서에서 그 속성을 만난 해석기가 속성 문서 하나만 보면 되게 한다(별도 어휘 문서를 두지 않는다, 2026-10-26 결정).
  ```
  REFERENCE-NODE( #n → URI http://wikiba.se/ontology#propertyType )
  TRIPLE( #n, P214, URI http://wikiba.se/ontology#ExternalId, normal )
  ```
  어휘는 위키데이터 RDF의 Wikibase 온톨로지 이름 그대로다(18종, files/geul/wikidata-datatype-mapping.md). 모르는 데이터 타입은 쓰지 않고 보고서 `skipped`에 남긴다. 문서당 약 50 WORD.
- 클레임은 TRIPLE(Edge Id, Predicate, Subject, Object, Rank)이고, 그 qualifier는 바로 뒤에 QUALIFIER(Edge Id, Property, 대상 TRIPLE의 Edge Id, Value)로 잇는다. qualifier 값은 본 값과 같은 규칙으로 바꾼다(날짜 원문 보존 포함). 출처는 담지 않는다(DB에 없다).
- qualifier 순서는 `triple_qualifiers`의 물리 행 순서(ctid)다. 같은 속성 안의 순서는 원본과 같지만 속성 사이 순서는 원본이 아니다(위 빈칸).
- TRANSFORMED-BY는 도구 Reference(EXTERNAL-ISSUER-1WORD, geul.org URL 정의 Reference 문서의 LOCAL 2) + 커밋 SHA1-DIGEST다(2026-10-26). 커밋은 커밋된 소스여야 한다: 작업 트리에 커밋되지 않은 변경이 있으면 그 커밋 ID를 쓰지 말고 먼저 커밋한다(도구는 아직 검사하지 않는다).

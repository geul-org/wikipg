# wikipg

위키데이터 JSON 덤프를 PostgreSQL(geuldev, 스키마 `wikidata`)에 무손실로 적재하는 도구 `wikipg`. 개체·레이블·클레임·qualifier·출처·sitelink·Lexeme을 덤프 순서와 원문 그대로 담고, 판본(파일·해시·적재 도구 커밋)을 `dump` 행에 남긴다.

- 저장소: [geul-org/wikipg](https://github.com/geul-org/wikipg). 2026-10-27 이름을 wikiport → wikipg로(명령 wikiload → wikipg), `go/`를 루트로 올렸고, GEUL 변환(옛 `cmd/wikiport`)은 [geul-org/gogeulwiki](https://github.com/geul-org/gogeulwiki)(`wikigeul`)로 옮겼다. 모듈 경로는 `github.com/geul-org/wikipg`.
- **무손실.** 덤프의 모든 줄을 `entity_raw`에 원문 그대로 두고, 구조화 표는 값을 버리지 않는다. jsonb·text가 못 담는 값(`\u0000`)은 JSON 원문을 `value_json`(statement·qualifier·reference_snak·label·description·alias·lemma·form_representation·sense_gloss)과 `title_json`(sitelink)에 둔다. 언어는 기본 전부(`-langs all`). 그래서 이 DB를 읽어 만든 문서의 출처에는 적재 도구를 적지 않는다(gogeulwiki).
- GEUL에 의존하지 않는다. 적재된 DB를 읽어 GEUL 문서로 바꾸는 것은 gogeulwiki가 한다.

```
cmd/wikipg/        덤프 → PostgreSQL 적재
internal/dump/     덤프 읽기(압축 해제·줄 나누기·개체 디코딩)
internal/load/     적재 행 만들기·COPY·판과 체크포인트 기록
sql/wikidata.sql   적재 스키마(1~3절)와 적재 뒤 정리·제약·색인(4절)
sql/repair_reference_snak.sql  repair-reference-snak 작업 표(처리한 hash, 체크포인트)
internal/repair/   reference_snak 복구(entity_raw 대조·교체)
sql/schema.go      wikidata.sql을 embed한 패키지(다른 모듈의 테스트가 import: github.com/geul-org/wikipg/sql)
```

## 적재한 판 (20261019, 적재 완료)

wikiload(지금 wikipg)로 geuldev의 스키마 `wikidata`에 적재한 판본이다. 위 20250908 DB와 달리 판본·해시를 추정하지 않는다. 출처 기록은 `dump` 행에서 읽는다.

| dump_id | 파일 | 개체 | 오류 | 적재 | 적재 도구 커밋 |
|---|---|---|---|---|---|
| 1 | `wikidata-20261019-all.json.bz2` | 121,971,281 | 3 | 2026-10-26 20:46 ~ 10-27 06:28 (KST) | `330a10fe3f2b84dd11994dbcaaa09888d4b16676` |
| 2 | `wikidata-20261021-lexemes.json.bz2` (483,529,071 byte, SHA-256 `c1a98139bb49c616fdb9876206aa285a9ab72bda7f5595a406f269c29c9b9512`) | 1,707,632 | 0 | 2026-10-27 06:29 ~ 06:33 | `a03df92ea5813f2389c58cdbce26aedd04e74539` |

- 오류 3건(`load_error`)은 모두 값에 `\u0000`이 있어 jsonb에 담지 못한 P14457 클레임이다(Q683111, Q5152113, Q4873200). 적재 당시 wikiload(wikiport `330a10f`)는 값을 NULL로 두었다(원문 줄은 `entity_raw`에 있음). 지금 wikipg는 `value_json`에 둔다. 기존 적재분은 `sql/wikidata.sql` 4.0절로 열을 더하고 세 값을 채운다.
- 두 행의 `loader`는 적재 당시 이름 `github.com/geul-org/wikiport/go/cmd/wikiload`다(저장소 wikiport, 명령 wikiload). gogeulwiki(wikigeul)는 옛 이름도 이 도구로 받는다.
- `sql/wikidata.sql` 4절: 2026-10-27 출처 중복 정리(이 정리가 reference_snak을 섞었다: 아래 "출처 snak 결함과 복구")와 PK 전부, `entity_dump_kind_num`·`statement_entity`·`entity_raw_dump_line` 완료. `alias_language_value`가 btree 행 한도(2704바이트)를 넘는 옛 긴 별칭(4751자)으로 실패해, 레이블·별칭·표제어·어형 검색 색인을 앞 250자 `lower(left(value, 250))`로 바꿨다(질의 식은 4.3 주석). 나머지 색인·FK 검증은 재개 대기.

다운로드 정보(dump 1):

| 항목 | 값 |
|---|---|
| 덤프 | Wikidata JSON 전체 덤프 **20261019** (`wikidata-20261019-all.json.bz2`) |
| 보관 파일 | `~/.clari/repos/geul/files/wikidata/wikidata-20261019-all.json.bz2`(2026-10-28 `/mnt/d/geul/wikidata/`에서 옮김), 103,488,697,311 byte, 받기 완료 2026-10-26 15:34 (KST), your.org 미러 |
| MD5 | `cfec3f8ab5fd8f6f0b70f58c590a02da` (공식 `md5sums`와 일치) |
| SHA-1 | `a05b8b809ab80576da6d838c67138ddf401055e6` (공식 `sha1sums`와 일치) |
| SHA-256 | `a249a19d12b1b0e8443248da3583090682be3f20c6bdd6ed1f7e1ae05b6b0774` (직접 계산, 같은 폴더 `.sha256`) |

## 출처 snak 결함과 복구 (2026-10-28)

- **결함.** 위키베이스 출처 hash는 snak 순서와 무관해서, 같은 hash가 개체마다 다른 snaks-order로 나온다. 적재기는 출현마다 `reference_snak` 행을 넣었고, 4.2 정리가 `DISTINCT ON (hash, dump_id, ordinal)`로 **번호마다 따로** 골라 서로 다른 출현의 행을 섞었다(표본 검증 E1: Q1–Q10000 출처 435,343개 중 420개 내용이 다름, 예 Q30 P463 출처에서 `P123 Q7809`가 빠지고 `P813`이 두 번). 스키마가 "hash → 순서 있는 snak 목록 하나"를 가정한 것이 틀렸다. `statement_reference`·`reference`는 영향이 없다(복구 도구가 대조한다).
- **출현 규칙(결정적).** `reference_snak`은 hash마다 **한 출현의 snak 목록을 통째로** 담는다. 그 출현은 덤프 순서상 첫 출현, 곧 `statement_reference`의 `(statement_id, ordinal)`이 가장 작은 출현이다: `entity_raw.line_no`가 가장 작은 개체 → 그 안의 적재기 statement 순서(속성 ID 문자열 순, 같은 속성 안은 배열 순; Lexeme → Form → Sense) → references 배열 순서. 같은 hash의 다른 출현은 snak 집합이 같고 순서만 다를 수 있다. 출현별 순서는 `entity_raw`에만 남는다(`sql/wikidata.sql` reference_snak 주석).
- **적재기 수정.** `reference_snak`에 스테이징 열 `occ_statement_id`·`occ_ordinal`(그 행을 쓴 출현)을 함께 적재한다. `RefSet`은 hash마다 기록한 가장 이른 출현을 기억해, 병렬 파서가 더 이른 출현을 늦게 만나도 다시 쓴다(첫 출현은 반드시 기록; 메모리는 종전처럼 상한). 4.2는 hash마다 가장 작은 출현의 행만 남기고 두 열을 뺀다. 재개 시 `ClearAfter`가 지운 줄의 출현 행도 지운다.
- **복구 도구** `wikipg repair-reference-snak` (이미 4.2를 거친 DB용). `entity_raw`를 줄 순서로 배치마다 읽어(작업자 병렬, 적용은 줄 순서) 출현마다 적재기와 같은 행을 만들고, hash마다 범위 안 첫 출현의 목록을 DB와 비교해 다르면 그 hash의 행을 트랜잭션 안에서 바꾼다. 처리한 hash와 출현은 `repair_reference_snak_done`, 재개 지점은 `repair_reference_snak_checkpoint`(`sql/repair_reference_snak.sql`, 쓰기 실행 때 만든다). 배치의 교체·처리 표·체크포인트가 한 트랜잭션이라 중단해도 정확히 이어진다. 같은 패스에서 `statement_reference`를 원본과 비교해 차이를 보고만 한다.

```sh
go build -o /tmp/wikipg ./cmd/wikipg
/tmp/wikipg repair-reference-snak -dry-run -entities Q1-Q10000      # 비교만(읽기 전용 세션), 보고
/tmp/wikipg repair-reference-snak                                    # 전 구간 쓰기(줄 2부터 끝까지), 중단 뒤 같은 명령으로 재개
/tmp/wikipg repair-reference-snak -from-line 2 -to-line 1000000 -batch 1000 -workers 9 -pause 0 -status 30s
```

  - 범위: `-from-line/-to-line`(기본 전체) 또는 `-entities Q1-Q10000`. 체크포인트는 범위 이름(`lines:0-0`, `entities:Q1-Q10000`)마다 따로다. 구간은 어떤 순서로 돌려도 전 구간을 마치면 hash마다 덤프 첫 출현이 된다(더 이른 출현이 오면 처리 표의 출현보다 앞서므로 다시 처리). 줄 순서대로 전 구간을 돌리면 hash마다 정확히 한 번 처리한다.
  - `-dry-run`은 `default_transaction_read_only=on` 세션으로 아무것도 쓰지 않는다(처리 표 대신 메모리, `-ref-limit`). 쓰기 실행은 스키마별 advisory lock으로 하나만 돈다.
  - 세션은 `jit=off`, `enable_seqscan=off`, `max_parallel_workers_per_gather=0`(동시 질의 수 = 작업자 수, 쓰기 때 적용기 1개 더). 질의는 모두 색인 조회다: entity_raw `(dump_id, line_no)`(구간은 범위 하나, 개체 목록은 줄마다 조회), statement_reference PK(줄마다 키 범위), reference_snak PK와 처리 표 PK(hash마다 한 번), 교체 DELETE(hash마다 PK 조회한 ctid). 목록은 모두 `unnest(배열) CROSS JOIN LATERAL (… OFFSET 0)`로 준다. 수만 개 hash를 `= ANY(배열)`로 주면 플래너가 행 수를 수천만으로 잡아 135GB 병렬 순차 탐색을 골랐다(2026-10-28 실측, dry-run 질의 하나가 8분 넘게 걸려 중단).
  - 기본 작업자 = CPU의 70%의 절반(28코어에서 9), 배치 1000줄. 4절 색인 생성 등 다른 큰 작업과 겹치면 `-workers 7`로 동시 질의를 8개 이하로 둔다. 보고: hash별 판정(same·order·content·missing), 출현별 비교(그 문서의 덤프 순서와 DB 비교, 표본 검증 E1·S10과 같은 단위), statement_reference 차이.
- **색인.** 행 교체(DELETE + INSERT)는 같은 트랜잭션에서 `reference_snak_pkey`·`reference_snak_value_entity`를 함께 갱신하므로 재색인이 필요 없다. PK가 hash 안 번호 중복을 막는다. 교체한 hash 수만큼 죽은 튜플이 생기니 끝난 뒤 `VACUUM (ANALYZE) wikidata.reference_snak`을 돌린다. 작업 표 두 개는 결과를 확인한 뒤 지운다.
- 복구 뒤 gogeulwiki 문서를 다시 만든다(표본 검증: 내용이 바뀌는 문서 214개, 순서까지 최대 370개).

## statement_guid 비고유 (2026-10-28)

4.3의 `statement_guid`는 UNIQUE가 아니다. 위키데이터 원본에 같은 GUID가 두 statement에 쓰인 경우가 있다(Q34433114 P953 두 개, `Q34433114$684DE268-387D-4E42-8BD8-394C5C36D10C`). statement의 키는 `statement_id`이고 guid는 찾기용이다. `sitelink_site_title` 비고유화와 같은 처리다.

## DB 테이블 (`geuldev.wikidata`)

정의는 `sql/wikidata.sql`이다. 원칙과 기존 스키마(0908 판 `geuldev.public`, 2026-10-26 삭제) 대비 개선은 그 머리말에 있다. 이 스키마를 읽는 쪽(gogeulwiki 등)이 쓰는 열은 그쪽 README에 적는다.

## 적재 도구 (`cmd/wikipg`)

덤프를 geuldev의 스키마 `wikidata`(`sql/wikidata.sql`)에 적재한다. 0908 판에서 뽑은 `ko_term`·`lexeme`는 스키마 `wikidata0908`로 옮겼다(2026-10-26). 같은 DB라 stdict 등과 조인할 수 있다. 기존 `geuldev.public`(geul-entity `postgres.go`, 2026-10-26 삭제)을 대신한다. 원칙과 기존 스키마 대비 개선은 `sql/wikidata.sql` 머리말에 있다.

```sh
cd draft/wikipg
go run ./cmd/wikipg -in ~/.clari/repos/geul/files/wikidata/wikidata-20261019-all.json.bz2 -commit "$(git rev-parse HEAD)"
```

- 스키마: `psql -d geuldev -f sql/wikidata.sql`은 1~3절(타입·테이블)만 만든다. 4절(재개 정리·출처 중복 정리·제약·색인)은 주석으로 되어 있고, 적재를 마친 뒤 주석을 풀어 실행한다.
- 테이블: `dump`(파일·URL·크기·MD5·SHA-1·SHA-256·적재 도구 커밋), `entity_raw`(덤프 줄 원문), `load_error`, `entity`, `label`, `description`, `alias`, `sitelink`, `statement`, `qualifier`, `reference`, `statement_reference`, `reference_snak`, Lexeme용 `lexeme`, `lemma`, `form`, `form_representation`, `sense`, `sense_gloss`. 개체 단위 테이블은 모두 `dump_id`를 가진다.
- 기본값은 기계를 30% 남긴다: CPU 70% 중 1/5 압축 해제(lbzip2), 2/5 JSON 파싱, 작성자(DB 연결) 4, 배치 8 MiB(메모리 약 1 GB). 2026-10-26 측정에서 작성자 4·8 MiB가 가장 빨랐다.
- Lexeme은 별도 덤프(`wikidata-YYYYMMDD-lexemes.json.bz2`)라 다른 판 번호로 적재한다: `-in …lexemes.json.bz2 -dump-id 2`. Form·Sense도 `entity`에 행(type `form`·`sense`, kind L)을 두어 그 statement(`entity_id` = `L4-F1`)가 가리킬 개체가 있다. 표제어·어형 표기는 낱말 자체라 `-langs`로 거르지 않고, 뜻풀이(gloss)는 거른다. Lexeme·Form·Sense의 statement 키는 한 줄 안에서 이어진다.
- 체크포인트에서 이어 읽는다(`-from-line 0`). 언어는 `-langs`(기본 `ko,en`, `all`은 전부).
- 플래그 전체: `go run ./cmd/wikipg -h`.
- 테스트는 `sql/wikidata.sql`을 임시 스키마에 만들어 돌린다. gogeulwiki의 DB 테스트는 `github.com/geul-org/wikipg/sql`의 `Schema`를 쓴다.

## 개발 규칙

- filefunc(`codebook.yaml`, `filefunc validate .` 위반 0), TSMA, `go build`·`go vet`·`go test ./...` 통과. DB 테스트는 PG* 환경변수로 geuldev에 임시 스키마를 만들어 돌리고, 접속이 안 되면 건너뛴다.

## 라이선스

우리 저작물은 CC0 1.0([LICENSE](LICENSE)). 제3자 구성요소는 각자의 라이선스를 따르며 고지는 [NOTICE](NOTICE). 위키데이터 덤프는 CC0.

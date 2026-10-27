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
- `sql/wikidata.sql` 4절: 2026-10-27 PK·색인 진행 중(entity·label·alias·qualifier·description PK, `entity_dump_kind_num` 완료, `statement_entity` 진행), 출처 중복 정리·FK·나머지 색인은 그 뒤.

다운로드 정보(dump 1):

| 항목 | 값 |
|---|---|
| 덤프 | Wikidata JSON 전체 덤프 **20261019** (`wikidata-20261019-all.json.bz2`) |
| 보관 파일 | `/mnt/d/geul/wikidata/wikidata-20261019-all.json.bz2`, 103,488,697,311 byte, 받기 완료 2026-10-26 15:34 (KST), your.org 미러 |
| MD5 | `cfec3f8ab5fd8f6f0b70f58c590a02da` (공식 `md5sums`와 일치) |
| SHA-1 | `a05b8b809ab80576da6d838c67138ddf401055e6` (공식 `sha1sums`와 일치) |
| SHA-256 | `a249a19d12b1b0e8443248da3583090682be3f20c6bdd6ed1f7e1ae05b6b0774` (직접 계산, 같은 폴더 `.sha256`) |

## DB 테이블 (`geuldev.wikidata`)

정의는 `sql/wikidata.sql`이다. 원칙과 기존 스키마(0908 판 `geuldev.public`, 2026-10-26 삭제) 대비 개선은 그 머리말에 있다. 이 스키마를 읽는 쪽(gogeulwiki 등)이 쓰는 열은 그쪽 README에 적는다.

## 적재 도구 (`cmd/wikipg`)

덤프를 geuldev의 스키마 `wikidata`(`sql/wikidata.sql`)에 적재한다. 0908 판에서 뽑은 `ko_term`·`lexeme`는 스키마 `wikidata0908`로 옮겼다(2026-10-26). 같은 DB라 stdict 등과 조인할 수 있다. 기존 `geuldev.public`(geul-entity `postgres.go`, 2026-10-26 삭제)을 대신한다. 원칙과 기존 스키마 대비 개선은 `sql/wikidata.sql` 머리말에 있다.

```sh
cd draft/wikipg
go run ./cmd/wikipg -in /mnt/d/geul/wikidata/wikidata-20261019-all.json.bz2 -commit "$(git rev-parse HEAD)"
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

MIT([LICENSE](LICENSE)). 제3자 고지는 [NOTICE](NOTICE). 위키데이터 덤프는 CC0.

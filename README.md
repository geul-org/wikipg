# 위키데이터 → GEUL 포팅

위키데이터 항목을 GEUL 문서로 옮겨 `.well-known/geul/` 아래에 공표한다. 출처·근거 기록(files/geul/sign.md, legacy.md)의 실전 시험대이기도 하다.

## 원본 판본

로컬 PostgreSQL(`geuldev`)에 적재한 위키데이터의 판본이다.

| 항목 | 값 |
|---|---|
| 덤프 | Wikidata JSON 전체 덤프 **20250908** (`wikidata-20250908-all.json`) |
| 판단 근거 | 아래 "판본 판단 근거" |
| 보관 파일 | `~/latest-all.json.bz2`, 98,022,279,070 byte, 받은 시각 2025-09-17 20:49 ~ 2025-09-18 00:32 (KST) |
| DB 적재 시각 | 2025-09-22 ~ 2025-09-23 (`entities.created_at`) |
| 적재 도구 | `github.com/geul-org/geul-entity` `artifacts/wikidata/postgres.go`, 커밋 `4225b3b4297796a79e95c4cda5aa54df3acc4069` |
| 적재 입력 | `/mnt/d/latest-all.json`(압축을 푼 파일). 지금은 없다(소실). 보관 파일과 같은 덤프로 본다 |

### 판본 판단 근거

- 파일 이름이 `latest-all.json.bz2`라 이름만으로는 판본을 알 수 없다. 덤프 사이트는 최근 판만 남겨 2025년 판의 크기·해시를 대조할 수 없었다.
- 덤프 앞부분 6,354개 개체의 `modified` 최댓값이 `2025-09-08T03:01:14Z`, `lastrevid` 최댓값이 `2402035052`다. 자주 편집되는 앞쪽 개체가 이 시각에서 멈춘다.
- DB의 가장 큰 QID는 `Q136198562`이고, 위키데이터 API 기준 첫 판이 `2025-09-10T01:34:49Z`(revid `2403031064`)다. 덤프는 여러 날에 걸쳐 만들어지므로 9월 8일(월)에 시작한 판과 맞는다.
- 9월 17일에 받았지만 그때 `latest`는 아직 0908 판을 가리키고 있었던 것으로 본다(0915 판은 완성 전).

### 빈칸 (legacy.md의 미상·소실·비공개)

| 항목 | 상태 |
|---|---|
| 항목별 판본(`lastrevid`) | **미적재.** DB에 없다. 덤프 판본으로 대신한다 |
| 위키데이터 참조(references) | **미적재.** `triple_references`가 비어 있다 |
| 적재 입력 파일 | **소실.** `/mnt/d/latest-all.json`이 없다 |
| 덤프 판본 | **추정.** 위 근거로 0908 판으로 판단했다 |
| 원본 해시 | **생략.** DB는 소실된 적재 입력에서 읽었으므로 보관 파일 해시는 DB 내용을 증명하지 못하고, 2025년 판은 덤프 사이트에 없어 대조할 수도 없다 |

## DB 테이블 (`geuldev.public`)

| 테이블 | 열 | 대략 행 수 |
|---|---|---|
| `entities` | `id`, `type`(item/property), `created_at`(적재 시각) | 1.17억 |
| `entity_labels` | `entity_id`, `language`, `label` | 7.3억 |
| `entity_aliases` | `entity_id`, `language`, `alias`, `alias_order` | 1.8억 |
| `entity_descriptions` | `entity_id`, `language`, `description` | 33억 |
| `triples` | `id`, `subject`, `property`, `object_value`, `object_type`, `rank` | 17억 |
| `triple_qualifiers` | `triple_id`, `property`, `value`, `datatype` | 4.3억 |
| `triple_references` | `triple_id`, `property`, `value`, `source_type` | 0 |
| `properties_meta` | `property_id`, `datatype`, `label_en`, `description_en`, `usage_count` | 1.3만 |

- `triples.object_value`는 위키데이터 datavalue를 그대로 담는다. 시간 예: `{"time":"+1952-03-11T00:00:00Z","timezone":0,"before":0,"after":0,"precision":11,"calendarmodel":"http://www.wikidata.org/entity/Q1985727"}`.
- `triples.object_type`에는 datavalue 형식(`wikibase-entityid`, `string`, `quantity`, `monolingualtext`, `time`, `globecoordinate`)과 datatype(`wikibase-item`, `globe-coordinate`)이 섞여 있다.

## 작업 범위 (2026-10-25 결정)

- **도구:** Go만 쓴다. DB에서 읽어 라이브러리 `core`로 인코딩하고 `resolve.Publish`로 공표한다. Python은 쓰지 않는다.
- **Issuer:**

  | 이름 | Issuer ID | 대상 |
  |---|---|---|
  | `wikidata-qid.geul.org` | `0x00010000` | 항목(Q) |
  | `wikidata-pid.geul.org` | `0x00010001` | 속성(P) |

- **문서 ID:** QID·PID의 숫자 부분과 같다(Q42 → `0x000000000000002a`). Issuer가 나뉘어 Q와 P의 같은 숫자가 겹치지 않는다.
- **문서 단위:** QID 하나당 `.geul` 문서 하나.
- **대상:** 우선 Q1 ~ Q1000. PID는 필요한 만큼 추가한다.
- **레코드:** TRIPLE 또는 STATEMENT, LABEL, ALIAS를 우선 작성한다.
- **값:** 항목(QID), 문자열, 외부 ID, **날짜**까지. 날짜의 정밀도·달력(율리우스력 등) 문제는 작업하면서 해결책을 함께 찾는다. 그 밖의 값 형식은 변환하지 않았다고 기록한다.
- **언어:** ko, en.
- **명세:** 아직 없는 명세(TRIPLE, STATEMENT, ALIAS 등)는 설계하면서 제안하고 검토를 받는다.

## 출처 기록 (초기 최소 세트, files/geul/legacy.md)

각 문서에 다음을 남긴다.

- 출처: Wikidata 덤프 20250908, 그 항목의 QID
- 받은 시각: 보관 파일을 받은 시각
- 원본 해시: 생략(위 빈칸 표). 생략했다는 사실을 기록한다
- 적재 도구: geul-entity `postgres.go` 커밋 `4225b3b`
- 변환 도구: 포팅 도구 저장소 + 커밋

## 포팅 도구 (`wiki/go`)

```sh
cd draft/t002/wiki/go
# DB 접속은 lib/pq의 PG* 환경변수(PGHOST, PGUSER, PGPASSWORD, PGDATABASE, PGSSLMODE)로 한다.
go run ./cmd/wikiport -root ../.. -from 1 -to 1000 -commit "$(git rev-parse HEAD)" -report report.tsv
```

| 플래그 | 뜻 | 기본값 |
|---|---|---|
| `-root` | `.well-known/geul`을 담은 t002 루트 | `.` |
| `-from`, `-to` | QID 범위 | 1, 1000 |
| `-langs` | 레이블·별칭 언어 | `ko,en` |
| `-commit` | TRANSFORMED-BY에 남길 포팅 도구 커밋 (필수) | |
| `-report` | 개체별 건수·건너뛴 사유·원문 보존 날짜의 사유 TSV (열: id, labels, aliases, statements, skipped, raw) | 없음 |
| `-dsn` | PostgreSQL DSN | 비우면 PG* 환경변수 |

- Q 범위를 먼저 옮기고, 그 문서들이 쓴 속성(Predicate, 값이 속성 개체인 것)을 이어서 옮긴 뒤 두 Issuer의 `documents.geul`을 쓴다. 색인은 이번 실행분과 그 Issuer 디렉터리에 이미 공표된 문서를 모두 담는다(범위를 나눠 실행해도 앞 실행의 이름이 지워지지 않는다).
- 공표는 `resolve.Dir.Publish`다. 내용이 같으면 새 버전을 만들지 않는다. 도구 커밋이 바뀌면 TRANSFORMED-BY가 바뀌어 새 버전이 된다.

## 값 변환

| 위키데이터 | GEUL |
|---|---|
| 항목 `Q…`, 속성 `P…` | WIKIDATA-QID, WIKIDATA-PID Reference |
| Lexeme `L…`, EntitySchema `E…` | 건너뜀 (Reference Codec 없음) |
| string, external-id | UTF8-TEXT |
| time, 그레고리력 일 정밀도(기원전 포함) | GREGORIAN-DATE |
| time, 그레고리력 월·연·10년 정밀도 | EDTF |
| time, 율리우스력 일·월·연 정밀도 | GEDCOM-DATE (`JULIAN 14 OCT 1066`) |
| time, 그 밖(세기 이상 정밀도, 시간대·오차 범위 등) | datavalue JSON 원문을 UTF8-TEXT로. 날짜는 건너뛰지 않는다 |
| novalue, somevalue | 건너뜀 (DB에 값이 비어 있다) |
| quantity, monolingualtext, globecoordinate 등 | 건너뜀 (이번 범위 밖) |

- 레이블·별칭은 UTF8-NFC-TEXT, 언어는 LANGUAGE-BCP47-CODE.
- 클레임은 STATEMENT(Edge Id, Predicate, Subject, Object, Rank). qualifier와 출처는 담지 않는다.
- TRANSFORMED-BY는 임시로 URI Value + 버전 텍스트다. 용어 정의 문서(문서 ID 0 색인)가 생기면 EXTERNAL-ISSUER로 바꾼다.

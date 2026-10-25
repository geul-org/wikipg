# 위키데이터 포팅 명세 제안

> 2026-10-25. 포팅 도구(`wiki/go`)가 쓰는, 아직 t002/docs에 없는 명세다. 검토 후 docs에 반영하거나 고친다. 도구는 이 제안대로 동작한다.

## 1. ALIAS (Edge `0x05`) — 채택, docs/edge/ALIAS.md로 옮김

LABEL 통합으로 비운 `0x05`에 둔다. 레이아웃은 [LABEL](../docs/edge/LABEL.md)과 같다.

```
TAG WORD          0001 0LLL L000 0101
[Length 확장 TAG]
Edge Id           LOCAL Reference 레코드
Language          LANGUAGE-SHORT-CODE 또는 LANGUAGE-BCP47-CODE Value 레코드 1개
Subject           Reference 레코드 1개
Object            Value 또는 LOCAL Reference 레코드 1개
```

- LABEL은 대표 이름(언어마다 많아야 하나), ALIAS는 다른 이름(언어마다 여럿)이다.
- 같은 Subject·언어의 ALIAS 순서는 문서 안 순서다(위키데이터 `alias_order`).

## 2. STATEMENT (Edge `0x0B`) 레이아웃

OVERVIEW 표에는 "TRIPLE과 그 qualifier·출처·rank·scope를 한 레코드로 묶은 축약형"으로만 있다. 이번 포팅은 rank까지만 담는다.

```
TAG WORD          0001 0LLL L000 1011
[Length 확장 TAG]
Edge Id           LOCAL Reference 레코드
Predicate         Reference 레코드 (WIKIDATA-PID 등)
Subject           Reference 레코드
Object            Reference 또는 Value 레코드
Rank              1 WORD (ENUM: 0 normal, 1 preferred, 2 deprecated)
```

- 모든 위키데이터 클레임에 rank가 있어 TRIPLE(`0x0A`) 대신 STATEMENT를 쓴다.
- **qualifier와 출처(references)는 이번에 담지 않았다.** qualifier는 반복 필드라 지금 FIELD 규격으로 선언할 수 없고(edge/DECLARE 미결정 "반복"), 출처는 DB에 적재되지 않았다(README 빈칸). 제안: qualifier는 STATEMENT의 Edge Id를 가리키는 별도 Edge(예: `QUALIFIER` Edge Id, 대상 STATEMENT, Predicate, Object)로 둔다. 서술 묶음 빈 코드 `0x0C`~`0x0F` 중 하나.

## 3. TRANSFORMED-BY 페이로드 — 원칙 결정, 도구 반영 전

결정(2026-10-25): 외부 대상은 (1) 맞는 Reference Codec, (2) 그 대상을 정의한 GEUL 문서(DOC-POINTER·EXTERNAL-ISSUER), (3) 없으면 자기 Issuer의 용어 정의 문서에 LOCAL로 정의해 가리킨다. 각 Issuer의 문서 ID `0`은 용어 정의 **색인** 문서로 예약한다(docs/ISSUER 용어 정의 색인).

- 지금 도구는 아직 임시로 "URI Value + UTF8-TEXT 버전"을 쓴다.
- 반영할 일: geul.org(`0x0000`)의 용어 정의 문서에 geul-entity 적재 도구와 wikiport를 LOCAL로 정의하고(URI는 그 LOCAL에 REFERENCE-NODE로 잇는다), 색인 문서 `0x0000/0x0000000000000000`에서 그 문서를 가리킨 뒤, TRANSFORMED-BY를 `EXTERNAL-ISSUER-1WORD`(`0x0000`, 용어 정의 문서, LOCAL) + 버전 텍스트로 바꾼다.

## 4. 출처 HEAD (포팅 문서)

| Meta | 값 |
|---|---|
| SPEC-VERSION | 2 |
| DOC-ID | DOC-POINTER(`0x00010000` 또는 `0x00010001`, 숫자, 0) |
| SOURCE | URI `https://dumps.wikimedia.org/wikidatawiki/entities/20250908/wikidata-20250908-all.json.bz2` |
| SOURCE | 그 개체의 WIKIDATA-QID/PID Reference |
| ACQUIRED-AT | POSIX-TIMESTAMP-SECONDS `1758109791` (2025-09-17T11:49:51Z, 덤프 파일을 받은 시각) |
| TRANSFORMED-BY | 적재 도구 geul-entity `postgres.go`, 커밋 `4225b3b…` |
| TRANSFORMED-BY | 포팅 도구 `wiki/go`, 커밋(실행 인자) |

- 실행 시각(벽시계)은 쓰지 않는다. 같은 입력과 같은 도구 커밋이면 같은 문서·같은 내용 해시가 나온다.
- 원본 해시는 쓰지 않는다(README 빈칸, 생략 사유).

## 5. 값 변환

| 위키데이터 | GEUL | 비고 |
|---|---|---|
| 항목 `Q…` | WIKIDATA-QID Reference | |
| 속성 `P…` | WIKIDATA-PID Reference | |
| Lexeme `L…`, EntitySchema `E…` | 건너뜀 | Reference Codec 없음 |
| string, external-id | UTF8-TEXT | 정규화 주장 없음 |
| time | GREGORIAN-DATE | 아래 조건일 때만 |
| novalue, somevalue | 건너뜀 | DB에 값이 비어 있다 |
| quantity, monolingualtext, globecoordinate 등 | 건너뜀 | 이번 범위 밖 |

- 레이블·별칭 텍스트는 UTF8-NFC-TEXT(Wikibase가 NFC로 저장). 언어는 LANGUAGE-BCP47-CODE(`en`, `ko`). GEUL 언어 레지스트리가 봉인되지 않아 LANGUAGE-SHORT-CODE는 배정된 코드가 없다.

### 날짜 (작업하면서 해결책을 함께 찾기로 한 부분)

GREGORIAN-DATE가 정확히 말할 수 있는 것만 변환했다: 정밀도 11(일), 그레고리력, timezone 0, before·after 0, 기원후, 실재하는 날, int32 범위.

Q1~Q1000의 time 클레임 1,176건 중 건너뛴 것:

| 사유 | 건수 |
|---|---|
| 정밀도: 연 | 436 |
| 율리우스력 | 68 |
| 정밀도: 세기 | 28 |
| 정밀도: 월 | 16 |
| 정밀도: 10년 | 9 |
| 정밀도: 천년, 그보다 거친 것(2, 3 …) | 18+ |

결정·반영(2026-10-25): 날짜는 **EDTF**(ISO 8601-2) 문자열을 담는 Value로 한다. TIME 계열 `0x24`/`0x25`(짝·홀 바이트 쌍), docs/value/time/EDTF.md. GREGORIAN-DATE로 정확히 쓸 수 있는 날(일 정밀도, 기원전 포함)은 GREGORIAN-DATE로 쓴다(정규형, schema V6). 검증은 `github.com/sfomuseum/go-edtf`(BSD-3-Clause)의 `parser.ParseString`으로 한다(`IsValid`는 문법만 봐서 `2023-02-29`도 통과시킨다).

- 변환(구현·연결됨): 일 → GREGORIAN-DATE, 월 `1952-03`, 연 `1952`, 10년 `195X` → EDTF. 기원전은 위키데이터 역사식 연도(-0044 = 기원전 44년)를 EDTF 천문식(-0043)으로 한 칸 옮긴다. ±9999를 넘는 연도는 `Y` 접두(go-edtf가 아직 지원하지 않아 검증 불가).
- Q1~Q1000 재실행: time 클레임 중 건너뛴 것은 율리우스력 223, 세기 이상 정밀도 4, 시간대·오차 범위 1뿐이다(달력을 정밀도보다 먼저 검사하므로 율리우스력 수가 늘었다).
- 미결정: 세기·천년 이상 정밀도를 EDTF 미상 자릿수(`19XX`)로 옮기는 규칙(위키데이터의 세기 표기 해석), 율리우스력(EDTF는 그레고리력만), before·after.

이전 제안(참고):

1. **정밀도 있는 날짜 Value.** 위키데이터 time과 같은 정보(연 int64, 월·일 0 허용, 정밀도, 달력)를 담는 형식을 ISSUER 형식(GEUL Issuer `0x0000`, TIME 계열 형식 번호)으로 둔다. 연·월 정밀도와 기원전, 율리우스력이 모두 담긴다.
2. **율리우스력**은 그레고리력 날짜로 바꾸지 않는다. 바꾸면 원본이 말한 날과 다른 날이 된다. 1번 형식에 달력 필드로 둔다.
3. before·after(오차 범위)는 1번 형식의 선택 필드로 둔다.

## 6. 문서 구성과 이름

- QID 하나당 문서 하나: `0x00010000/0x000000000000<QID 16진>/`. 속성은 `0x00010001/…`.
- 속성 문서는 Q 문서가 Predicate나 값으로 쓴 속성만 만든다.
- 이름 색인 `documents.geul`: `q42`, `p31` 식.
- Edge Id: LABEL → ALIAS → STATEMENT 순으로 1부터. STATEMENT는 DB의 triple ID 순.

## 7. 시험 실행 결과 (Q1~Q1000, 임시 트리)

| | Q | P |
|---|---|---|
| ID | 1,000 | 3,707 |
| 문서 | 966 | 3,675 |
| LABEL | 1,819 | 5,026 |
| ALIAS | 4,317 | 17,118 |
| STATEMENT | 196,248 | 91,796 |
| 건너뜀 상위 | quantity 40,143, monolingualtext 12,748, globecoordinate 1,315 | monolingualtext 1,890, quantity 1,768 |

- 파일 18,566개(문서, 버전 목록, `h/` 사본, 색인), 약 100MB. 스키마 검사 위반 0건. 소요 약 10초.

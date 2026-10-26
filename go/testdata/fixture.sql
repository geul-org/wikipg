-- fixture.sql — a few entities in the wikidata schema (wiki/wikidata.sql) for wikiport's DB tests.
-- Run after sections 1–3 with search_path set to the test schema. Q6 is left out (never created).
INSERT INTO dump (dump_id, dump_date, file_name, url, size_bytes, sha256, downloaded_at, loader, loader_commit) VALUES
  (1, '2026-10-19', 'wikidata-20261019-all.json.bz2', 'https://dumps.wikimedia.org/wikidatawiki/entities/20261019/wikidata-20261019-all.json.bz2',
   103488697311, 'a249a19d12b1b0e8443248da3583090682be3f20c6bdd6ed1f7e1ae05b6b0774', '2026-10-26 06:34:43+00',
   'github.com/geul-org/wikiport/go/cmd/wikiload', 'bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb'),
  (2, '2026-10-19', 'other.json.bz2', 'https://example.org/other', 1, '00', '2026-10-26 06:34:43+00', 'elsewhere/loader', 'bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb');

INSERT INTO entity (entity_id, line_no, kind, num, type, datatype, dump_id) VALUES
  ('Q1', 2, 'Q', 1, 'item', NULL, 1), ('Q2', 3, 'Q', 2, 'item', NULL, 1), ('Q3', 4, 'Q', 3, 'item', NULL, 1),
  ('Q42', 5, 'Q', 42, 'item', NULL, 1), ('P31', 6, 'P', 31, 'property', 'wikibase-item', 1),
  ('P580', 7, 'P', 580, 'property', 'time', 1), ('Q99', 8, 'Q', 99, 'item', NULL, 2);

INSERT INTO label (entity_id, dump_id, language, value) VALUES
  ('Q1', 1, 'en', 'Universe'), ('Q2', 1, 'en', 'Earth'), ('Q42', 1, 'en', 'Douglas Adams'), ('Q42', 1, 'ko', '더글러스 애덤스'),
  ('P31', 1, 'en', 'instance of'), ('Q42', 2, 'en', 'other dump');
INSERT INTO description (entity_id, dump_id, language, value) VALUES
  ('Q42', 1, 'en', 'English writer and humorist'), ('Q42', 1, 'ko', '영국의 작가');
INSERT INTO alias (entity_id, dump_id, language, ordinal, value) VALUES
  ('Q42', 1, 'en', 1, 'Douglas Noël Adams'), ('Q42', 1, 'en', 0, 'DNA');

INSERT INTO statement (statement_id, guid, entity_id, property, ordinal, rank, snaktype, datatype, value_type, value, value_entity, dump_id) VALUES
  (101, 'Q1$a', 'Q1', 'P31', 0, 'normal', 'value', 'wikibase-item', 'wikibase-entityid', '{"id": "Q36906466", "numeric-id": 36906466, "entity-type": "item"}', 'Q36906466', 1),
  (201, 'Q2$a', 'Q2', 'P31', 0, 'normal', 'value', 'wikibase-item', 'wikibase-entityid', '{"id": "Q3504248", "numeric-id": 3504248, "entity-type": "item"}', 'Q3504248', 1),
  (4201, 'Q42$a', 'Q42', 'P31', 0, 'preferred', 'value', 'wikibase-item', 'wikibase-entityid', '{"id": "Q5", "numeric-id": 5, "entity-type": "item"}', 'Q5', 1),
  (4202, 'Q42$b', 'Q42', 'P569', 0, 'normal', 'value', 'time', 'time', '{"time": "+1952-03-11T00:00:00Z", "after": 0, "before": 0, "timezone": 0, "precision": 11, "calendarmodel": "http://www.wikidata.org/entity/Q1985727"}', NULL, 1),
  (4203, 'Q42$c', 'Q42', 'P1559', 0, 'normal', 'value', 'monolingualtext', 'monolingualtext', '{"text": "Douglas Adams", "language": "en"}', NULL, 1),
  (4204, 'Q42$d', 'Q42', 'P214', 0, 'normal', 'value', 'external-id', 'string', '"113230702"', NULL, 1),
  (4205, 'Q42$e', 'Q42', 'P40', 0, 'normal', 'novalue', 'wikibase-item', NULL, NULL, NULL, 1),
  (4206, 'Q42$f', 'Q42', 'P2067', 0, 'normal', 'value', 'quantity', 'quantity', '{"amount": "+73", "unit": "http://www.wikidata.org/entity/Q11570"}', NULL, 1),
  (9901, 'Q99$a', 'Q99', 'P31', 0, 'normal', 'value', 'wikibase-item', 'wikibase-entityid', '{"id": "Q5", "numeric-id": 5, "entity-type": "item"}', 'Q5', 2);

INSERT INTO qualifier (statement_id, ordinal, property, value_ordinal, snaktype, datatype, value_type, value, value_entity) VALUES
  (4201, 1, 'P642', 0, 'value', 'wikibase-item', 'wikibase-entityid', '{"id": "Q1", "numeric-id": 1, "entity-type": "item"}', 'Q1'),
  (4201, 0, 'P580', 0, 'value', 'time', 'time', '{"time": "+1979-00-00T00:00:00Z", "after": 0, "before": 0, "timezone": 0, "precision": 9, "calendarmodel": "http://www.wikidata.org/entity/Q1985727"}', NULL),
  (4201, 2, 'P582', 0, 'novalue', 'time', NULL, NULL, NULL);

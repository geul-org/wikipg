-- migrate_checkpoint_key.sql — 체크포인트 표의 키를 run에서 (dump_id, run)으로 옮긴다 (2026-10-29)
--
-- 2026-10-29 이전의 reference_snak_order_checkpoint·repair_reference_snak_checkpoint는 PK가 run(구간
-- 이름)뿐이었다. 저장이 ON CONFLICT (run) DO UPDATE SET line_no라서, 다른 판을 같은 구간 이름(전체
-- 구간은 어느 판이든 lines:0-0)으로 쓰면 앞 판 행(dump_id는 그대로)의 line_no를 덮어썼다. 재개는
-- run과 dump_id로 읽으므로 다른 판의 위치에서 시작하거나(줄 건너뜀) 처음부터 다시 했다.
-- 지금 도구는 (dump_id, run)을 키로 쓰고(ON CONFLICT (dump_id, run)), 시작할 때 키가 그렇지 않으면 멈춘다.
--
-- 이 SQL은 있는 표만 옮긴다(없으면 건너뛴다). 옛 표에도 dump_id 열은 NOT NULL로 있어(행을 처음 넣은 판)
-- 행의 dump_id는 그대로 쓴다. 다만 덮어쓰기가 있었다면 line_no가 다른 판의 것일 수 있다. 그래서 행마다
-- dump_id가 dump 표에 있고 line_no가 그 판의 마지막 줄(dump.checkpoint_line)을 넘지 않는지 확인해,
-- 어긋나면 아무것도 바꾸지 않고 멈춘다. 옮긴 뒤 행을 NOTICE로 보인다. 이미 옮긴 표는 그대로 둔다(다시
-- 돌려도 된다). 한 트랜잭션이다.
--
-- 실행(search_path를 대상 스키마로):
--   PGOPTIONS='-c search_path=wikidata' psql -w -v ON_ERROR_STOP=1 -f sql/migrate_checkpoint_key.sql
DO $$
DECLARE
    t    text;
    pk   name;
    cols text;
    r    record;
BEGIN
    FOREACH t IN ARRAY ARRAY['reference_snak_order_checkpoint', 'repair_reference_snak_checkpoint'] LOOP
        IF to_regclass(t) IS NULL THEN
            RAISE NOTICE '%: 없음, 건너뜀', t;
            CONTINUE;
        END IF;
        SELECT c.conname, (SELECT string_agg(a.attname, ',' ORDER BY a.attname) FROM pg_attribute a
                           WHERE a.attrelid = c.conrelid AND a.attnum = ANY (c.conkey))
          INTO pk, cols
          FROM pg_constraint c WHERE c.conrelid = to_regclass(t) AND c.contype = 'p';
        IF cols = 'dump_id,run' THEN
            RAISE NOTICE '%: 이미 (dump_id, run), 건너뜀', t;
            CONTINUE;
        END IF;
        IF to_regclass('dump') IS NOT NULL THEN
            FOR r IN EXECUTE format('SELECT k.dump_id, k.run, k.line_no, d.checkpoint_line FROM %I k LEFT JOIN dump d USING (dump_id) '
                                    'WHERE d.dump_id IS NULL OR k.line_no > d.checkpoint_line', t) LOOP
                RAISE EXCEPTION '%: dump % run % line % — 판이 없거나 그 판의 마지막 줄(%)을 넘는다; 손으로 확인한다',
                    t, r.dump_id, r.run, r.line_no, r.checkpoint_line;
            END LOOP;
        END IF;
        IF pk IS NOT NULL THEN
            EXECUTE format('ALTER TABLE %I DROP CONSTRAINT %I', t, pk);
        END IF;
        EXECUTE format('ALTER TABLE %I ADD PRIMARY KEY (dump_id, run)', t);
        FOR r IN EXECUTE format('SELECT dump_id, run, line_no FROM %I ORDER BY dump_id, run', t) LOOP
            RAISE NOTICE '%: dump % run % line %', t, r.dump_id, r.run, r.line_no;
        END LOOP;
        RAISE NOTICE '%: 키 (%) → (dump_id, run)', t, coalesce(cols, '없음');
    END LOOP;
END
$$;

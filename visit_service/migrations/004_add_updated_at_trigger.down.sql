-- 004_add_update_triggers.down.sql

-- Отключаем вывод сообщений
SET client_min_messages TO WARNING;

-- 1. Удаляем все триггеры для таблиц с полем updated_at
DO $$
DECLARE
    t text;
    trigger_count integer := 0;
BEGIN
    RAISE NOTICE 'Удаление триггеров update_updated_at...';
    
    FOR t IN 
        SELECT table_name 
        FROM information_schema.columns 
        WHERE column_name = 'updated_at' 
        AND table_schema = 'public'
    LOOP
        EXECUTE format('DROP TRIGGER IF EXISTS update_%I_updated_at ON %I;', t, t);
        GET DIAGNOSTICS trigger_count = ROW_COUNT;
        
        IF trigger_count > 0 THEN
            RAISE NOTICE '  ✓ Триггер для таблицы % удален', t;
        END IF;
    END LOOP;
    
    RAISE NOTICE 'Удаление триггеров завершено';
END;
$$ language 'plpgsql';

-- 2. Проверяем, есть ли еще триггеры, использующие функцию
DO $$
DECLARE
    function_usage integer;
BEGIN
    -- Проверяем, используется ли функция где-то еще
    SELECT COUNT(*) INTO function_usage
    FROM pg_trigger
    WHERE tgname LIKE '%updated_at%';
    
    IF function_usage = 0 THEN
        RAISE NOTICE 'Функция больше не используется, можно удалять';
    ELSE
        RAISE NOTICE 'ВНИМАНИЕ: Функция все еще используется % триггерами', function_usage;
    END IF;
END;
$$ language 'plpgsql';

-- 3. Удаляем функцию (только если она существует)
DROP FUNCTION IF EXISTS update_updated_at_column();

-- 4. Проверяем результат
DO $$
BEGIN
    RAISE NOTICE 'Откат миграции завершен';
    RAISE NOTICE 'Функция update_updated_at_column() удалена';
END;
$$ language 'plpgsql';

-- Возвращаем стандартный уровень сообщений
RESET client_min_messages;
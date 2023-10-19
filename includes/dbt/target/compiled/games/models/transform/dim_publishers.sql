-- dim_publishers.sql

-- Crea la dimension de la tabla
WITH publishers_cte AS (
    SELECT DISTINCT
        app_id,
        publishers AS publishers_name
    FROM `pristine-flames-400818`.`games`.`raw_games`
    WHERE publishers IS NOT NULL
)

SELECT
    ROW_NUMBER() OVER() AS publisher_id,
    app_id,
    publishers_name
FROM
    publishers_cte
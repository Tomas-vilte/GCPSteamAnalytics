-- dim_publishers.sql

-- Crea la dimension de la tabla
WITH publishers_cte AS (
    SELECT
        DISTINCT publishers AS publishers_name
    FROM {{ source('games', 'raw_games') }}
)

SELECT
    ROW_NUMBER() OVER () AS publishers_id,
    publishers_name
FROM
    publishers;
-- dim_developers.sql

-- Crea la dimension de la tabla
WITH developers_cte AS (
    SELECT
        developers AS developers_name,
        app_id
    FROM {{ source('games', 'raw_games') }}
    WHERE developers IS NOT NULL
)

SELECT
    ROW_NUMBER() OVER() AS developer_id,
    app_id,
    developers_name
FROM
    developers_cte
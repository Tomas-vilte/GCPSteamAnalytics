-- dim_platforms.sql

WITH platforms_cte AS (
    SELECT
        windows, mac, linux
    FROM {{ source('games', 'raw_games') }}
)

SELECT
    ROW_NUMBER() OVER() AS platform_id,
    windows,
    mac,
    linux
FROM
    platforms_cte

-- dim_release.sql

WITH release_cte AS (
    SELECT
        release_date, coming_soon
    FROM
        {{ source('games', 'raw_games') }}
)

SELECT
    ROW_NUMBER() OVER() AS release_id,
    release_date,
    coming_soon
FROM release_cte
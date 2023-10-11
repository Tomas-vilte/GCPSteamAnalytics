-- dim_games.sql

WITH games_cte AS (
    SELECT
        app_id,
        name,
        description,
        fullgame_app_id,
        fullgame_name,
        type,
        is_free,
        genre_id,
        type_genre
    FROM
        {{ source('games', 'raw_games') }}
)

SELECT
    g.*
FROM games_cte g
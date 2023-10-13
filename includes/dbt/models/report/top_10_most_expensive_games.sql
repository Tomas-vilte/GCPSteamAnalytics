-- Top 10 juegos más caros
SELECT
    app_id,
    game_name,
    max_price,
    formatted_initial_price,
FROM
    {{ ref('fct_game_pricing') }}
ORDER BY max_price DESC
LIMIT 10
-- Top 10 juegos más caros
SELECT
    app_id,
    MAX(initial_price) AS max_price
FROM
    {{ ref('fct_game_pricing') }}
GROUP BY app_id
ORDER BY max_price DESC
LIMIT 10;
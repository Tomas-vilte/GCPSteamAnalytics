-- Top 10 juegos con más impuestos
SELECT
    app_id,
    game_name,
    SUM(total_sales_with_tax) AS total_taxes
FROM
    {{ ref('fct_game_pricing') }}
GROUP BY app_id, game_name
ORDER BY total_taxes DESC, game_name DESC
LIMIT 10
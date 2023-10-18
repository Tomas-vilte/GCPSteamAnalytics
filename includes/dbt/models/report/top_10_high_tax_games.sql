-- Top 10 juegos con más impuestos
SELECT
    app_id,
    game_name,
    SUM(total_sales_with_tax) AS total_taxes,
    formatted_price_with_tax
FROM
    {{ ref('fct_game_pricing') }}
GROUP BY app_id, game_name, formatted_price_with_tax
ORDER BY total_taxes DESC, game_name DESC
LIMIT 10
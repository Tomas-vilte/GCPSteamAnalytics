-- Top 10 descuentos más grandes
SELECT
    app_id,
    SUM(discount_amount) AS total_discounts
FROM
    {{ ref('fct_game_pricing') }}
GROUP BY app_id
ORDER BY total_discounts DESC
LIMIT 10;
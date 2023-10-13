-- Top 10 descuentos más grandes
SELECT
    app_id,
    game_name,
    total_discounts
FROM
    {{ ref('fct_game_pricing') }}
ORDER BY total_discounts DESC
LIMIT 10
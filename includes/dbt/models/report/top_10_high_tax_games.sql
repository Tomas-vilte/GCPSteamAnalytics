-- Top 10 juegos con más impuestos
SELECT
    app_id,
    SUM(price_with_tax) - SUM(final_price) AS total_taxes
FROM
    {{ ref('fct_sales') }}
GROUP BY app_id
ORDER BY total_taxes DESC
LIMIT 10;
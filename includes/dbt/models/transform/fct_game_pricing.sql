-- fct_game_pricing.sql

-- Crea una tabla de hechos para el análisis de precios de juegos
WITH game_sales AS (
    SELECT
        app_id,
        SUM(price_with_tax) AS total_sales,
        SUM(discount_amount) AS total_discounts,
        COUNT(*) AS num_sales,
        MAX(discount_pct) AS max_discount,
        MIN(initial_price) AS min_price,
        MAX(initial_price) AS max_price,
        SUM(price_with_tax) AS total_sales_with_tax,
        SUM(final_price) AS total_sales_with_tax_discounted
    FROM
        {{ ref('dim_price_overview') }}
    GROUP BY app_id
)

SELECT
   g.*
FROM game_sales g
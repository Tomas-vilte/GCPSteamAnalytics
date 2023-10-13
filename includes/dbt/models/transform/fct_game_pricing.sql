-- fct_game_pricing.sql

-- Crea una tabla de hechos para el análisis de precios de juegos
WITH game_sales AS (
    SELECT
        app_id,
        formatted_initial_price,
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
    fct.app_id,
    fct.formatted_initial_price,
    fct.total_sales,
    fct.total_discounts,
    fct.num_sales,
    fct.max_discount,
    fct.min_price,
    fct.max_price,
    fct.total_sales_with_tax,
    fct.total_sales_with_tax_discounted,
    games.name AS game_name,
    games.description AS game_description,
    games.is_free AS game_is_free,
    games.fullgame_app_id,
    games.fullgame_name,
    games.type,
    games.genre_id,
    games.type_genre,
FROM
    game_sales AS fct
INNER JOIN
    {{ ref('dim_games') }} AS games
ON
    fct.app_id = games.app_id
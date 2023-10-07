-- dim_price_overview.sql

-- Calculamos el precio final con impuesto considerando el descuento
WITH price_overview_cte AS (
    SELECT
        *,
        CASE
            WHEN discount_percent > 0 THEN initial_price * (1 - discount_percent / 100)
            ELSE 0.0  
        END AS discounted_price,
        CONCAT('ARS$ ',
            CASE
                WHEN discount_percent > 0 THEN CAST(ROUND(initial_price * (1 - discount_percent / 100), 2) AS STRING)
                ELSE 'ARS$ 0.00'  
            END
        ) AS formatted_discounted_price,
        CONCAT('ARS$ ', CAST(ROUND((initial_price * (1 - discount_percent / 100)) * 1.75, 2) AS STRING)) AS formatted_price_with_tax
    FROM
    {{ source('games', 'raw_games') }}
)

SELECT
    ROW_NUMBER() OVER() AS purchase_id,
    *
FROM price_overview_cte
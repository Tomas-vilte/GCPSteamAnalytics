-- dim_price_overview.sql

-- Calculamos el precio final con impuesto considerando el descuento
WITH price_overview_cte AS (
    SELECT
        initial_price,
        final_price,
        discount_percent AS discount_pct,
        formatted_initial_price,
        formatted_final_price,
        -- Monto del Descuento formateado como decimal con dos decimales
       CAST(
        (CASE
            WHEN discount_percent > 0 THEN initial_price * (discount_percent / 100)
            ELSE 0.0
        END) / 100 AS NUMERIC
    ) AS discount_amount,

    -- Precio del Juego con Impuesto (75%)
    (CASE
    WHEN discount_percent > 0 THEN
        CAST((initial_price * (1 - discount_percent / 100) * 1.75) AS NUMERIC)
    ELSE
        CAST(initial_price * 1.75 AS NUMERIC)
    END) AS price_with_tax
    FROM
    {{ source('games', 'raw_games') }}
)

SELECT
    ROW_NUMBER() OVER() AS purchase_id,
    initial_price,
    final_price,
    discount_pct,
    discount_amount,
    formatted_initial_price,
    formatted_final_price,
    price_with_tax,
    CONCAT('ARS$ ', REPLACE(FORMAT('%.2f', price_with_tax), '.', ',')) AS formatted_price_with_tax
FROM price_overview_cte
WHERE 
    (initial_price IS NOT NULL)
    AND (final_price IS NOT NULL)
    AND (discount_pct IS NOT NULL)
    AND (discount_amount IS NOT NULL)
    AND (formatted_initial_price IS NOT NULL)
    AND (formatted_final_price IS NOT NULL)
    AND (price_with_tax IS NOT NULL)
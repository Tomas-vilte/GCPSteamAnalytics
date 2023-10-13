-- Informe de Ingresos por Plataforma
SELECT
    s.app_id,
    p.platform_id,
    s.total_sales,
    SUM(CASE WHEN p.windows = 1 THEN s.price_with_tax ELSE 0 END) AS total_sales_windows,
    SUM(CASE WHEN p.mac = 1 THEN s.price_with_tax ELSE 0 END) AS total_sales_mac,
    SUM(CASE WHEN p.linux = 1 THEN s.price_with_tax ELSE 0 END) AS total_sales_linux
FROM
    {{ ref('fct_game_pricing') }} s
JOIN
    {{ ref('dim_platforms') }} p ON s.platform_id = p.platform_id
GROUP BY
    s.app_id, p.platform_id, s.total_sales
ORDER BY
    total_sales DESC
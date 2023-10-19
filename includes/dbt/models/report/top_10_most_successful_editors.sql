SELECT
    dp.publisher_id,
    dp.publishers_name,
    SUM(dpo.initial_price) AS total_initial_price,
    dpo.formatted_initial_price
FROM
    {{ ref('dim_publishers') }} dp
JOIN
    {{ ref('dim_price_overview') }} dpo
ON
    dp.app_id = dpo.app_id
GROUP BY
    dp.publisher_id, dp.publishers_name, dpo.formatted_initial_price
ORDER BY
    total_initial_price DESC
LIMIT 10
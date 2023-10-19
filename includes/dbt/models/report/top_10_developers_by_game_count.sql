SELECT
    d.developer_id,
    d.developers_name,
    COUNT(g.app_id) AS total_games_developed
FROM
    {{ ref('dim_developers') }} d
LEFT JOIN
    {{ ref('dim_games') }} g
ON
    d.app_id = g.app_id
GROUP BY
    d.developer_id, d.developers_name
ORDER BY
    total_games_developed DESC
LIMIT 10

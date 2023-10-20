SELECT
    launch_date,
    COUNT(*) AS num_games_launched
FROM
    (
        SELECT
            CASE
                WHEN release_date = 'Próximamente' THEN DATE '2024-02-15'
                WHEN release_date = 'Por confirmarse' THEN DATE '2023-11-25'
                WHEN SAFE.PARSE_DATE('%d %b %Y', release_date) IS NOT NULL THEN SAFE.PARSE_DATE('%d %b %Y', release_date)
                ELSE NULL
            END AS launch_date
        FROM
            {{ ref('dim_release') }}
    ) AS parsed_data
WHERE
    launch_date IS NOT NULL
GROUP BY
    launch_date
HAVING
    num_games_launched >= 5  
ORDER BY
    launch_date
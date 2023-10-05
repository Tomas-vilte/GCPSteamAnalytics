-- dim_support_language.sql

WITH support_language_cte AS (
    SELECT
        interface_languages,
        subtitles_languages,
        fullaudio_languages
    FROM {{ source('games', 'raw_games') }}
)

SELECT
    ROW_NUMBER() OVER() AS language_id,
    interface_languages,
    subtitles_languages,
    fullaudio_languages
FROM support_language_cte
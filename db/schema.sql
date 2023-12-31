USE steamAnalytics;

-- Crear la tabla game
CREATE TABLE IF NOT EXISTS game (
    id INT AUTO_INCREMENT UNIQUE PRIMARY KEY,
    app_id INT NOT NULL UNIQUE,
    name TEXT,
    status TEXT NOT NULL,
    valid BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


-- Crear un índice en el campo app_id de la tabla game
CREATE INDEX idx_appid ON game (app_id);


-- Crear la tabla games_details si aún no existe
CREATE TABLE IF NOT EXISTS games_details (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
    app_id INT UNIQUE,
    name TEXT,
    description TEXT,
    fullgame_app_id INT,
    fullgame_name TEXT,
    type TEXT,
    publishers TEXT, 
    developers TEXT, 
    is_free BOOLEAN,
    interface_languages TEXT, 
    fullaudio_languages TEXT, 
    subtitles_languages TEXT, 
    windows BOOLEAN,
    mac BOOLEAN,
    linux BOOLEAN,
    genre_id TEXT,
    type_genre TEXT,
    release_date TEXT,
    coming_soon BOOLEAN,
    currency TEXT,
    initial_price INT,
    final_price INT,
    discount_percent INT,
    formatted_initial_price TEXT,
    formatted_final_price TEXT
);

CREATE INDEX idx_appid ON games_details (app_id);

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
    app_id INT NOT NULL,
    review_type TEXT,
    recommendation_id INT UNIQUE,
    steam_id TEXT,
    num_games_owned INT,
    num_reviews INT,
    playtime_forever INT,
    playtime_last_two_weeks INT,
    playtime_at_review INT,
    last_played INT,
    language TEXT,
    review_text TEXT,
    timestamp_created INT,
    timestamp_updated INT,
    voted_up BOOLEAN,
    votes_up INT,
    votes_funny INT,
    comment_count INT,
    steam_purchase BOOLEAN,
    received_for_free BOOLEAN,
    written_during_early_access BOOLEAN
);

CREATE INDEX idx_appid ON reviews (app_id);
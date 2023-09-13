USE steamAnalytics;

-- Crear la tabla game
CREATE TABLE IF NOT EXISTS game (
    id SERIAL PRIMARY KEY,
    app_id INT UNIQUE,
    name TEXT,
    status TEXT,
    valid BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
    fullAudio_languages TEXT, 
    subtitles_languages TEXT, 
    windows BOOLEAN,
    mac BOOLEAN,
    linux BOOLEAN,
    genre_id TEXT,
    type_genre TEXT,
    release_date DATE,
    coming_soon BOOLEAN,
    currency TEXT,
    initial_price FLOAT,
    final_price FLOAT,
    discount_percent INT,
    formatted_initial_price TEXT,
    formatted_final_price TEXT
);

CREATE INDEX idx_appid ON games_details (app_id);
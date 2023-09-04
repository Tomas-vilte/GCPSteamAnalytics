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

-- Crear la tabla games_details
CREATE TABLE IF NOT EXISTS games_details (
    id SERIAL PRIMARY KEY,
    app_id INT UNIQUE,
    description TEXT,
    type TEXT,
    name TEXT,
    publishers TEXT, 
    developers TEXT, 
    is_Free BOOLEAN,
    interface_languages TEXT, 
    fullAudio_languages TEXT, 
    subtitles_languages TEXT, 
    windows BOOLEAN,
    mac BOOLEAN,
    linux BOOLEAN,
    release_date TEXT,
    coming_soon BOOLEAN,
    currency TEXT,
    discount_percent INT,
    initial_formatted TEXT,
    final_formatted TEXT
);

CREATE INDEX idx_appid ON games_details (app_id);
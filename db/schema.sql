USE steamAnalytics;

-- Crear la tabla game
CREATE TABLE IF NOT EXISTS game (
    id SERIAL PRIMARY KEY,
    app_id INT,
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
    SteamAppid INT,
    Description TEXT,
    Type TEXT,
    Name TEXT,
    Publishers TEXT[],
    Developers TEXT[],
    isFree BOOLEAN,
    InterfaceLanguages TEXT[],
    FullAudioLanguages TEXT[],
    SubtitlesLanguages TEXT[],
    Windows BOOLEAN,
    Mac BOOLEAN,
    Linux BOOLEAN,
    Date DATE,
    ComingSoon BOOLEAN,
    Currency TEXT,
    DiscountPercent INT,
    InitialFormatted TEXT,
    FinalFormatted TEXT,
    game_id INT REFERENCES games(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_appid ON games_details (SteamAppid);
USE steamAnalytics;

-- Crear la tabla games
CREATE TABLE IF NOT EXISTS games (
    appid INT PRIMARY KEY,
    name VARCHAR(500) NOT NULL
);

-- Crear un índice en el campo appid de la tabla games
CREATE INDEX idx_appid ON games (appid);

-- Crear la tabla state_table
CREATE TABLE IF NOT EXISTS state_table (
    last_appid INT NOT NULL
);

INSERT INTO state_table (last_appid) VALUES (5);

-- Crear la tabla empty_appids
CREATE TABLE IF NOT EXISTS empty_appids (
    appid INT PRIMARY KEY NOT NULL
);

-- Crear un índice en el campo appid de la tabla empty_appids
CREATE INDEX idx_appid_empty ON empty_appids (appid);
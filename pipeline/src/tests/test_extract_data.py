import pytest
from unittest.mock import Mock
from pipeline.src.extract_data import extract_data_games_details
from pipeline.src.cloudsql_connection import DatabaseConnection


@pytest.fixture
def mock_db_connection():
    mock_connection = Mock()
    db_connection = DatabaseConnection()
    db_connection.create_connection = Mock(return_value=mock_connection)
    return mock_connection


def test_extract_data_games_details(mock_db_connection):
    simulated_data = [
        (
            1,
            2572820,
            "Zra Stories",
            "Zra Stories es un juego de exploración basado en historias que combina "
            "aventura con elementos detectivescos. Juega en el papel de una joven "
            "cuidadora de la naturaleza, una maga con la capacidad de controlar los "
            "fenómenos naturales. Un encargo ordinario se convierte en una intensa "
            "investigación al llegar a la isla de Zra.",
            0,
            "",
            "game",
            "Mykhail Konokh",
            "Mykhail Konokh",
            0,
            "No hay soporte para este tipo de idioma",
            "No hay soporte para este tipo de idioma",
            "No hay soporte para este tipo de idioma",
            1,
            0,
            0,
            "25, 23",
            "Aventura, Indie",
            "Por confirmarse",
            1,
            "",
            0.0,
            0.0,
            0,
            "",
            "",
        )
    ]
    simulated_column_names = ["id", "name", "description"]

    cursor_mock = Mock()
    mock_db_connection.cursor.return_value = cursor_mock

    cursor_mock.fetchall.return_value = simulated_data
    cursor_mock.description = [(name,) for name in simulated_column_names]

    data, column_names = extract_data_games_details()

    assert data[0] == simulated_data[0]
    assert column_names[0] == simulated_column_names[0]


def test_extract_data_error(mock_db_connection):
    mock_db_connection.cursor.side_effect = Exception("Error simulado")

    # Llamar a la función
    data, column_names = extract_data_games_details()

    assert data == []
    assert column_names == []


if __name__ == "__main__":
    pytest.main()

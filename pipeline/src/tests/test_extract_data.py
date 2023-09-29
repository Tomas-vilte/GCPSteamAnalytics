import unittest
from unittest.mock import Mock, patch, MagicMock
from pipeline.src.extract_data import extract_data_games_details
from pipeline.src.cloudsql_connection import DatabaseConnection


class TestExtractDataGamesDetails(unittest.TestCase):

    def setUp(self):
        self.mock_connection = Mock()

    def test_extract_data_games_details(self):
        # Crear una instancia de DatabaseConnection con un mock de conexión
        db_connection = DatabaseConnection()
        db_connection.create_connection = Mock(return_value=self.mock_connection)

        # Definir un conjunto de datos simulado y nombres de columna simulados
        simulated_data = [
            (1, 2572820, 'Zra Stories', 'Zra Stories es un juego de exploración basado en historias que combina '
                                        'aventura con elementos detectivescos. Juega en el papel de una joven '
                                        'cuidadora de la naturaleza, una maga con la capacidad de controlar los '
                                        'fenómenos naturales. Un encargo ordinario se convierte en una intensa '
                                        'investigación al llegar a la isla de Zra.',
             0,
             '',
             'game',
             'Mykhail Konokh',
             'Mykhail Konokh',
             0,
             'No hay soporte para este tipo de idioma',
             'No hay soporte para este tipo de idioma',
             'No hay soporte para este tipo de idioma',
             1,
             0,
             0,
             '25, 23',
             'Aventura, Indie',
             'Por confirmarse',
             1,
             '',
             0.0,
             0.0,
             0,
             '',
             '')]
        simulated_column_names = ['id', 'name', 'description']

        cursor_mock = Mock()
        self.mock_connection.cursor.return_value = cursor_mock

        cursor_mock.fetchall.return_value = simulated_data
        cursor_mock.description = [(name,) for name in simulated_column_names]

        data, column_names = extract_data_games_details()

        # Verificar que los datos coincidan con el primer elemento de los datos simulados
        self.assertEqual(data[0], simulated_data[0])
        self.assertEqual(column_names[0], simulated_column_names[0])

    def test_extract_data_error(self):
        with unittest.mock.patch('pipeline.src.cloudsql_connection.DatabaseConnection') as mock_db_connection:
            mock_db_connection.return_value.get_connection.side_effect = Exception("Error simulado")

            # Llamar a la función
            data, column_names = extract_data_games_details()

            # Verificar que data y column_names sean listas vacías
            self.assertEqual(data, [])
            self.assertEqual(column_names, [])


if __name__ == '__main__':
    unittest.main()

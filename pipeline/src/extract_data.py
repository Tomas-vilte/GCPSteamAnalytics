from pipeline.src.cloudsql_connection import DatabaseConnection


def extract_data_games_details() -> list:
    connection = DatabaseConnection().get_connection()
    data = []
    # Realiza operaciones de consulta o inserción aquí si la conexión se estableció con éxito.
    if connection:
        cursor = connection.cursor()
        cursor.execute("SELECT * FROM games_details")
        rows = cursor.fetchall()

        for row in rows:
            data.append(row)

        cursor.close()
        connection.close()
    return data


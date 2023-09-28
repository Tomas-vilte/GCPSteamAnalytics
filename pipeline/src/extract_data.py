from pipeline.src.cloudsql_connection import connect_to_cloud_sql


def extract_data_games_details() -> list:
    connection = connect_to_cloud_sql()
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



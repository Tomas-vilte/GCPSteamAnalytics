from pipeline.src.cloudsql_connection import connect_to_cloud_sql

if __name__ == "__main__":
    connection = connect_to_cloud_sql()

    # Realiza operaciones de consulta o inserción aquí si la conexión se estableció con éxito.
    if connection:
        cursor = connection.cursor()
        cursor.execute("SELECT * FROM game")
        rows = cursor.fetchall()

        for row in rows:
            print(row)

        cursor.close()
        connection.close()


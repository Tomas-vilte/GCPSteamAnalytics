import mysql.connector
from pipeline.src.config.configs import load_env_variables
from pipeline.src.logger.custom_logger import logs


class DatabaseConnection:
    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super(DatabaseConnection, cls).__new__(cls)
            cls._instance.conn = cls._instance.create_connection()
        return cls._instance

    def create_connection(self) -> mysql.connector.connection_cext.CMySQLConnection | None:
        environment: dict = load_env_variables()
        try:
            # Crea una conexión a Cloud SQL
            conn = mysql.connector.connect(
                user=environment["DB_USER"],
                password=environment["DB_PASS"],
                database=environment["DB_NAME"],
                host=environment["DB_HOST"]
            )

            if conn.is_connected():
                logs.info(f'Conexión exitosa a la base de datos {environment["DB_NAME"]} en Cloud SQL.')
                return conn
            else:
                logs.error('No se pudo establecer la conexión.')
                return None

        except mysql.connector.errors.Error as e:
            logs.error(f'Error al conectar a la base de datos: {str(e)}')
            return None

    def get_connection(self) -> mysql.connector.connection_cext.CMySQLConnection | None:
        return self.conn

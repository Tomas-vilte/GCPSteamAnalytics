import pymysql
from typing import Optional
from dags.src.config.configs import load_env_variables
from dags.src.logger.custom_logger import logs


class DatabaseConnection:
    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super(DatabaseConnection, cls).__new__(cls)
            cls._instance.conn = cls._instance.create_connection()
        return cls._instance

    def create_connection(
            self,
    ) -> Optional[pymysql.connections.Connection]:
        environment: dict = load_env_variables()
        try:
            # Crea una conexión a la base de datos MySQL
            conn = pymysql.connect(
                user=environment["DB_USER"],
                password=environment["DB_PASS"],
                database=environment["DB_NAME"],
                host=environment["DB_HOST"],
            )

            if conn.open:
                logs.info(
                    f'Conexión exitosa a la base de datos {environment["DB_NAME"]}.'
                )
                return conn
            else:
                logs.error("No se pudo establecer la conexión.")
                return None

        except pymysql.Error as e:
            logs.error(f"Error al conectar a la base de datos: {str(e)}")
            return None

    def get_connection(self) -> Optional[pymysql.connections.Connection]:
        return self.conn

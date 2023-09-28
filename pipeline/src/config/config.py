import os
from pipeline.src.logger.custom_logger import logs


def load_env_variables() -> dict:
    """
    Carga las variables de entorno relacionadas con la configuración de la base de datos.
    :return: Un diccionario con las variables de entorno de la base de datos.
    :raises ValueError: Si alguna de las variables de entorno está vacía o no está establecida.
    """
    db_variables: dict = {
        "DB_HOST": get_env_variable("DB_HOST"),
        "DB_PASS": get_env_variable("DB_PASS"),
        "DB_NAME": get_env_variable("DB_NAME"),
        "InstanceConnectionName": get_env_variable("INSTANCE_CONNECTION_NAME")
    }
    return db_variables


def get_env_variable(key: str) -> str:
    """
      Obtiene el valor de una variable de entorno.
      :param key: La clave de la variable de entorno.
      :return: El valor de la variable de entorno.
      """
    value: str = os.getenv(key)
    if value == "":
        logs.error(f"Error: {value} variable de entorno no establecida")
        raise ValueError(f"Error: {key} variable de entorno no establecida")
    return value

import os
from pathlib import Path
from dotenv import load_dotenv
from pipeline.src.logger.custom_logger import logs

dir: Path = Path(__file__).resolve().parent


def load_env_variables() -> dict:
    """
    Carga las variables de entorno relacionadas con la configuración de la base de datos.
    :return: Un diccionario con las variables de entorno de la base de datos.
    :raises ValueError: Si alguna de las variables de entorno está vacía o no está establecida.
    """
    db_variables: dict = {
        "DB_PASS": get_env_variable("DB_PASS"),
        "DB_NAME": get_env_variable("DB_NAME"),
        "DB_USER": get_env_variable("DB_USER"),
        "DB_HOST": get_env_variable("DB_HOST")
    }
    return db_variables


dotenvPath = Path(f'{dir}/config.env')
load_dotenv(dotenvPath)


def get_env_variable(key: str) -> str:
    """
    Obtiene el valor de una variable de entorno.
    :param key: La clave de la variable de entorno.
    :return: El valor de la variable de entorno.
    """
    value: str = os.environ.get(key)
    if value is None:
        logs.error(f"Error: {key} variable de entorno no establecida")
        raise ValueError(f"Error: {key} variable de entorno no establecida")
    logs.info(f"Variables de entorno establecidas con exito {key}")
    return value

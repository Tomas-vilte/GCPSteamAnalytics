import logging
from logging import Logger
from pathlib import Path

dir: Path = Path(__file__).resolve().parent.parent.parent
logPath = Path(f"{dir}/logs/myapp.log")


def setup_logger(name: str) -> Logger:
    """
       Configura y devuelve un objeto de registro personalizado.

       Args:
           name (str): El nombre del logger.

       Returns:
           Logger: Un objeto Logger configurado.
    """
    log = logging.getLogger(name)
    log.setLevel(logging.DEBUG)

    formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')

    # Configura el manejador de registro para escribir en un archivo
    file_handler = logging.FileHandler(logPath)
    file_handler.setLevel(logging.DEBUG)
    file_handler.setFormatter(formatter)

    # Configura el manejador de registro para mostrar en la consola
    console_handler = logging.StreamHandler()
    console_handler.setLevel(logging.INFO)
    console_handler.setFormatter(formatter)

    # Agrega los manejadores al logger
    log.addHandler(file_handler)
    log.addHandler(console_handler)

    return log


logs = setup_logger("mi_logger")

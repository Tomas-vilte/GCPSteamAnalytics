import csv
import datetime
from pathlib import Path
from typing import Tuple, List
from cloudsql_connection import DatabaseConnection
from logger.custom_logger import logs

dir: Path = Path(__file__).resolve().parent.parent.parent
dataPath: Path = Path(f"{dir}/includes/dataset/")


def extract_data_games_details() -> Tuple[List, List]:
    data: list = []
    column_names: list = []
    try:
        with DatabaseConnection().get_connection() as connection:
            if connection:
                with connection.cursor() as cursor:
                    cursor.execute("SELECT * FROM games_details")
                    column_names = [column[0] for column in cursor.description]
                    rows = cursor.fetchall()

                    for row in rows:
                        data.append(row)

    except Exception as e:
        logs.info(f"Error al extraer datos: {str(e)}")
    return data, column_names


def save_data_to_csv(data: list, column_names: list) -> None:
    try:
        column_names_lower = [name.lower() for name in column_names]
        with open(
            f"{dataPath}/games_details_{datetime.date.today()}.csv", "w", newline=""
        ) as csv_file:
            csv_writer = csv.writer(csv_file)
            csv_writer.writerow(column_names_lower)
            csv_writer.writerows(data)
    except Exception as e:
        logs.error(f"Error al guardar los datos en CSV:{str(e)}")
    logs.info(f"Se cargaron los datos en el archivo CSV: {csv_file.name}")

from pathlib import Path
from airflow.decorators import dag, task
from datetime import datetime
from airflow.providers.google.cloud.transfers.local_to_gcs import LocalFilesystemToGCSOperator
from airflow.providers.google.cloud.operators.bigquery import BigQueryCreateEmptyDatasetOperator
from astro import sql as aql
from astro.files import File
from astro.sql.table import BaseTable, Metadata
from astro.constants import FileType

dir: Path = Path(__file__).resolve().parent.parent


@dag(
    start_date=datetime(2023, 1, 1),
    schedule='@daily',
    catchup=False,
    tags=["steam_analytics"],
    description="data pipeline ELT",
)
def games_details():
    upload_csv_to_gcs = LocalFilesystemToGCSOperator(
        task_id="upload_csv_to_gcs",
        src=f"{dir}/include/dataset/games_details_2023-09-29.csv",
        dst="raw/games_details_2023-09-29.csv",
        bucket="steamanalytics",
        gcp_conn_id="gcp",
        mime_type="text/csv",
    )

    create_details_dataset = BigQueryCreateEmptyDatasetOperator(
        task_id="create_details_dataset",
        dataset_id="games",
        gcp_conn_id="gcp",
    )

    gcs_to_raw = aql.load_file(
        task_id="gcs_to_raw",
        input_file=File(
            "gs://steamanalytics/raw/games_details_2023-09-29.csv",
            conn_id="gcp",
            filetype=FileType.CSV,
        ),
        output_table=BaseTable(
            name="raw_games",
            conn_id="gcp",
            metadata=Metadata(schema="games")

        ),
        use_native_support=False,
    )


games_details()

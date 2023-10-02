from pathlib import Path
from airflow.decorators import dag, task
from datetime import datetime
from airflow.providers.google.cloud.transfers.local_to_gcs import LocalFilesystemToGCSOperator
from airflow.providers.google.cloud.operators.bigquery import BigQueryCreateEmptyDatasetOperator


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


games_details()

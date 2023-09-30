from airflow.decorators import dag, task
from datetime import datetime
from airflow.providers.google.cloud.transfers.local_to_gcs import (
    LocalFilesystemToGCSOperator,
)


@dag(
    start_date=datetime(2023, 1, 1),
    schedule=None,
    catchup=False,
    tags=["steam_analytics"],
)
def games_details():
    upload_csv_to_gcs = LocalFilesystemToGCSOperator(
        task_id="upload_csv_to_gcs",
        src="../include/dataset/games_details_2023-09-29.csv",
        dst="gs://steam-analytics/games_details_2023-09-29.csv",
        bucket="steam-analytics",
        gcp_conn_id="gcp",
        mime_type="text/csv",
    )


games_details()

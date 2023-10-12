from airflow.decorators import dag
from datetime import datetime
from airflow.operators.python import PythonVirtualenvOperator
from airflow.providers.google.cloud.transfers.local_to_gcs import LocalFilesystemToGCSOperator
from airflow.providers.google.cloud.operators.bigquery import BigQueryCreateEmptyDatasetOperator
from airflow.models.baseoperator import chain
from astro import sql as aql
from astro.files import File
from astro.sql.table import BaseTable, Metadata
from astro.constants import FileType
from cosmos.airflow.task_group import DbtTaskGroup
from cosmos.config import RenderConfig
from cosmos.constants import LoadMode
from includes.dbt.cosmos_config import DBT_PROJECT_CONFIG, DBT_CONFIG
from includes.soda_quality.check_function import check


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
        src="/opt/airflow/includes/dataset/games_details_2023-10-9.csv",
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

    check_load = PythonVirtualenvOperator(
        task_id='soda_core_scan_demodata',
        python_callable=check,
        requirements=["-i https://pypi.cloud.soda.io", "soda-core-bigquery"],
        system_site_packages=False,
        op_args=["check_load", "sources"]
    )

    transform = DbtTaskGroup(
        group_id="transform",
        project_config=DBT_PROJECT_CONFIG,
        profile_config=DBT_CONFIG,
        render_config=RenderConfig(
            load_method=LoadMode.DBT_LS,
            select=['path:models/transform']
        )
    )

    check_transform = PythonVirtualenvOperator(
        task_id="test_transform",
        python_callable=check,
        requirements=["-i https://pypi.cloud.soda.io", "soda-core-bigquery"],
        system_site_packages=False,
        op_args=["check_transform", "transform"]
    )

    report = DbtTaskGroup(
        group_id="report",
        project_config=DBT_PROJECT_CONFIG,
        profile_config=DBT_CONFIG,
        render_config=RenderConfig(
            load_method=LoadMode.DBT_LS,
            select=["path:models/report"]
        )
    )

    check_report = PythonVirtualenvOperator(
        task_id="test_report",
        python_callable=check,
        requirements=["-i https://pypi.cloud.soda.io", "soda-core-bigquery"],
        system_site_packages=False,
        op_args=["check_report", "report"]
    )

    chain(
        upload_csv_to_gcs,
        create_details_dataset,
        gcs_to_raw,
        check_load,
        transform,
        check_transform,
        report,
        check_report,
    )


games_details()

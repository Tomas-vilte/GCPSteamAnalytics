from cosmos import ProfileConfig, ProjectConfig
from pathlib import Path

DBT_CONFIG = ProfileConfig(
    profile_name="games",
    target_name="dev",
    profiles_yml_filepath=Path("/opt/airflow/includes/dbt/profiles.yml"),
)

DBT_PROJECT_CONFIG = ProjectConfig(
    dbt_project_path=Path("/opt/airflow/includes/dbt/")
)

FROM apache/airflow:2.7.1

COPY credentials.json /opt/airflow

RUN pip install pytest
RUN pip install astro-sdk-python
RUN pip install apache-airflow
RUN pip install dbt-bigquery
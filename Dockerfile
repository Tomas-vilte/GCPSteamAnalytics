FROM apache/airflow:2.7.1

COPY credentials.json /opt/airflow
RUN pip install --upgrade pip
RUN pip install --upgrade setuptools
RUN pip cache purge
RUN pip install pytest
RUN pip install astro-sdk-python
RUN pip install apache-airflow
RUN pip install dbt-bigquery
RUN pip install astronomer-cosmos
RUN pip install pymysql
RUN pip install python-dotenv

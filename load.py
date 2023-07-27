import pandas as pd
from sqlalchemy import create_engine, Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

# Configuración de la conexión a MySQL
mysql_config = {
    "host": "localhost",
    "user": "root",
    "password": "root",
    "database": "steamAnalytics",
}

# Ruta del archivo .csv
csv_file_path = "/home/tomi/GCPSteamAnalytics/juegos.csv"

# Nombre de la tabla en la base de datos donde se cargarán los datos
table_name = "games"

# Crear la base de datos y la tabla utilizando SQLAlchemy
Base = declarative_base()

class TuTabla(Base):
    __tablename__ = table_name
    appid = Column(Integer, primary_key=True)
    name = Column(String(255))

# Leer el archivo .csv utilizando pandas
df = pd.read_csv(csv_file_path)

# Reemplazar los valores NaN con una cadena vacía
df.fillna('', inplace=True)

# Establecer la conexión con la base de datos
engine = create_engine(f"mysql+mysqlconnector://{mysql_config['user']}:{mysql_config['password']}@{mysql_config['host']}/{mysql_config['database']}")
Base.metadata.create_all(engine)
Session = sessionmaker(bind=engine)
session = Session()

# Insertar los datos en la tabla utilizando SQLAlchemy
try:
    for index, row in df.iterrows():
        data = TuTabla(appid=row['appid'], name=row['name'])
        session.add(data)

    session.commit()
    session.close()
    print("Los datos se han cargado correctamente en la tabla.")
except Exception as e:
    session.rollback()
    session.close()
    print(f"Ha ocurrido un error: {e}")

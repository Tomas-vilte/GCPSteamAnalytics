import requests
import csv

def get_data():
    response = requests.get("http://localhost:5000/datos").json()
    data = response["data"]["store_items"]
    return data

if __name__ == "__main__":
    data_list = get_data()

    # Guardar los datos en un archivo CSV
    csv_file = "output.csv"
    csv_columns = ["appid", "name"]  # Agrega aquí más columnas si es necesario

    with open(csv_file, 'w', newline='', encoding='utf-8') as csvfile:
        writer = csv.DictWriter(csvfile, fieldnames=csv_columns)
        writer.writeheader()
        for data in data_list:
            # Filtrar el diccionario para que solo contenga las claves definidas en csv_columns
            filtered_data = {key: data[key] for key in csv_columns}
            writer.writerow(filtered_data)

    print("Datos obtenidos y guardados en output.csv")

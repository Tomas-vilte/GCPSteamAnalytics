from flask import Flask, jsonify

app = Flask(__name__)

# Cargar los datos desde el archivo JSON
import json

def cargar_datos():
    with open("../games.json", "r") as archivo:
        datos = json.load(archivo)
    return datos

# Endpoint para obtener todos los datos
@app.route('/datos', methods=['GET'])
def obtener_datos():
    datos = cargar_datos()
    return jsonify(datos)

if __name__ == '__main__':
    app.run(debug=True)

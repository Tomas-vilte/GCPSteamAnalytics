import requests
import concurrent.futures

# URL del endpoint al que deseas enviar las solicitudes
url = "https://steamapigateway-7tl86y4z.uc.gateway.dev/gameDetails/?appid=730"

# Número de solicitudes concurrentes que deseas enviar
num_requests = 140

# Contadores para solicitudes exitosas y fallidas
successful_requests = 0
failed_requests = 0

# Función para enviar una solicitud HTTP a la URL especificada
def send_request(url):
    global successful_requests, failed_requests
    try:
        response = requests.get(url)
        if response.status_code == 200:
            successful_requests += 1
            print(response.json())
            print(f"Solicitud exitosa ({successful_requests} exitosas)")
        else:
            failed_requests += 1
            print(f"Solicitud fallida con código de estado {response.status_code} ({failed_requests} fallidas)")
    except Exception as e:
        failed_requests += 1
        print(f"Error en la solicitud: {str(e)} ({failed_requests} fallidas)")

# Crear un pool de hilos para enviar solicitudes concurrentes
with concurrent.futures.ThreadPoolExecutor(max_workers=num_requests) as executor:
    # Enviar las solicitudes concurrentemente
    futures = [executor.submit(send_request, url) for _ in range(num_requests)]

    # Esperar a que todas las solicitudes se completen
    concurrent.futures.wait(futures)

print(f"Todas las solicitudes han sido procesadas. {successful_requests} exitosas, {failed_requests} fallidas.")

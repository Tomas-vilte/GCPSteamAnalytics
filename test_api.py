import requests
import concurrent.futures

url = "http://localhost:8080/gameDetails/?appid=730"

num_requests = 101

def send_request(url):
    try:
        response = requests.get(url)
        if response.status_code == 200:
            print("Solicitud exitosa")
        else:
            print(f"Solicitud fallida con código de estado {response.status_code}")
    except Exception as e:
        print(f"Error en la solicitud: {str(e)}")

with concurrent.futures.ThreadPoolExecutor(max_workers=num_requests) as executor:
    futures = [executor.submit(send_request, url) for _ in range(num_requests)]

    concurrent.futures.wait(futures)

print("Todas las solicitudes han sido procesadas.")

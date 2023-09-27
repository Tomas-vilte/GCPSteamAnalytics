## Arquitectura de la API

![ArquitecturaAPI](/diagram/ArquitecturaAPI.png)

Aca se describe la arquitectura de la API, que se basa en servicios de Google Cloud Platform (GCP) para ofrecer una API de juegos eficientes y confiable. La arquitectura consta de los siguientes componentes clave:

1. `Google API Gateway`: Google API Gateway sirve como la puerta de enlace principal para la API. Todas las solicitudes de los usuarios se enrutan a través de API Gateway, que luego las dirige a las funciones apropiadas en función de la URL proporcionada por el usuario. API Gateway se encuentra fuera de la VPC, actuando como punto de entrada público.

2. `Cloud Functions`: La lógica está implementada en Cloud Functions. Alojadas dentro de una VPC. Cada función corresponde a un endpoint específico de la API y realiza las operaciones necesarias en función de la solicitud del usuario. Estas funciones interactúan con varios servicios de GCP y se encargan de consultar la API de Steam, actualizar el estado de los juegos en Cloud SQL y gestionar la caché en MemoryStore para Redis.

3. `Cloud SQL`: Use Cloud SQL para almacenar y administrar datos estructurados dentro de la VPC. Esto incluye detalles de juegos y estados de procesamiento. Las Cloud Functions pueden acceder y modificar la base de datos de Cloud SQL según sea necesario para actualizar estados y recuperar datos.

4. `MemoryStore para Redis`: MemoryStore para Redis es nuestro sistema de almacenamiento en caché, ubicado dentro de una VPC. Utilice Redis para almacenar en caché datos consultados con frecuencia, como detalles de juegos o reseñas, con el objetivo de mejorar el rendimiento y reducir la latencia de las respuestas. Esto optimiza la eficiencia de nuestra API.

## Flujo de una Solicitud
Cuando un usuario realiza una solicitud a la API, el flujo general se desarrolla de la siguiente manera:

1. La solicitud llega a través de Google API Gateway (fuera de la VPC).

2. API Gateway enruta la solicitud a la Cloud Function correspondiente (dentro de la VPC).

3. La Cloud Function verifica si los datos están en la memoria caché de Redis (dentro de la VPC).

4. Si los datos están en caché, se devuelven directamente al usuario.

5. Si no están en caché, la Cloud Function realiza una solicitud a la API de Steam para obtener datos actualizados.

6. Luego, la Cloud Function actualiza o consulta Cloud SQL según sea necesario (ambos dentro de la VPC).

7. Los resultados se almacenan en caché en Redis si se espera que sean consultados nuevamente en el futuro.
# GCPSteamAnalytics

## En este repo estare haciendo ELT

# Documentación de la API

Esta es la documentación de la API de Steam-Analytics. Aca vas a encontrar toda la documentacion de los endpoints disponibles y cómo utilizarlos.

## Arquitectura de la API

Si deseas comprender la arquitectura de la API, clickea aca [explicacion de la arquitectura](/arquitecturaAPI.md).


### Límite de velocidad (Rate Limit)

Implemente un límite de velocidad en las solicitudes a la API. El límite actual es de 100 solicitudes por minuto. Si excedes ese límite, vas a recibir una respuesta de error con el código de estado 429.
![Texto alternativo](https://media.tenor.com/jxpZqsCMEWsAAAAC/paraaa-fisura.gif)



## Tabla de Contenidos

- [Obtener Detalles de un Juego por appID](#obtener-detalles-de-un-juego-por-appid)
- [Obtener Lista de Juegos](#obtener-lista-de-juegos)
- [Obtener Reviews de un Juego](#obtener-reviews-de-un-juego)
- [Procesar juegos de Steam](#procesar-juegos-de-steam)
- [Procesar reviews de Steam](#procesar-reviews-de-steam)
- [Explicacion Flujo de Trabajo ELT](#explicacion-flujo-de-trabajo-elt)

## Obtener Detalles de un Juego por appID

### Descripción

Este endpoint te permite obtener los detalles de un juego específico utilizando su `appID`.

- Método: GET

- Endpoint: https://steamapigateway-7tl86y4z.uc.gateway.dev/gameDetails/?appid=730

### Parámetros de Consulta

- `appid` (integer): El ID de la aplicación del juego que deseas consultar. por ej el juego con el appid 730

### Respuesta Exitosa
```json
{
    "id": 106,
    "app_id": 730,
    "name": "Counter-Strike: Global Offensive",
    "description": "Counter-Strike: Global Offensive (CS:GO) amplía el juego de acción por equipos del que fue pionero cuando salió hace más de 20 años. CS:GO incluye nuevos mapas, personajes, armas y modos de juego, y ofrece versiones actualizadas del contenido clásico de Counter-Strike (de_dust2, etc.).",
    "fullgame": {
        "appid": 0,
        "fullgame_name": ""
    },
    "type": "game",
    "publishers": [
        "Valve"
    ],
    "developers": [
        "Valve",
        "Hidden Path Entertainment"
    ],
}
```

## Obtener Lista de Juegos

### Descripción

Este endpoint te permite obtener una lista de juegos.

- Método: GET

- Endpoint: https://steamapigateway-7tl86y4z.uc.gateway.dev/games?page=1&page_size=10&filter=game

### Parámetros de Consulta

- `page` (integer, opcional): El número de página que deseas consultar (predeterminado: 1).

- `page_size` (integer, opcional): El número máximo de juegos por página (predeterminado: 10).

- `filter` (string, opcional): Filtro para tipo de juego (dlc o game) (predeterminado: game).

### Respuesta Exitosa

```json
{
    "metadata": {
        "page": 1,
        "size": 10,
        "total": 107,
        "total_pages": 11
    },
    "games": [
        {
            "id": 1,
            "app_id": 2572820,
            "name": "Zra Stories",
            "description": "Zra Stories es un juego de exploración basado en historias que combina aventura con elementos detectivescos. Juega en el papel de una joven cuidadora de la naturaleza, una maga con la capacidad de controlar los fenómenos naturales. Un encargo ordinario se convierte en una intensa investigación al llegar a la isla de Zra.",
            "fullgame": {
                "appid": 0,
                "fullgame_name": ""
            },
            "type": "game",
            "publishers": [
                "Mykhail Konokh"
            ],
            "developers": [
                "Mykhail Konokh"
            ],
            "is_free": false,
        }
    ]
}
```

## Obtener Reviews de un Juego

### Descripción

Este endpoint te permite obtener reseñas de un juego específico.

- Método: GET

- Endpoint: https://steamapigateway-7tl86y4z.uc.gateway.dev/getReviews/?appid=730&review_type=negative&limit=1

### Parámetros de Consulta

- `appid`: El ID del juego de la que deseas obtener reseñas.

- `review_type`: Puede ser "negative" o "positive" para filtrar reseñas negativas o positivas, respectivamente.

- `limit` (opcional): El numero max de reseñas que deseas obtener (por defecto, se devuelven 10 reseñas si no se especifica).

### Respuesta Exitosa
```json
{
    "metadata": {
        "size": 1,
        "total_review": 30,
        "type_review": "negative"
    },
    "reviews": [
        {
            "app_id": 730,
            "review_type": "negative",
            "recommendation_id": 146379159,
            "author": {
                "steam_id": "76561198020273432",
                "num_games_owned": 0,
                "num_reviews": 35,
                "playtime_forever": 1572,
                "playtime_last_two_weeks": 10,
                "playtime_at_review": 1572,
                "last_played": 1694842983
            },
            "language": "latam",
            "review_text": "UNA PIJA-.",
            "timestamp_created": 1694843029,
            "timestamp_updated": 1694843029,
            "voted_up": false,
            "votes_up": 0,
            "votes_funny": 0,
            "comment_count": 0,
            "steam_purchase": true,
            "received_for_free": false,
            "written_during_early_access": false
        }
    ]
}
```

## Procesar juegos de Steam

### Descripción

Este endpoint proporciona acceso a datos y análisis relacionados con Steam y sus juegos. Puede utilizarse para obtener detalles de juegos,

- Método: POST

- Endpoint: https://steamapigateway-7tl86y4z.uc.gateway.dev/processGames?limit=50

### Parámetros de Consulta

 - `limit` (Obligatorio): Determina cuántos juegos se desean obtener en la respuesta.

### Respuesta Exitosa
```json
{
    "message": "50 registros de datos procesados"
}
```

## Procesar reviews de Steam

### Descripción

Este endpoint se utiliza para procesar las reseñas de usuarios de Steam.

- Método: POST

- Endpoint: https://steamapigateway-7tl86y4z.uc.gateway.dev/processReviews/?review_type=positive&appid=730&limit=5


### Parámetros de Consulta

- `review_type` (Opcional, predeterminado: "negative"): Determina el tipo de reseña (positiva o negativa).

- `appid` (Obligatorio, predeterminado: "10"): El ID del juego al que pertenecen las reseñas.

- `limit` (Opcional, predeterminado: "10"): Límite de resultados a procesar.

### Respuesta Exitosa
```json
{
    "query_summary": {
        "num_reviews": 5
    },
    "success": 1,
    "reviews": [
        {
            "recommendationid": "146818391",
            "author": {
                "steamid": "76561199552951075",
                "num_games_owned": 0,
                "num_reviews": 1,
                "playtime_forever": 744,
                "playtime_last_two_weeks": 744,
                "playtime_at_review": 744,
                "last_played": 1695425996
            },
            "language": "latam",
            "review": "GOOODODOD",
            "timestamp_created": 1695426268,
            "timestamp_updated": 1695426268,
            "voted_up": true,
            "votes_up": 1,
            "votes_funny": 0,
            "comment_count": 0,
            "steam_purchase": true,
            "received_for_free": false,
            "written_during_early_access": false
        },
    ]
}
```


## Flujo de Trabajo de ELT, Validación y Generación de Informes

Este proyecto tiene como objetivo demostrar un flujo de trabajo de Extracción, Transformación y Carga (ELT) que garantiza la calidad de los datos en cada etapa y facilita la generación de informes confiables. Aquí se describe en detalle cómo funciona el flujo de trabajo y las herramientas que utilice.

## Explicacion Flujo de Trabajo ELT
![Arquitectura Data Pipeline](/diagram/Data%20pipeline%20arquitectura.png)

1. **Extracción de Datos de Cloud SQL**: Comenzamos extrayendo datos de nuestras tablas en Cloud SQL. Estos datos actúan como la fuente principal para nuestro análisis.

2. **Guardar Datos como CSV**: Los datos extraídos se guardan en formato CSV. Este formato es adecuado para su posterior procesamiento y carga.

3. **Cargar Datos en BigQuery**: Luego, cargamos los datos en BigQuery. Esta plataforma de análisis escalable nos permite realizar transformaciones y consultas complejas.

4. **Validación de Calidad con Soda Data (Primera Etapa)**: Utilizamos Soda Data para realizar controles de calidad en los datos recién cargados en BigQuery. Esto asegura que los datos sean precisos y cumplan con nuestras expectativas.

5. **Transformación con dbt**: Utilizamos dbt (Data Build Tool) para realizar transformaciones en los datos. Estas transformaciones pueden incluir la creación de modelos dimensionales y tablas de hecho para facilitar la generación de informes.

6. **Validación de Calidad con Soda Data (Segunda Etapa)**: Nuevamente, empleamos Soda Data para verificar la calidad de los datos después de las transformaciones realizadas con dbt.

7. **Ejecución de Modelos dbt para Generación de Informes**: Los modelos dbt se ejecutan para generar informes y vistas basados en las transformaciones previas. Estos informes son fundamentales para nuestro análisis de datos.

8. **Validación de Calidad con Soda Data (Tercera Etapa)**: Una vez más, utilizamos Soda Data para garantizar que los datos utilizados en los informes cumplan con nuestros estándares de calidad.

9. **Visualización con Looker Studio**: Finalmente, empleamos Looker Studio para visualizar y explorar los datos a través de paneles de control y reportes. Looker Studio proporciona una plataforma poderosa para la creación de informes interactivos.

10. **Validación Continua**: A lo largo de todo el flujo de trabajo, realizamos controles de calidad en múltiples etapas utilizando Soda Data para garantizar la integridad de los datos en cada fase.

## Requisitos y Configuración

- Antes de comenzar, asegúrate de que tienes acceso a las siguientes herramientas y servicios:
  - Docker y Docker Compose: Asegúrate de tener Docker y Docker Compose instalados en tu sistema. Puedes encontrar instrucciones detalladas sobre cómo instalarlos en [el sitio web oficial de Docker](https://docs.docker.com/get-docker/) y [Docker Compose](https://docs.docker.com/compose/install/).
  - Configuración de GCP con Apache Airflow: Este proyecto utiliza Google Cloud Platform (GCP) para la extracción y carga de datos. Asegúrate de tener una cuenta de GCP y configurar la conexión de GCP con Apache Airflow siguiendo las directrices de [GCP y Apache Airflow](https://cloud.google.com/composer/docs/how-to/managing/connections).
  - Cuenta en Soda: Para realizar controles de calidad de datos, necesitarás una cuenta en Soda Data. Si aún no tienes una cuenta, puedes registrarte en [el sitio web de Soda Data](https://www.soda.io/).

  - Cloud SQL para la fuente de datos o MySQL en tu local.
  - Google Cloud Storage para el almacenamiento intermedio de archivos CSV.
  - BigQuery para el procesamiento y almacenamiento de datos.
  - Soda Data para las validaciones de calidad.
  - dbt para la transformación de datos.
  - Looker Studio para la visualización.

## Ejecución

Para ejecutar este flujo de trabajo, sigue los siguientes pasos:

1. Asegúrate de tener Docker y Docker Compose instalados en tu sistema.

2. Abre una terminal y navega al directorio raíz de este proyecto.

3. Ejecuta el siguiente comando para iniciar los contenedores de Docker Compose en segundo plano:

   ```bash
   docker-compose up -d

4. Esto iniciará los servicios necesarios para ejecutar el flujo de trabajo. Después de que los contenedores se hayan iniciado correctamente, puedes ejecutar el proceso de ELT utilizando un Dockerfile personalizado que contiene los paquetes y configuraciones necesarios. Ejecuta el siguiente comando:

    ```bash
    cd GCPSteamAnalytics/
    docker build . --tag extended_airflow:2.7.1


## Informes y Consultas DBT

Este directorio contiene una serie de informes y consultas SQL desarrollados en DBT para analizar los datos de juegos y ventas. Cada informe se enfoca en un aspecto específico de los datos y proporciona información valiosa para la toma de decisiones y el análisis de rendimiento.

## Informes Disponibles

1. `revenue_by_platform_and_game.sql`: Este informe analiza los ingresos desglosados por plataforma y juego. Proporciona una visión general de cuánto ingreso se genera en cada plataforma para cada juego.


2. `top_10_discounted_games.sql`: Este informe identifica y muestra los diez juegos con los mayores descuentos en términos de porcentaje de descuento. Es útil para comprender cuáles son los juegos más rebajados en el catálogo.


3. `top_10_high_tax_games.sql`: Este informe muestra los diez juegos que tienen el impuesto más alto aplicado a sus precios. Ayuda a identificar los juegos con mayores impuestos y, por lo tanto, precios más elevados.


4. `top_10_most_expensive_games.sql`: Este informe enumera los diez juegos con los precios más altos en el catálogo. Proporciona información sobre los juegos más caros disponibles.

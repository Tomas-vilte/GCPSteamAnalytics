# GCPSteamAnalytics

## En este repo estare haciendo ELT

# Documentación de la API

Esta es la documentación de la API de Steam-Analytics. Aca vas a encontrar toda la documentacion de los endpoints disponibles y cómo utilizarlos.

### Límite de velocidad (Rate Limit)

Implemente un límite de velocidad en las solicitudes a la API. El límite actual es de 100 solicitudes por minuto. Si excedes ese límite, vas a recibir una respuesta de error con el código de estado 429.
![Texto alternativo](https://media.tenor.com/jxpZqsCMEWsAAAAC/paraaa-fisura.gif)



## Tabla de Contenidos

- [Obtener Detalles de un Juego por appID](#obtener-detalles-de-un-juego-por-appid)
- [Obtener Lista de Juegos](#obtener-lista-de-juegos)
- [Obtener Reviews de un Juego](#obtener-reviews-de-un-juego)
- [Procesar juegos de Steam](#procesar-juegos-de-steam)
- [Procesar reviews de Steam](#procesar-reviews-de-steam)

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
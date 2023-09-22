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
- [Obtener Reviews](#obtener-reviews)

## Obtener Detalles de un Juego por appID

### Descripción

Este endpoint te permite obtener los detalles de un juego específico utilizando su `appID`.

- Endpoint: https://myapi-7tl86y4z.uc.gateway.dev/gameDetails/10

### Parámetros de la URL

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

- Endpoint: https://myapi-7tl86y4z.uc.gateway.dev/games?page=1&page_size=10&filter=dlc

### Parámetros de la URL

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

- Endpoint: Proximamente

### Parámetros de la URL

- `appid`: El ID del juego de la que deseas obtener reseñas.

- `type_reviews`: Puede ser "negative" o "positive" para filtrar reseñas negativas o positivas, respectivamente.

- `limit` (opcional): El numero max de reseñas que deseas obtener (por defecto, se devuelven 10 reseñas si no se especifica).
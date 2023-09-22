# GCPSteamAnalytics

## En este repo estare haciendo ELT

# Documentación de la API

Esta es la documentación de la API de Steam-Analytics. Aca vas a encontrar toda la documentacion de los endpoints disponibles y cómo utilizarlos.

## Tabla de Contenidos

- [Obtener Detalles de un Juego por appID](#obtener-detalles-de-un-juego-por-appid)
- [Obtener Lista de Juegos](#obtener-lista-de-juegos)

## Obtener Detalles de un Juego por appID

### Descripción

Este endpoint te permite obtener los detalles de un juego específico utilizando su `appID`.

https://myapi-7tl86y4z.uc.gateway.dev/gameDetails/10

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

https://myapi-7tl86y4z.uc.gateway.dev/games?page=1&page_size=10&filter=dlc

### Parámetros de la URL

- `page` (integer, opcional): El número de página que deseas consultar (predeterminado: 1).

- `limit` (integer, opcional): El número máximo de juegos por página (predeterminado: 10).

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
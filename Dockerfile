# Utiliza una imagen base que tenga el entorno de Go instalado
FROM golang:1.20

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia todo el contenido de tu proyecto al directorio de trabajo en el contenedor
COPY . .

# Compila tu aplicación
RUN go build -o steam-analytics cmd/main.go


# Exponer el puerto en el que tu aplicación escucha (ajusta según tu aplicación)
EXPOSE 8080

# Especifica el comando que se ejecutará cuando el contenedor se inicie
CMD ["./steam-analytics"]
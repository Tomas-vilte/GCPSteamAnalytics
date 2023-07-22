package utilities

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

// UploadFileToGCS es una función que sube un archivo local a Google Cloud Storage (GCS).
// Recibe tres parámetros:
//   - filePath: Ruta completa del archivo local que queremos subir a GCS.
//   - bucket: Nombre del bucket de GCS al que se subirá el archivo.
//   - object: Ruta dentro del bucket donde se almacenará el archivo.
//     Puede incluir directorios para organizar el archivo en GCS.
//     Ejemplo: "carpeta_destino/nombre_archivo.txt"
//
// La función devuelve un error si ocurre algún problema durante la subida, o nil si fue exitosa.
func UploadFileToGCS(filePath, bucket, object string) error {
	// Iniciamos el cliente de Google Cloud Storage.
	// Este cliente nos permitirá interactuar con GCS.
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("error al crear el cliente de Google Cloud Storage: %w", err)
	}
	defer client.Close()

	// Abrimos el archivo local que queremos subir.
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo %s: %w", filePath, err)
	}
	defer f.Close()

	// Configuramos un contexto con un tiempo límite de 50 segundos para la operación de subida.
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Obtenemos una referencia al objeto en GCS donde queremos almacenar el archivo.
	o := client.Bucket(bucket).Object(object)

	// Configuramos una condición para asegurarnos de que el objeto no exista previamente en GCS.
	// Esto evita sobrescribir datos por accidente si el archivo ya está en el bucket.
	o = o.If(storage.Conditions{DoesNotExist: true})

	// Creamos un escritor (writer) de GCS para el objeto.
	wc := o.NewWriter(ctx)

	// Copiamos el contenido del archivo local al escritor de GCS, lo que hace la subida.
	if _, err := io.Copy(wc, f); err != nil {
		return fmt.Errorf("error al copiar el contenido del archivo al escritor de GCS: %w", err)
	}

	// Cerramos el escritor de GCS después de la subida.
	if err := wc.Close(); err != nil {
		return fmt.Errorf("error al cerrar el escritor de GCS: %w", err)
	}

	// ¡La subida fue exitosa! Devolvemos nil para indicar que no hubo errores.
	fmt.Printf("Archivo %v subido a GCS.\n", filePath)
	return nil
}

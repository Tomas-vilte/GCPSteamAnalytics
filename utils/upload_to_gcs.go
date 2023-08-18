package utils

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
)

// UploadFileToGCS carga un archivo en Google Cloud Storage.
// bucketName es el nombre del depósito.
// fileName es el nombre del archivo.
// content es el contenido del archivo a cargar en formato de cadena.
func UploadFileToGCS(content string, bucketName, fileName string) error {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("error al crear el cliente de Google Cloud Storage: %v", err)
	}

	defer client.Close()

	bucket := client.Bucket(bucketName)
	obj := bucket.Object(fileName)

	wc := obj.NewWriter(ctx)
	if _, err := io.WriteString(wc, content); err != nil {
		wc.Close()
		return fmt.Errorf("error al escribir el contenido del archivo: %v", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("error al cerrar el escritor: %v", err)
	}

	return nil
}

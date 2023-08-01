package zipper

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

// Genera los bytes de un archivo comprimido con los recursos de la aplicacion
func ZipFolder(folder string) ([]byte, error) {
	// Creamos un buffer de bytes para guardar el archivo zip en memoria
	buf := new(bytes.Buffer)

	// Crea un nuevo archivo zip
	zipWriter := zip.NewWriter(buf)

	// Función para agregar archivos y subdirectorios al zip
	addFilesToZip := func(pth string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Si es un directorio, no lo agregamos al zip
		if info.IsDir() {
			return nil
		}

		// Abrimos el archivo
		fileToZip, err := os.Open(pth)
		if err != nil {
			return err
		}
		defer fileToZip.Close()

		// Creamos el header del archivo en el zip
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name, err = filepath.Rel(folder, pth)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate

		// Creamos el archivo dentro del zip
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// Copiamos el contenido del archivo original al archivo del zip
		_, err = io.Copy(writer, fileToZip)
		if err != nil {
			return err
		}

		return nil
	}

	// Recorremos todos los archivos y subdirectorios de la carpeta
	err := filepath.Walk(folder, addFilesToZip)
	if err != nil {
		return nil, err
	}

	// Cerramos el archivo zip
	err = zipWriter.Close()
	if err != nil {
		return nil, err
	}

	// Devolvemos los bytes del archivo zip en memoria
	return buf.Bytes(), nil
}

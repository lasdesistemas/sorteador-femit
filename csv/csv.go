package csv

import (
	"bufio"
	"encoding/csv"
	"os"
)

func ObtenerParticipantesValidos(archivoParticipantes, archivoInscriptes string) [][]string {

	archivo, err := os.OpenFile(archivoParticipantes, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer archivo.Close()

	reader := csv.NewReader(bufio.NewReader(archivo))

	registros, errRead := reader.ReadAll()

	if errRead != nil {
		panic(errRead)
	}

	return registros
}

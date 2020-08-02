package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func ObtenerParticipantesSorteo(archivoInscriptesSorteo, archivoInscriptesConf string) [][]string {

	var participantesSorteo [][]string
	inscriptesConf := leerCsv(archivoInscriptesConf, ';')
	inscriptesSorteo := leerCsv(archivoInscriptesSorteo, ',')

	for _, inscripteSorteo := range inscriptesSorteo {
		if participanteValido(inscripteSorteo, inscriptesConf) {
			participantesSorteo = append(participantesSorteo, inscripteSorteo)
		}
	}

	fmt.Printf("Personas inscriptas al sorteo: %d\n", len(inscriptesSorteo))
	fmt.Printf("Personas que participan del sorteo: %d\n", len(participantesSorteo))

	escribirCsv("participantes_sorteo.csv", participantesSorteo)

	return participantesSorteo
}

func participanteValido(inscripteSorteo []string, inscriptesConf [][]string) bool {

	mailInscripteSorteo := strings.ToLower(strings.Trim(inscripteSorteo[1], " "))

	for _, inscripteConf := range inscriptesConf {
		mailInscripteConf := strings.ToLower(strings.Trim(inscripteConf[12], " "))
		if strings.Compare(mailInscripteConf, mailInscripteSorteo) == 0 {
			return true
		}
	}

	return false

}

func leerCsv(rutaArchivo string, delimitador rune) [][]string {
	archivo, err := os.OpenFile(rutaArchivo, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer archivo.Close()

	reader := csv.NewReader(bufio.NewReader(archivo))
	reader.Comma = delimitador

	registros, errRead := reader.ReadAll()

	if errRead != nil {
		panic(errRead)
	}

	return registros
}

func escribirCsv(rutaArchivo string, registros [][]string) error {

	archivo, err := os.OpenFile(rutaArchivo, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer archivo.Close()

	writer := csv.NewWriter(bufio.NewWriter(archivo))
	writer.Comma = ';'

	return writer.WriteAll(registros)

}

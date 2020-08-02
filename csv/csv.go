package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func ObtenerParticipantesValidos(archivoInscriptesSorteo, archivoInscriptesConf string) [][]string {

	var participantesValidos [][]string
	inscriptesConf := leerCsv(archivoInscriptesConf, ';')
	inscriptesSorteo := leerCsv(archivoInscriptesSorteo, ',')

	for _, inscripteSorteo := range inscriptesSorteo {
		if participanteValido(inscripteSorteo, inscriptesConf) {
			participantesValidos = append(participantesValidos, inscripteSorteo)
		}
	}

	fmt.Printf("Personas inscriptas al sorteo: %d\n", len(inscriptesSorteo))
	fmt.Printf("Participantes con mail válido: %d\n", len(participantesValidos))

	return participantesValidos
}

func participanteValido(inscripteSorteo []string, inscriptesConf [][]string) bool {

	mailInscripteSorteo := strings.ToLower(strings.Trim(inscripteSorteo[1], " "))

	for _, inscripteConf := range inscriptesConf {
		mailInscripteConf := strings.ToLower(strings.Trim(inscripteConf[12], " "))
		fmt.Printf("mail1: %s - mail2: %s\n", mailInscripteSorteo, mailInscripteConf)
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

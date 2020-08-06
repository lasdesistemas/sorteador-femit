package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// ObtenerParticipantesSorteo Obtiene la lista de participantes del sorteo
func ObtenerParticipantesSorteo(archivoInscriptesSorteo, archivoInscriptesConf string) [][]string {

	var participantesSorteo [][]string
	inscriptesConf := leerCsv(archivoInscriptesConf, ';')
	inscriptesSorteo := leerCsv(archivoInscriptesSorteo, ',')

	for _, inscripteSorteo := range inscriptesSorteo {
		if puedeParticipar(inscripteSorteo, inscriptesConf) {
			participantesSorteo = append(participantesSorteo, inscripteSorteo)
		}
	}

	err := escribirCsv("participantes_sorteo.csv", participantesSorteo)

	if err != nil {
		fmt.Println("Warning: No se pudo escribir el archivo csv de resguardo de participantes")
	}

	return participantesSorteo
}

// Una persona puede participar del sorteo solo si está inscripta a la conferencia y no es hombre cis
func puedeParticipar(inscripteSorteo []string, inscriptesConf [][]string) bool {

	var seInscribioALaConf bool
	var inscripteConf []string
	mailInscripteSorteo := strings.ToLower(strings.Trim(inscripteSorteo[1], " "))

	for _, inscripteConf = range inscriptesConf {
		mailInscripteConf := strings.ToLower(strings.Trim(inscripteConf[12], " "))
		seInscribioALaConf = strings.Compare(mailInscripteConf, mailInscripteSorteo) == 0
		if seInscribioALaConf {
			genero := strings.ToLower(strings.Trim(inscripteConf[15], " "))
			if strings.Compare("hombre cis", genero) != 0 {
				return true
			}
			fmt.Printf("------------> NO PARTICIPA: inscripto como hombre cis - %s %s \n", inscripteConf[13], inscripteConf[14])
			break
		}
	}

	if !seInscribioALaConf {
		fmt.Printf("------------> NO PARTICIPA: no está inscripte a FemIT - %s\n", inscripteSorteo[0])
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
		return err
	}

	defer archivo.Close()

	writer := csv.NewWriter(bufio.NewWriter(archivo))
	writer.Comma = ';'

	return writer.WriteAll(registros)

}

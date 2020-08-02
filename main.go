package main

import (
	"fmt"
	"strings"

	"github.com/lasdesistemas/sorteador-femit/csv"
	"github.com/lasdesistemas/sorteador-femit/sorteo"
)

var archivoParticipantes string = "participantes.csv"
var archivoInscriptes string = "inscriptes.csv"

func main() {

	participantesValidos := csv.ObtenerParticipantesValidos(archivoParticipantes, archivoInscriptes)

	ganadore := sorteo.PersonaGanadora(participantesValidos)

	fmt.Printf("La persona ganadora es: %s - %s\n", ganadore[0], ocultarMail(ganadore[1]))
}

func ocultarMail(mail string) string {

	cadenas := strings.Split(mail, "@")
	usuario := cadenas[0]
	dominio := cadenas[1]

	for i := range usuario {
		if i != 0 && i != len(usuario)-1 {
			usuario = ocultarCaracter(usuario, i)
		} else if len(usuario) == 1 || len(usuario) == 2 {
			usuario = ocultarCaracter(usuario, i)
		}
	}

	return usuario + "@" + dominio
}

func ocultarCaracter(cadena string, posicion int) string {
	usuario := []rune(cadena)
	usuario[posicion] = '*'
	return string(usuario)
}

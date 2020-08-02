package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/lasdesistemas/sorteador-femit/csv"
)

var archivo string = "sorteo.csv"

func main() {

	csv := csv.New(archivo)
	participantes := csv.Procesar()

	ganadore := sortear(participantes)

	fmt.Printf("La persona ganadora es: %s - %s\n", ganadore[0], ocultarMail(ganadore[1]))
}

func sortear(participantes [][]string) []string {
	semilla := rand.NewSource(time.Now().UnixNano())
	random := rand.New(semilla)
	elegidx := random.Intn(len(participantes))
	fmt.Printf("La fila elegida es: %d\n", elegidx+1)
	return participantes[elegidx]
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

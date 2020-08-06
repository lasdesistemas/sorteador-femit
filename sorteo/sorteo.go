package sorteo

import (
	"fmt"
	"math/rand"
	"time"
)

func PersonaGanadora(participantes [][]string) ([]string, error) {

	semilla := rand.NewSource(time.Now().UnixNano())
	random := rand.New(semilla)

	if participantes != nil {
		elegidx := random.Intn(len(participantes))
		fmt.Printf("\nLa fila elegida es: %d\n", elegidx+1)
		return participantes[elegidx], nil
	}

	return nil, fmt.Errorf("No hay participantes válidos para realizar el sorteo")
}

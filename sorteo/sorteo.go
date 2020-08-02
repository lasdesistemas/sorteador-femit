package sorteo

import (
	"fmt"
	"math/rand"
	"time"
)

func PersonaGanadora(participantes [][]string) []string {
	semilla := rand.NewSource(time.Now().UnixNano())
	random := rand.New(semilla)
	elegidx := random.Intn(len(participantes))
	fmt.Printf("La fila elegida es: %d\n", elegidx+1)
	return participantes[elegidx]
}

package csv

import (
	"bufio"
	"encoding/csv"
	"os"
)

type Csv struct {
	rutaArchivo string
}

func New(archivo string) *Csv {
	return &Csv{rutaArchivo: archivo}
}

type Participante struct {
	Nombre string
}

func (c *Csv) Procesar() [][]string {

	archivo, err := os.OpenFile(c.rutaArchivo, os.O_RDWR|os.O_CREATE, os.ModePerm)

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

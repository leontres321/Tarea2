package client

import (
	"fmt"
	//	"io/ioutil"
	"math"
	"os"
	//	"strconv"
)

func Run(int_eleccion int8) {
	var eleccion string

	if int_eleccion == 1 {
		eleccion = "carga"
	} else {
		eleccion = "descarga"
	}

	fmt.Printf("Corriendo cliente como: %s", eleccion)

	if int_eleccion == 1 {
		Carga()
	} else {
		Descarga()
	}
}

func Carga() {
	const FILECHUNK = 1 * (1 << 20)

	var nombre string
	var dir_libros string = "./libros/"

	fmt.Printf("\nIngrese el nombre del libro para subir: ")
	fmt.Scanf("%s", &nombre)

	file, err := os.Open(dir_libros + nombre)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	totalParts := uint64(math.Ceil(float64(fileSize) / float64(FILECHUNK)))

	//Para seguir el tutorial al pie de la letra mi nino
	fmt.Printf("\nDividiendo el archivo en %d partes de tamano %d B\n", totalParts, FILECHUNK)
	/*
		for i := uint64(0); i < totalParts; i++ {
			partSize := int(math.Min(FILECHUNK, float64(fileSize-int64(i*FILECHUNK))))
			partBuffer := make([]byte, partSize)

			file.Read(partBuffer)

			fileName := nombre + strconv.FormatUint(i, 10)
			_, err := os.Create(fileName)

			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

			ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
		}
	*/
}

func Descarga() {
}

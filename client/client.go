package client

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

func Run(int_eleccion int8, algoritmo int8) {
	var eleccion string

	if int_eleccion == 1 {
		eleccion = "carga"
	} else {
		eleccion = "descarga"
	}

	fmt.Printf("Corriendo cliente como: %s", eleccion)

	if int_eleccion == 1 {
		Carga(algoritmo)
	} else {
		Descarga(algoritmo)
	}
}

func Carga(algoritmo int8) {
	const FILECHUNK = 250000 //1 * (1 << 18)

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
	fmt.Printf("\nDividiendo el archivo en %d partes de tamano %dkB\n", totalParts, FILECHUNK/1000)

	for i := uint64(0); i < totalParts; i++ {
		partSize := int(math.Min(FILECHUNK, float64(fileSize-int64(i*FILECHUNK))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		fileName := nombre + "_" + strconv.FormatUint(i+1, 10)
		_, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

		//Aca se puede enviar el chunk
	}

}

func Descarga(algoritmo int8) {
	var nombre string
	var partes uint64

	fmt.Println("\nIngrese el nombre del libro a bajar: ")
	fmt.Scanf("%s", &nombre)

	//Preguntar a NameNode si existe y preguntar cantidad de partes

	//Asumiendo que existe
	_, err := os.Create(nombre)

	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	file, err := os.OpenFile(nombre, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	var writePosition int64 = 0

	for j := uint64(0); j < partes; j++ {
		currentChunkFileName := nombre + "_" + strconv.FormatUint(j+1, 10)

		newFileChunk, err := os.Open(currentChunkFileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(5)
		}

		defer newFileChunk.Close()

		chunkInfo, err := newFileChunk.Stat()

		if err != nil {
			fmt.Println(err)
			os.Exit(6)
		}

		var chunkSize int64 = chunkInfo.Size()
		chunkBufferBytes := make([]byte, chunkSize)

		writePosition = writePosition + chunkSize

		reader := bufio.NewReader(newFileChunk)
		_, err = reader.Read(chunkBufferBytes)

		if err != nil {
			fmt.Println(err)
			os.Exit(7)
		}

		_, err = file.Write(chunkBufferBytes)

		if err != nil {
			fmt.Println(err)
			os.Exit(8)
		}

		file.Sync()

		chunkBufferBytes = nil
	}
	fmt.Println("Libro descargado")

	file.Close()
}

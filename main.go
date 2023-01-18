package csvsimultaneo

import (
	"encoding/csv"
	"log"
	"os"
)

//CsvLine is an auxiliar type
//for sending arrays of strings and a filename to write those strings to
type CsvLine struct {
	Text [][]string
	File string
}

//writeExcel just a function to write string arrays to csv file
//closing the file just before ending the function
func writeExcel(texto *[][]string, archivo string) {
	var (
		ExcelT      *os.File
		ExcelWriter *csv.Writer
	)
	ExcelT, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panic(err.Error())
	}
	defer ExcelT.Close()
	ExcelWriter = csv.NewWriter(ExcelT)
	defer ExcelWriter.Flush()
	if err := ExcelWriter.WriteAll(*texto); err != nil {
		log.Fatalln("error escribiendo dato a archivo ", err)

	}

}

//WriteCsv function receives a channel as parameter of typep CsvLine
// takes care of writing to csv file without allowing data tracing
func WriteCsv(textChan <-chan CsvLine) {

	for {
		select {
		case texto := <-textChan:
			writeExcel(&texto.Text, texto.File)
		}
	}
}

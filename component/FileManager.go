package component

import (
	"log"
	"os"
)

func ReadFile(filePath string) string {
	fileContentAsBytes, errorReadFile := os.ReadFile(filePath)
	if errorReadFile != nil {
		log.Println("** ERROR AL INTENTAR LEER EL ARCHIVO **")
		log.Println(errorReadFile.Error())
		return ""
	}
	log.Println("** ARCHIVO LEIDO CORRECTAMENTE **")
	fileContenAsString := string(fileContentAsBytes)
	return fileContenAsString
}

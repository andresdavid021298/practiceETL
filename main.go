package main

import (
	"ETLProject/service"
	"log"
)

func main() {
	log.Println("** INICIO ETL NOMINA **")
	nominaService := service.NominaService{}
	processNominaResponse := nominaService.ProcesarNominaPorArchivo()
	log.Println("** RESPUESTA DEL PROCESO DE NOMINA: {", processNominaResponse, "} **")
	log.Println("** FIN ETL NOMINA **")
}

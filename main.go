package main

import (
	"ETLProject/config"
	"ETLProject/service"
	"log"
)

func main() {
	config.InitAWS()
	log.Println("** INICIO ETL NOMINA **")
	nominaService := service.NominaService{}
	processNominaResponse := nominaService.ProcesarNominaPorArchivo()
	log.Println("** RESPUESTA DEL PROCESO DE NOMINA: {", processNominaResponse, "} **")
	log.Println("** FIN ETL NOMINA **")
}

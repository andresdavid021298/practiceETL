package service

import (
	"ETLProject/component"
	"ETLProject/model"
	"ETLProject/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type NominaService struct {
}

func (service NominaService) ProcesarNominaPorArchivo() string {
	pathFileNomina := util.LoadProperty("nomina.file.path")
	fileContent := component.ReadFile(pathFileNomina)
	nominaRecords := strings.Split(fileContent, "\n")
	log.Println("** REGISTROS LEIDOS DEL ARCHIVO DE NOMINA:", len(nominaRecords), "**")
	empleados := convertirLineasStringToListaEmpleados(nominaRecords)
	empleadosPorDepartamento := dividirEmpleadosPorDepartamento(empleados)
	var sb strings.Builder
	for depto, lista := range empleadosPorDepartamento {
		sb.WriteString(fmt.Sprintf("Departamento: %s - %d empleado(s)\n", depto, len(lista)))
	}
	mailTo := util.LoadProperty("to.mail.address")
	util.SendMail(mailTo, "Reporte de Empleados por Departamento", sb.String())
	return "Ok"
}

func convertirLineasStringToListaEmpleados(lines []string) []model.Empleado {
	var empleados []model.Empleado
	for _, line := range lines {
		empleadoData := strings.Split(line, ";")
		if hasEmptyInformation(line) {
			continue
		}
		idEmpleado, _ := strconv.Atoi(empleadoData[0])
		nombre := empleadoData[1]
		cargo := empleadoData[2]
		departamento := empleadoData[3]
		fechaIngreso := empleadoData[4]
		tipoContrato := empleadoData[5]
		salarioBase, _ := strconv.Atoi(empleadoData[6])
		salarioNeto, _ := strconv.Atoi(empleadoData[7])
		email := empleadoData[8]
		empleado := model.NewEmpleado(int64(idEmpleado), nombre, cargo, departamento, fechaIngreso, tipoContrato, int64(salarioBase), int64(salarioNeto), email)
		empleados = append(empleados, empleado)
	}
	return empleados
}

func hasEmptyInformation(line string) bool {
	empleadoData := strings.Split(line, ";")
	for _, data := range empleadoData {
		if strings.TrimSpace(data) == "" {
			return true
		}
	}
	return false
}

func dividirEmpleadosPorDepartamento(empleados []model.Empleado) map[string][]model.Empleado {
	empleadosPorDepartamentoMap := make(map[string][]model.Empleado)
	for _, e := range empleados {
		empleadosPorDepartamentoMap[e.Departamento] = append(empleadosPorDepartamentoMap[e.Departamento], e)
	}
	return empleadosPorDepartamentoMap
}

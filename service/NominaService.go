package service

import (
	"ETLProject/component"
	"ETLProject/model"
	"ETLProject/util"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type NominaService struct {
}

func (service NominaService) ProcesarNominaPorArchivo() string {
	empleados := leerEmpleadosSegunArchivo()
	empleadosPorDepartamento := dividirEmpleadosPorDepartamento(empleados)
	var sb strings.Builder
	for depto, lista := range empleadosPorDepartamento {
		sb.WriteString(fmt.Sprintf("Departamento: %s - %d empleado(s)\n", depto, len(lista)))
	}
	mailTo := util.LoadProperty("to.mail.address")
	util.SendMail(mailTo, "Reporte de Empleados por Departamento", sb.String())
	bucketName := util.LoadProperty("aws.s3.bucket.nomina")
	seCreoExcelExitoso := crearExcelPorDepartamento(empleadosPorDepartamento)
	if seCreoExcelExitoso {
		component.UploadFileToS3(bucketName, util.LoadProperty("nomina.file.excel.path"))
	}
	return "Ok"
}

func leerEmpleadosSegunArchivo() []model.Empleado {
	pathFileNomina := util.LoadProperty("nomina.file.path")
	fileContent := component.ReadFile(pathFileNomina)
	nominaRecords := strings.Split(fileContent, "\n")
	log.Println("** REGISTROS LEIDOS DEL ARCHIVO DE NOMINA:", len(nominaRecords), "**")
	empleados := convertirLineasStringToListaEmpleados(nominaRecords)
	return empleados
}

func convertirLineasStringToListaEmpleados(lines []string) []model.Empleado {
	var empleados []model.Empleado
	for _, line := range lines {
		empleadoData := strings.Split(line, ";")
		if tieneDataVacia(line) {
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

func tieneDataVacia(line string) bool {
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

func crearExcelPorDepartamento(empleadosPorDepartamento map[string][]model.Empleado) bool {
	log.Println("** CREANDO ARCHIVO EXCEL POR DEPARTAMENTO **")

	excelFile := excelize.NewFile()

	excelFile.SetCellValue("Sheet1", "A1", "Departamento")
	excelFile.SetCellValue("Sheet1", "B1", "Cantidad Empleados")

	numeroFila := 2
	for departamento, empleados := range empleadosPorDepartamento {
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%d", numeroFila), departamento)
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%d", numeroFila), len(empleados))
		numeroFila++
	}

	nominaExcelPath := util.LoadProperty("nomina.file.excel.path")
	if err := excelFile.SaveAs(nominaExcelPath); err != nil {
		log.Println("** ERROR GENERANDO EL ARCHIVO EXCEL **")
		return false
	}

	log.Println("** ARCHIVO NOMINA GENERADO CORRECTAMENTE **")
	return true
}

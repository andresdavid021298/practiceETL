package model

import "fmt"

type Empleado struct {
	IdEmpleado   int64
	Nombre       string
	Cargo        string
	Departamento string
	FechaIngreso string
	TipoContrato string
	SalarioBase  int64
	SalarioNeto  int64
	Email        string
}

func NewEmpleado(idEmpleado int64, nombre string, cargo string, departamento string, fechaIngreso string, tipoContrato string, salarioBase int64, salarioNeto int64, email string) Empleado {
	return Empleado{
		IdEmpleado:   idEmpleado,
		Nombre:       nombre,
		Cargo:        cargo,
		Departamento: departamento,
		FechaIngreso: fechaIngreso,
		TipoContrato: tipoContrato,
		SalarioBase:  salarioBase,
		SalarioNeto:  salarioNeto,
		Email:        email,
	}
}

func (e Empleado) ToString() string {
	return fmt.Sprintf(
		"ID: %d, Nombre: %s, Cargo: %s, Departamento: %s, Ingreso: %s, Contrato: %s, Salario Base: %d, Salario Neto: %d, Email: %s",
		e.IdEmpleado, e.Nombre, e.Cargo, e.Departamento, e.FechaIngreso, e.TipoContrato, e.SalarioBase, e.SalarioNeto, e.Email)
}

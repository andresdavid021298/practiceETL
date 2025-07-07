# 🛠️ ETL de Nómina en Go

Este proyecto es una herramienta ETL (Extract, Transform, Load) desarrollada en Go, que procesa información de nómina desde un archivo `.csv`, limpia los datos, organiza la información por departamento y realiza dos acciones finales:

- Envía por correo electrónico un resumen de la nómina.
- Carga un archivo Excel con el detalle a un bucket en AWS S3.

---

## 📁 Estructura del Archivo CSV

El archivo de entrada contiene la siguiente estructura:

```csv
(SIN ENCABEZADOS) idEmpleado,nombre,cargo,departamento,fechaIngreso,tipoContrato,salarioBase,salarioNeto,email (SIN ENCABEZADOS)
1,Juan Pérez,Desarrollador,IT,2020-01-10,Indefinido,3000000,2500000,juan.perez@empresa.com
...
```
## 📦 Paquetes
| Funcionalidad             | Paquete                        | Instalación                                                                                      |
| ------------------------- | ------------------------------ | ------------------------------------------------------------------------------------------------ |
| Variables de entorno      | `github.com/joho/godotenv`     | `go get github.com/joho/godotenv`                                                                |
| Lectura de archivos       | `os` (incluido en Go)          | N/A                                                                                              |
| Validaciones y utilidades | `strings` (incluido en Go)     | N/A                                                                                              |
| Envío de correos          | `gopkg.in/gomail.v2`           | `go get gopkg.in/gomail.v2`                                                                      |
| Creación de Excel         | `github.com/xuri/excelize/v2`  | `go get github.com/xuri/excelize/v2`                                                             |
| Cliente AWS S3            | `github.com/aws/aws-sdk-go-v2` | `go get github.com/aws/aws-sdk-go-v2/config`<br>`go get github.com/aws/aws-sdk-go-v2/service/s3` |

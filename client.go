package main

import (
	"fmt"
	"net/rpc"
)

type general_cliente struct {
	Nombre string
	Materia string
	Calificacion float64
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println ( "1) Agregar/Asignar alumno y materia" )
		fmt.Println ( "2) Ver materias por Alumno" )
		fmt.Println ( "3) Ver alumnos por materia" )
		fmt.Println ( "4) Promedio de Materia" )
		fmt.Println ( "5) Promedio de Alumno" )
		fmt.Println ( "6) Promedio General" )
		fmt.Println ( "0) Salir" )
		fmt.Scanln(&op)

		switch op {
		case 1:
			var nombre string
			fmt.Print ( "Nombre alumno: " )
			fmt.Scan(&nombre)
			var materia string
			fmt.Print ( "Nombre de la materia: " )
			fmt.Scanln(&materia)
			var calificacion float64
			fmt.Print( "Calificacion: " )
			fmt.Scan(&calificacion)
			alumnoNuevo := general_cliente{nombre, materia, calificacion}
			var result string
			err = c.Call( "Server.NuevoAlumno" , alumnoNuevo, &result)
			if err != nil {
				fmt. Println (err)
			} else {
				fmt. Println (result)
}
		case 2:
			
		case 3:
		
		case 4:

		case 5:

		case 6:

		case 0:
			return
		}
	}
}

func main() {
	client()
}
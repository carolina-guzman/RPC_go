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
		fmt.Println("---------------MENU-----------------")
		fmt.Println("1. Agregar Alumno con materia" )
		fmt.Println("2. Materias de un alumno" )
		fmt.Println("3. Mostrar alumnos por materia" )
		fmt.Println("4. Promedio de Materia" )
		fmt.Println("5. Promedio de Alumno" )
		fmt.Println("6. Promedio General" )
		fmt.Println("0. Salir" )
		fmt.Scanln(&op)

		switch op {
		case 1:
			fmt.Println("----Agregar Alumno----")
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
				fmt.Println (err)
			} else {
				fmt.Println (result)
			}
		case 2:
			fmt.Println("----Materias de un Alumno----")
			var nombre string
			fmt.Print ( "Nombre alumno: " )
			fmt.Scan(&nombre)
			var result string
			err = c.Call( "Server.MateriasPorAlumno" , nombre, &result)
			if err != nil {
				fmt.Println (err)
			} else {
				fmt.Println (result)
			}	
		case 3:
			fmt.Println("----Mostrar alumnos por materia----")
			var materia string
			fmt.Print ( "Nombre de la materia: " )
			fmt.Scan(&materia)
			var result string
			err = c.Call( "Server.AlumnosPorMateria" , materia , &result)
			if err != nil {
				fmt.Println (err)
			} else {
				fmt.Println (result)
			}
		case 4:
			fmt.Println("----Promedio de materia----")
			var materia string
			fmt.Print ( "Nombre de la materia: " )
			fmt.Scan(&materia)
			var result string
			err = c.Call( "Server.PromedioPorMateria" , materia , &result)
			if err != nil {
				fmt.Println (err)
			} else {
				fmt.Println (result)
			}
		case 5:
			fmt.Println("----Promedio de Alumno----")
			var nombre string
			fmt.Print ( "Nombre del alumno: " )
			fmt.Scan(&nombre)
			var result string
			err = c.Call( "Server.PromedioPorAlumno" , nombre , &result)
			if err != nil {
				fmt.Println (err)
			} else {
				fmt.Println (result)
			}
		case 6:
			fmt.Println("----Promedio de todos los alumnos----")
			var materia string
			fmt.Print ( "Nombre de la materia: " )
			fmt.Scan(&materia)
			var result string
			err = c.Call( "Server.PromedioGeneral" , materia , &result)
			if err != nil {
				fmt.Println (err)
			} else {
				fmt.Println (result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
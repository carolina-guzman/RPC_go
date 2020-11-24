package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)



type general struct {
	Nombre string
	Materia string
	Calificacion float64
}

type materia struct {
	Alumno string
	calificacion float64
}

type alumno struct {
	Materia string
	calificacion float64
}

var materias map [ string ][]materia
var alumnos map [ string ][]alumno

type Server struct{}

func ValidarNoRepetido(nombre string , materia string ) string {
	for mat , als := range materias {
		if mat == materia {
			for _ , al := range als {
				if al.Alumno == nombre {
					return "Ya existe esta calificacion"
				}
			}
		}
	}
	return "error"
}

func (this *Server) NuevoAlumno (a general, reply *string) error {
	err := ValidarNoRepetido (a.Nombre, a.Materia)
	if err != "error" {
		return errors. New (err)
	}
	alumnos[a.Nombre] = append (alumnos[a.Nombre], alumno{a.Materia, a.Calificacion})
	materias[a.Materia] = append (materias[a.Materia], materia{a.Nombre, a.Calificacion})
	* reply = a.Nombre + "asignado a " + a.Materia 
	return nil
}

func (this *Server) MateriasPorAlumno(name string, reply *string) error {
		* reply = ""
		_ , existe := alumnos[name]
		if existe{
			for _ , mat := range alumnos[name] {
				* reply += mat.Materia + " \n "
			}
			return nil
		}
		return errors.New( "Alumno "+name+ "no registrado" )
}	

func (this *Server) AlumnosPorMateria (nombre string , reply * string ) error {
		* reply = ""
		_ , existe := materias[nombre]
		if existe {
			for _ , mat := range materias[nombre] {
				*reply += mat.Alumno + " \n "
			}
			return nil
		}
		return errors.New( "La materia "+nombre+"no existe" )
}


func (this *Server) PromedioPorAlumno (nombre string , reply * float64 ) error {
	* reply = 0.0
	var counter float64
	counter = 0.0
	_ , existe := alumnos[nombre]
	if existe {
		for _ , mat := range alumnos[nombre] {
			*reply += mat.calificacion
			counter++
		}
		* reply = *reply / counter
		return nil
	}
	return errors. New ( "El alumno "+nombre+"no existe" )
}

func (this *Server) PromedioPorMateria(materiaNombre string, reply *float64) error {
	* reply = 0.0
	var counter float64
	counter = 0.0
	_ , existe := materias[materiaNombre]
	if existe {
		for _ , mat := range materias[materiaNombre] {
			*reply += mat.calificacion
			counter++
		}
		* reply = *reply / counter
		return nil
	}
	return errors. New ( "No existe la materia en el registro" )
}


func (this *Server) PromedioGeneral (nombre string , reply * float64 ) error {
		* reply = 0.0
		var total float64
		var counter float64
		var mat float64
		for name := range alumnos {
			counter = 0
			mat = 0
			for _ , alumno := range alumnos[name] {
				counter += alumno.calificacion
				mat++
			}
			total += counter / mat
		}
		* reply = total / float64 ( len (alumnos))
		return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {

	go server()

	var input string
	fmt.Scanln(&input)
}

  
package main

import (
	
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)



type Server struct{
	Materias map[string]map[string]float64

	Alumnos map[string]map[string]float64

	SeInicializo bool
}

func (this* Server) Constructor(s string, reply *string) error {
	if this.SeInicializo == false{
		this.Materias = make(map[string]map[string]float64)
		this.Alumnos = make(map[string]map[string]float64)
		
		this.SeInicializo = true
	}

	*reply = "todoBien"

	return nil
}

func (this *Server) AgregarCalificacion(s []string, reply *string) error {
	v, err := this.Materias[s[0]]
	
	*reply = "Se agrego correctamente"

	if err == false{
		alumno := make(map[string]float64)
		
		f2, _ := strconv.ParseFloat(s[2], 8)
		alumno[s[1]] = f2

		// creacion de una materia
		this.Materias[s[0]] = alumno
	} else {
		_, err2 := v[s[1]]
		if err2 == false{
			alumno := make(map[string]float64)

			f2, _ := strconv.ParseFloat(s[2], 8)

			for auxAlumno, calificacion := range this.Materias[s[0]] {
				alumno[auxAlumno] = calificacion
			}


			alumno[s[1]] = f2
			
			this.Materias[s[0]] = alumno
			
		} else {
			*reply = "Ya tiene calificacion"
		}
	}

	v2, err2 := this.Alumnos[s[1]]
	
	*reply = "Se agrego correctamente"
	
	if err2 == false{
		clase := make(map[string]float64)
		
		f2, _ := strconv.ParseFloat(s[2], 8)
		clase[s[0]] = f2

		// creacion de una materia
		this.Alumnos[s[1]] = clase
	} else {
		_, err4 := v2[s[0]]
		if err4 == false{
			clase := make(map[string]float64)

			f2, _ := strconv.ParseFloat(s[2], 8)

			for auxClase, calificacion := range this.Alumnos[s[1]] {
				clase[auxClase] = calificacion
			}


			clase[s[0]] = f2
			
			this.Alumnos[s[1]] = clase
			
		} else {
			*reply = "Ya tiene calificacion"
		}
	}
	
	return nil
}

func (this *Server) PromedioAlumno(nombre string, reply *float64) error {
	var promedio float64
	var i int64
	promedio = 0
	i = 0

	for _, calificacion := range this.Alumnos[nombre] {
		promedio = promedio + calificacion
		i = i + 1
	}
	promedio = promedio / float64(i)
	*reply = promedio
	return nil
}

func (this *Server) PromedioMateria(nombre string, reply *float64) error {
	var promedio float64
	var i int64
	promedio = 0
	i = 0

	for _, calificacion := range this.Materias[nombre] {
		promedio = promedio + calificacion
		i = i + 1
	}
	promedio = promedio / float64(i)
	*reply = promedio
	return nil
}

func (this *Server) PromedioGeneral(f int64,reply *float64) error {
	var promedio float64
	var promedioGeneral float64
	var i int64
	var j int64
	i = 0
	j = 0
	promedio = 0
	promedioGeneral = 0

	for nombreAlumno := range this.Alumnos {
		i = 0
		promedio = 0
		for _, calificacion := range this.Alumnos[nombreAlumno] {
			promedio = promedio + float64(calificacion)
			i = i + 1
		}
		promedio = promedio / float64(i)
		promedioGeneral = promedioGeneral + promedio
		j = j + 1
	}
	promedioGeneral = promedioGeneral / float64(j)
	*reply = promedioGeneral
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
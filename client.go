package main

import (
	"fmt"
	"net/rpc"
)

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	var result string
	err = c.Call("Server.Constructor", "Se hizo bien", &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	for {
		fmt.Println("1) Agregar calificacion")
		fmt.Println("2) Mostrar promedio de un alumno")
		fmt.Println("3) Mostrar promedio general")
		fmt.Println("4) Mostrar promedio de una materia")
		fmt.Println("0) Exit")
		fmt.Scanln(&op)

		

		switch op {
		case 1:
			var materia string
			var nombre string
			var calificacion string
			var s []string
			var nada []string
			fmt.Print("Escriba el nombre de la materia: ")
			fmt.Scanln(&materia)
			fmt.Print("Escriba el nombre del alumno: ")
			fmt.Scanln(&nombre)
			fmt.Print("Escriba la calificacion de un alumno: ")
			fmt.Scanln(&calificacion)

			s = append(s,materia)
			s = append(s,nombre)
			s = append(s,calificacion)

			var result string
			err = c.Call("Server.AgregarCalificacion", s, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
			s = nada
		case 2:
			var nombre string
			fmt.Print("Escribe el nombre del alumno: ")
			fmt.Scanln(&nombre)

			var result float64
			err = c.Call("Server.PromedioAlumno", nombre, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio del alumno es ", result)
			}
		case 3:
			var aux int64
			var result float64
			err = c.Call("Server.PromedioGeneral", aux ,&result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio general es ", result)
			}
		case 4:
			var nombre string
			fmt.Print("Escribe el nombre de la materia: ")
			fmt.Scanln(&nombre)

			var result float64
			err = c.Call("Server.PromedioMateria", nombre, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio en la materia es ", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
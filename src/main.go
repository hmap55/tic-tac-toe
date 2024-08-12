package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Tablero struct {
	dimTablero [3][3]string
}

func initTablero(tablero *Tablero) Tablero {

	tablero.dimTablero = [3][3]string{{" - ", " - ", " - "},
		{" - ", " - ", " - "},
		{" - ", " - ", " - "}}
	return *tablero

}

func dibujarTablero(tablero Tablero) {
	fmt.Println("====================")
	for _, valor := range tablero.dimTablero {
		for i, interno := range valor {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(interno)
		}
		fmt.Println()
	}

	fmt.Println("======================")

}

func validarCoordenadas(coordenada string, tablero *Tablero, jugador Jugador) error {

	coordenadaF := strings.ReplaceAll(coordenada, " ", "")
	re := regexp.MustCompile(`^([1-3])([1-3])$`)
	fila, _ := strconv.Atoi(string(coordenadaF[0]))
	columna, _ := strconv.Atoi(string(coordenadaF[1]))

	if !re.MatchString(strings.TrimSpace(coordenadaF)) {
		return errors.New("coordenada por fuera del tablero")
	}

	if len(strings.TrimSpace(coordenadaF)) > 2 {
		fmt.Println("ingreso incorrecto")
		return errors.New("error en longitud de coordenada")
	}

	if tablero.dimTablero[fila-1][columna-1] != " - " {
		return errors.New("casilla ya ha sido marcada")
	}
	tablero.dimTablero[fila-1][columna-1] = jugador.marca

	return nil

}

func validarGanador(tablero Tablero) bool {

	var str string
	var str2 string
	var str3 string
	var diagonal1 string
	var diagonal2 string

	//validando filas
	for _, valor := range tablero.dimTablero {

		slice := valor[:]
		for i, s := range slice {
			slice[i] = strings.ReplaceAll(s, " ", "")
		}
		result := strings.Join(slice, "")
		if result == "XXX" || result == "OOO" {
			return true
		}
	}

	//fin validar filas

	//validando columnas
	for _, valor := range tablero.dimTablero {

		for i, interno := range valor {
			if i == 0 {
				str += interno
			} else if i == 1 {
				str2 += interno
			} else {
				str3 += interno
			}
		}

	}
	str = strings.ReplaceAll(str, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")
	str3 = strings.ReplaceAll(str3, " ", "")

	if str == "XXX" || str == "OOO" {
		return true
	}
	if str2 == "XXX" || str2 == "OOO" {
		return true
	}
	if str3 == "XXX" || str3 == "OOO" {
		return true
	}
	//fin validar columnas

	//validar diagonales
	for i, valor := range tablero.dimTablero {
		for j, v := range valor {
			if i == j {
				diagonal1 += v
			}
			if (i == 0 && j == 2) || (i == 1 && j == 1) || (i == 2 && j == 0) {
				diagonal2 += v
			}
		}

	}
	diagonal1 = strings.ReplaceAll(diagonal1, " ", "")
	diagonal2 = strings.ReplaceAll(diagonal2, " ", "")
	if diagonal1 == "XXX" || diagonal1 == "OOO" {
		return true
	}
	if diagonal2 == "XXX" || diagonal2 == "OOO" {
		return true
	}

	// fin validar diagonales

	return false

}

// structura jugadores
type Jugador struct {
	Nombre     string
	puntuacion int
	marca      string
}

func main() {

	//inicializar tablero
	var tablero Tablero
	tableroJuego := initTablero(&tablero)
	var jugador1 Jugador
	var jugador2 Jugador
	var jugadas int

	//mensaje de bienvenida
	fmt.Println("TIC-TAC-TOE")
	//entrada datos jugadores
	fmt.Println("ingrese nombre de primer jugador")
	fmt.Scanln(&jugador1.Nombre)
	jugador1.marca = " X "
	fmt.Println("ingrese nombre de segundo jugador")
	fmt.Scanln(&jugador2.Nombre)
	jugador2.marca = " O "

	//explicacion e ingreso de coordenadas
	fmt.Println("marque las coordenadas en el tablero segun el diagrama: ")
	fmt.Println(" 1  2  3")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	//dibujamos tablero inicial
	dibujarTablero(tableroJuego)
	//variable de entrada de datos
	bufer := bufio.NewReader(os.Stdin)
	//establecemos primer turno
	turno := jugador1

	//for del juego
	for {

		//validar maximo de jugadas posibles
		if jugadas == 9 {
			println("Empate, no hay mas jugadas posibles")
			return
		}

		//variable e ingreso de coordenadas
		var coordenada string
		fmt.Println("turno jugador: ", turno.Nombre)
		fmt.Println("ingrese coordenada x y separadas por espacio")
		coordenada, _ = bufer.ReadString('\n')
		//validar coordenadas ingresadas
		err := validarCoordenadas(coordenada, &tableroJuego, turno)
		if err != nil {
			fmt.Println("error en ingreso: ", err)
			continue
		}
		//se dibuja tablero con la nueva marca
		dibujarTablero(tableroJuego)
		//validamos si hay ganador
		ganador := validarGanador(tableroJuego)
		if ganador {
			println("Ganador jugador ", turno.Nombre)
			return
		}

		//cambiamos de turno
		if turno == jugador1 {
			turno = jugador2
		} else {
			turno = jugador1
		}
		jugadas++

	}

}

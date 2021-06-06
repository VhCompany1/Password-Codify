package main

import (
	"fmt"
	"os"
)

/*Main Vars*/
var opening = "646" //Secret code that will have to be entered in the fourth argument. Leave 0 to not use
var key = os.Args[4]
var module = os.Args[1]
var typecode = os.Args[2]
var pass = os.Args[3]
var cl = 000
var byteres = []byte{}
var ap int = 0
var res = []int{}

/*----------------------------------------------------*/

/*Function: Main*/
func main() {
	if opening != "0" {
		door(0)
	}

	if module == "code" {
		code()
	} else if module == "decode" {
		decode()
	} else {
		error(2)
	}
}

/*----------------------------------------------------*/

/*Function: Key verify*/
func door(c int) {
	if opening != key {
		error(4)
		return
	} else {

	}
}

/*----------------------------------------------------*/

/*Function: Encoding function*/
func code() {
	if typecode == "number" {
		d := []int{981, 561, 156, 654, 732, 822, 136, 515, 984, 166}

		for i := 0; i <= 9; i++ {
			if len(pass) < 10 || len(pass) <= 9 {
				if len(pass) == i {
					cl = d[i]
				}
			} else {
				cl = 109
			}
		}

		for i := 0; i < len(pass); i++ {
			byteres = append(byteres, pass[i])
		}

		for i := 0; i < len(pass); i++ {
			var numberbyte int = int(byteres[i])

			ap = numberbyte * 8
			ap = ap + 35
			ap = ap * 55

			res = append(res, ap)

			//numberbytelist := []int{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}

		}

		fmt.Printf("Sua senha codificada fica: %v-", cl)

		for i := 0; i < len(pass); i++ {
			fmt.Print(res[i])
		}

		if cl != 109 {
			fmt.Print("\n !AVISO:sua senha tem menos de 10 caracteres isso não é o recomendado")
		}
	} else if typecode == "string" {
		d := []int{981, 561, 156, 654, 732, 822, 136, 515, 984, 166}

		for i := 0; i <= 9; i++ {
			if len(pass) < 10 || len(pass) <= 9 {
				if len(pass) == i {
					cl = d[i]
				}
			} else {
				cl = 109
			}
		}

		for i := 0; i < len(pass); i++ {
			byteres = append(byteres, pass[i])
		}

		for i := 0; i < len(pass); i++ {
			var numberbyte int = int(byteres[i])

			ap = numberbyte * 8
			ap = ap + 35
			ap = ap * 55

			res = append(res, ap)

			//numberbytelist := []int{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}

		}

		fmt.Printf("Sua senha codificada fica: %v-", cl)

		for i := 0; i < len(pass); i++ {
			fmt.Print(res[i])
		}

		if cl != 109 {
			fmt.Print("\n !AVISO:sua senha tem menos de 10 caracteres isso não é o recomendado")
		}
	} else {
		error(1)
	}
}

/*----------------------------------------------------*/

/*Function: Decoding function*/
func decode() {
	if typecode == "number" {
		//decode := ""
		lenpass := len(pass) / 5

		if len(pass)%5 != 0 {
			error(3)

			return
		}

		ci := 0

		for i := 0; i < lenpass; i++ {

			for c := 0; c < 5; c++ {
				fmt.Print(string(pass[ci]))
				ci++

			}

			fmt.Print("-")

			//fmt.Printf("Sua senha decodificada fica: %v", decode)
		}
	} else if typecode == "string" {

	} else {
		error(1)
	}
}

/*----------------------------------------------------*/

/*Function: Errors*/
func error(e int) {
	/*Errors*/
	var erros = []string{
		"_",                                  //0
		"Type não encontrado",                //1
		"Modo de codificação não encontrado", //2
		"Sua senha não esta corretamente codificada",    //3
		"Você não é o proprietario deste decodificador", //4
	}
	/*----------------------------------------------------*/

	fmt.Printf("Erro[%v]:%v\n", e, erros[e])
}

/*----------------------------------------------------*/

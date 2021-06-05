package main

import (
	"fmt"
	"os"
)

func main() {
	module := os.Args[1]
	typecode := os.Args[2]
	pass := os.Args[3]
	cl := 000
	var byteres = []byte{}
	var ap int = 0
	var res = []int{}

	if module == "code" {
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
	} else if module == "decode" {
		if typecode == "number" {
			//decode := ""
			lenpass := len(pass) / 5

			if len(pass)%5 != 0 {
				fmt.Print("Sua senha não esta corretamente codificada")

				return
			}

			ci := 0
			//var caracteres = []int{}

			for i := 0; i < lenpass; i++ {

				for c := 0; c < 5; c++ {
					fmt.Print(string(pass[ci]))
					ci++

				}
				//caracteres = append(caracteres)

				fmt.Print("-")

				//fmt.Println(caracteres)
			}

			//fmt.Printf("Sua senha decodificada fica: %v", decode)
		} else if typecode == "string" {

		} else {
			error(1)
		}
	} else {
		error(2)
	}
}

func error(e int) {
	var erros = []string{
		"_",                                  //0
		"type não encontrado",                //1
		"modo de codificação não encontrado", //2
	}

	fmt.Printf("Erro[%v]:%v\n", e, erros[e])
}

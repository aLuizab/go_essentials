package main

import "fmt"

func main() {
	fmt.Println("Hello, Go Essentials!") //string
	fmt.Println(2)                       //inteiros
	fmt.Println(2 + 5)                   //expressoes
	fmt.Println(true)                    //booleanos
	fmt.Println("Hello, " + "World!")    //concatenar strings
	fmt.Println(2 * 5)                   //operações matemáticas
	fmt.Println(len("Hello, World!"))    //tamanho de uma string

	fmt.Printf("Hello, i'm %v and i'm an integrer\n", 2)
}

// pasta = folders podem ter multiplus arquivos
// podem ter multiplas funçoes dentro dela (espalhadas em multiplus arquivos)
// toda pasta é um pacote
// funçoes pertencem a um pacote

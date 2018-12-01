package main

import "fmt"

func printf() {
	/*
		Le mot-clé var permet de déclarer une variable. L'opérateur = permet
		de l'initialiser tout de suite, en inférant son type implicite.
		C'est possible d'utiliser une déclaration/assignation courte (short assignement statement)
		avec l'opérateur := qui donne word := "World" (pas besoin d'utiliser le mot-clé var).
	*/
	var word = "World"
	fmt.Printf("Hello %s\n", word)
}

func printfShortAssignement() {
	/*
		Attention à l'opérateur :=
		Il permet de déclarer une variable, donc il n'est pas possible
		de l'utiliser de nouveau sur la même variable.
		:= declaration + assignation
		= assignation sur une variable déjà déclarée (avec := ou var)
	*/
	word := "World"
	word = "Short"
	fmt.Printf("Hello %s\n", word)
}

func printImplicitType(value interface{}) {
	/*
		Il n'est pas obligatoire de préciser le type de la variable
		lors de l'assignation. Pour les types primitifs, le compilateur
		va créer le type correspondant à la syntaxe.
	*/
	fmt.Printf("Type: %T Value: %v\n", value, value)
}

func main() {
	printf()
	printfShortAssignement()
	printImplicitType("str")
	printImplicitType(1)
	printImplicitType(1.4)
	printImplicitType(-4)
	printImplicitType(false)
}

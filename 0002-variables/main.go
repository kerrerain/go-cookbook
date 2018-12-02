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
	fmt.Printf("Hello %s\n", word) // Hello World
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
	fmt.Printf("Hello %s\n", word) // Hello Short
}

/*
	Le type interface est un type indéfini. Il ne correspond
	à rien de précis.
*/
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

	/*
		bool Booléen

		string Chaîne de charactères

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr
		Entiers. Utiliser int par défaut, à moins d'avoir une raison spécifique.

		byte // alias for uint8

		rune // alias for int32
		// represents a Unicode code point

		float32 float64

		complex64 complex128

	*/
	printImplicitType('a')      // Type: int32 (alias rune) Value: 97 (U+0061) sur 1 byte (car encodé en UTF-8 ~ ASCII + Unicode)
	printImplicitType('ä')      // Type: int32 (alias rune) Value: 228 (U+00E4) sur 2 bytes (idem)
	printImplicitType('\u00E4') // Type: int32 (alias rune) Value: 228 (U+00E4) sur 2 bytes (idem)
	printImplicitType("str")    // Type: string Value: str
	printImplicitType(1)        // Type: int Value: 1
	printImplicitType(1.4)      // Type: float64 Value: 1.4
	printImplicitType(-4)       // Type: int Value: -4
	printImplicitType(false)    // Type: bool Value: false
}

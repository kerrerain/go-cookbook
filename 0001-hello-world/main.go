/*
Les commandes exécutables doivent déclarer un package "main".
C'est différent dans le cas d'une librairie, où on utilisera package <librairie>
*/
package main

import "fmt"

func main() {
	/*
		Le package fmt implémente des fonction d'E/S analogues à printf et scanf.
	*/
	fmt.Println("Hello world")
}

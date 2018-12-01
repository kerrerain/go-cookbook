/*
Executable commands must always use package main.
That's different for a library: package <library's name>
*/
package main

import "fmt"

func main() {
	/*
		Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
		The format 'verbs' are derived from C's but are simpler.
	*/
	fmt.Println("Hello world")
}

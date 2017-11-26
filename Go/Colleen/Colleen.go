package main

import (
	"fmt"
)

/*
** Com2
 */

func getChar() string {
	return "package main%[3]c%[3]cimport (%[3]c%[2]c%[1]cfmt%[1]c%[3]c)%[3]c%[3]c/*%[3]c** Com2%[3]c */%[3]c%[3]cfunc getChar() string {%[3]c%[2]creturn %[1]c%[4]s%[1]c%[3]c}%[3]c%[3]cfunc main() {%[3]c%[2]c/*%[3]c%[2]c** Com1%[3]c%[2]c */%[3]c%[2]cfmt.Printf(getChar(), 34, 9, 10, getChar())%[3]c}%[3]c"
}

func main() {
	/*
	** Com1
	 */
	fmt.Printf(getChar(), 34, 9, 10, getChar())
}

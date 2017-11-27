package main

import (
	"fmt"
	"os"
)

/*
** Com1
 */

func main() {
	f, err := os.Create("Grace_kid.go")
	if err != nil {
		panic("")
	}
	defer f.Close()
	str := "package main%[3]c%[3]cimport (%[3]c%[1]c%[2]cfmt%[2]c%[3]c%[1]c%[2]cos%[2]c%[3]c)%[3]c%[3]c/*%[3]c** Com1%[3]c */%[3]c%[3]cfunc main() {%[3]c%[1]cf, err := os.Create(%[2]cGrace_kid.go%[2]c)%[3]c%[1]cif err != nil {%[3]c%[1]c%[1]cpanic(%[2]c%[2]c)%[3]c%[1]c}%[3]c%[1]cdefer f.Close()%[3]c%[1]cstr := %[2]c%[4]s%[2]c%[3]c%[1]cfmt.Fprintf(f, str, 9, 34, 10, str)%[3]c}%[3]c"
	fmt.Fprintf(f, str, 9, 34, 10, str)
}

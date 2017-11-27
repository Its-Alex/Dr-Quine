package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	count := 5
	_, file, _, _ := runtime.Caller(0)
	if strings.Index(file, "Sully_0.go") != -1 {
		return
	}
	count--
	f, err := os.Create("Sully_" + strconv.Itoa(count) + ".go")
	if err != nil {
		panic("")
	}
	defer f.Close()
	str := "package main%[3]c%[3]cimport (%[3]c%[1]c%[2]cfmt%[2]c%[3]c%[1]c%[2]cos%[2]c%[3]c%[1]c%[2]cos/exec%[2]c%[3]c%[1]c%[2]cruntime%[2]c%[3]c%[1]c%[2]cstrconv%[2]c%[3]c%[1]c%[2]cstrings%[2]c%[3]c)%[3]c%[3]cfunc main() {%[3]c%[1]ccount := %[5]d%[3]c%[1]c_, file, _, _ := runtime.Caller(0)%[3]c%[1]cif strings.Index(file, %[2]cSully_0.go%[2]c) != -1 {%[3]c%[1]c%[1]creturn%[3]c%[1]c}%[3]c%[1]ccount--%[3]c%[1]cf, err := os.Create(%[2]cSully_%[2]c + strconv.Itoa(count) + %[2]c.go%[2]c)%[3]c%[1]cif err != nil {%[3]c%[1]c%[1]cpanic(%[2]c%[2]c)%[3]c%[1]c}%[3]c%[1]cdefer f.Close()%[3]c%[1]cstr := %[2]c%[4]s%[2]c%[3]c%[1]cfmt.Fprintf(f, str, 9, 34, 10, str, count)%[3]c%[1]ccmd := exec.Command(%[2]cgo%[2]c, %[2]crun%[2]c, %[2]cSully_%[2]c+strconv.Itoa(count)+%[2]c.go%[2]c)%[3]c%[1]cerr = cmd.Start()%[3]c%[1]cif err != nil {%[3]c%[1]c%[1]cpanic(%[2]c%[2]c)%[3]c%[1]c}%[3]c%[1]cerr = cmd.Wait()%[3]c%[1]cif err != nil {%[3]c%[1]c%[1]cpanic(%[2]c%[2]c)%[3]c%[1]c}%[3]c}%[3]c"
	fmt.Fprintf(f, str, 9, 34, 10, str, count)
	cmd := exec.Command("go", "run", "Sully_"+strconv.Itoa(count)+".go")
	err = cmd.Start()
	if err != nil {
		panic("")
	}
	err = cmd.Wait()
	if err != nil {
		panic("")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rea := bufio.NewReader(os.Stdin)
	fmt.Print("ingrese su sueldo:")
	res, _ := rea.ReadString('\n')

	res = strings.TrimSpace(res)
	res = strings.Replace(res, ",", "", -1)

	sueldo, e := strconv.Atoi(res)
	if e != nil {
		fmt.Println("error, ingrese un numero valido")
		os.Exit(1)
	}

	imp := sueldo * 15 / 100
	f := sueldo - imp

	fmt.Printf("sueldo basico:$%d\n ", sueldo)
	fmt.Printf("impuesto:$%d\n  ", imp)
	fmt.Printf("sueldo final:$%d\n  ", f)

}

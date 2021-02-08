package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU()) // Vai mostrar quantos processadores físicos existem nesta máquina
}

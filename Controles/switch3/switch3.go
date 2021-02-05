package main

import (
	"fmt"
	"time"
)

/*
interface pode retornar tipo genérico
Como ver o tipo que vem da interface passada no argumento?
*/
func tipo(i interface{}) string {
	switch i.(type) { //tipo da interface i
	case int:
		return "inteiro"
	case float32, float64:
		return "decimal"
	case string:
		return "string"
	case bool:
		return "boolean"
	case func():
		return "função"
	default:
		return "nao sei"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(2))
	fmt.Println(tipo("hey"))
	fmt.Println(tipo(true))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
	t := time.Now().Hour()
	fmt.Println(tipo(t))
	fmt.Println("HORAS:", t)
}

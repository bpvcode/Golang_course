package math

import (
	"fmt"
	"strconv"
)

//Media é responsável por calcular a média
func Media(num ...float64) float64 {
	total := 0.0
	for _, valor := range num {
		total += valor
	}
	media := total / float64(len(num))

	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64) //Para arredondar valor, mas transformá-lo em float64
	/*if err != nil {
		panic("PANICO - erro")
	}*/
	return mediaArredondada

}

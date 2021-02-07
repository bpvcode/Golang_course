package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"` // FORMA DE DIZER COMO APARECERÁ NO JSON
	Name  string   `json:"name"`
	Preco float64  `json:"preço"`
	Tag   []string `json:"tags"` //Slice do tipo string
}

/*
	NOTA MT IMPORTANTE:
	* Variável com letra MAIÚSCULA para poder ser exportado para json
	* Maiúscula - Significado public
	* Minúscula - Significado private
*/

func main() {

	// ====================== Struct to JSON ======================
	p1 := product{
		ID:    1,
		Name:  "Produto 1",
		Preco: 1.99,
		Tag:   []string{"Promoção", "Sem stock"}, //Tem de ter o tipo atrás ([]string)
	}

	p1Json, err := json.Marshal(p1) //Marshal retorna dois valores (slice de bytes []byte ou erro)
	if err != nil {
		panic("PANICO: Não veio o json de p1")
	}
	fmt.Println(string(p1Json)) //Cast de slice de bytes para string

	// ====================== JSON to Struct ======================
	var p2 product

	jsonString := `{"id":2,"name":"Produto 2","preço":5.45,"tags":["Eletrónica","Em stock"]}` //Texto em formato json

	json.Unmarshal([]byte(jsonString), p2) //ERRADO!
	fmt.Println(p2)

	json.Unmarshal([]byte(jsonString), &p2) //CORRECTO!
	fmt.Println(p2)
	/*
		Temos de passar um slice de bytes (json a transformar para struct) e a struct que receberá este json (p2)
		Passamos mesmo a referência para que altere mesmo os valores de p2
		RESUMO: Ele lê o json e seta os atributos e valores para dentro da struct p2
	*/

	fmt.Println(p2) //Como podemos ver alterou a struct e manteve porque acedemos à referência de p2
}

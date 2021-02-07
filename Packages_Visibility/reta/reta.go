package main

import "math"

/*
Iniciando um método/variável com letra maiúscula, por convenção, ele será 'public' (será visível dentro e fora do package)
Se for minúscula, será 'private' (só será visível dentro do PACKAGE - não do arquivo mas sim do package)

Por mais arquivos que tenhamos dentro de um package, quando for compilado, será criado uma estrutura, a estrutura do package, pelo que:
NÃO PODEMOS TER FUNÇÕES DUPLICADAS EM VÁRIOS ARQUIVOS DENTRO DO MESMO PACKAGE
A visibilidade privada não se refere ao arquivo mas sim ao package
*/

//Visivel FORA do package
type Ponto struct {
	x float64
	y float64
}

//Visivel DENTRO do package
func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return cx, cy
}

//Visivel FORA do package
func Distance(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

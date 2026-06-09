package main

import "math"

// Iniciando com letra maiúscula é público (visível fora do pacote)
// Iniciando com letra minúscula é privado (visível apenas dentro do pacote)

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

// catetos calcula a hipotenusa
func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia calcula a distância entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Hypot(cx, cy)
}

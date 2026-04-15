package main

import (
	"fmt"
	
)

func arvore(pen *Pen, dist float64) {
    // Comece definindo o seu Caso Base aqui
	if dist < 5 {
		pen.SetRGB(0, 5, 0)
		pen.FillCircle(3)
		return
	}
	ang := 27.5
	fator := 0.75

	pen.SetLineWidth(dist / 10)
	
	if(dist > 30) {
		pen.SetRGB(101, 67, 33)
	} else {
		pen.SetRGB(50, 0, 50)
	}

	pen.Walk(dist)
	pen.Right(ang)
	arvore(pen, dist*fator)
	pen.Left(ang)

	pen.Left(ang)
	arvore(pen, dist*fator)
	pen.Right(ang)

	pen.Walk(-dist)
	

    
    // Desenhe o tronco, vire a caneta e mergulhe na recursão
}
func circulo( pen *Pen) {
	pen.SetRGB(0, 120, 105)

	for i := 0; i < 5; i++ {
		pen.SetPosition(100 + float64(i)*90, 300)
		pen.DrawCircle(40)
	}

}

func main() {
	// Cria o canvas
	pen := NewPen(600, 600)   
	
	// Configurações iniciais
	pen.SetHeading(90)
	pen.SetPosition(300, 550)

	// Chama sua função
	arvore(pen, 150)

	// Salva o arquivo (sem tentar capturar retorno de erro)
	pen.SavePNG("tree.png")
	fmt.Println("Processo finalizado!")
}

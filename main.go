package main

import "github.com/josafa-dieb/knn/modules"

func main() {

	dados_treino := [][]float64{
		{10, 10, 1},
		{10, 30, 1},
		{15, 25, 1},
		{20, 15, 1},
		{30, 30, 2},
		{35, 25, 2},
		{40, 15, 2},
		{40, 40, 2}}
	novo_exemplo := []float64{30, 20}

	modules.KNearestNeighbors(novo_exemplo, dados_treino, 3)
}

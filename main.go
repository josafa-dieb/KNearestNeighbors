package main

import (
	"fmt"

	"github.com/josafa-dieb/knn/modules"
	"github.com/josafa-dieb/knn/utils"
)

func main() {

	data := utils.Load_dataset_file("./dataset/Iris.csv")
	dados_treino := utils.Load_data_to_training(data)
	novo_exemplo := []float64{4.6, 3.6, 1, 0.2}

	value := modules.KNearestNeighbors(novo_exemplo, dados_treino, 3)
	fmt.Println(value)
}

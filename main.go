package main

import (
	"fmt"

	"github.com/josafa-dieb/knn/modules"
	"github.com/josafa-dieb/knn/utils"
)

func main() {

	var metrics utils.Metrics
	training, test := utils.Load_dataset_file("./dataset/Iris.csv")
	data_to_training := utils.Load_data_to_training(training)
	for i := 0; i < len(test); i++ {
		var flower utils.Flower
		flower = test[i]
		data_to_testing := utils.Load_data_to_test(test[i])
		result := modules.KNearestNeighbors(data_to_testing, data_to_training, &metrics, 3)

		if flower.Specie == "Iris-setosa" && result == 1 {
			metrics.Correct_predictions += 1
		} else if flower.Specie == "Iris-versicolor" && result == 2 {
			metrics.Correct_predictions += 1
		} else if flower.Specie == "Iris-virginica" && result == 3 {
			metrics.Correct_predictions += 1
		}

		switch result {
		case 1:
			fmt.Println("Iris-setosa")
			break
		case 2:
			fmt.Println("Iris-versicolor")
			break
		case 3:
			fmt.Println("Iris-virginica")
			break
		}
	}
	fmt.Printf("Acruácia: %f", (float64(metrics.Correct_predictions) / float64(metrics.Predictions)))

}

package main

import (
	"fmt"

	"github.com/josafa-dieb/knn/modules"
	"github.com/josafa-dieb/knn/utils"
)

func main() {

	flowerTypeMap := map[string]int{
		"Iris-setosa":     1,
		"Iris-versicolor": 2,
		"Iris-virginica":  3,
	}

	var metrics utils.Metrics
	training, test := utils.Load_dataset_file("./dataset/Iris.csv")
	data_to_training := utils.Load_data_to_training(training)
	for i := 0; i < len(test); i++ {
		var flower utils.Flower
		flower = test[i]
		data_to_testing := utils.Load_data_to_test(test[i])
		result := modules.KNearestNeighbors(data_to_testing, data_to_training, &metrics, 3)

		if value, ok := flowerTypeMap[flower.Specie]; ok {
			if value == result {
				if flower.Specie == "Iris-setosa" && result == 1 || flower.Specie == "Iris-versicolor" && result == 2 || flower.Specie == "Iris-virginica" && result == 3 {
					metrics.True_positive += 1
				} else {
					metrics.True_negative += 1
				}
			} else {
				if flower.Specie == "Iris-setosa" && result == 1 || flower.Specie == "Iris-versicolor" && result == 2 || flower.Specie == "Iris-virginica" && result == 3 {
					metrics.False_negative += 1
				} else {
					metrics.False_positive += 1
				}
			}
		}

	}

	accuracy := float64((metrics.True_positive + metrics.True_negative) / len(test))

	precision := func() float64 {
		if (metrics.True_positive + metrics.False_positive) != 0 {
			return float64(metrics.True_positive / (metrics.True_positive + metrics.False_positive))
		} else {
			return 0
		}
	}()

	recall := func() float64 {
		if (metrics.True_positive + metrics.False_negative) != 0 {
			return float64(metrics.True_positive / (metrics.True_positive + metrics.False_negative))
		} else {
			return 0
		}
	}()

	fmt.Printf("Acurácia: %f", accuracy)
	fmt.Println()
	fmt.Printf("Precisão: %f", precision)
	fmt.Println()
	fmt.Printf("Recall do algorítmo: %f", recall)
	fmt.Println()
}

package main

import (
	"fmt"
	"os"

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
	file_trained, err := os.OpenFile("./result/treinamento.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Erro ao tentar carregar o arquivo.", err)
	}
	for _, row := range training {
		fmt.Fprintf(file_trained, "%d,%.2f,%.2f,%.2f,%.2f,%s\n", row.Id, row.Sepal_length, row.Sepal_width, row.Petal_length, row.Petal_width, row.Specie)
	}

	file_tested, err := os.OpenFile("./result/testes.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Erro ao tentar carregar o arquivo.", err)
	}
	for _, row := range test {
		fmt.Fprintf(file_tested, "%d,%.2f,%.2f,%.2f,%.2f,%s\n", row.Id, row.Sepal_length, row.Sepal_width, row.Petal_length, row.Petal_width, row.Specie)
	}

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

	accuracy := (metrics.True_positive + metrics.True_negative) / float64(len(test))

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

	fmt.Printf("Acurácia: %.2f", accuracy)
	fmt.Println()
	fmt.Printf("Precisão: %.2f", precision)
	fmt.Println()
	fmt.Printf("Recall do algorítmo: %.2f", recall)
	fmt.Println()
}

package utils

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Flower struct {
	Id           int
	Sepal_length float64
	Sepal_width  float64
	Petal_length float64
	Petal_width  float64
	Specie       string
}

func Load_dataset_file(dir string) ([]Flower, []Flower) {
	file, err := os.Open(dir)

	if err != nil {
		fmt.Println("Erro ao tentar carregar o arquivo.", err)
		return nil, nil
	}

	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Erro ao ler as linhas do arquivo.", err)
		return nil, nil
	}

	flowers_to_training := make([]Flower, 0)
	flowers_to_test := make([]Flower, 0)
	size := len(lines)
	size_to_traning := int(0.8 * float32(size))
	//size_to_test := size - size_to_traning

	count := 0
	for {
		random_index := rand.Intn(len(lines))
		row := lines[random_index]
		id, _ := strconv.Atoi(row[0])
		sepal_length, _ := strconv.ParseFloat(row[1], 64)
		sepal_width, _ := strconv.ParseFloat(row[2], 64)
		petal_length, _ := strconv.ParseFloat(row[3], 64)
		petal_width, _ := strconv.ParseFloat(row[4], 64)
		specie := row[5]

		flower := Flower{id, sepal_length, sepal_width, petal_length, petal_width, specie}

		flowers_to_training = append(flowers_to_training, flower)

		// remover o elemento
		lines = append(lines[:random_index], lines[random_index+1:]...)

		count++
		if count == size_to_traning {
			break
		}
	}
	for _, value := range lines {
		id, _ := strconv.Atoi(value[0])
		sepal_length, _ := strconv.ParseFloat(value[1], 64)
		sepal_width, _ := strconv.ParseFloat(value[2], 64)
		petal_length, _ := strconv.ParseFloat(value[3], 64)
		petal_width, _ := strconv.ParseFloat(value[4], 64)
		specie := value[5]
		flower := Flower{id, sepal_length, sepal_width, petal_length, petal_width, specie}
		flowers_to_test = append(flowers_to_test, flower)
	}

	return flowers_to_training, flowers_to_test
}

func Load_data_to_training(values []Flower) [][]float64 {

	float_values := make([][]float64, 0)

	for _, flower := range values {
		var flower_type int
		switch flower.Specie {
		case "Iris-setosa":
			flower_type = 1
			break
		case "Iris-versicolor":
			flower_type = 2
			break
		case "Iris-virginica":
			flower_type = 3
			break
		}
		values := []float64{flower.Sepal_length, flower.Sepal_width, flower.Petal_length, flower.Petal_width, float64(flower_type)}
		float_values = append(float_values, values)
	}
	return float_values

}
func Load_data_to_test(flower Flower) []float64 {

	float_values := make([]float64, 0)

	var flower_type int
	switch flower.Specie {
	case "Iris-setosa":
		flower_type = 1
		break
	case "Iris-versicolor":
		flower_type = 2
		break
	case "Iris-virginica":
		flower_type = 3
		break
	}
	float_values = append(float_values, flower.Sepal_length, flower.Sepal_width, flower.Petal_length, flower.Petal_width, float64(flower_type))
	return float_values

}

package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Flower struct {
	id           int
	sepal_length float64
	sepal_width  float64
	petal_length float64
	petal_width  float64
	specie       string
}

func Load_dataset_file(dir string) []Flower {
	file, err := os.Open(dir)

	if err != nil {
		fmt.Println("Erro ao tentar carregar o arquivo.", err)
		return nil
	}

	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Erro ao ler as linhas do arquivo.", err)
		return nil
	}

	flowers := make([]Flower, 0)

	for i := 1; i < len(lines); i++ {
		row := lines[i]
		id, _ := strconv.Atoi(row[0])
		sepal_length, _ := strconv.ParseFloat(row[1], 64)
		sepal_width, _ := strconv.ParseFloat(row[2], 64)
		petal_length, _ := strconv.ParseFloat(row[3], 64)
		petal_width, _ := strconv.ParseFloat(row[4], 64)
		specie := row[5]

		flower := Flower{id, sepal_length, sepal_width, petal_length, petal_width, specie}

		flowers = append(flowers, flower)

	}
	return flowers
}

func Load_data_to_training(values []Flower) [][]float64 {

	float_values := make([][]float64, 0)

	for _, flower := range values {
		var flower_type int
		switch flower.specie {
		case "Iris-setosa":
			flower_type = 1
		case "Iris-versicolor":
			flower_type = 2
		case "Iris-virginica":
			flower_type = 3
		}
		values := []float64{flower.sepal_length, flower.sepal_width, flower.petal_length, flower.petal_width, float64(flower_type)}
		float_values = append(float_values, values)
	}
	return float_values

}

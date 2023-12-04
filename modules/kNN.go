package modules

import (
	"sort"

	"github.com/josafa-dieb/knn/utils"
)

type KeyValue struct {
	Key   int
	Value float64
}

func KNearestNeighbors(novo_exemplo []float64, dados_treino [][]float64, metrics *utils.Metrics, k int32) int {

	// calcular as distâncias
	distancias := make(map[int]float64, 0)
	for i := 0; i < len(dados_treino); i++ {
		distancias[i] = Dist(novo_exemplo, dados_treino[i])
	}
	// ordenar as distancias
	distancia_ordenada := make([]KeyValue, 0, len(distancias))
	for k, v := range distancias {
		distancia_ordenada = append(distancia_ordenada, KeyValue{k, v})
	}
	sort.SliceStable(distancia_ordenada, func(i, j int) bool {
		return distancia_ordenada[i].Value < distancia_ordenada[j].Value
	})

	// fmt.Println(distancia_ordenada)
	vizinhos := make([]KeyValue, 0)
	for i := 0; i < int(k); i++ {
		vizinhos = append(vizinhos, distancia_ordenada[i])
	}
	classe := make(map[int]int, 0)

	for i := 0; i < len(vizinhos); i++ {
		indice := vizinhos[i].Key
		classe_vizinha := int(dados_treino[indice][len(dados_treino[indice])-1])
		classe[classe_vizinha]++
		metrics.Predictions += 1
	}

	var result int
	var max_count int

	for class_type, count := range classe {
		if count > max_count {
			result = class_type
			max_count = count
		}
	}

	return result
}

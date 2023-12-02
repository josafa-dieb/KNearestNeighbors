package modules

import (
	"fmt"
	"sort"
)

type KeyValue struct {
	Key   int
	Value float64
}

func KNearestNeighbors(novo_exemplo []float64, dados_treino [][]float64, k int32) {

	// calcular as distâncias
	distancias := make(map[int]float64, 0)
	for i := 0; i < len(dados_treino); i++ {
		distancias[i] = Dist(novo_exemplo, dados_treino[i])
	}
	// ordenar as distancias
	vizinhos := make([]KeyValue, 0, len(distancias))
	for k, v := range distancias {
		vizinhos = append(vizinhos, KeyValue{k, v})
	}
	sort.SliceStable(vizinhos, func(i, j int) bool {
		return vizinhos[i].Value < vizinhos[j].Value
	})

	classe := make(map[int]int, 0)
	classe[1] = 0
	classe[2] = 0
	fmt.Println(vizinhos)
}

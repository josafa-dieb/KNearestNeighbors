package modules

import (
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
	s := make([]KeyValue, 0, len(distancias))
	for k, v := range distancias {
		s = append(s, KeyValue{k, v})
	}
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].Value < s[j].Value
	})

}

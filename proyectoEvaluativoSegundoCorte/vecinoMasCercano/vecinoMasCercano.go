package vecinoMasCercano

import (
	"math"
	"sync"
)

// Message estructura para la comunicación de nodos
type Message struct {
	Node   int
	Weight int
}

// VecinoMasCercano encuentra el vecino más cercano no visitado
func VecinoMasCercano(nodoActual int, visitados map[int]bool, costos [][]int, ch chan<- Message, wg *sync.WaitGroup) {
	defer wg.Done()
	distanciaMinima := math.MaxInt
	vecinoCercano := -1

	for nodo := range costos {
		if nodo != nodoActual && !visitados[nodo] {
			dist := costos[nodoActual][nodo]
			if dist < distanciaMinima {
				distanciaMinima = dist
				vecinoCercano = nodo
			}
		}
	}
	ch <- Message{Node: vecinoCercano, Weight: distanciaMinima}
}

// MetodoDelCartero encuentra la ruta óptima usando el algoritmo del vecino más cercano
func MetodoDelCartero(nodos []string, costos [][]int) ([][2]int, int) {
	visitados := make(map[int]bool)
	ruta := [][2]int{}
	costoTotal := 0
	nodoActual := 0
	visitados[nodoActual] = true

	for len(visitados) < len(nodos) {
		ch := make(chan Message)
		var wg sync.WaitGroup
		wg.Add(1)
		go VecinoMasCercano(nodoActual, visitados, costos, ch, &wg)
		wg.Wait()
		close(ch)

		mensaje := <-ch
		vecino := mensaje.Node
		costo := mensaje.Weight

		ruta = append(ruta, [2]int{nodoActual, vecino})
		visitados[vecino] = true
		nodoActual = vecino
		costoTotal += costo
	}

	// Agrega el último enlace para volver al punto de partida
	ruta = append(ruta, [2]int{nodoActual, 0})
	costoTotal += costos[nodoActual][0]

	return ruta, costoTotal
}

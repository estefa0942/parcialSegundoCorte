package main

import (
	"fmt"

	"github.com/IsabelRamirezs/proyectoEvaluativoSegundoCorte/insercionMasCercana"
	"github.com/IsabelRamirezs/proyectoEvaluativoSegundoCorte/vecinoMasCercano"
)

func main() {
	nodos := []string{"Pereira", "Armenia", "Medellín", "Cartago"}
	costos := [][]int{
		{0, 45, 60, 25},
		{45, 0, 90, 50},
		{60, 90, 0, 70},
		{25, 50, 70, 0},
	}

	rutaOptimizada, costoTotal := vecinoMasCercano.MetodoDelCartero(nodos, costos)
	fmt.Println("Ruta optimizada usando Vecino Más Cercano:")
	for _, enlace := range rutaOptimizada {
		punto1, punto2 := enlace[0], enlace[1]
		fmt.Printf("De %s a %s\n", nodos[punto1], nodos[punto2])
	}
	fmt.Printf("Costo total: %d\n", costoTotal)

	coordenadasPuntos := [][]int{
		{0, 1, 1},
		{1, 2, 2},
		{2, 3, 3},
		{3, 4, 4},
	}
	matrizAdyacencia := insercionMasCercana.ConstruirMatrizAdyacencia(coordenadasPuntos)
	tour, fo := insercionMasCercana.InsercionMasCercana(coordenadasPuntos, matrizAdyacencia)
	insercionMasCercana.ImprimirTour(coordenadasPuntos, tour, fo)
}

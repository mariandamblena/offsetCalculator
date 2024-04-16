package main

/*
import (
	"fmt"
	"log"

	stats "gonum.org/v1/gonum/stat"
)

func main() {
	// Paso 1: Recopilación de datos (simulado para propósitos de demostración)
	// Supongamos que tienes dos conjuntos de datos llamados sensor1 y sensor2
	// Cada conjunto de datos es una slice de float64 que contiene las mediciones de temperatura

	// Los datos de ejemplo
	sensor1 := []float64{23.4, 24.5, 25.6, 26.7, 27.8}
	sensor2 := []float64{22.9, 24.2, 25.5, 26.8, 28.1}

	// Paso 3: Cálculo del offset medio
	var offsetMean float64
	if !stats.EqualLength(sensor1, sensor2) {
		log.Fatal("Los conjuntos de datos no tienen la misma longitud")
	}
	for i := range sensor1 {
		offsetMean += sensor2[i] - sensor1[i]
	}
	offsetMean /= float64(len(sensor1))

	// Paso 4: Visualización de datos (No hay visualización en Go en este ejemplo)

	// Paso 5: Análisis estadístico
	var (
		x = make([]float64, len(sensor1))
		y = make([]float64, len(sensor2))
	)
	copy(x, sensor1)
	copy(y, sensor2)

	// Ajuste de la regresión lineal
	var (
		reg  stats.Regression
		beta = []float64{0}
	)
	reg.Init(nil, 1, nil)
	for i, xv := range x {
		reg.Train(stats.regression.DataPoint(y[i], []float64{xv}))
	}
	ok := reg.Solve()
	if !ok {
		log.Fatal("No se pudo ajustar el modelo de regresión")
	}
	copy(beta, reg.Parameters)

	fmt.Println("Coeficiente de la regresión (slope):", beta[0])
}

*/

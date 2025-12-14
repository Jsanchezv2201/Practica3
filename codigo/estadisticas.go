package main

import (
	"fmt"
	"strings"
	"time"
)

type Estadisticas struct {
	TiempoTotal       time.Duration
	CochesPorTipo     map[string]int
	TiemposPorTipo    map[string]time.Duration
	TiempoInicio      time.Time
	CochesCompletados int
	Metodo            string // "RWMutex" o "WaitGroup"
}

func NuevasEstadisticas() *Estadisticas {
	return &Estadisticas{
		CochesPorTipo:  make(map[string]int),
		TiemposPorTipo: make(map[string]time.Duration),
		TiempoInicio:   time.Now(),
	}
}

func (e *Estadisticas) RegistrarCoche(tipo string, tiempo time.Duration) {
	e.CochesPorTipo[tipo]++
	e.TiemposPorTipo[tipo] += tiempo
	e.CochesCompletados++

	if tiempo > e.TiempoTotal {
		e.TiempoTotal = tiempo
	}
}

func (e *Estadisticas) ImprimirResumen() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("📊 ESTADÍSTICAS %s (con prioridad real)\n", e.Metodo)
	fmt.Println(strings.Repeat("=", 60))

	totalCoches := 0
	for _, count := range e.CochesPorTipo {
		totalCoches += count
	}

	tiempoTotalSim := time.Since(e.TiempoInicio)

	fmt.Printf("🚗 Total coches procesados: %d\n", totalCoches)
	fmt.Printf("⏱️  Tiempo total simulación: %v\n", tiempoTotalSim.Round(100*time.Millisecond))

	// Análisis de rendimiento
	if totalCoches > 0 {
		tiempoPromedioTotal := tiempoTotalSim / time.Duration(totalCoches)
		eficiencia := float64(totalCoches) / tiempoTotalSim.Minutes()

		fmt.Println("\n🎯 MÉTRICAS DE RENDIMIENTO:")
		fmt.Println(strings.Repeat("-", 60))
		fmt.Printf("📊 Tiempo promedio por coche: %v\n", tiempoPromedioTotal.Round(100*time.Millisecond))
		fmt.Printf("⚡ Eficiencia: %.2f coches/minuto\n", eficiencia)
	}
}
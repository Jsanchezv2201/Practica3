package main

import (
	"fmt"
	"time"
)

type Estadisticas struct {
	TiempoTotal    time.Duration
	CochesPorTipo  map[string]int
	TiemposPorTipo map[string]time.Duration
	TiempoInicio   time.Time
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
	
	// Actualizar tiempo total (el máximo entre todos los coches)
	if tiempo > e.TiempoTotal {
		e.TiempoTotal = tiempo
	}
}

func (e *Estadisticas) ImprimirResumen() {
	fmt.Println("\n📊 ESTADÍSTICAS FINALES:")
	fmt.Println("=======================")
	
	totalCoches := 0
	for _, count := range e.CochesPorTipo {
		totalCoches += count
	}
	
	fmt.Printf("🚗 Coches totales: %d\n", totalCoches)
	fmt.Printf("⏱️  Tiempo total simulación: %v\n", time.Since(e.TiempoInicio).Round(100*time.Millisecond))
	
	fmt.Println("\n📈 POR TIPO DE INCIDENCIA:")
	fmt.Println("-----------------------")
	
	for tipo, count := range e.CochesPorTipo {
		if count > 0 {
			tiempoPromedio := e.TiemposPorTipo[tipo] / time.Duration(count)
			emoji := "🔧"
			if tipo == "eléctrica" {
				emoji = "⚡"
			} else if tipo == "carrocería" {
				emoji = "🚙"
			}
			
			fmt.Printf("%s %s: %d coches | Tiempo promedio: %v | Tiempo total: %v\n", 
				emoji, tipo, count, 
				tiempoPromedio.Round(100*time.Millisecond),
				e.TiemposPorTipo[tipo].Round(100*time.Millisecond))
		}
	}
	
	// Calcular eficiencia (coches por minuto)
	tiempoTotal := time.Since(e.TiempoInicio)
	if tiempoTotal > 0 {
		eficiencia := float64(totalCoches) / tiempoTotal.Minutes()
		fmt.Printf("\n⚡ Eficiencia: %.2f coches/minuto\n", eficiencia)
	}
	
	// Análisis de rendimiento
	fmt.Println("\n🎯 ANÁLISIS DE RENDIMIENTO:")
	fmt.Println("-------------------------")
	
	if totalCoches > 0 {
		tiempoPromedioTotal := time.Since(e.TiempoInicio) / time.Duration(totalCoches)
		fmt.Printf("⏳ Tiempo promedio por coche: %v\n", tiempoPromedioTotal.Round(100*time.Millisecond))
		
		// Identificar el tipo más rápido y más lento
		var tipoMasRapido, tipoMasLento string
		var tiempoMasRapido, tiempoMasLento time.Duration
		
		primero := true
		for tipo := range e.CochesPorTipo {
			tiempoPromedio := e.TiemposPorTipo[tipo] / time.Duration(e.CochesPorTipo[tipo])
			if primero || tiempoPromedio < tiempoMasRapido {
				tiempoMasRapido = tiempoPromedio
				tipoMasRapido = tipo
			}
			if primero || tiempoPromedio > tiempoMasLento {
				tiempoMasLento = tiempoPromedio
				tipoMasLento = tipo
			}
			primero = false
		}
		
		fmt.Printf("🐇 Más rápido: %s (%v)\n", tipoMasRapido, tiempoMasRapido.Round(100*time.Millisecond))
		fmt.Printf("🐢 Más lento: %s (%v)\n", tipoMasLento, tiempoMasLento.Round(100*time.Millisecond))
	}
}
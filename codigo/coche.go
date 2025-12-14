package main

import (
	"math/rand"
	"time"
	"fmt"
)

// Coche representa un vehículo en el taller con prioridad
type Coche struct {
	ID             int
	TipoIncidencia string
	Prioridad      string
	TiempoPorFase  time.Duration
	TiempoInicio   time.Time
}

// PrioridadInt devuelve un valor numérico para comparar prioridades
func (c *Coche) PrioridadInt() int {
	switch c.Prioridad {
	case "alta":
		return 0
	case "media":
		return 1
	case "baja":
		return 2
	default:
		return 3
	}
}

// generarCoches crea los coches según el enunciado P3
func generarCoches(a, b, c int) []*Coche {
	var coches []*Coche
	id := 1

	// Categoría A: Mecánica - Prioridad Alta
	for i := 0; i < a; i++ {
		// Variación ±20% (0.8 a 1.2)
		variacion := 0.8 + rand.Float64()*0.4
		tiempoBase := 5 * time.Second
		
		coches = append(coches, &Coche{
			ID:             id,
			TipoIncidencia: "mecánica",
			Prioridad:      "alta",
			TiempoPorFase:  time.Duration(float64(tiempoBase) * variacion),
		})
		id++
	}

	// Categoría B: Eléctrica - Prioridad Media
	for i := 0; i < b; i++ {
		variacion := 0.8 + rand.Float64()*0.4
		tiempoBase := 3 * time.Second
		
		coches = append(coches, &Coche{
			ID:             id,
			TipoIncidencia: "eléctrica",
			Prioridad:      "media",
			TiempoPorFase:  time.Duration(float64(tiempoBase) * variacion),
		})
		id++
	}

	// Categoría C: Carrocería - Prioridad Baja
	for i := 0; i < c; i++ {
		variacion := 0.8 + rand.Float64()*0.4
		tiempoBase := 1 * time.Second
		
		coches = append(coches, &Coche{
			ID:             id,
			TipoIncidencia: "carrocería",
			Prioridad:      "baja",
			TiempoPorFase:  time.Duration(float64(tiempoBase) * variacion),
		})
		id++
	}

	// Mezclar aleatoriamente (requisito del enunciado)
	rand.Shuffle(len(coches), func(i, j int) {
		coches[i], coches[j] = coches[j], coches[i]
	})

	return coches
}

func (c *Coche) String() string {
	emoji := "🟡"
	switch c.Prioridad {
	case "alta":
		emoji = "🔴"
	case "baja":
		emoji = "🟢"
	}
	
	tipoEmoji := "🔧"
	switch c.TipoIncidencia {
	case "eléctrica":
		tipoEmoji = "⚡"
	case "carrocería":
		tipoEmoji = "🚙"
	}
	
	return fmt.Sprintf("%s%s Coche %2d", emoji, tipoEmoji, c.ID)
}
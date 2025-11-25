package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type TallerRWMutex struct {
	plazas       chan struct{}
	mecanicos    chan struct{}
	limpieza     chan struct{}
	revision     chan struct{}
	mu           sync.RWMutex
	wg           sync.WaitGroup
	tiempoInicio time.Time
	estadisticas *Estadisticas
}

type Coche struct {
	ID            int
	TipoIncidencia string
	Prioridad     string
	TiempoPorFase time.Duration
	TiempoInicio  time.Time
}

func NuevoTallerRWMutex(plazas, mecanicos, limpieza, revision int) *TallerRWMutex {
	return &TallerRWMutex{
		plazas:       make(chan struct{}, plazas),
		mecanicos:    make(chan struct{}, mecanicos),
		limpieza:     make(chan struct{}, limpieza),
		revision:     make(chan struct{}, revision),
		tiempoInicio: time.Now(),
		estadisticas: NuevasEstadisticas(),
	}
}

func SimularTallerRWMutex(taller *TallerRWMutex, cochesA, cochesB, cochesC int) {
	totalCoches := cochesA + cochesB + cochesC
	taller.wg.Add(totalCoches)
	
	coches := generarCoches(cochesA, cochesB, cochesC)
	
	fmt.Printf("🏁 INICIANDO SIMULACIÓN con %d coches\n", totalCoches)
	fmt.Printf("   🔴 Mecánica (2s/fase) | 🟡 Eléctrica (1.5s/fase) | 🟢 Carrocería (1s/fase)\n\n")
	
	for _, coche := range coches {
		coche.TiempoInicio = time.Now()
		go taller.procesarCoche(coche)
	}
	
	taller.wg.Wait()
	
	// Mostrar estadísticas al final
	fmt.Printf("\n🎉 TODOS LOS %d COCHES TERMINADOS\n", totalCoches)
	taller.estadisticas.ImprimirResumen()
}

func (t *TallerRWMutex) procesarCoche(coche *Coche) {
	defer t.wg.Done()

	// FASE 1: 🅿️ Plazas
	t.logFase(coche, "🚗 ENTRANDO", "al taller")
	t.plazas <- struct{}{}
	time.Sleep(coche.TiempoPorFase)
	<-t.plazas

	// FASE 2: 🔧 Mecánicos
	t.logFase(coche, "🔧 REPARANDO", "la incidencia")
	t.mecanicos <- struct{}{}
	time.Sleep(coche.TiempoPorFase)
	<-t.mecanicos

	// FASE 3: 🧽 Limpieza
	t.logFase(coche, "🧽 LIMPIANDO", "el coche")
	t.limpieza <- struct{}{}
	time.Sleep(coche.TiempoPorFase)
	<-t.limpieza

	// FASE 4: ✅ Revisión
	t.logFase(coche, "✅ REVISANDO", "final")
	t.revision <- struct{}{}
	time.Sleep(coche.TiempoPorFase)
	<-t.revision
	
	// Registrar estadísticas al finalizar
	duracion := time.Since(coche.TiempoInicio)
	t.estadisticas.RegistrarCoche(coche.TipoIncidencia, duracion)
	t.logCompletado(coche, duracion)
}

func (t *TallerRWMutex) logFase(coche *Coche, accion, descripcion string) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	emojiPrioridad := "🔴"
	if coche.Prioridad == "media" {
		emojiPrioridad = "🟡"
	} else if coche.Prioridad == "baja" {
		emojiPrioridad = "🟢"
	}
	
	fmt.Printf("[%4v] %s Coche %2d %s %s\n",
		time.Since(t.tiempoInicio).Round(100*time.Millisecond),
		emojiPrioridad,
		coche.ID,
		accion,
		descripcion)
}

func (t *TallerRWMutex) logCompletado(coche *Coche, duracion time.Duration) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	emojiTipo := "🔧"
	if coche.TipoIncidencia == "eléctrica" {
		emojiTipo = "⚡"
	} else if coche.TipoIncidencia == "carrocería" {
		emojiTipo = "🚙"
	}
	
	fmt.Printf("[%4v] 🎉 Coche %2d TERMINADO %s %s (tiempo: %v)\n",
		time.Since(t.tiempoInicio).Round(100*time.Millisecond),
		coche.ID,
		emojiTipo,
		coche.TipoIncidencia,
		duracion.Round(100*time.Millisecond))
}

func generarCoches(a, b, c int) []*Coche {
	var coches []*Coche
	id := 1
	
	// Categoría A: Mecánica - Prioridad Alta (2s)
	for i := 0; i < a; i++ {
		coches = append(coches, &Coche{
			ID:            id,
			TipoIncidencia: "mecánica",
			Prioridad:     "alta",
			TiempoPorFase: 2 * time.Second,
		})
		id++
	}
	
	// Categoría B: Eléctrica - Prioridad Media (1.5s)
	for i := 0; i < b; i++ {
		coches = append(coches, &Coche{
			ID:            id,
			TipoIncidencia: "eléctrica", 
			Prioridad:     "media",
			TiempoPorFase: 1500 * time.Millisecond,
		})
		id++
	}
	
	// Categoría C: Carrocería - Prioridad Baja (1s)
	for i := 0; i < c; i++ {
		coches = append(coches, &Coche{
			ID:            id,
			TipoIncidencia: "carrocería",
			Prioridad:     "baja", 
			TiempoPorFase: 1 * time.Second,
		})
		id++
	}
	
	// Mezclar aleatoriamente
	rand.Shuffle(len(coches), func(i, j int) {
		coches[i], coches[j] = coches[j], coches[i]
	})
	
	return coches
}
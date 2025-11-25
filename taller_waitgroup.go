package main

import (
	"fmt"
	"sync"
	"time"
)

type TallerWaitGroup struct {
	plazas       chan struct{}
	mecanicos    chan struct{}
	limpieza     chan struct{}
	revision     chan struct{}
	wg           sync.WaitGroup
	tiempoInicio time.Time
	estadisticas *Estadisticas
}

func NuevoTallerWaitGroup(plazas, mecanicos, limpieza, revision int) *TallerWaitGroup {
	return &TallerWaitGroup{
		plazas:       make(chan struct{}, plazas),
		mecanicos:    make(chan struct{}, mecanicos),
		limpieza:     make(chan struct{}, limpieza),
		revision:     make(chan struct{}, revision),
		tiempoInicio: time.Now(),
		estadisticas: NuevasEstadisticas(),
	}
}

func SimularTallerWaitGroup(cochesA, cochesB, cochesC int) {
	taller := NuevoTallerWaitGroup(2, 2, 1, 1)
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

func (t *TallerWaitGroup) procesarCoche(coche *Coche) {
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

func (t *TallerWaitGroup) logFase(coche *Coche, accion, descripcion string) {
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

func (t *TallerWaitGroup) logCompletado(coche *Coche, duracion time.Duration) {
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
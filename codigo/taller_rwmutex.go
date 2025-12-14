package main

import (
	"fmt"
	"sync"
	"time"
)

type TallerRWMutex struct {
	plazas       *GestorRecurso
	mecanicos    *GestorRecurso
	limpieza     *GestorRecurso
	revision     *GestorRecurso
	mu           sync.RWMutex
	wg           sync.WaitGroup
	tiempoInicio time.Time
	estadisticas *Estadisticas
}

func NuevoTallerRWMutex(plazas, mecanicos, limpieza, revision int) *TallerRWMutex {
	estadisticas := NuevasEstadisticas()
	estadisticas.Metodo = "RWMutex"
	
	return &TallerRWMutex{
		plazas:       NuevoGestorRecurso(plazas),
		mecanicos:    NuevoGestorRecurso(mecanicos),
		limpieza:     NuevoGestorRecurso(limpieza),
		revision:     NuevoGestorRecurso(revision),
		tiempoInicio: time.Now(),
		estadisticas: estadisticas,
	}
}

func SimularTallerRWMutex(taller *TallerRWMutex, cochesA, cochesB, cochesC int) {
	totalCoches := cochesA + cochesB + cochesC
	taller.wg.Add(totalCoches)

	coches := generarCoches(cochesA, cochesB, cochesC)

	fmt.Printf("\n🏁 INICIANDO SIMULACIÓN RWMUTEX con %d coches\n", totalCoches)
	fmt.Printf("   🔴 Mecánica: %d coches (5s/fase con variación ±20%%)\n", cochesA)
	fmt.Printf("   🟡 Eléctrica: %d coches (3s/fase con variación ±20%%)\n", cochesB)
	fmt.Printf("   🟢 Carrocería: %d coches (1s/fase con variación ±20%%)\n", cochesC)
	fmt.Println("   🎯 ALTA(🔴) > MEDIA(🟡) > BAJA(🟢) - PRIORIDAD REAL")

	for _, coche := range coches {
		coche.TiempoInicio = time.Now()
		go taller.procesarCoche(coche)
	}

	taller.wg.Wait()

	fmt.Printf("\n🎉 TODOS LOS %d COCHES TERMINADOS (RWMutex con Prioridad)\n", totalCoches)
	taller.estadisticas.ImprimirResumen()
}

func (t *TallerRWMutex) procesarCoche(coche *Coche) {
	defer t.wg.Done()

	// FASE 1: Entrada al taller
	t.logEvento(coche, 1, "Esperando plaza")
	t.plazas.Solicitar(coche)
	t.logEvento(coche, 1, "Ocupando plaza")
	time.Sleep(coche.TiempoPorFase)
	t.plazas.Liberar()

	// FASE 2: Reparación
	t.logEvento(coche, 2, "Esperando mecánico")
	t.mecanicos.Solicitar(coche)
	t.logEvento(coche, 2, "Siendo reparado")
	time.Sleep(coche.TiempoPorFase)
	t.mecanicos.Liberar()

	// FASE 3: Limpieza
	t.logEvento(coche, 3, "Esperando limpieza")
	t.limpieza.Solicitar(coche)
	t.logEvento(coche, 3, "Siendo limpiado")
	time.Sleep(coche.TiempoPorFase)
	t.limpieza.Liberar()

	// FASE 4: Revisión final
	t.logEvento(coche, 4, "Esperando revisión")
	t.revision.Solicitar(coche)
	t.logEvento(coche, 4, "Siendo revisado")
	time.Sleep(coche.TiempoPorFase)
	t.revision.Liberar()
	
	// Registrar estadísticas
	duracion := time.Since(coche.TiempoInicio)
	t.estadisticas.RegistrarCoche(coche.TipoIncidencia, duracion)
	t.logCompletado(coche, duracion)
}

func (t *TallerRWMutex) logEvento(coche *Coche, fase int, estado string) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	// Emoji de prioridad
	emojiPrioridad := "🟡"
	switch coche.Prioridad {
	case "alta":
		emojiPrioridad = "🔴"
	case "baja":
		emojiPrioridad = "🟢"
	}
	
	// Emoji de estado
	emojiEstado := "⚙️  "
	switch estado {
	case "Esperando plaza", "Esperando mecánico", "Esperando limpieza", "Esperando revisión":
		emojiEstado = "⏳ "
	case "Ocupando plaza":
		emojiEstado = "🚗 "
	}
	
	fmt.Printf("Tiempo:%3vs | %s Coche %2d (%s) | Fase %d | %s%s\n",
		time.Since(t.tiempoInicio).Round(time.Second).Seconds(),
		emojiPrioridad,
		coche.ID,
		coche.TipoIncidencia,
		fase,
		emojiEstado,
		estado)
}

func (t *TallerRWMutex) logCompletado(coche *Coche, duracion time.Duration) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	fmt.Printf("-----------------------------------------------------------\n")
	fmt.Printf("Tiempo:%3vs | 🎉 Coche  %2d (%s) COMPLETADO en %v\n",
		time.Since(t.tiempoInicio).Round(time.Second).Seconds(),
		coche.ID,
		coche.TipoIncidencia,
		duracion.Round(time.Second))
	fmt.Printf("-----------------------------------------------------------\n")
}
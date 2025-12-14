package main

import (
	"container/heap"
	"sync"
	"time"
)

// ItemCola representa un elemento en la cola de prioridad
type ItemCola struct {
	Coche      *Coche
	Prioridad  int       // 0=alta, 1=media, 2=baja
	Timestamp  time.Time // Para desempatar misma prioridad
	Index      int       // Índice en el heap
}

// ColaPrioridad implementa heap.Interface
type ColaPrioridad []*ItemCola

func (cp ColaPrioridad) Len() int { return len(cp) }

func (cp ColaPrioridad) Less(i, j int) bool {
	// Primero por prioridad (menor número = mayor prioridad)
	if cp[i].Prioridad != cp[j].Prioridad {
		return cp[i].Prioridad < cp[j].Prioridad
	}
	// Si misma prioridad, por orden de llegada (FIFO)
	return cp[i].Timestamp.Before(cp[j].Timestamp)
}

func (cp ColaPrioridad) Swap(i, j int) {
	cp[i], cp[j] = cp[j], cp[i]
	cp[i].Index = i
	cp[j].Index = j
}

func (cp *ColaPrioridad) Push(x interface{}) {
	n := len(*cp)
	item := x.(*ItemCola)
	item.Index = n
	*cp = append(*cp, item)
}

func (cp *ColaPrioridad) Pop() interface{} {
	old := *cp
	n := len(old)
	item := old[n-1]
	item.Index = -1 // marcado como eliminado
	*cp = old[0 : n-1]
	return item
}

// GestorRecurso maneja un recurso con prioridad
type GestorRecurso struct {
	Capacidad int
	Ocupados  int
	Cola      *ColaPrioridad
	Mu        sync.Mutex
	Cond      *sync.Cond
}

func NuevoGestorRecurso(capacidad int) *GestorRecurso {
	g := &GestorRecurso{
		Capacidad: capacidad,
		Ocupados:  0,
		Cola:      &ColaPrioridad{},
	}
	g.Cond = sync.NewCond(&g.Mu)
	heap.Init(g.Cola)
	return g
}

// Solicitar intenta obtener el recurso con prioridad
func (g *GestorRecurso) Solicitar(coche *Coche) {
	g.Mu.Lock()

	// Crear ítem para la cola
	item := &ItemCola{
		Coche:     coche,
		Prioridad: coche.PrioridadInt(),
		Timestamp: time.Now(),
	}
	
	// Añadir a la cola
	heap.Push(g.Cola, item)
	
	// Esperar hasta que:
	// 1. Haya capacidad disponible
	// 2. Seamos el primero en la cola (mayor prioridad)
	for {
		// Verificar si podemos acceder al recurso
		if g.Ocupados < g.Capacidad && g.Cola.Len() > 0 {
			// Verificar si somos el primero en la cola
			primero := (*g.Cola)[0]
			if primero.Coche.ID == coche.ID {
				// ¡Somos el primero! Sacarnos de la cola y tomar el recurso
				heap.Pop(g.Cola)
				g.Ocupados++
				g.Mu.Unlock()
				return
			}
		}
		// Si no podemos acceder ahora, esperar
		g.Cond.Wait()
	}
}

// Liberar libera el recurso
func (g *GestorRecurso) Liberar() {
	g.Mu.Lock()
	defer g.Mu.Unlock()
	
	g.Ocupados--
	// Despertar a todos los que esperan
	g.Cond.Broadcast()
}
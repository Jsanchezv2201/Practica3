package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("🚗 === TALLER MECÁNICO - PRÁCTICA 3 ===")
	fmt.Println("1. Ejecutar Tests con RWMutex")
	fmt.Println("2. Ejecutar Tests con WaitGroup")
	fmt.Println("3. Comparativa de Métodos")
	fmt.Print("👉 Seleccione opción: ")
	
	var opcion int
	fmt.Scan(&opcion)
	
	switch opcion {
	case 1:
		ejecutarTestsRWMutex()
	case 2:
		ejecutarTestsWaitGroup()
	case 3:
		compararMetodos()
	default:
		fmt.Println("❌ Opción no válida")
	}
}

func ejecutarTestsRWMutex() {
	fmt.Println("\n🔧 === EJECUTANDO TESTS CON RWMUTEX ===")
	
	tests := []struct{
		nombre string
		cochesA, cochesB, cochesC int
	}{
		{"Test 1 - Balanceado", 2, 2, 2},
		{"Test 2 - Mayoría A", 3, 1, 1},
		{"Test 3 - Mayoría C", 1, 1, 3},
	}
	
	for _, test := range tests {
		fmt.Printf("\n🎯 %s\n", test.nombre)
		fmt.Printf("   🔴 Mecánica (Alta): %d coches\n", test.cochesA)
		fmt.Printf("   🟡 Eléctrica (Media): %d coches\n", test.cochesB) 
		fmt.Printf("   🟢 Carrocería (Baja): %d coches\n", test.cochesC)
		
		taller := NuevoTallerRWMutex(2, 2, 1, 1)
		SimularTallerRWMutex(taller, test.cochesA, test.cochesB, test.cochesC)
		fmt.Printf("✅ %s COMPLETADO\n\n", test.nombre)
	}
}

func ejecutarTestsWaitGroup() {
	fmt.Println("\n🔧 === EJECUTANDO TESTS CON WAITGROUP ===")
	
	tests := []struct{
		nombre string
		cochesA, cochesB, cochesC int
	}{
		{"Test 1 - Balanceado", 2, 2, 2},
		{"Test 2 - Mayoría A", 3, 1, 1},
		{"Test 3 - Mayoría C", 1, 1, 3},
	}
	
	for _, test := range tests {
		fmt.Printf("\n🎯 %s\n", test.nombre)
		fmt.Printf("   🔴 Mecánica (Alta): %d coches\n", test.cochesA)
		fmt.Printf("   🟡 Eléctrica (Media): %d coches\n", test.cochesB)
		fmt.Printf("   🟢 Carrocería (Baja): %d coches\n", test.cochesC)
		
		SimularTallerWaitGroup(test.cochesA, test.cochesB, test.cochesC)
		fmt.Printf("✅ %s COMPLETADO\n\n", test.nombre)
	}
}

func compararMetodos() {
	fmt.Println("\n📊 === COMPARATIVA DE MÉTODOS ===")
	fmt.Println("Ejecutando ambos métodos para Test Balanceado...")
	
	fmt.Println("\n--- RWMUTEX ---")
	taller := NuevoTallerRWMutex(2, 2, 1, 1)
	SimularTallerRWMutex(taller, 2, 1, 1)
	
	fmt.Println("\n--- WAITGROUP ---")
	SimularTallerWaitGroup(2, 1, 1)
	
	fmt.Println("\n✅ Comparativa completada")
}
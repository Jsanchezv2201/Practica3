package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("🚗 TALLER MECÁNICO - PRÁCTICA 3 - SISTEMAS DISTRIBUIDOS")
	fmt.Println("\n  TIEMPOS: Mecánica(5s±20%) | Eléctrica(3s±20%) | Carrocería(1s±20%)")
	fmt.Println("  PRIORIDAD: ALTA(🔴) > MEDIA(🟡) > BAJA(🟢) con colas de prioridad")
	fmt.Println(strings.Repeat("=", 70))

	for {
		fmt.Println("\n🔧 MENÚ PRINCIPAL - PRÁCTICA 3:")
		fmt.Println()
		fmt.Println("1. Ejecutar TODOS los TESTS del enunciado + COMPARATIVA")
		fmt.Println("   • 3 tests: (10-10-10) | (20-5-5) | (5-5-20)")
		fmt.Println("   • 2 métodos: RWMutex vs WaitGroup")
		fmt.Println()
		fmt.Println("2. Ejecutar TESTS específicos")
		fmt.Println()
		fmt.Println("0. Salir")
		fmt.Print("\n👉 Seleccione opción: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			ejecutarTodosTestsEnunciado() // OPCIÓN PRINCIPAL
		case 2:
			menuTestsEspecificos()         // SUBMENÚ para pruebas
		case 0:
			fmt.Println("\n👋 ¡Gracias por usar el sistema! Hasta luego.")
			return
		default:
			fmt.Println("❌ Opción no válida. Intente nuevamente.")
		}
	}
}

func ejecutarTodosTestsEnunciado() {
    fmt.Println("\n" + strings.Repeat("=", 80))
    fmt.Println("🚀 EJECUTANDO TODOS LOS TESTS DEL ENUNCIADO P3")
    fmt.Println("📄 Página 2: Comparativa RWMutex vs WaitGroup (OBLIGATORIO)")
    fmt.Println(strings.Repeat("=", 80))
    
    fmt.Printf("\n🔧 CONFIGURACIÓN ACTUAL DEL TALLER:\n")
    fmt.Printf("   • Plazas de espera: %d\n", ConfigTaller.NumPlazas)
    fmt.Printf("   • Mecánicos: %d\n", ConfigTaller.NumMecanicos)
    fmt.Printf("   • Puestos limpieza: %d\n", ConfigTaller.NumLimpieza)
    fmt.Printf("   • Puestos revisión: %d\n\n", ConfigTaller.NumRevision)

    // Preguntar número de simulaciones
    var sims int
    fmt.Print("\n🔢 ¿Cuántas veces desea ejecutar cada test (sims)? ")
    fmt.Scan(&sims)
    if sims <= 0 {
        sims = 1
    }
    
    fmt.Printf("📋 Ejecutando cada test %d veces\n\n", sims)
    
    fmt.Println("📋 TABLA DE TESTS REQUERIDOS POR EL ENUNCIADO:")
    fmt.Println("┌─────────┬──────────────┬──────────────┬──────────────┐")
    fmt.Println("│ Test #  │ Mecánica(A)  │ Eléctrica(B) │ Carrocería(C)│")
    fmt.Println("├─────────┼──────────────┼──────────────┼──────────────┤")
    fmt.Println("│    1    │      10      │      10      │      10      │")
    fmt.Println("│    2    │      20      │      5       │      5       │")
    fmt.Println("│    3    │      5       │      5       │      20      │")
    fmt.Println("└─────────┴──────────────┴──────────────┴──────────────┘")
    
    // Almacenar resultados para el resumen final
    resultados := make([][2]time.Duration, 3)
    
    // TEST 1: Balanceado (10-10-10)
    fmt.Println("\n" + strings.Repeat("═", 60))
    fmt.Println("🧪 TEST 1: Balanceado (10 mecánica, 10 eléctrica, 10 carrocería)")
    fmt.Println(strings.Repeat("═", 60))
    resultados[0] = ejecutarTestComparativo("Balanceado", 10, 10, 10, sims,
        ConfigTaller.NumPlazas, ConfigTaller.NumMecanicos, 
        ConfigTaller.NumLimpieza, ConfigTaller.NumRevision)    

    // TEST 2: Mayoría Mecánica (20-5-5)
    fmt.Println("\n" + strings.Repeat("═", 60))
    fmt.Println("🧪 TEST 2: Mayoría Mecánica (20-5-5)")
    fmt.Println(strings.Repeat("═", 60))
    resultados[1] = ejecutarTestComparativo("Mayoría Mecánica", 20, 5, 5, sims,
        ConfigTaller.NumPlazas, ConfigTaller.NumMecanicos,
        ConfigTaller.NumLimpieza, ConfigTaller.NumRevision)
    
    // TEST 3: Mayoría Carrocería (5-5-20)
    fmt.Println("\n" + strings.Repeat("═", 60))
    fmt.Println("🧪 TEST 3: Mayoría Carrocería (5-5-20)")
    fmt.Println(strings.Repeat("═", 60))
    resultados[2] = ejecutarTestComparativo("Mayoría Carrocería", 5, 5, 20, sims,
        ConfigTaller.NumPlazas, ConfigTaller.NumMecanicos,
        ConfigTaller.NumLimpieza, ConfigTaller.NumRevision)
    
    // RESUMEN FINAL - COMPARATIVA
    fmt.Println("\n" + strings.Repeat("=", 80))
    fmt.Printf("🎯 RESUMEN Y COMPARATIVA FINAL (%d sims por test)\n", sims)
    fmt.Println(strings.Repeat("=", 80))
    
    fmt.Println("| Test                        | RWMutex (prom) | WaitGroup (prom) | Diferencia    | Más rápido    |")
    fmt.Println("|-----------------------------|----------------|------------------|---------------|---------------|")
    
    for i, testName := range []string{
        "Balanceado (10-10-10)",
        "Mayoría Mecánica (20-5-5)",
        "Mayoría Carrocería (5-5-20)",
    } {
        tiempoMutex := resultados[i][0]
        tiempoWait := resultados[i][1]
        
        masRapido := "WaitGroup"
        diferencia := tiempoMutex - tiempoWait
        if tiempoMutex < tiempoWait {
            masRapido = "RWMutex"
            diferencia = tiempoWait - tiempoMutex
        }
        
        fmt.Printf("| %-27s | %-14v | %-16v | %-13v | %-13s |\n",
            testName,
            tiempoMutex.Round(time.Millisecond),
            tiempoWait.Round(time.Millisecond),
            diferencia.Round(time.Millisecond),
            masRapido)
    }
}

func ejecutarTestComparativo(nombre string, a, b, c, sims, plazas, mecanicos, limpieza, revision int) [2]time.Duration {
    fmt.Printf("\n🔄 Ejecutando %s - %d simulación(es)...\n", nombre, sims)
    
    var totalMutex, totalWait time.Duration
    
    for i := 0; i < sims; i++ {
        if sims > 1 {
            fmt.Printf("  📊 Simulación %d/%d:\n", i+1, sims)
        }
        
        if sims == 1 {
            fmt.Print("    🔄 RWMutex... ")
        }
        tallerMutex := NuevoTallerRWMutex(plazas, mecanicos, limpieza, revision) 
        inicioMutex := time.Now()
        SimularTallerRWMutex(tallerMutex, a, b, c)
        duracionMutex := time.Since(inicioMutex)
        totalMutex += duracionMutex
        if sims == 1 {
            fmt.Printf("completado en %v\n", duracionMutex.Round(time.Millisecond))
        } else {
            fmt.Printf("      • RWMutex: %v\n", duracionMutex.Round(time.Millisecond))
        }
        
        if sims == 1 {
            fmt.Print("    🔗 WaitGroup... ") 
        }
        tallerWait := NuevoTallerWaitGroup(plazas, mecanicos, limpieza, revision) 
        inicioWait := time.Now()
        SimularTallerWaitGroup(tallerWait, a, b, c)
        duracionWait := time.Since(inicioWait)
        totalWait += duracionWait
        if sims == 1 {
            fmt.Printf("completado en %v\n", duracionWait.Round(time.Millisecond))
        } else {
            fmt.Printf("      • WaitGroup: %v\n", duracionWait.Round(time.Millisecond))
        }
    }
    
    // Calcular promedios
    avgMutex := totalMutex / time.Duration(sims)
    avgWait := totalWait / time.Duration(sims)
    
    fmt.Printf("\n📊 RESULTADOS PROMEDIO %s (%d simulación(es)):\n", nombre, sims)
    fmt.Printf("   • RWMutex promedio:  %v\n", avgMutex.Round(time.Millisecond))
    fmt.Printf("   • WaitGroup promedio: %v\n", avgWait.Round(time.Millisecond))
    
    if avgMutex < avgWait {
        porcentaje := (float64(avgWait-avgMutex) / float64(avgWait)) * 100
        fmt.Printf("   🏆 RWMutex es %.1f%% más rápido (promedio)\n", porcentaje)
    } else if avgWait < avgMutex {
        porcentaje := (float64(avgMutex-avgWait) / float64(avgMutex)) * 100
        fmt.Printf("   🏆 WaitGroup es %.1f%% más rápido (promedio)\n", porcentaje)
    } else {
        fmt.Printf("   ⚖️  Ambos métodos tienen el mismo tiempo promedio\n")
    }
    
    return [2]time.Duration{avgMutex, avgWait}
}

// CAMBIA ejecutarTestEspecifico para usar ConfigTaller
func ejecutarTestEspecifico(nombre string, a, b, c int) {
    // Preguntar número de simulaciones
    var sims int
    fmt.Print("\n🔢 ¿Cuántas veces desea ejecutar este test (sims)? ")
    fmt.Scan(&sims)
    if sims <= 0 {
        sims = 1
    }

    fmt.Printf("\n🔧 Configuración: Plazas=%d, Mecánicos=%d, Limpieza=%d, Revisión=%d\n",
    ConfigTaller.NumPlazas, ConfigTaller.NumMecanicos, 
    ConfigTaller.NumLimpieza, ConfigTaller.NumRevision)
    
    fmt.Printf("\n🧪 %s - Seleccione método (ejecutando %d veces):\n", nombre, sims)
    fmt.Println("1. RWMutex")
    fmt.Println("2. WaitGroup")
    fmt.Print("\n👉 Método: ")
    
    var metodo int
    fmt.Scan(&metodo)
    
    switch metodo {
    case 1:
        fmt.Printf("\n🔄 Ejecutando %d simulación(es) con RWMutex...\n", sims)
        var total time.Duration
        for i := 0; i < sims; i++ {
            if sims > 1 {
                fmt.Printf("  📊 Simulación %d/%d: ", i+1, sims)
            }
            taller := NuevoTallerRWMutex(
                ConfigTaller.NumPlazas, 
                ConfigTaller.NumMecanicos,
                ConfigTaller.NumLimpieza, 
                ConfigTaller.NumRevision) // ✅
            inicio := time.Now()
            SimularTallerRWMutex(taller, a, b, c)
            duracion := time.Since(inicio)
            total += duracion
        }
        if sims > 1 {
            promedio := total / time.Duration(sims)
            fmt.Printf("\n✅ %d simulaciones completadas. Promedio: %v\n", 
                sims, promedio.Round(time.Millisecond))
        } else {
            fmt.Printf("✅ Completado en %v\n", total.Round(time.Millisecond))
        }
    case 2:
        fmt.Printf("\n🔗 Ejecutando %d simulación(es) con WaitGroup...\n", sims)
        var total time.Duration
        for i := 0; i < sims; i++ {
            if sims > 1 {
                fmt.Printf("  📊 Simulación %d/%d: ", i+1, sims)
            }
            taller := NuevoTallerWaitGroup(
                ConfigTaller.NumPlazas,
                ConfigTaller.NumMecanicos,
                ConfigTaller.NumLimpieza,
                ConfigTaller.NumRevision) // ✅
            inicio := time.Now()
            SimularTallerWaitGroup(taller, a, b, c)
            duracion := time.Since(inicio)
            total += duracion
        }
        if sims > 1 {
            promedio := total / time.Duration(sims)
            fmt.Printf("\n✅ %d simulaciones completadas. Promedio: %v\n", 
                sims, promedio.Round(time.Millisecond))
        } else {
            fmt.Printf("✅ Completado en %v\n", total.Round(time.Millisecond))
        }
    default:
        fmt.Println("❌ Método no válido")
    }
}

func menuTestsEspecificos() {
	for {
		fmt.Println("\n" + strings.Repeat("=", 60))
		fmt.Println("🧪 TESTS ESPECÍFICOS - PARA PRUEBAS Y ANÁLISIS")
		fmt.Println(strings.Repeat("=", 60))
		fmt.Println("\n1. Test Balanceado (10-10-10)")
		fmt.Println("   • 10 mecánica ")
		fmt.Println("   • 10 eléctrica ")
		fmt.Println("   • 10 carrocería ")
		
		fmt.Println("\n2. Test Mayoría Mecánica (20-5-5)")
		fmt.Println("   • 20 mecánica ")
		fmt.Println("   • 5 eléctrica")
		fmt.Println("   • 5 carrocería")
		
		fmt.Println("\n3. Test Mayoría Carrocería (5-5-20)")
		fmt.Println("   • 5 mecánica")
		fmt.Println("   • 5 eléctrica")
		fmt.Println("   • 20 carrocería ")
		
		fmt.Println("\n4. Volver al menú principal")
		fmt.Print("\n👉 Seleccione test: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			ejecutarTestEspecifico("Balanceado", 10, 10, 10)
		case 2:
			ejecutarTestEspecifico("Mayoría Mecánica", 20, 5, 5)
		case 3:
			ejecutarTestEspecifico("Mayoría Carrocería", 5, 5, 20)
		case 4:
			return
		default:
			fmt.Println("❌ Opción no válida")
		}
	}
}
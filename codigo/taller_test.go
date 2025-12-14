package main

import (
	"testing"
	"time"
)

// Tests según tabla del enunciado 

func TestBalanceado_10_10_10(t *testing.T) {
	plazas := ConfigTaller.NumPlazas
	mecanicos := ConfigTaller.NumMecanicos
	limpieza := ConfigTaller.NumLimpieza
	revision := ConfigTaller.NumRevision

	t.Logf("\n📊 Ejecutando Test Balanceado (10-10-10)")
	t.Logf("🔧 Configuración: Plazas=%d, Mecánicos=%d, Limpieza=%d, Revisión=%d",
		plazas, mecanicos, limpieza, revision)

	// Test RWMutex
	t.Log("🔧 Probando RWMutex...")
	start := time.Now()
	taller1 := NuevoTallerRWMutex(plazas, mecanicos, limpieza, revision)
	SimularTallerRWMutex(taller1, 10, 10, 10)
	duration1 := time.Since(start)

	timeout := 5 * time.Minute
	if duration1 > timeout {
		t.Errorf("❌ RWMutex Balanceado (10-10-10) tardó demasiado: %v", duration1)
	} else {
		t.Logf("✅ RWMutex Balanceado (10-10-10) completado en %v", duration1)
	}

	// Test WaitGroup
	t.Log("🔗 Probando WaitGroup...")
	start = time.Now()
	taller2 := NuevoTallerWaitGroup(plazas, mecanicos, limpieza, revision)
	SimularTallerWaitGroup(taller2, 10, 10, 10)
	duration2 := time.Since(start)

	if duration2 > timeout {
		t.Errorf("❌ WaitGroup Balanceado (10-10-10) tardó demasiado: %v", duration2)
	} else {
		t.Logf("✅ WaitGroup Balanceado (10-10-10) completado en %v", duration2)
	}

	// Comparación
	t.Logf("📈 Comparativa: RWMutex=%v vs WaitGroup=%v",
		duration1.Round(time.Second), duration2.Round(time.Second))
}

func TestMayoriaA_20_5_5(t *testing.T) {
	plazas := ConfigTaller.NumPlazas
	mecanicos := ConfigTaller.NumMecanicos
	limpieza := ConfigTaller.NumLimpieza
	revision := ConfigTaller.NumRevision

	t.Logf("\n📊 Ejecutando Test Mayoría Mecánica (20-5-5)")
	t.Logf("🔧 Configuración: Plazas=%d, Mecánicos=%d, Limpieza=%d, Revisión=%d",
		plazas, mecanicos, limpieza, revision)

	// Test RWMutex
	t.Log("🔧 Probando RWMutex...")
	start := time.Now()
	taller1 := NuevoTallerRWMutex(plazas, mecanicos, limpieza, revision)
	SimularTallerRWMutex(taller1, 20, 5, 5)
	duration1 := time.Since(start)

	timeout := 6 * time.Minute
	if duration1 > timeout {
		t.Errorf("❌ RWMutex Mayoría Mecánica (20-5-5) tardó demasiado: %v", duration1)
	} else {
		t.Logf("✅ RWMutex Mayoría Mecánica (20-5-5) completado en %v", duration1)
	}

	// Test WaitGroup
	t.Log("🔗 Probando WaitGroup...")
	start = time.Now()
	taller2 := NuevoTallerWaitGroup(plazas, mecanicos, limpieza, revision)
	SimularTallerWaitGroup(taller2, 20, 5, 5)
	duration2 := time.Since(start)

	if duration2 > timeout {
		t.Errorf("❌ WaitGroup Mayoría Mecánica (20-5-5) tardó demasiado: %v", duration2)
	} else {
		t.Logf("✅ WaitGroup Mayoría Mecánica (20-5-5) completado en %v", duration2)
	}

	// Comparación
	t.Logf("📈 Comparativa: RWMutex=%v vs WaitGroup=%v",
		duration1.Round(time.Second), duration2.Round(time.Second))
}

func TestMayoriaC_5_5_20(t *testing.T) {
	plazas := ConfigTaller.NumPlazas
	mecanicos := ConfigTaller.NumMecanicos
	limpieza := ConfigTaller.NumLimpieza
	revision := ConfigTaller.NumRevision

	t.Logf("\n📊 Ejecutando Test Mayoría Carrocería (5-5-20)")
	t.Logf("🔧 Configuración: Plazas=%d, Mecánicos=%d, Limpieza=%d, Revisión=%d",
		plazas, mecanicos, limpieza, revision)

	// Test RWMutex
	t.Log("🔧 Probando RWMutex...")
	start := time.Now()
	taller1 := NuevoTallerRWMutex(plazas, mecanicos, limpieza, revision)
	SimularTallerRWMutex(taller1, 5, 5, 20)
	duration1 := time.Since(start)

	timeout := 4 * time.Minute
	if duration1 > timeout {
		t.Errorf("❌ RWMutex Mayoría Carrocería (5-5-20) tardó demasiado: %v", duration1)
	} else {
		t.Logf("✅ RWMutex Mayoría Carrocería (5-5-20) completado en %v", duration1)
	}

	// Test WaitGroup
	t.Log("🔗 Probando WaitGroup...")
	start = time.Now()
	taller2 := NuevoTallerWaitGroup(plazas, mecanicos, limpieza, revision)
	SimularTallerWaitGroup(taller2, 5, 5, 20)
	duration2 := time.Since(start)

	if duration2 > timeout {
		t.Errorf("❌ WaitGroup Mayoría Carrocería (5-5-20) tardó demasiado: %v", duration2)
	} else {
		t.Logf("✅ WaitGroup Mayoría Carrocería (5-5-20) completado en %v", duration2)
	}

	// Comparación
	t.Logf("📈 Comparativa: RWMutex=%v vs WaitGroup=%v",
		duration1.Round(time.Second), duration2.Round(time.Second))
}

// Test de integración que ejecuta todos los tests del enunciado
func TestEnunciadoCompleto(t *testing.T) {
	t.Run("Balanceado_10_10_10", TestBalanceado_10_10_10)
	t.Run("MayoriaA_20_5_5", TestMayoriaA_20_5_5)
	t.Run("MayoriaC_5_5_20", TestMayoriaC_5_5_20)
}
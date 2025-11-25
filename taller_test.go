package main

import (
    "testing"
    "time"
)

func TestRWMutex(t *testing.T) {
    tests := []struct {
        name      string
        cochesA   int
        cochesB   int
        cochesC   int
        expected  time.Duration // tiempo máximo esperado
    }{
        {"Test1_Balanceado", 10, 10, 10, 2 * time.Minute},
        {"Test2_MayoriaA", 20, 5, 5, 3 * time.Minute},
        {"Test3_MayoriaC", 5, 5, 20, 90 * time.Second},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            start := time.Now()
            
            taller := NuevoTallerRWMutex(5, 3, 2, 2)
            SimularTallerRWMutex(taller, tt.cochesA, tt.cochesB, tt.cochesC)
            
            duration := time.Since(start)
            if duration > tt.expected {
                t.Errorf("Test %s tardó demasiado: %v", tt.name, duration)
            } else {
                t.Logf("✅ %s completado en %v", tt.name, duration)
            }
        })
    }
}

func TestWaitGroup(t *testing.T) {
    tests := []struct {
        name     string
        cochesA  int
        cochesB  int
        cochesC  int
        expected time.Duration
    }{
        {"Test1_Balanceado", 10, 10, 10, 2 * time.Minute},
        {"Test2_MayoriaA", 20, 5, 5, 3 * time.Minute},
        {"Test3_MayoriaC", 5, 5, 20, 90 * time.Second},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            start := time.Now()
            
            SimularTallerWaitGroup(tt.cochesA, tt.cochesB, tt.cochesC)
            
            duration := time.Since(start)
            if duration > tt.expected {
                t.Errorf("Test %s tardó demasiado: %v", tt.name, duration)
            } else {
                t.Logf("✅ %s completado en %v", tt.name, duration)
            }
        })
    }
}

func BenchmarkRWMutex(b *testing.B) {
    for i := 0; i < b.N; i++ {
        taller := NuevoTallerRWMutex(5, 3, 2, 2)
        SimularTallerRWMutex(taller, 10, 10, 10)
    }
}

func BenchmarkWaitGroup(b *testing.B) {
    for i := 0; i < b.N; i++ {
        SimularTallerWaitGroup(10, 10, 10)
    }
}
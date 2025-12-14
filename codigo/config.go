package main

// Configuración del taller - MODIFICAR AQUÍ PARA CAMBIAR LOS RECURSOS
var ConfigTaller = struct {
    NumPlazas    int
    NumMecanicos int
    NumLimpieza  int
    NumRevision  int
}{
    NumPlazas:    5,    // Número de plazas de espera
    NumMecanicos: 3,    // Número de mecánicos disponibles
    NumLimpieza:  2,    // Puestos de limpieza simultáneos
    NumRevision:  2,    // Puestos de revisión simultáneos
}
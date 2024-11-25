package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Generamos los datos
	proposed := generateDates(100000)

	// Prueba sin preasignación de capacidad
	fmt.Println("\n=== Sin Preasignación de Capacidad ===")
	printMemStats("Antes")
	start := time.Now()
	vdates := withoutPreallocation(proposed)
	fmt.Printf("Tiempo: %v, Elementos válidos: %d\n", time.Since(start), len(vdates))
	printMemStats("Después")
}

// Función sin preasignación de capacidad
func withoutPreallocation(proposed []time.Time) []time.Time {
	vdates := []time.Time{}
	for _, dd := range proposed {
		if canUseDeliveryDate(dd) {
			vdates = append(vdates, dd)
		}
	}
	return vdates
}

// Simulamos una validación de fecha
func canUseDeliveryDate(date time.Time) bool {
	// Por ejemplo, solo aceptamos fechas impares
	return date.Day()%2 != 0
}

// Genera un slice de fechas consecutivas para simular datos de entrada
func generateDates(count int) []time.Time {
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	dates := make([]time.Time, count)
	for i := 0; i < count; i++ {
		dates[i] = startDate.AddDate(0, 0, i) // Fechas consecutivas
	}
	return dates
}

// Utilidad para imprimir estadísticas de memoria
func printMemStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s - Memoria asignada: %v MB, Total asignada: %v MB, Heap en uso: %v MB\n",
		label,
		bToMb(m.Alloc),
		bToMb(m.TotalAlloc),
		bToMb(m.HeapAlloc),
	)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

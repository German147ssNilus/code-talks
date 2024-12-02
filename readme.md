## Memory Allocation en Slices

•	make(len): Al usar make([]T, len), estás diciendo “quiero un slice con esta longitud exacta”. Go asigna memoria para todos los elementos inmediatamente, incluso si no todos se usan.

•	make(0, cap): Al usar make([]T, 0, cap), estás diciendo “no quiero ningún elemento todavía, pero planeo agregar hasta cap elementos más tarde”. Esto permite agregar elementos solo cuando sea necesario.


Fuentes:

https://go.dev/doc/effective_go#slices

https://stackoverflow.com/questions/41668053/cap-vs-len-of-slice-in-golang

## **Resultados**


### **Con Preasignación de Capacidad**
```bash
go run with_prealloc.go

=== Con Preasignación de Capacidad ===
Antes - Memoria asignada: 2 MB, Total asignada: 2 MB, Heap en uso: 2 MB
Tiempo: 1.229833ms, Elementos válidos: 50992
Después - Memoria asignada: 4 MB, Total asignada: 4 MB, Heap en uso: 4 MB
```

## **Resultados**

### **Sin Preasignación de Capacidad**
```bash
go run without_prealloc.go

=== Sin Preasignación de Capacidad ===
Antes - Memoria asignada: 2 MB, Total asignada: 2 MB, Heap en uso: 2 MB
Tiempo: 3.288125ms, Elementos válidos: 50992
Después - Memoria asignada: 7 MB, Total asignada: 8 MB, Heap en uso: 7 MB
```


Antes
Los 2 MB iniciales vienen de:
- El runtime de Go
- Las estructuras básicas del programa
- El slice proposed con 100,000 fechas
- Variables del sistema


Despues sin prealloc

El incremento a 7 MB se debe a:
Los 2 MB iniciales
Múltiples realocaciones del slice mientras crece
Memoria extra reservada por el crecimiento dinámico
Memoria temporal usada durante las copias.


Con prealloc

Los 2 MB iniciales
Una única alocación del slice final
No hay memoria extra por realocaciones


```bash
// Sin preasignación: Crece múltiples veces
vdates := []time.Time{}  // Empieza vacío
// Crece: 0 -> 1 -> 2 -> 4 -> 8 -> 16 -> 32 -> 64 -> 128 -> ...

// Con preasignación: Una única alocación
vdates := make([]time.Time, 0, len(proposed))  // Reserva todo el espacio de una vez
// Tamaño: 0 -> final (sin pasos intermedios)
```

Conclusión:
El sistema tiene que copiar datos de un lugar de memoria a otro.

## https://german-mendieta.notion.site/Code-Talks-Effective-Go-1497d62b2c3480d8962fcd1b42b8a906?pvs=73

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

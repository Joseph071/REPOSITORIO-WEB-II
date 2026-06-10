# Cafetería ULEAM — Proyecto base del taller

API REST de la Cafetería Universitaria. Es el punto de partida del taller de hoy.

## ¿Qué hay ya implementado?

Las dos entidades **Producto** y **Categoria** están completas: modelo, storage,
handlers y rutas. La arquitectura es la de siempre: Chi + estructura
`cmd/` + `internal/`, y un struct `Server` que recibe el almacenamiento por
inyección de dependencias.

El almacenamiento vive detrás de la interfaz `storage.Almacen`, y ya hay **dos
implementaciones** que la cumplen:

- `Memoria` — guarda en slices, en memoria.
- `AlmacenSQLite` — persiste con GORM sobre SQLite.

El `Server` y los handlers dependen de la **interfaz**, no de una implementación
concreta. Por eso se puede cambiar el backend sin tocar un handler.

## Cómo correr

```bash
go mod tidy                  # resuelve dependencias y genera go.sum
go run ./cmd/cafeteria-api   # arranca en http://localhost:8080
```

Probar:

```bash
curl http://localhost:8080/api/v1/productos
curl http://localhost:8080/api/v1/categorias
```

> Este código fue verificado en **sintaxis** (`gofmt`) en el entorno de
> generación, pero las dependencias no se pudieron descargar ahí. Corré
> `go build ./... && go vet ./...` en tu máquina antes de usarlo.

## Hoy

Seguí las instrucciones del **handout del taller**. La consigna te va guiando
paso a paso.

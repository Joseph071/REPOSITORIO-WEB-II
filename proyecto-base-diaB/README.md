# Proyecto base — Día B · Semana 4

**Aplicaciones Web II · TDI-601 · ULEAM**

API en memoria de una plataforma de turismo multi-idioma. Es el código base
para el Laboratorio C3 (90 min, individual, sin IA).

## Dominio

3 entidades relacionadas:

```
   Turista                  Negocio
      |                        |
      |       CheckIn          |
      └──────TuristaID────────►| (FK)
              NegocioID────────┘ (FK)
```

| Entidad | Métodos del repositorio |
|---|---|
| `Negocio` | `Guardar`, `BuscarPorID`, `Listar`, `Eliminar` (4) |
| `Turista` | `Guardar`, `BuscarPorID`, `Listar` (3) |
| `CheckIn` | `Guardar`, `BuscarPorTurista`, `Listar` (3) |

Total: **10 métodos** que necesitan tests.

## Estructura

```
proyecto-base-diaB/
├── cmd/turismo/main.go              ← demo ejecutable con datos seed
├── internal/
│   ├── models/
│   │   ├── negocio.go               ← struct Negocio + sets válidos
│   │   ├── turista.go               ← struct Turista
│   │   └── checkin.go               ← struct CheckIn (con FKs)
│   ├── errs/errs.go                 ← ErrNoEncontrado, ErrYaExiste, ErrDatosInvalidos
│   └── storage/
│       ├── negocio_repository.go    ← interface + memoria
│       ├── turista_repository.go    ← interface + memoria
│       └── checkin_repository.go    ← interface + memoria con validación cruzada
├── go.mod
└── README.md
```

## Cómo correrlo

```bash
# 1. Verificá que compila
go build ./...

# 2. Corré la demo
go run ./cmd/turismo
```

Si imprime los 5 negocios, 4 turistas y 4 check-ins, está listo.

## Lo importante para el taller

**1. NO modifiques los archivos `.go` existentes.** Solo agregás archivos
`*_test.go` al lado de los archivos que querés probar.

**2. Cada test va en el mismo paquete que el código que prueba.** Es decir:
- `negocio_repository_test.go` va en `internal/storage/`
- NO se crea una carpeta `tests/`

**3. Validación cruzada en `CheckIn`.** Cuando guardás un check-in, el
repositorio verifica que el `TuristaID` y el `NegocioID` realmente existan
en sus respectivos repositorios. Esto significa que tus tests del repo de
check-ins necesitan **sembrar primero** un turista y un negocio.

## Comandos útiles

```bash
# Correr todos los tests
go test ./...

# Ver cobertura del paquete storage
go test -cover ./internal/storage

# Generar reporte HTML de cobertura
go test -coverprofile=cover.out ./internal/storage
go tool cover -html=cover.out
```

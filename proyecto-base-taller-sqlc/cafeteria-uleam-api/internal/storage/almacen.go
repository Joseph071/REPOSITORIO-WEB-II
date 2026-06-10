package storage

import "cafeteria-uleam-api/internal/models"

// Almacen define QUÉ sabe hacer un almacén de la cafetería, sin decir CÓMO.
//
// Memoria (slices) ya cumple esta interfaz sin cambios — por el duck typing
// que vimos en S3 — y AlmacenSQLite (GORM) la cumple igual. El Server depende
// de esta interfaz, no de una implementación concreta: por eso podemos cambiar
// el backend de almacenamiento sin tocar un solo handler.
type Almacen interface {
	// Productos
	ListarProductos() []models.Producto
	BuscarProductoPorID(id int) (models.Producto, bool)
	CrearProducto(p models.Producto) models.Producto
	ActualizarProducto(id int, datos models.Producto) (models.Producto, bool)
	BorrarProducto(id int) bool

	// Categorias
	ListarCategorias() []models.Categoria
	BuscarCategoriaPorID(id int) (models.Categoria, bool)
	CrearCategoria(c models.Categoria) models.Categoria
	ActualizarCategoria(id int, datos models.Categoria) (models.Categoria, bool)
	BorrarCategoria(id int) bool
}

// Chequeo en tiempo de compilación: si Memoria dejara de cumplir Almacen,
// el proyecto NO compila. Red de seguridad opcional.
var _ Almacen = (*Memoria)(nil)

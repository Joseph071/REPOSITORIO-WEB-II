// Package errs define los errores tipados del dominio de Turismo.
//
// Estos errores son la frontera entre lo que retornan los repositorios y lo
// que los handlers HTTP convertirán en códigos de respuesta. Se usan con
// errors.Is en los tests y en el código consumidor.
package errs

import "errors"

var (
	// ErrNoEncontrado indica que un recurso solicitado por ID o por algún otro
	// criterio no existe en el repositorio.
	ErrNoEncontrado = errors.New("recurso no encontrado")

	// ErrYaExiste indica que se intentó guardar un recurso con un ID que ya
	// está ocupado por otro recurso.
	ErrYaExiste = errors.New("el recurso ya existe")

	// ErrDatosInvalidos indica que los datos proporcionados no cumplen las
	// reglas de validación: campos obligatorios vacíos, valores fuera de
	// rango, formatos incorrectos, etc.
	ErrDatosInvalidos = errors.New("datos inválidos")
)

package storage

import (
	"errors"
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

func TestGuardar_TableDriven_Turista(t *testing.T) {
	repo := NewTuristaMemoria()

	// Pre-condición: sembramos un negocio para poder probar "ID duplicado".
	turistaBase := models.Turista{
		ID: 1, Nombre: "Valeria", Nacionalidad: "Veneca",
		IdiomaPreferido:  "en",
	}
	if err := repo.Guardar(turistaBase); err != nil {
		// t.Fatalf detiene el test inmediatamente. Si el setup falla,
		// no tiene sentido seguir corriendo el resto de los casos.
		t.Fatalf("setup falló: %v", err)
	}

	// La tabla de casos. Cada elemento es un escenario completo:
	// nombre del subtest, datos de entrada, error esperado.
	casos := []struct {
		nombre    string
		entrada   models.Turista
		esperaErr error
	}{
		{
			nombre: "caso feliz - Turista válido",
			entrada: models.Turista{
				ID: 100, Nombre: "jostin", Nacionalidad:"Ecuador",
				IdiomaPreferido: "es",
			},
			esperaErr: nil,
		},
		{
			nombre: "nombre vacío falla",
			entrada: models.Turista{
				ID: 101, Nombre: "", Nacionalidad: "peruano",
				IdiomaPreferido:  "es",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "Nacionalidad no válido falla",
			entrada: models.Turista{
				ID: 102, Nombre: "turista X", Nacionalidad: "",
				IdiomaPreferido: "es",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "lista de idiomas vacía falla",
			entrada: models.Turista{
				ID: 103, Nombre: "Turista Y", Nacionalidad: "Europeo",
				IdiomaPreferido:  "",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "idioma no soportado falla",
			entrada: models.Turista{
				ID: 104, Nombre: "juan", Nacionalidad: "tour",
				IdiomaPreferido:  "ja",// ja=japonés no está en la lista
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "ID duplicado falla",
			entrada: models.Turista{
				ID: 1, Nombre: "pal", Nacionalidad: "peruano",
				IdiomaPreferido:  "en",
			},
			esperaErr: errs.ErrYaExiste,
		},
	}

	// Iteramos sobre los casos y corremos un subtest por cada uno.
	// t.Run permite que cada subtest se reporte por separado y que se
	// puedan correr individualmente con `go test -run`.
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			err := repo.Guardar(c.entrada)

			// errors.Is es la forma idiomática de comparar errores
			// tipados. NUNCA uses err == c.esperaErr ni
			// err.Error() == "..." — son frágiles.
			if !errors.Is(err, c.esperaErr) {
				t.Errorf("Guardar(%q): esperaba error=%v, obtuvo error=%v",
					c.entrada.Nombre, c.esperaErr, err)
			}
		})
	}
}

func TestBuscarPorID_TuristaExiste(t *testing.T) {
	repo := NewTuristaMemoria()

	// Arrange: creamos y guardamos un negocio.
	esperado := models.Turista{
		ID: 42, Nombre: "xxx", Nacionalidad: "dsada",
		IdiomaPreferido: "es",
	}
	if err := repo.Guardar(esperado); err != nil {
		t.Fatalf("setup falló: %v", err)
	}

	// Act: buscamos el negocio por su ID.
	obtenido, err := repo.BuscarPorID(42)

	// Assert: no debe haber error y debe coincidir con lo guardado.
	if err != nil {
		t.Fatalf("no esperaba error: %v", err)
	}
	if obtenido.ID != esperado.ID {
		t.Errorf("ID: esperaba %d, obtuvo %d", esperado.ID, obtenido.ID)
	}
	if obtenido.Nombre != esperado.Nombre {
		t.Errorf("Nombre: esperaba %q, obtuvo %q", esperado.Nombre, obtenido.Nombre)
	}
	
}

func TestTuristaMemoria_Eliminar(t *testing.T) {
 // Negocio que sembramos en cada caso para tener algo que eliminar
 base := models.Turista{
 ID: 1, Nombre: "jostin", Nacionalidad: "ecuador",
 IdiomaPreferido: "es",
 }
 casos := []struct {
 nombre string
 idEliminar int // ID que se intenta eliminar
 errEsperado error
 }{
 {
 nombre: "elimina un turista existente",
 idEliminar: 1,
 errEsperado: nil,
 },
 {
 nombre: "ID inexistente retorna ErrNoEncontrado",
 idEliminar: 999,
 errEsperado: errs.ErrNoEncontrado,
 },
 }
 for _, c := range casos {
 t.Run(c.nombre, func(t *testing.T) {
 // Arrange: cada caso con su propio repo + el negocio sembrado
 repo := NewTuristaMemoria()
 if err := repo.Guardar(base); err != nil {
 t.Fatalf("setup falló: %v", err)
 }
 // Act
 err := repo.Eliminar(c.idEliminar)
 // Assert
 if !errors.Is(err, c.errEsperado) {
 t.Errorf("esperaba error %v, obtuvo %v", c.errEsperado, err)
 }
 })
 }
}

func TestTurista_Listar(t *testing.T) {

	// Arrange
	repo := NewTuristaMemoria()

	turista1 := models.Turista{
		ID:              1,
		Nombre:          "juan",
		Nacionalidad:    "peruana",
		
		IdiomaPreferido: "es",
	}

	turista2 := models.Turista{
		ID:              2,
		Nombre:          "pablo",
		Nacionalidad:            "colombiana",
		
		IdiomaPreferido: "en",
	}

	if err := repo.Guardar(turista1); err != nil {
		t.Fatalf("error guardando turista1: %v", err)
	}

	if err := repo.Guardar(turista2); err != nil {
		t.Fatalf("error guardando turista2: %v", err)
	}

	// Act
	lista := repo.Listar()

	// Assert
	if len(lista) != 2 {
		t.Errorf("esperaba 2 turistas, obtuvo %d", len(lista))
	}
}
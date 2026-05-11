// Este archivo contiene DOS tests RESUELTOS como ejemplo. Sirven como
// modelo de cómo escribir los tests que faltan: los 8 métodos restantes
// del taller no tienen tests todavía y debés escribirlos vos siguiendo
// estos patrones.
//
// Lo que aprendés leyendo este archivo:
//
//  1. TestGuardar_TableDriven — patrón table-driven con t.Run y subtests.
//     Aplicalo a métodos que tienen MÚLTIPLES casos de validación.
//
//  2. TestBuscarPorID_CheckInExiste — test simple de un solo caso.
//     Aplicalo a métodos con UN comportamiento esperado.
//
// IMPORTANTE: este archivo es solo un ejemplo. Vos vas a crear archivos
// nuevos como turista_repository_test.go y checkin_repository_test.go
// para los otros 8 métodos.
package storage

import (
	"errors"
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

// TestGuardar_TableDriven cubre los 6 escenarios de CheckIn.Guardar usando
// el patrón table-driven idiomático de Go.
//
// Los 6 casos cubren:
//
//  1. Caso feliz — un negocio válido se guarda sin error
//  2. Nombre vacío — debe fallar con ErrDatosInvalidos
//  3. Tipo no válido — debe fallar con ErrDatosInvalidos
//  4. Idiomas vacío — debe fallar con ErrDatosInvalidos
//  5. Idioma no soportado — debe fallar con ErrDatosInvalidos
//  6. ID duplicado — debe fallar con ErrYaExiste
//
// El primer caso siembra el repo. El sexto caso reusa ese mismo repo para
// probar el ID duplicado. Por eso el repo se construye UNA SOLA VEZ fuera
// del bucle, no dentro.
func TestCheckInGuardar_TableDriven(t *testing.T) {
	turistas := NewTuristaMemoria()
	negocios := NewNegocioMemoria()
	repo := NewCheckInMemoria(turistas, negocios)

	// Arrange: creamos y guardamos un negocio.

	turista := models.Turista{
		ID: 1, Nombre: "Juan", Nacionalidad: "Ecuatoriana", IdiomaPreferido: "es",
	}
	negocio := models.Negocio{
		ID: 1, Nombre: "Café Manabita", Tipo: "restaurante",
		Ciudad: "Manta", IdiomasHablados: []string{"es", "en"}, Activo: true,
	}
	esperado := models.CheckIn{
		ID: 1, TuristaID: 1, NegocioID: 1, Fecha: "25/10/2023", Calificacion: 5,
	}
	if err := turistas.Guardar(turista); err != nil {
		t.Fatalf("setup falló: %v", err)
	}
	if err := negocios.Guardar(negocio); err != nil {
		t.Fatalf("setup falló: %v", err)
	}
	if err := repo.Guardar(esperado); err != nil {
		t.Fatalf("setup falló: %v", err)
	}

	// La tabla de casos. Cada elemento es un escenario completo:
	// nombre del subtest, datos de entrada, error esperado.
	casos := []struct {
		nombre    string
		entrada   models.CheckIn
		esperaErr error
	}{
		{
			nombre: "caso feliz - checkin válido",
			entrada: models.CheckIn{
				ID: 100, TuristaID: 1, NegocioID: 1, Fecha: "25/10/2023", Calificacion: 5,
			},
			esperaErr: nil,
		},
	
		{
			nombre: "nombre fecha Vacia",
			entrada: models.CheckIn{
				ID: 101, TuristaID: 101, NegocioID: 101, Fecha: "", Calificacion: 5,
			},
			esperaErr: errs.ErrDatosInvalidos,
		},{
			nombre: "Calificación = 0 (fuera de 1-5)",
			entrada: models.CheckIn{
				ID: 101, TuristaID: 101, NegocioID: 101, Fecha: "25/10/2023", Calificacion: 0,
			},
			esperaErr: errs.ErrDatosInvalidos,
		},{
			nombre: "Calificación = 6 (fuera de 1-5)",
			entrada: models.CheckIn{
				ID: 101, TuristaID: 101, NegocioID: 101, Fecha: "25/10/2023", Calificacion: 6,
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "TuristaID que NO existe",
			entrada: models.CheckIn{
				ID: 101, TuristaID: 999 , NegocioID: 101, Fecha: "25/10/2023", Calificacion: 5,
			},
			esperaErr: errs.ErrNoEncontrado,
		},{
			nombre: "NegocioID que NO existe",
			entrada: models.CheckIn{
				ID: 101, TuristaID: 101, NegocioID: 999 , Fecha: "25/10/2023", Calificacion: 5,
			},
			esperaErr: errs.ErrNoEncontrado,
		},
		{
			nombre: "ID duplicado falla",
			entrada: models.CheckIn{
				ID: 1, TuristaID: 1, NegocioID: 1, Fecha: "25/10/2023", Calificacion: 5,
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
					c.entrada.ID, c.esperaErr, err)
			}
		})
	}
}

// TestBuscarPorID_CheckInExiste verifica el caso feliz de BuscarPorID.
//
// Este es un test SIMPLE de un solo caso. No necesita el patrón
// table-driven porque solo hay un comportamiento esperado a verificar.
//
// Los OTROS casos de BuscarPorID (ID negativo, ID inexistente) deberían
// ir en otro test, posiblemente table-driven, que VOS tenés que escribir.
func TestBuscarPorIDTurista_CheckInExiste(t *testing.T) {
	turistas := NewTuristaMemoria()
	negocios := NewNegocioMemoria()
	repo := NewCheckInMemoria(turistas, negocios)

	// Arrange: creamos y guardamos un negocio.

	turista := models.Turista{
		ID: 1, Nombre: "Juan", Nacionalidad: "Ecuatoriana", IdiomaPreferido: "es",
	}
	negocio := models.Negocio{
		ID: 1, Nombre: "Café Manabita", Tipo: "hotel",
		Ciudad: "Manta", IdiomasHablados: []string{"es", "en"}, Activo: true,
	}
	esperado := models.CheckIn{
		ID: 1, TuristaID: 1, NegocioID: 1, Fecha: "25/10/2023", Calificacion: 5,
	}
	if err := turistas.Guardar(turista); err != nil {
		t.Fatalf("setup falló: %v", err)
	}
	if err := negocios.Guardar(negocio); err != nil {
		t.Fatalf("setup falló: %v", err)
	}
	if err := repo.Guardar(esperado); err != nil {
		t.Fatalf("setup falló: %v", err)
	}

	// Act: buscamos el checkin turista ID.
	obtenidos, err := repo.BuscarPorTurista(1)
	obtenido := obtenidos[0]

	// Assert: no debe haber error y debe coincidir con lo guardado.
	if err != nil {
		t.Fatalf("no esperaba error: %v", err)
	}
	if obtenido.ID != esperado.ID {
		t.Errorf("ID: esperaba %d, obtuvo %d", esperado.ID, obtenido.ID)
	}
	if obtenido.TuristaID != esperado.TuristaID {
		t.Errorf("TuristaID: esperaba %q, obtuvo %q", esperado.TuristaID, obtenido.TuristaID)
	}
	if obtenido.NegocioID != esperado.NegocioID {
		t.Errorf("NegocioID: esperaba %q, obtuvo %q", esperado.NegocioID, obtenido.NegocioID)
	}
}

func TestCheckInMemoria_Listar(t *testing.T){
	turistas := NewTuristaMemoria()
	negocios := NewNegocioMemoria()
	repo := NewCheckInMemoria(turistas, negocios)

	turista := models.Turista{
		ID: 1, Nombre: "Juan", Nacionalidad: "Ecuatoriana", IdiomaPreferido: "es",
	}
	negocio := models.Negocio{
		ID: 1, Nombre: "Café Manabita", Tipo: "hotel",
		Ciudad: "Manta", IdiomasHablados: []string{"es", "en"}, Activo: true,
	}
	esperado := models.CheckIn{
		ID: 1, TuristaID: 1, NegocioID: 1, Fecha: "25/10/2023", Calificacion: 5,
	}
	if err := turistas.Guardar(turista); err != nil {
		t.Fatalf("setup falló: %v", err)
	}
	if err := negocios.Guardar(negocio); err != nil {
		t.Fatalf("setup falló: %v", err)
	}
	if err := repo.Guardar(esperado); err != nil {
		t.Fatalf("setup falló: %v", err)
	}

	lista := repo.Listar()

	// Assert
	if len(lista) != 1 {
		t.Errorf("esperaba 1 checkin 1 id negocios y 1 id turista, obtuvo %d", len(lista))
	}

	

}
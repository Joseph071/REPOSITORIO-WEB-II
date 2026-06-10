package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"cafeteria-uleam-api/internal/models"
)

// ListarCategorias atiende GET /api/v1/categorias.
func (s *Server) ListarCategorias(w http.ResponseWriter, _ *http.Request) {
	categorias := s.Storage.ListarCategorias()
	RespondJSON(w, http.StatusOK, categorias)
}

// ObtenerCategoria atiende GET /api/v1/categorias/{id}.
func (s *Server) ObtenerCategoria(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	categoria, encontrado := s.Storage.BuscarCategoriaPorID(id)
	if !encontrado {
		RespondError(w, http.StatusNotFound, "categoría no encontrada")
		return
	}

	RespondJSON(w, http.StatusOK, categoria)
}

// CrearCategoria atiende POST /api/v1/categorias.
func (s *Server) CrearCategoria(w http.ResponseWriter, r *http.Request) {
	var nueva models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(nueva.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}

	creada := s.Storage.CrearCategoria(nueva)
	RespondJSON(w, http.StatusCreated, creada)
}

// ActualizarCategoria atiende PUT /api/v1/categorias/{id}.
func (s *Server) ActualizarCategoria(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos models.Categoria
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		RespondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if strings.TrimSpace(datos.Nombre) == "" {
		RespondError(w, http.StatusBadRequest, "el campo nombre es obligatorio")
		return
	}

	actualizada, encontrada := s.Storage.ActualizarCategoria(id, datos)
	if !encontrada {
		RespondError(w, http.StatusNotFound, "categoría no encontrada")
		return
	}

	RespondJSON(w, http.StatusOK, actualizada)
}

// BorrarCategoria atiende DELETE /api/v1/categorias/{id}.
func (s *Server) BorrarCategoria(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if !s.Storage.BorrarCategoria(id) {
		RespondError(w, http.StatusNotFound, "categoría no encontrada")
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}

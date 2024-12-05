package handler

import (
	"errors"
	"regexp"
	"strconv"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/dentista"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/pkg/web"
	"github.com/gin-gonic/gin"
	"strings"
)

type dentistaHandler struct {
	s dentista.DentistaService
}

// // Creo un nuevo controller de dentistas
func NewDentistaHandler(s dentista.DentistaService) *dentistaHandler {
	return &dentistaHandler{
		s: s,
	}
}


// GetByMatricula godoc
// @Summary Buscar dentista por Matrícula
// @Description Obtiene un dentista a partir de su Matrícula
// @Tags Dentistas
// @Accept json
// @Produce json
// @Param matricula path string true "Matrícula del Dentista"
// @Success 200 {object} domain.Dentista
// @Failure 404 {object} web.errorResponse
// @Router /dentistas/matricula/{matricula} [get]
func (d *dentistaHandler) GetByMatricula() gin.HandlerFunc {
	return func(c *gin.Context) {
		matricula := c.Param("matricula") 
		
		/*
		if err != nil {
			web.Failure(c, 400, errors.New("Id invalido"))
			return
		}*///ver si validar el dni 
		dentista, err := d.s.GetByMatricula(matricula)
		if err != nil {
			web.Failure(c, 404, errors.New("Dentista no encontrado"))
			return
		}
		web.Success(c, 200, dentista)//Buscar paciente por dni
	}
}

// GetAll godoc
// @Summary Obtener todos los dentistas
// @Description Obtiene una lista de todos los dentistas
// @Tags Dentistas
// @Produce json
// @Success 200 {array} domain.Dentista
// @Router /dentistas [get]
func (d *dentistaHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}*/
		dentistas, _ := d.s.GetAll()
		web.Success(c, 200, dentistas)
	}
}


// GetByID godoc
// @Summary Buscar dentista por ID
// @Description Obtiene un dentista a partir de su ID
// @Tags Dentistas
// @Accept json
// @Produce json
// @Param id path int true "ID del Dentista"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas/{id} [get]
func (h *dentistaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentista, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentista not found"))
			return
		}
		web.Success(c, 200, dentista)
	}
}

// validateEmptys valida que los campos no esten vacio
func validateEmptysDentista(dentista *domain.Dentista) (bool, error) {
	// Verificar que los campos no estén vacíos
	if dentista.Nombre == "" {
		return false, errors.New("los campos no pueden estar vacíos")
	}
	

	// Valida formato del Matrícula (solo números y longitud entre 7 y 10 caracteres)
	matriculaRegex := regexp.MustCompile(`^\d{7,10}$`)
	if !matriculaRegex.MatchString(dentista.Matricula) {
		return false, errors.New("la matrícula debe contener solo números y tener entre 7 y 10 dígitos")
	}
	return true, nil

}

// Post godoc
// @Summary Agregar dentista
// @Description Crea un nuevo dentista en el sistema
// @Tags Dentistas
// @Accept json
// @Produce json
// @Param dentista body domain.Dentista true "Datos del dentista"
// @Success 201 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Router /dentistas [post]
func (h *dentistaHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentista domain.Dentista

		err := c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysDentista(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Create(dentista)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

// Delete godoc
// @Summary Eliminar dentista
// @Description Elimina un dentista existente
// @Tags Dentistas
// @Param id path int true "ID del Dentista"
// @Success 200 {object} map[string]string
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentistas/{id} [delete]
func (h *dentistaHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("id inválido"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			if strings.Contains(err.Error(), "no encontrado") {
				web.Failure(c, 404, errors.New("dentista no encontrado"))
			} else {
				web.Failure(c, 500, err)
			}
			return
		}
		// Enviar mensaje de éxito
		web.Success(c, 200, gin.H{"message": "Dentista borrado exitosamente"})
	}
}


// Put godoc
// @Summary Modificar dentista
// @Description Actualiza todos los campos de un dentista
// @Tags Dentistas
// @Accept json
// @Produce json
// @Param id path int true "ID del Dentista"
// @Param dentista body domain.Dentista true "Datos del dentista"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Router /dentistas/{id} [put]
func (h *dentistaHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentista not found"))
			return
		}

		var dentista domain.Dentista
		err = c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysDentista(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Update(id, dentista)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

// Patch godoc
// @Summary Actualizar parcialmente dentista
// @Description Modifica uno o varios campos de un dentista existente
// @Tags Dentistas
// @Accept json
// @Produce json
// @Param id path int true "ID del Dentista"
// @Param dentista body object true "Campos opcionales del dentista"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Router /dentistas/{id} [patch]
func (h *dentistaHandler) Patch() gin.HandlerFunc {
    type Request struct {
        Nombre    string `json:"nombre,omitempty"`
        Apellido  string `json:"apellido,omitempty"`
        Matricula string `json:"matricula,omitempty"`
    }
    return func(c *gin.Context) {
        var r Request
        idParam := c.Param("id")
        id, err := strconv.Atoi(idParam)
        if err != nil {
            web.Failure(c, 400, errors.New("invalid id"))
            return
        }

        // Verificar si el dentista existe
        _, err = h.s.GetByID(id)
        if err != nil {
            web.Failure(c, 404, err) // Use the error message returned
            return
        }

        // Validar el JSON recibido
        if err := c.ShouldBindJSON(&r); err != nil {
            web.Failure(c, 400, errors.New("invalid json"))
				
			return
        }

        // Crear el objeto de actualización
        update := domain.Dentista{
            Nombre:    r.Nombre,
            Apellido:  r.Apellido,
            Matricula: r.Matricula,
        }

        // Aplicar la actualización
        d, err := h.s.Patch(id, update)
        if err != nil {
            web.Failure(c, 409, err)
            return
        }

        web.Success(c, 200, d)
    }
}

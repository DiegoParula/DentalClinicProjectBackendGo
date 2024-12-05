package handler

import (
	"errors"
	"fmt"
	"regexp"

	"strconv"
	"time"

	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/paciente"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/pkg/web"
	"github.com/gin-gonic/gin"
	
)

type pacienteHandler struct {
	s paciente.Service
}

// Creo un nuevo controller de pacientes
func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{s}
}

// BuscarPorID godoc
// @Summary Buscar paciente por ID
// @Description Obtiene un paciente a partir de su ID
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param id path int true "ID del Paciente"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse 
// @Failure 404 {object} web.errorResponse 
// @Router /pacientes/{id} [get]
func (h *pacienteHandler) BuscarPorID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(c, 400, errors.New("Id invalido"))
			return
		}
		paciente, err := h.s.BuscarPorID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Paciente no encontrado"))
			return
		}
		web.Success(c, 200, paciente)
	}
}

// BuscarPorDNI godoc
// @Summary Buscar paciente por DNI
// @Description Obtiene un paciente a partir de su DNI
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param dni path string true "DNI del Paciente"
// @Success 200 {object} domain.Paciente
// @Failure 404 {object} web.errorResponse 
// @Router /pacientes/dni/{dni} [get]
func (h *pacienteHandler) BuscarPorDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		dni := c.Param("dni")

		paciente, err := h.s.BuscarPorDNI(dni)
		if err != nil {
			web.Failure(c, 404, errors.New("Paciente no encontrado"))
			return
		}
		web.Success(c, 200, paciente)
	}
}

// validateEmptys valida que los campos no esten vacio
func validateEmptys(paciente *domain.Paciente) (bool, error) {
	// Verificar que los campos no estén vacíos
	if paciente.Nombre == "" {
		return false, errors.New("los campos no pueden estar vacíos")
	}
	

	// Valida formato del DNI (solo números y longitud entre 7 y 10 caracteres)
	dniRegex := regexp.MustCompile(`^\d{7,10}$`)
	if !dniRegex.MatchString(paciente.DNI) {
		return false, errors.New("el DNI debe contener solo números y tener entre 7 y 10 dígitos")
	}
	return true, nil

}

// Post godoc
// @Summary Agregar paciente
// @Description Crea un nuevo paciente en el sistema, no se le pasa fecha porque toma la fecha actual del sistema
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param paciente body domain.Paciente true "Datos del paciente"
// @Success 201 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse 
// @Router /pacientes [post]
func (h *pacienteHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paciente domain.Paciente

		// Imprime datos recibidos para depurar
		fmt.Printf("Datos recibidos: %+v\n", paciente)
		

		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			//c.JSON(400, gin.H{"status": 400, "code": "Bad Request", "message": err.Error()})
			return
		}


		valid, err := validateEmptys(&paciente)
		if !valid {
			//c.JSON(400, gin.H{"status": 400, "code": "Bad Request", "message": err.Error()})
			web.Failure(c, 400, err)
			return
		}
		// Asignamos la fecha actual a fechaAlta
		paciente.FechaAlta = time.Now()
		///ver de crear para validar la fecha
		p, err := h.s.Agregar(paciente)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}

		
		//paciente.ID = p.ID//para que retorno el paciente con el id creado y no 0
	web.Success(c, 201, p)
	}
}

// Delete godoc
// @Summary Eliminar paciente
// @Description Elimina un paciente existente
// @Tags Pacientes
// @Param id path int true "ID del Paciente"
// @Success 204 
// @Failure 400 {object} web.errorResponse 
// @Failure 404 {object} web.errorResponse 
// @Router /pacientes/{id} [delete]
func (h *pacienteHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}

		err = h.s.Eliminar(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Put godoc
// @Summary Modificar paciente
// @Description Actualiza todos los campos de un paciente existente no le pasamos fecha porque no se puede actualizar
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param id path int true "ID del Paciente"
// @Param paciente body domain.Paciente true "Datos del paciente"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse 
// @Failure 404 {object} web.errorResponse 
// @Router /pacientes/{id} [put]
func (h *pacienteHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {

		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Id invalido"))
			return
		}
		_, err = h.s.BuscarPorID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Paciente no encontrado"))
			return
		}

		var paciente domain.Paciente
		err = c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		p, err := h.s.Modificar(id, paciente)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch godoc
// @Summary Actualizar parcialmente paciente
// @Description Modifica uno o varios campos de un paciente existente no le pasamos fecha porque no se puede actualizar ya que es la fecha de creacion y no deberia 
// @Tags Pacientes
// @Accept json
// @Produce json
// @Param id path int true "ID del Paciente"
// @Param paciente body domain.Paciente true "Campos opcionales del paciente"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse 
// @Failure 404 {object} web.errorResponse 
// @Failure 409 {object} web.errorResponse 
// @Router /pacientes/{id} [patch]
func (h *pacienteHandler) Patch() gin.HandlerFunc {
	type Request struct {
		ID        int       `json:"id"`
		Nombre    string    `json:"nombre,omitempty"`
		Apellido  string    `json:"apellido,omitempty"`
		Direccion string    `json:"direccion,omitempty"`
		DNI       string    `json:"dni,omitempty"`
		FechaAlta time.Time `json:"fechaAlta,omitempty"`
	}
	return func(c *gin.Context) {

		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Id invalido"))
			return
		}
		_, err = h.s.BuscarPorID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Paciente no encontrado"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Paciente{
			Nombre:    r.Nombre,
			Apellido:  r.Apellido,
			Direccion: r.Direccion,
			DNI:       r.DNI,
			FechaAlta: r.FechaAlta,
		}

		p, err := h.s.Modificar(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}

}

// GetAll godoc
// @Summary      Obtener todos los pacientes
// @Description  Obtiene una lista de todos los pacientes.
// @Tags         Pacientes
// @Produce      json
// @Success      200  {array}   domain.Paciente
// @Router       /pacientes [get]
func (h *pacienteHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		pacientes, _ := h.s.Listar()
		web.Success(c, 200, pacientes)
	}
}

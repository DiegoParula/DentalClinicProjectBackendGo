package handler

import (
	"errors"
	"strconv"
	"time"

	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/domain"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/turno"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/pkg/web"
	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	s turno.ServiceTurno
}

// Creo un nuevo controller de turnos
func NewTurnoHandler(s turno.ServiceTurno) *turnoHandler {
	return &turnoHandler{s}
}

// BuscarPorID godoc
// @Summary Obtener un turno por ID
// @Description Obtiene los detalles de un turno específico por su ID
// @Tags turnos
// @Accept json
// @Produce json
// @Param id path int true "ID del Turno"
// @Success 200 {object} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos/{id} [get]
func (h *turnoHandler) BuscarPorID() gin.HandlerFunc {

	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(c, 400, errors.New("Id invalido"))
			return
		}
		turno, err := h.s.BuscarPorID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Turno no encontrado"))
			return
		}
		web.Success(c, 200, turno)
	}
}


// POST godoc
// @Summary Crear un nuevo turno
// @Description Agrega un nuevo turno al sistema, el paciente y dentista deben existir
// @Tags turnos
// @Accept json
// @Produce json
// @Param turno body domain.Turno true "Objeto Turno"
// @Success 201 {object} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Router /turnos [post]
func (h *turnoHandler) POST() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno
		err := c.ShouldBindJSON(&turno)
		if err != nil {
			//c.JSON(400, gin.H{"status": 400, "code": "Bad Request", "message": err.Error()})
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		//falta validar campos vacios
		//turno.Fecha = time.Now()

		t, err := h.s.Agregar(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}


		web.Success(c, 201, t)

	}
}

// POSTPorDNIyMatricula godoc
// @Summary Crear un nuevo turno por DNI y Matrícula
// @Description Agrega un nuevo turno al sistema usando DNI, Matrícula para búsqueda y también cargar fecha y descripción del turno
// @Tags turnos
// @Accept json
// @Produce json
// @Param requestData body domain.RequestData true "Datos de la Solicitud"
// @Success 201 {object} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Router /turnos/dm [post]
func (h *turnoHandler) POSTPorDNIyMatricula() gin.HandlerFunc {
	
	var requestData domain.RequestData

	return func(c *gin.Context) {
		//var turno domain.Turno
		err := c.ShouldBindJSON(&requestData)
		if err != nil {
			//c.JSON(400, gin.H{"status": 400, "code": "Bad Request", "message": err.Error()})
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		//falta validar campos vacios
		//turno.Fecha = time.Now()

		t, err := h.s.AgregarPorDNIyMatricula(requestData.DNI, requestData.Matricula, requestData.Fecha, requestData.Descripcion)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, t)

	}
}


// PUT godoc
// @Summary Actualizar un turno
// @Description Actualiza un turno existente en el sistema
// @Tags turnos
// @Accept json
// @Produce json
// @Param id path int true "ID del Turno"
// @Param turno body domain.Turno true "Objeto Turno Actualizado"
// @Success 200 {object} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos/{id} [put]
func (h *turnoHandler) PUT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Id invalido"))
			return
		}

		_, err = h.s.BuscarPorID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Turno no encontrado"))
			return
		}

		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		t, err := h.s.Actualizar(id, turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, t)

	}
}


// PATCH godoc
// @Summary Actualizar parcialmente un turno
// @Description Actualiza parcialmente un turno existente en el sistema, solo se actualizan los campos proporcionados
// @Tags turnos
// @Accept json
// @Produce json
// @Param id path int true "ID del Turno"
// @Param turno body domain.Turno true "Objeto Turno Parcial"
// @Success 200 {object} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Router /turnos/{id} [patch]
func (h *turnoHandler) PATCH() gin.HandlerFunc {
	type Request struct {
		ID       int              `json:"id"`
		Paciente *domain.Paciente `json:"paciente,omitempty"` //Usamos punteros para que permita usar el omitempty
		Dentista *domain.Dentista `json:"dentista,omitempty"` //Usamos punteros para que permita usar el omitempty
		//PacienteID int `json:"pacienteId" binding:"required"`
		//DentistaID int `json:"dentistaId" binding:"required"`
		Fecha       time.Time `json:"fecha,omitempty"`
		Descripcion string    `json:"descripcion,omitempty"`
	}

	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Id invalido"))
			return
		}

		/*_, err = h.s.BuscarPorID(id)
		  if err != nil {
		      web.Failure(c, 404, errors.New("Turno no encontrado"))
		      return
		  }*/

		err = c.ShouldBindJSON(&r)
		if err != nil {
			//c.JSON(400, gin.H{"status": 400, "code": "Bad Request", "message": err.Error()})
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		// Buscar el turno actual
		update, err := h.s.BuscarPorID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Turno no encontrado"))
			return
		}

		// Actualizar el turno con los datos proporcionados
		if r.Paciente != nil {
			update.Paciente = *r.Paciente
		}
		if r.Dentista != nil {
			update.Dentista = *r.Dentista
		}
		if !r.Fecha.IsZero() {
			update.Fecha = r.Fecha
		}
		if r.Descripcion != "" {
			update.Descripcion = r.Descripcion
		}
		/*
					update := domain.Turno{
						//ID: r.ID,
			            Paciente: *r.Paciente,
			            Dentista: *r.Dentista,
			            //PacienteID: r.PacienteID,
			            //DentistaID: r.DentistaID,
			            Fecha: r.Fecha,
			            Descripcion: r.Descripcion,
					}*/

		t, err := h.s.Actualizar(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, t)
	}
}


// DELETE godoc
// @Summary Eliminar un turno
// @Description Elimina un turno existente del sistema
// @Tags turnos
// @Accept json
// @Produce json
// @Param id path int true "ID del Turno"
// @Success 204
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos/{id} [delete]
func (h *turnoHandler) DELETE() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Id invalido"))
			return
		}

		_, err = h.s.BuscarPorID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Turno no encontrado"))
			return
		}

		err = h.s.Eliminar(id)
		if err != nil {
			web.Failure(c, 404, errors.New("Error al eliminar"))
			return
		}
		web.Success(c, 204, nil)
	}
}


// GETPorDNIPaciente godoc
// @Summary Obtener turnos por DNI del paciente
// @Description Obtiene todos los turnos de un paciente específico por su DNI
// @Tags turnos
// @Accept json
// @Produce json
// @Param dni query string true "DNI del Paciente"
// @Success 200 {array} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos [get]
func (h *turnoHandler) GETPorDNIPaciente() gin.HandlerFunc {
	return func(c *gin.Context) {
		dni := c.Query("dni")
		if dni == "" {
			web.Failure(c, 400, errors.New("DNI invalido"))
			return
		}

		turnos, err := h.s.BuscarPorDNIPaciente(dni)
		if err != nil {
			//c.JSON(400, gin.H{"status": 400, "code": "Bad Request", "message": err.Error()})
			web.Failure(c, 404, errors.New("Turnos no encontrados"))
			return
		}
		web.Success(c, 200, turnos)
	}



}
// POSTPorDNIyMatriculaOpcion2 godoc
// @Summary Crear un nuevo turno por DNI y Matrícula
// @Description Agrega un nuevo turno al sistema usando DNI, Matrícula para búsqueda y también cargar fecha y descripción del turno
// @Tags turnos
// @Accept json
// @Produce json
// @Param requestData body domain.RequestData true "Datos de la Solicitud"
// @Success 201 {object} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Router /turnos/dm [post]
func (h *turnoHandler) POSTPorDNIyMatriculaOpcion2() gin.HandlerFunc {
	
	var requestData domain.RequestData

	return func(c *gin.Context) {
		//var turno domain.Turno
		err := c.ShouldBindJSON(&requestData)
		if err != nil {
			//c.JSON(400, gin.H{"status": 400, "code": "Bad Request", "message": err.Error()})
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		//falta validar campos vacios
		//turno.Fecha = time.Now()

		t, err := h.s.AgregarPorDyM(requestData.DNI, requestData.Matricula, requestData.Fecha, requestData.Descripcion)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, t)

	}
}
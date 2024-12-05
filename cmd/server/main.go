package main

import (
	"log"
	"os"

	"github.com/DiegoParula/SerranaMarset-DiegoParula/cmd/server/handler"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/docs"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/dentista"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/paciente"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/internal/turno"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/pkg/middleware"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
	"github.com/DiegoParula/SerranaMarset-DiegoParula/pkg/store/sqlstore"
)

// @title API de Gestión de Turnos en Clínica Odontológica
// @version 1.0
// @description Esta es una API para gestionar pacientes, dentistas y turnos en una clínica odontológica.
// @host localhost:8080

func main() {
	
	if err := godotenv.Load(".env"); err != nil {
		panic("Error cargando .env: " + err.Error())
	}
	

	db, err := store.NewDatabaseConnection()
	if err != nil {
        log.Fatal(err)
    }
	defer db.Close()

	// paciente 
	storage := sqlstore.NewSqlStorePaciente(db)

	pacienteRepository := paciente.NewRepository(storage)
	servicePaciente := paciente.NuevoServicio(pacienteRepository)
	pacienteHandler := handler.NewPacienteHandler(servicePaciente)


	//dentista
	
	storageDentista := sqlstore.NewSqlDentista(db)
    dentistaRepository := dentista.NewDentistaRespository(storageDentista)
    serviceDentista := dentista.DentistaService(dentistaRepository)
    dentistaHandler := handler.NewDentistaHandler(serviceDentista)


	// turno
	storageTurno := sqlstore.NewSqlStoreTurno(db)
    turnoRepository := turno.NewRepositoryTurno(storageTurno)
    serviceTurno := turno.NewServiceTurno(turnoRepository, pacienteRepository, dentistaRepository)
    turnoHandler := handler.NewTurnoHandler(serviceTurno)

    // router 
    
    //r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())


    turnos := r.Group("/turnos")
    {
        
        turnos.GET("/:id", turnoHandler.BuscarPorID())
        turnos.GET("/", turnoHandler.GETPorDNIPaciente())
        turnos.POST("/", middleware.Authentication(), turnoHandler.POST())
		turnos.POST("/dm", middleware.Authentication(), turnoHandler.POSTPorDNIyMatricula())
		turnos.POST("/dnimat", middleware.Authentication(), turnoHandler.POSTPorDNIyMatriculaOpcion2())
		turnos.PATCH("/:id", middleware.Authentication(), turnoHandler.PATCH())
        turnos.PUT("/:id", middleware.Authentication(), turnoHandler.PUT())
        turnos.DELETE("/:id", middleware.Authentication(), turnoHandler.DELETE())
	}

	

	pacientes := r.Group("/pacientes")
	{
		pacientes.GET("/", pacienteHandler.GetAll())
        pacientes.GET("/:id", pacienteHandler.BuscarPorID())
		pacientes.GET("/dni/:dni", pacienteHandler.BuscarPorDNI())
        pacientes.POST("/", middleware.Authentication(), pacienteHandler.Post())
        pacientes.PATCH("/:id", middleware.Authentication(), pacienteHandler.Patch())
		pacientes.PUT("/:id", middleware.Authentication(), pacienteHandler.Put())
        pacientes.DELETE("/:id", middleware.Authentication(), pacienteHandler.Delete())
	}

	dentistas := r.Group("/dentistas")
	{ 
		dentistas.GET("/:id", dentistaHandler.GetByID())
        dentistas.GET("/", dentistaHandler.GetAll())
		dentistas.GET("/matricula/:matricula", dentistaHandler.GetByMatricula())
        dentistas.POST("/", middleware.Authentication(), dentistaHandler.Post())
        dentistas.PATCH("/:id", middleware.Authentication(), dentistaHandler.Patch())
		dentistas.PUT("/:id",middleware.Authentication(), dentistaHandler.Put())
        dentistas.DELETE("/:id", middleware.Authentication(), dentistaHandler.Delete())
	}
		docs.SwaggerInfo.Host = os.Getenv("HOST")
   		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
	
}
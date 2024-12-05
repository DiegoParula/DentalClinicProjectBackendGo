Sistema de Reserva de Turnos para Clínica Odontológica
Este proyecto implementa una API RESTful en Go (Golang) para administrar la reserva de turnos en una clínica odontológica. Incluye funcionalidades de CRUD para gestionar odontólogos, pacientes y turnos, así como características avanzadas de seguridad y documentación de la API.

Objetivo
Desarrollar un sistema que permita administrar eficientemente la información de una clínica odontológica, incluyendo odontólogos, pacientes y sus turnos, utilizando las mejores prácticas de desarrollo.

Características principales
1. Administración de odontólogos (Dentista)
Se permite gestionar la información de los odontólogos mediante las siguientes operaciones:

POST: Agregar un nuevo dentista.
GET: Traer un dentista por su ID.
PUT: Actualizar la información completa de un dentista.
PATCH: Actualizar campos específicos de un dentista.
DELETE: Eliminar un dentista.
2. Administración de pacientes (Paciente)
Se pueden gestionar los datos de los pacientes con estas operaciones:

POST: Agregar un nuevo paciente.
GET: Traer un paciente por su ID.
PUT: Actualizar la información completa de un paciente.
PATCH: Actualizar campos específicos de un paciente.
DELETE: Eliminar un paciente.
3. Gestión de turnos (Turno)
El sistema permite registrar turnos asignados a pacientes con odontólogos, incluyendo:

POST: Agregar un turno.
GET: Traer un turno por su ID.
PUT: Actualizar la información completa de un turno.
PATCH: Actualizar campos específicos de un turno.
DELETE: Eliminar un turno.
POST (por DNI y matrícula): Asignar un turno usando el DNI del paciente y la matrícula del dentista.
GET (por DNI): Obtener turnos de un paciente específico usando el DNI como query parameter, incluyendo detalles como fecha, hora, descripción, paciente y dentista.
4. Seguridad mediante middleware
Las operaciones sensibles como POST, PUT, PATCH y DELETE están protegidas mediante autenticación implementada con middleware.
5. Documentación de la API
La API está completamente documentada con Swagger, permitiendo una fácil exploración y pruebas.
Requerimientos técnicos
Arquitectura del proyecto
El diseño sigue un enfoque orientado a paquetes con las siguientes capas:

Dominio de entidades de negocio: Representa las entidades principales (Dentista, Paciente, Turno).
Capa de acceso a datos (Repository): Implementación de repositorios para interactuar con la base de datos.
Capa de servicio (Service): Lógica de negocio y reglas de la aplicación.
Capa de handler: Gestión de solicitudes HTTP y comunicación con los servicios.
Base de datos: Compatible con bases de datos relacionales (como MySQL o H2) y no relacionales (como MongoDB).
Tecnologías utilizadas
Lenguaje: Go (Golang).
Base de datos: Relacional o no relacional (ejemplo: MySQL, H2 o MongoDB).
Swagger: Documentación interactiva de la API.
Middleware personalizado: Para autenticación y seguridad.


Configuración del Proyecto
1. Crear y Correr la Base de Datos
El archivo build_database.sql contiene las instrucciones SQL necesarias para crear las tablas y configurar la base de datos en MySQL. 
2. Configurar Credenciales de MySQL
El archivo database.go contiene la lógica de conexión a la base de datos. Debes actualizar este archivo con las credenciales correctas de tu instancia local de MySQL.
Abre database.go y localiza la función que se encarga de la conexión a la base de datos.
Modifica las líneas donde se definen el usuario y la contraseña:
dsn := "[usuario]:[contraseña]@tcp(localhost:3306)/[nombre_base_de_datos]"
Reemplaza [usuario], [contraseña] y [nombre_base_de_datos] con la información correspondiente de tu entorno local.
3. Importar Endpoints en Postman
Para facilitar la prueba de los endpoints, hemos incluido una colección de Postman en el archivo postman_collection.json. Sigue estos pasos para importar la colección:
Abre Postman y selecciona la opción Importar.
Selecciona el archivo postman_collection.json.
Ahora tendrás disponibles todos los endpoints del proyecto para su prueba.
4. Acceder a la Documentación de Swagger
La documentación de los endpoints está disponible mediante Swagger. Puedes acceder a ella navegando a la siguiente URL una vez que el servidor esté en funcionamiento:
http://localhost:8080/docs/index.html

Ejecución del Proyecto
Asegúrate de haber configurado correctamente la base de datos y las credenciales en database.go.
Ejecuta la API con el siguiente comando:
go run main.go
Accede a la documentación Swagger en http://localhost:8080/docs/index.html

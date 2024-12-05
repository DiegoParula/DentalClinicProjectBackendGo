-- Crear la base de datos
CREATE DATABASE IF NOT EXISTS clinica_odontologica;
USE clinica_odontologica;

-- Crear tabla para Dentistas
CREATE TABLE IF NOT EXISTS dentistas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    matricula VARCHAR(50) NOT NULL
);

-- Crear tabla para Pacientes
CREATE TABLE IF NOT EXISTS pacientes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    direccion VARCHAR(255),
    dni VARCHAR(20) NOT NULL,
    fecha_alta DATE NOT NULL
);

-- Crear tabla para Turnos
CREATE TABLE IF NOT EXISTS turnos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fecha_hora DATETIME NOT NULL,
    descripcion TEXT,
    paciente_id INT,
    dentista_id INT,
    FOREIGN KEY (paciente_id) REFERENCES pacientes(id) ON DELETE CASCADE,
    FOREIGN KEY (dentista_id) REFERENCES dentistas(id) ON DELETE CASCADE
);

-- Insertar datos de prueba en Dentistas
INSERT INTO dentistas (nombre, apellido, matricula) VALUES 
('Juan', 'Pérez', 'D12345'),
('María', 'Gómez', 'D54321'),
('Carlos', 'Sánchez', 'D67890');

-- Insertar datos de prueba en Pacientes
INSERT INTO pacientes (nombre, apellido, direccion, dni, fecha_alta) VALUES 
('Ana', 'Martínez', 'Calle Falsa 123', '12345678', '2023-01-15'),
('Luis', 'García', 'Avenida Siempre Viva 742', '87654321', '2023-03-10'),
('Sofía', 'Rodríguez', 'Calle Luna 456', '56789012', '2023-06-05');

-- Insertar datos de prueba en Turnos
INSERT INTO turnos (fecha_hora, descripcion, paciente_id, dentista_id) VALUES 
('2024-09-10 10:00:00', 'Consulta inicial', 1, 1),
('2024-09-11 11:30:00', 'Control de caries', 2, 2),
('2024-09-12 09:00:00', 'Limpieza dental', 3, 3);
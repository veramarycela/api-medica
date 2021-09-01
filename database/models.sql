CREATE TABLE IF NOT EXISTS paciente (
    id_p serial NOT NULL,
    nombre VARCHAR,
    apellido VARCHAR,
    fecha_nac DATE,
    direccion VARCHAR, 
    telefono VARCHAR,    
    CONSTRAINT pk_paciente PRIMARY KEY(id_p)
);

CREATE TABLE IF NOT EXISTS especialidad (
    id_e VARCHAR NOT NULL,    
    nombre VARCHAR,   
    CONSTRAINT pk_especialidad PRIMARY KEY(id_e)    
);

CREATE TABLE IF NOT EXISTS medico (
    id_m int NOT NULL,
    id_especialidad VARCHAR NOT NULL,
    nombre VARCHAR,
    apellido VARCHAR,
    direccion VARCHAR, 
    telefono VARCHAR, 
    CONSTRAINT pk_medico PRIMARY KEY(id_m),
    CONSTRAINT fk_especialidad FOREIGN KEY(id_especialidad) REFERENCES especialidad(id_e)
);

CREATE TABLE IF NOT EXISTS historia (
    id_h VARCHAR NOT NULL,
    id_medico int NOT NULL,
    id_paciente VARCHAR NOT NULL,
    fecha DATE,
    motivo TEXT,
    diagnostico TEXT, 
    receta TEXT, 
    CONSTRAINT pk_historia PRIMARY KEY(id_h),
    CONSTRAINT fk_medico FOREIGN KEY(id_medico) REFERENCES medico(id_m),
    CONSTRAINT fk_paciente FOREIGN KEY(id_paciente) REFERENCES paciente(id_p)
);


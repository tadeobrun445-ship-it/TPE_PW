CREATE TABLE avistamientos (
    id SERIAL PRIMARY KEY,

    fecha_hora TIMESTAMP WITH TIME ZONE NOT NULL,

    ubicacion TEXT NOT NULL,

    descripcion TEXT NOT NULL,

    imagen TEXT,

    hubo_destrozos BOOLEAN NOT NULL DEFAULT FALSE,

    detalle_destrozos TEXT,

    nombre_reportante VARCHAR(255),

    email_reportante VARCHAR(255),

    created_at TIMESTAMP WITH TIME ZONE
        NOT NULL DEFAULT CURRENT_TIMESTAMP
);

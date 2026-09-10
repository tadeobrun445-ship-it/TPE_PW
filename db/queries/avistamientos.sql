-- name: CreateAvistamiento :one
INSERT INTO avistamientos (
    fecha_hora,
    ubicacion,
    descripcion,
    imagen,
    hubo_destrozos,
    detalle_destrozos,
    nombre_reportante,
    email_reportante
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;


-- name: GetAvistamiento :one
SELECT *
FROM avistamientos
WHERE id = $1;


-- name: ListAvistamientos :many
SELECT *
FROM avistamientos
ORDER BY fecha_hora DESC;


-- name: UpdateAvistamiento :exec
UPDATE avistamientos
SET
    fecha_hora = $2,
    ubicacion = $3,
    descripcion = $4,
    imagen = $5,
    hubo_destrozos = $6,
    detalle_destrozos = $7,
    nombre_reportante = $8,
    email_reportante = $9
WHERE id = $1;


-- name: DeleteAvistamiento :exec
DELETE FROM avistamientos
WHERE id = $1;

# TPE Programación Web

Aplicación web para el registro y consulta de avistamientos de jabalíes.

En esta entrega se incorpora la capa de persistencia utilizando PostgreSQL y `sqlc`, junto con tests automatizados de las operaciones CRUD.

## Requisitos

Para ejecutar los tests es necesario disponer de:

* Go
* Docker
* Docker Compose
* Make

No es necesario crear manualmente la base de datos ni instalar `sqlc` globalmente. El proceso de testing se encarga de preparar automáticamente el entorno necesario.

## Persistencia

La entidad principal persistida es `Avistamiento`.

La definición del esquema se encuentra en:

```text
db/schema/schema.sql
```

Las consultas SQL utilizadas por `sqlc` se encuentran en:

```text
db/queries/avistamientos.sql
```

La configuración de `sqlc` se encuentra en:

```text
sqlc.yaml
```

El código Go generado se almacena en:

```text
db/sqlc/
```

La documentación detallada sobre las decisiones de persistencia se encuentra en:

```text
docs/persistencia.md
```

## Ejecución de los tests

Desde la raíz del proyecto ejecutar:

```bash
make test
```

Este comando realiza automáticamente las siguientes tareas:

1. Genera el código Go utilizando `sqlc`.
2. Compila el proyecto.
3. Elimina contenedores y volúmenes de pruebas anteriores.
4. Levanta una instancia PostgreSQL mediante Docker Compose.
5. Espera hasta que PostgreSQL esté disponible.
6. Inicializa el esquema de la base de datos.
7. Ejecuta los tests con el paquete `testing` de Go.
8. Elimina los contenedores y volúmenes utilizados al finalizar.

La limpieza del entorno también se realiza si ocurre un error durante la ejecución.

## Tests de persistencia

Los tests realizan un ciclo CRUD completo sobre la entidad `Avistamiento`:

* creación;
* consulta por identificador;
* actualización;
* listado;
* eliminación.

La base de datos utilizada para los tests es temporal y se ejecuta dentro de Docker, evitando depender de una base PostgreSQL previamente configurada en la computadora.

## Estructura principal

```text
.
├── db/
│   ├── queries/
│   │   └── avistamientos.sql
│   ├── schema/
│   │   └── schema.sql
│   └── sqlc/
│       ├── avistamientos.sql.go
│       ├── avistamientos_test.go
│       ├── db.go
│       └── models.go
├── docs/
│   └── persistencia.md
├── compose.yaml
├── Makefile
├── sqlc.yaml
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Generación manual con sqlc

Normalmente no es necesario ejecutar este paso manualmente porque `make test` ya lo realiza.

En caso de querer regenerar únicamente el código de acceso a datos se puede ejecutar:

```bash
go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.31.1 generate
```

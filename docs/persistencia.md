# Persistencia de datos

## Descripción general

El proyecto implementa una plataforma para registrar avistamientos de jabalíes.

Para la persistencia de los datos se utiliza PostgreSQL como sistema gestor de base de datos y `sqlc` como herramienta para generar código Go tipado a partir de consultas SQL.

La entidad principal persistida es `Avistamiento`.

## Modelo de datos

La tabla `avistamientos` contiene los siguientes atributos:

* `id`: identificador único del avistamiento. Es generado automáticamente por PostgreSQL.
* `fecha_hora`: fecha y hora en la que ocurrió el avistamiento.
* `ubicacion`: lugar donde ocurrió el avistamiento.
* `descripcion`: descripción del avistamiento.
* `imagen`: referencia opcional a una imagen asociada.
* `hubo_destrozos`: indica si el avistamiento estuvo asociado a destrozos.
* `detalle_destrozos`: descripción opcional de los destrozos.
* `nombre_reportante`: nombre opcional de la persona que realizó el reporte.
* `email_reportante`: correo electrónico opcional de la persona que realizó el reporte.
* `created_at`: fecha y hora en la que el registro fue incorporado al sistema.

## Diferencia entre `fecha_hora` y `created_at`

Estos campos representan dos momentos diferentes.

`fecha_hora` pertenece al dominio del avistamiento e indica cuándo ocurrió el hecho.

`created_at` es un metadato del sistema e indica cuándo ese avistamiento fue registrado en la aplicación.

Esto permite distinguir, por ejemplo, un avistamiento ocurrido el día anterior de un reporte que fue incorporado recién hoy. También proporciona trazabilidad y permite ordenar o filtrar registros según el momento en que fueron cargados.

## Campos opcionales

Los campos `imagen`, `detalle_destrozos`, `nombre_reportante` y `email_reportante` pueden almacenar `NULL`.

Esto se debe a que no siempre existe una imagen o un detalle de destrozos y el usuario puede realizar un reporte sin proporcionar información personal.

En el código generado por `sqlc`, estos valores se representan mediante tipos como `sql.NullString`.

## Acceso a datos con sqlc

Las consultas SQL se encuentran en:

`db/queries/avistamientos.sql`

Se implementan las operaciones CRUD principales:

* creación de un avistamiento;
* búsqueda de un avistamiento por identificador;
* listado de avistamientos;
* actualización de un avistamiento;
* eliminación de un avistamiento.

El esquema de la base de datos se encuentra en:

`db/schema/schema.sql`

La configuración de `sqlc` se encuentra en:

`sqlc.yaml`

A partir de estos archivos se ejecuta:

`sqlc generate`

Esto genera automáticamente código Go tipado dentro de `db/sqlc`.

Los archivos generados no deben modificarse manualmente, ya que pueden volver a generarse a partir del esquema y las consultas SQL.

## Base de datos de pruebas

Los tests utilizan una instancia PostgreSQL ejecutada mediante Docker Compose.

La base utilizada para las pruebas se denomina:

`tpe_pw_test`

Al inicializar el contenedor, el archivo `db/schema/schema.sql` se monta en el directorio de inicialización de PostgreSQL, por lo que la estructura de la base se crea automáticamente.

Los tests realizan un ciclo CRUD completo sobre la entidad `Avistamiento`, comprobando las operaciones de creación, consulta, actualización, listado y eliminación.

## Ejecución automatizada

La ejecución de los tests se automatiza mediante el `Makefile`.

El comando:

`make test`

realiza las tareas necesarias para disponer de un entorno reproducible:

* genera nuevamente el código de `sqlc`;
* compila el proyecto;
* elimina contenedores y volúmenes anteriores;
* inicia PostgreSQL mediante Docker Compose;
* espera hasta que la base de datos esté disponible;
* ejecuta los tests escritos con el paquete `testing` de Go;
* elimina finalmente los contenedores y volúmenes utilizados.

La limpieza final se realiza incluso si ocurre un error durante la ejecución, evitando dejar recursos Docker asociados a los tests.

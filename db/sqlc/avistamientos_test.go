package db

import (
	"database/sql"
	"errors"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestAvistamientosCRUD(t *testing.T) {

	// La URL podrá venir desde el Makefile.
	// Si no está definida, usamos esta conexión local por defecto.
	databaseURL := os.Getenv("TEST_DATABASE_URL")

	if databaseURL == "" {
		databaseURL =
			"postgres://postgres:postgres@localhost:5433/tpe_pw_test?sslmode=disable"
	}

	database, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatalf("Error al abrir la base de datos: %v", err)
	}
	defer database.Close()

	if err := database.Ping(); err != nil {
		t.Fatalf("Error al conectarse a la base de pruebas: %v", err)
	}

	// Dejamos la tabla limpia antes de ejecutar el flujo CRUD.
	_, err = database.Exec("DELETE FROM avistamientos")
	if err != nil {
		t.Fatalf("Error al limpiar avistamientos: %v", err)
	}

	queries := New(database)

	var creado Avistamiento

	t.Run("CreateAvistamiento", func(t *testing.T) {

		creado, err = queries.CreateAvistamiento(
			t.Context(),
			CreateAvistamientoParams{
				FechaHora:        time.Now(),
				Ubicacion:        "Tandil",
				Descripcion:      "Avistamiento de prueba",
				Imagen:           sql.NullString{},
				HuboDestrozos:    false,
				DetalleDestrozos: sql.NullString{},
				NombreReportante: sql.NullString{
					String: "Tomas",
					Valid:  true,
				},
				EmailReportante: sql.NullString{
					String: "tomas@test.com",
					Valid:  true,
				},
			},
		)

		if err != nil {
			t.Fatalf("CreateAvistamiento dio error: %v", err)
		}

		if creado.ID == 0 {
			t.Errorf("Se esperaba que PostgreSQL asignara un ID")
		}

		if creado.Ubicacion != "Tandil" {
			t.Errorf(
				"Ubicacion incorrecta: esperada Tandil, obtenida %s",
				creado.Ubicacion,
			)
		}
	})

	t.Run("GetAvistamiento", func(t *testing.T) {

		encontrado, err := queries.GetAvistamiento(
			t.Context(),
			creado.ID,
		)

		if err != nil {
			t.Fatalf("GetAvistamiento dio error: %v", err)
		}

		if encontrado.ID != creado.ID {
			t.Errorf(
				"ID incorrecto: esperado %d, obtenido %d",
				creado.ID,
				encontrado.ID,
			)
		}

		if encontrado.Descripcion != creado.Descripcion {
			t.Errorf(
				"Descripcion incorrecta: esperada %s, obtenida %s",
				creado.Descripcion,
				encontrado.Descripcion,
			)
		}
	})

	t.Run("UpdateAvistamiento", func(t *testing.T) {

		err := queries.UpdateAvistamiento(
			t.Context(),
			UpdateAvistamientoParams{
				ID:            creado.ID,
				FechaHora:     creado.FechaHora,
				Ubicacion:     "Tandil - Cerro El Centinela",
				Descripcion:   "Avistamiento actualizado",
				Imagen:        creado.Imagen,
				HuboDestrozos: true,
				DetalleDestrozos: sql.NullString{
					String: "Daños en un alambrado",
					Valid:  true,
				},
				NombreReportante: creado.NombreReportante,
				EmailReportante:  creado.EmailReportante,
			},
		)

		if err != nil {
			t.Fatalf("UpdateAvistamiento dio error: %v", err)
		}

		actualizado, err := queries.GetAvistamiento(
			t.Context(),
			creado.ID,
		)

		if err != nil {
			t.Fatalf(
				"Error al recuperar el avistamiento actualizado: %v",
				err,
			)
		}

		if actualizado.Ubicacion != "Tandil - Cerro El Centinela" {
			t.Errorf(
				"La ubicacion no se actualizó correctamente: %s",
				actualizado.Ubicacion,
			)
		}

		if !actualizado.HuboDestrozos {
			t.Errorf("Se esperaba HuboDestrozos = true")
		}

		creado = actualizado
	})

	t.Run("ListAvistamientos", func(t *testing.T) {

		avistamientos, err := queries.ListAvistamientos(
			t.Context(),
		)

		if err != nil {
			t.Fatalf("ListAvistamientos dio error: %v", err)
		}

		encontrado := false

		for _, avistamiento := range avistamientos {
			if avistamiento.ID == creado.ID {
				encontrado = true
				break
			}
		}

		if !encontrado {
			t.Errorf(
				"El avistamiento con ID %d no apareció en la lista",
				creado.ID,
			)
		}
	})

	t.Run("DeleteAvistamiento", func(t *testing.T) {

		err := queries.DeleteAvistamiento(
			t.Context(),
			creado.ID,
		)

		if err != nil {
			t.Fatalf("DeleteAvistamiento dio error: %v", err)
		}

		_, err = queries.GetAvistamiento(
			t.Context(),
			creado.ID,
		)

		if !errors.Is(err, sql.ErrNoRows) {
			t.Errorf(
				"Se esperaba sql.ErrNoRows luego del Delete, pero se obtuvo: %v",
				err,
			)
		}
	})
}

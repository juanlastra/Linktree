// funciones y structura

package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ConfiSitio struct {
	Name    string `yaml:"name"`
	Bio     string `yaml:"bio"`
	Picture string `yaml:"picture"`
	Meta    Meta   `yaml:"meta"`
	Links   []Link `yaml:"links"`
	Theme   string `yaml:"theme"`
}

type Meta struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Lang        string `yaml:"lang"`
	Author      string `yaml:"author"`
	SiteUrl     string `yaml:"siteUrl"`
}

type Link struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

// cargar configuración

func Cargarconfig(ruta string) (*ConfiSitio, error) {
	data, err := os.ReadFile(ruta)

	if err != nil {
		return nil, err
	}

	var config ConfiSitio

	err = yaml.Unmarshal(data, &config)

	if err != nil {

		return nil, fmt.Errorf("error al deserializar los datos: %v", err)
	}

	return &config, nil
}

func GenerarHTML(config *ConfiSitio) error {

	// definir ruta a la plantilla
	pathplantilla := "./Recursos/index.html"

	// cargar plantilla html

	tmpl, err := template.ParseFiles(pathplantilla)

	if err != nil {
		log.Fatalf("Error al leer la plantilla HTML %v", err)
		return err
	}

	// crea el nuevo archivo Index
	Salidaarchivo, err := os.Create("./Resultados/salida.html")

	if err != nil {
		return err
	}

	defer Salidaarchivo.Close()

	// definir datos a pasar a la plantilla

	data := struct {
		Config *ConfiSitio
	}{
		Config: config,
	}

	jsonBytes, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		log.Fatalf("Error al convertir a Json: %v", err)
	}

	fmt.Println(string(jsonBytes))

	return tmpl.Execute(Salidaarchivo, data)
}

func Copiararchivo(files []string, pathdestino string) error {

	for _, file := range files {
		data, err := os.ReadFile(file)

		if err != nil {
			return fmt.Errorf("no se leer el archivo %s: error %w",
				file,
				err)
		}

		rutadestino := filepath.Join(pathdestino, filepath.Base(file))

		archivosalida, err := os.Create(rutadestino)

		if err != nil {
			return fmt.Errorf("fallo al crear el archivo %s: %w", rutadestino, err)
		}

		defer archivosalida.Close()

		if _, err := io.Copy(archivosalida, bytes.NewReader(data)); err != nil {
			return fmt.Errorf("fallo al copiar datos en %s: %w", rutadestino,
				err)
		}

		archivosalida.Close()

	}

	return nil
}

func CopiaRecurso() error {

	// crear directorios

	dirs := []string{"./Recursos/css", "./Recursos/js", "./Recursos/icons"}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("fallo al crear %s, %w", dir, err)
		}
	}

	Rutas_destino := []struct {
		inicio  string
		destino string
	}{
		{inicio: "./Recursos/*.css", destino: "./Recursos/css"},
		{inicio: "./Recursos/*.js", destino: "./Recursos/js"},
		{inicio: "./Recursos/*.ico", destino: "./Recursos/icons"},
	}

	for _, rutas := range Rutas_destino {
		files, err := filepath.Glob(fmt.Sprintf(rutas.inicio))

		if err != nil {
			return fmt.Errorf("no se encuentra un archivo %s: %w",
				rutas.destino, err)
		}

		if err := Copiararchivo(files, rutas.destino); err != nil {
			return fmt.Errorf("error al copiar archivo en %s, error: %w",
				rutas.destino, err)
		}
	}

	return nil
}

//func minificar()

package main

import (
	config "Go_juan_linkstree/Scripts"
	"log"
)

// funciones a utilizar

func main() {

	// cargar configuración

	infoyml, err := config.Cargarconfig("./Recursos/config.yml")

	if err != nil {
		log.Fatalf("Error cargando la configuración: %v", err)
	}

	//jsonBytes, err := json.MarshalIndent(infoyml, "", " ")

	//if err != nil {
	//	log.Fatalf("Error al convertir a Json: %v", err)
	//}

	//fmt.Println(string(jsonBytes))

	// generar plantilla html

	err = config.GenerarHTML(infoyml)

	if err != nil {
		log.Fatalf("error al generar la plantilla html: %v", err)
	}

	err = config.CopiaRecurso()

	if err != nil {
		log.Fatalf("no se pudieron copiar los recursos: %v", err)
	}
}

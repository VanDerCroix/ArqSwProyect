package main

import (
	datos "./db"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	//"strconv"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/ejemplo1", Ejemplo1_handler},
		&rest.Route{"GET", "/ejemplo2", Ejemplo2_handler},
		// ---- ARQUITECTURABD ---
		/*&rest.Route{"GET", "/facultades", Facultades_handler},
		&rest.Route{"GET", "/facultades/:id", Facultad_handler},
		&rest.Route{"GET", "/escuelas", Escuelas_handler},
		&rest.Route{"GET", "/escuelasxfacu/:id", EscuelasxFacu_handler},
		&rest.Route{"GET", "/administradores", Administradores_handler},
		&rest.Route{"GET", "/noticias", Noticias_handler},
		&rest.Route{"GET", "/ubicacion", Ubicacion_handler},
		&rest.Route{"GET", "/areasuniversidad", AreasUniversidad_handler},*/

		// ---- CALIDAD SOS ---
		&rest.Route{"GET", "/usuario", Usuario_handler},
		&rest.Route{"GET", "/usuario/:dni/contacto", Contacto_handler},
		&rest.Route{"GET", "/centrosatencion", Centros_Atencion_handler},
		&rest.Route{"POST", "/postusuario", PostUsuario_handler},
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9988", api.MakeHandler()))
}

// ---- EJEMPLOS ---
func Ejemplo1_handler(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{
		"nombre": "barry allen",
		"correo": "aelosmit@gmail.com",
	})
}
func Ejemplo2_handler(w rest.ResponseWriter, r *rest.Request) {
/*	idubi, err := strconv.Atoi(r.PathParam("id"))

	if err != nil {
		w.WriteJson(map[string]string{"Error": err.Error()})
	} else {
		ubicacion := datos.ConsultaUbicacion(idubi)
		w.WriteJson(ubicacion)
	}*/
	// datos.Query()
	w.WriteJson(map[string]string{
		"nombre": "2",
		"correo": "asdit@gmail.com",
	})
}

// ---- ARQUITECTURABD ---
/*func Facultades_handler(w rest.ResponseWriter, r *rest.Request) {
	facultad := datos.ConsultaFacultades()
	w.WriteJson(facultad)
}

func Facultad_handler(w rest.ResponseWriter, r *rest.Request) {
	facuid := r.PathParam("id")
	facultad := datos.ConsultaFacultadDetalles(facuid)
	w.WriteJson(facultad)
}*/

// ---- CALIDAD SOS ---
func Usuario_handler(w rest.ResponseWriter, r *rest.Request) {
	usuarios := datos.ConsultaUsuarios()
	w.WriteJson(usuarios)
}

func Contacto_handler(w rest.ResponseWriter, r *rest.Request) {
	userdni := r.PathParam("dni")
	dni, err := strconv.Atoi(userdni)
	contactos := datos.ConsultaContactos(dni)
	w.WriteJson(contactos)
}

func Centros_Atencion_handler(w rest.ResponseWriter, r *rest.Request){
centros_atencion := datos.ConsultaCentrosAtencion()
	w.WriteJson(centros_atencion)	
}

func PostUsuario_handler(w rest.ResponseWriter, r *rest.Request) {
	usuario := new(datos.Usuario)
	err := r.DecodeJsonPayload(&usuario)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datos.InsertarUsuario(*usuario)
	w.WriteJson(&usuario)
}
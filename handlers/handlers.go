package handlers

import "net/http"

//PathInicio Ruta raíz
const PathInicio string = "/"

//PathJSFiles Ruta a la carpeta de scripts de javascript
const PathJSFiles string = "/js/"

//PathEnvioPeticion Ruta de envío de peticiones
const PathEnvioPeticion string = "/envio"

//PathListadoPeticiones Ruta de obtención de las peticiones de hoy
const PathListadoPeticiones string = "/lista"

//PathEnvioIdioma Ruta de envío de idiomas
const PathEnvioIdioma string = "/envioIdioma"

//PathListadoIdiomas Ruta de obtención de las idiomas de hoy
const PathListadoIdiomas string = "/listadoIdiomas"

//PathEnvioTraduccion Ruta de envío de idiomas
const PathEnvioTraduccion string = "/envioTraduccion"

//ManejadorHTTP encapsula como tipo la función de manejo de peticiones HTTP, para que sea posible almacenar sus referencias en un diccionario
type ManejadorHTTP = func(w http.ResponseWriter, r *http.Request)

//Lista es el diccionario general de las peticiones que son manejadas por nuestro servidor
var Manejadores map[string]ManejadorHTTP

func init() {
	Manejadores = make(map[string]ManejadorHTTP)
	Manejadores[PathInicio] = IndexFile
	Manejadores[PathJSFiles] = JsFile
	Manejadores[PathEnvioPeticion] = Insert
	Manejadores[PathListadoPeticiones] = List
	Manejadores[PathEnvioIdioma] = InsertIdioma
	Manejadores[PathEnvioTraduccion] = Insert
	Manejadores[PathListadoIdiomas] = ListIdioma
}

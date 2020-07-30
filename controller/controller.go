package controller

import (
	"net/http"

	"github.com/RomuloDurante/baseNewServer/controller/delete"
	"github.com/RomuloDurante/baseNewServer/controller/get"
	"github.com/RomuloDurante/baseNewServer/controller/post"
	"github.com/RomuloDurante/baseNewServer/controller/put"
)

//Controller ...
type Controller struct {
	Method string
	Path   string
	Query  map[string][]string
}

func startController(r *http.Request) *Controller {
	return &Controller{
		Method: r.Method,
		Path:   r.URL.Path,
		Query:  r.URL.Query(),
	}

}

// HandleController func ...
func HandleController(w http.ResponseWriter, r *http.Request) {
	//TODO: make channel here

	controller := startController(r)

	switch controller.Method {
	case "GET":
		//err = Get(w, r, c)
		get.StartService()
		break
	case "POST":
		//err = Post(w, r, c)
		post.StartService()
		break
	case "PUT":
		//err = Put(w, r, c)
		put.StartService()
		break
	case "DELETE":
		//err = Delete(w, r, c)
		delete.StartService()
		break
	}

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

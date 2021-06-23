package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"rest_api/server"
	"time"
)

type Server server.Server

func (ts *Server) provinceHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/province/" {
		// Request is plain "/province/", without trailing ID.
		if req.Method == http.MethodPost {
			ts.createProvinceHandler(w, req)
			/*} else if req.Method == http.MethodGet {
				ts.getAllProvincesHandler(w, req)
			} else if req.Method == http.MethodDelete {
				ts.deleteAllProvincesHandler(w, req)*/
		} else {
			http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /province/, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	} else { /*
			// Request has an ID, as in "/province/<id>".
			path := strings.Trim(req.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /province/<id> in province handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if req.Method == http.MethodDelete {
				ts.deleteProvinceHandler(w, req, int(id))
			} else if req.Method == http.MethodGet {
				ts.getProvinceHandler(w, req, int(id))
			} else { */
		http.Error(w, fmt.Sprintf("expect method GET or DELETE at /task/<id>, got %v", req.Method), http.StatusMethodNotAllowed)
		return
		//}
	}
}

func (ts *Server) createProvinceHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling province create at %s\n", req.URL.Path)

	// Types used internally in this handler to (de-)serialize the request and
	// response from/to JSON.
	type RequestTask struct {
		Text string    `json:"text"`
		Tags []string  `json:"tags"`
		Due  time.Time `json:"due"`
	}

	type ResponseId struct {
		Id int `json:"id"`
	}

	// Enforce a JSON Content-Type.
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var rt RequestTask
	if err := dec.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := ts.store.CreateTask(rt.Text, rt.Tags, rt.Due)
	js, err := json.Marshal(ResponseId{Id: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

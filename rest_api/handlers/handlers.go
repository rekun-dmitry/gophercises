package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"rest_api/provinces"
	"strconv"
	"strings"
)

type Server struct {
	store *provinces.ProvinceStore
}

func NewServer() *Server {
	store := provinces.New()
	return &Server{store: store}
}

func (ps *Server) EconomicProvinceHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/province/economic/" {
		// Request is plain "/province/", without trailing ID.
		if req.Method == http.MethodPost {
			ps.createProvinceHandler(w, req)
		} else if req.Method == http.MethodGet {
			ps.getAllProvincesHandler(w, req)
		} else if req.Method == http.MethodDelete {
			ps.deleteAllProvincesHandler(w, req)
		} else {
			http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /province/economic/, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	} else {
		// Request has an ID, as in "/province/<id>".
		path := strings.Trim(req.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 3 {
			http.Error(w, "expect /province/economic/<id> in province handler", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(pathParts[2])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if req.Method == http.MethodDelete {
			ps.deleteProvinceHandler(w, req, int(id))
		} else if req.Method == http.MethodGet {
			ps.getProvinceHandler(w, req, int(id))
		} else {
			http.Error(w, fmt.Sprintf("expect method GET or DELETE at /province/<id>, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}

func (ts *Server) createProvinceHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling province create at %s\n", req.URL.Path)

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
	var rt provinces.Province
	if err := dec.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := ts.store.CreateProvince(rt)
	js, err := json.Marshal(ResponseId{Id: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ts *Server) getAllProvincesHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get all provinces at %s\n", req.URL.Path)

	allProvinces := ts.store.GetAllProvinces()
	js, err := json.Marshal(allProvinces)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ps *Server) deleteAllProvincesHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling delete all provinces at %s\n", req.URL.Path)
	ps.store.DeleteAllProvinces()
}

func (ps *Server) getProvinceHandler(w http.ResponseWriter, req *http.Request, id int) {
	log.Printf("handling get province at %s\n", req.URL.Path)

	province, err := ps.store.GetProvince(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(province)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ps *Server) deleteProvinceHandler(w http.ResponseWriter, req *http.Request, id int) {
	log.Printf("handling delete province at %s\n", req.URL.Path)

	err := ps.store.DeleteProvince(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

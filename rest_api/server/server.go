package server

import "rest_api/provinces"

type Server struct {
	store *provinces.ProvinceStore
}

func NewServer() *Server {
	store := provinces.New()
	return &Server{store: store}
}

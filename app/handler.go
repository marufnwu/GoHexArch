package app

import (
	"encoding/json"
	"net/http"

	"github.com/marufnwu/HexaArchProject/services"
)

type MyHandler struct{
	service services.MyServiceInterface
}

func (h MyHandler) getAllUser(w http.ResponseWriter, r *http.Request){
	users :=h.service.GetUsers()

	json.NewEncoder(w).Encode(users)

}

func NewHandler(service services.MyService) MyHandler{
	return MyHandler{service: service}
}
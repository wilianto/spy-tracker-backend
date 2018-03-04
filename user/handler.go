package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//HTTPHandler is a delivery layer for users via HTTP
type HTTPHandler interface {
	PostUser(w http.ResponseWriter, r *http.Request)
}

type userHTTPHandler struct {
	userService Service
}

func (h *userHTTPHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var (
		req  *postUserRequest
		resp *postUserResponse
	)
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		resp = &postUserResponse{Error: err.Error()}
		writeResponse(w, http.StatusBadRequest, resp)
		return
	}

	usr := &User{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
	}

	userID, err := h.userService.Register(usr)
	if err != nil {
		resp = &postUserResponse{Error: err.Error()}
		writeResponse(w, http.StatusUnprocessableEntity, resp)
		return
	}

	resp = &postUserResponse{ID: userID}
	writeResponse(w, http.StatusCreated, resp)
	return
}

func writeResponse(w http.ResponseWriter, code int, body *postUserResponse) {
	respBody, _ := json.Marshal(body)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

type postUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type postUserResponse struct {
	ID    int64  `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

//NewHTTPHandler initialize user handler with specific service
func NewHTTPHandler(r *mux.Router, userService Service) {
	handler := userHTTPHandler{userService}
	r.HandleFunc("/users", handler.PostUser).Methods("POST")
}

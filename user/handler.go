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
		resp = &postUserResponse{Errors: errorToStringMapper([]error{err})}
		writeResponse(w, http.StatusBadRequest, resp)
		return
	}

	usr := &User{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
	}

	userID, errs := h.userService.Register(usr)
	if len(errs) != 0 {
		resp = &postUserResponse{Errors: errorToStringMapper(errs)}
		writeResponse(w, http.StatusUnprocessableEntity, resp)
		return
	}

	resp = &postUserResponse{ID: userID}
	writeResponse(w, http.StatusCreated, resp)
	return
}

type postUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type postUserResponse struct {
	ID     int64    `json:"id,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

//NewHTTPHandler initialize user handler with specific service
func NewHTTPHandler(r *mux.Router, userService Service) {
	handler := userHTTPHandler{userService}
	r.HandleFunc("/users", handler.PostUser).Methods("POST")
}

func writeResponse(w http.ResponseWriter, code int, body *postUserResponse) {
	respBody, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(respBody)
}

func errorToStringMapper(errs []error) []string {
	errStrings := []string{}
	for _, err := range errs {
		errStrings = append(errStrings, err.Error())
	}
	return errStrings
}

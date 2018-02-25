package user

import (
	"encoding/json"
	"net/http"
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
		respBody, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(respBody)
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
		respBody, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(respBody)
		return
	}

	resp = &postUserResponse{ID: userID}
	respBody, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)
	return
}

type postUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type postUserResponse struct {
	ID    int64  `json:"id"`
	Error string `json:"error"`
}

//NewHTTPHandler initialize user handler with specific service
func NewHTTPHandler(userService Service) HTTPHandler {
	return &userHTTPHandler{userService}
}

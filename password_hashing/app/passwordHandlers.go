package app

import (
	"net/http"
	"password_hashing/dto"
	"password_hashing/errs"
	"password_hashing/service"
	"strconv"
	"sync/atomic"
)

// Define and keep track of the password ids
//source: https://stackoverflow.com/questions/27917750/how-to-define-a-global-counter-in-golang-http-server
var id int64 = 0

// increments the number of the id and returns the new value
func incId() int64 {
	return atomic.AddInt64(&id, 1)
}

// returns the current value
func getId() int64 {
	return atomic.LoadInt64(&id)
}

type PasswordHandler struct {
	service service.PasswordService
}

func (ph PasswordHandler) NewPassword(w http.ResponseWriter, r *http.Request) {
	incId()
	passwordId := getId()
	var request = dto.NewHashRequest{}

	//Build the request object
	request.Id = passwordId
	err := r.ParseForm()
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err)
	}
	passwordString := r.Form.Get("password")
	if passwordString == "" {
		appError := errs.NewValidationError("No password provided")
		writeResponse(w, appError.Code, appError.AsMessage())
		return
	}
	request.PasswordString = passwordString

	//Process Password
	response, appError := ph.service.NewHash(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, response.HashId)
	}
}

func (ph PasswordHandler) FindBy(w http.ResponseWriter, id string) {
	var request = dto.NewHashRequest{}
	hashId, err := strconv.Atoi(id)
	if err != nil {
		appError := errs.NewValidationError("Please provide valid identifier. (Numbers only)")
		writeResponse(w, appError.Code, appError.AsMessage())
		return
	}
	request.Id = int64(hashId)
	hash, appError := ph.service.FindById(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, hash.HashString)
	}
}

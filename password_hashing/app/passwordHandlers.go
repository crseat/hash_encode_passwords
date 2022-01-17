package app

import (
	"net/http"
	"password_hashing/dto"
	"password_hashing/errs"
	"password_hashing/service"
	"strconv"
	"sync"
	"time"
)

type PasswordHandler struct {
	passwordService service.PasswordService
	statsService    service.StatsService
}

func (ph PasswordHandler) NewPassword(w http.ResponseWriter, r *http.Request, startTime time.Time, wg *sync.WaitGroup) {
	var request = dto.NewHashRequest{}

	//Build the request object
	//request.Id = passwordId
	err := r.ParseForm()
	if err != nil {
		//We want to update the average even if the request is bad because that still takes some amount of microseconds
		//to process
		ph.passwordService.UpdateAverage(startTime)
		writeResponse(w, http.StatusBadRequest, err)
		return
	}
	passwordString := r.Form.Get("password")
	if passwordString == "" {
		appError := errs.NewValidationError("No password provided")
		ph.passwordService.UpdateAverage(startTime)
		writeResponse(w, appError.Code, appError.AsMessage())
		return
	}
	request.PasswordString = passwordString

	//Process Password
	response, appError := ph.passwordService.NewHash(request, startTime, wg)
	if appError != nil {
		ph.passwordService.UpdateAverage(startTime)
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		//We don't want to update the average time here because this is only the initial response not including the
		//actual hashing.
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
	hash, appError := ph.passwordService.FindById(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, hash.HashString)
	}
}

func (ph PasswordHandler) GetStats(w http.ResponseWriter) {
	stats, appError := ph.statsService.GetStats()
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, stats)
	}
}

package app

import (
	"net/http"
	"password_hashing/dto"
	"password_hashing/service"
	"sync"
	"time"
)

type PasswordHandler struct {
	passwordService service.PasswordService
	statsService    service.StatsService
}

// NewPassword takes in the ResponseWriter the Request, the start time, and a pointer to the wait group. Validates the
// password string and then sends it on to the password service to process.
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
	request.PasswordString = passwordString
	appError := request.Validate()
	if appError != nil {
		ph.passwordService.UpdateAverage(startTime)
		writeResponse(w, http.StatusBadRequest, appError.AsMessage())
		return
	}

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

// FindBy takes in the ResponseWriter and the hash identifier. Validates the identifier then passes it on to the
// password service to process
func (ph PasswordHandler) FindBy(w http.ResponseWriter, id string) {
	var request = dto.NewHashRequest{}
	hashId, appError := request.ValidateId(id)
	if appError != nil {
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

// GetStats takes in the ResponseWriter and hen calls the stats service to retrieve the current stats.
func (ph PasswordHandler) GetStats(w http.ResponseWriter) {
	stats, appError := ph.statsService.GetStats()
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, stats)
	}
}

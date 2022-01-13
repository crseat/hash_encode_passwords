package service

import (
	"password_hashing/dto"
	"password_hashing/errs"
)

type StatsService interface {
	Stats(request dto.NewStatsRequest) (*dto.NewHashResponse, *errs.AppError)
}

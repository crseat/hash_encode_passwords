package service

import (
	"password_hashing/domain"
	"password_hashing/dto"
	"password_hashing/errs"
)

type StatsService interface {
	GetStats() (*dto.NewStatsResponse, *errs.AppError)
}

type DefaultStatsService struct {
	Repo domain.StatsRepository
}

func (service DefaultStatsService) GetStats() (*dto.NewStatsResponse, *errs.AppError) {
	stats, err := service.Repo.GetStats()
	if err != nil {
		return nil, err
	}
	response := stats.ToNewStatsResponseDto()
	return &response, nil
}

func NewStatsService(statsRepository domain.StatsRepository) DefaultStatsService {
	return DefaultStatsService{statsRepository}
}

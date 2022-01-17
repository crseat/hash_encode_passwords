package domain

import (
	"password_hashing/dto"
	"password_hashing/errs"
)

type Stats struct {
	Total   int64
	Average int64
}

//ToNewStatsResponseDto takes a Stats object and converts it into an appropriate response to the client.
func (stats Stats) ToNewStatsResponseDto() dto.NewStatsResponse {
	return dto.NewStatsResponse{
		Total:   stats.Total,
		Average: stats.Average,
	}
}

type StatsRepository interface {
	GetStats() (*Stats, *errs.AppError)
}

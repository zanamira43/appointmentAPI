package response

import (
	"math"
)

func Pagination(data interface{}, total int64, page, limit int) (*PaginatedResponse, error) {

	// Calculate pagination metadata
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	hasNext := page < totalPages
	hasPrev := page > 1

	res := PaginatedResponse{
		Data:       data,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: int(total / (int64(limit))),
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}

	return &res, nil

}

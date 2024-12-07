package utils

import (
	"net/url"
	"strconv"
	"strings"
)

type Filter struct {
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
	Q         string `json:"q"`         // Search query
	SortBy    string `json:"sortBy"`    // Sort by field
	SortOrder string `json:"sortOrder"` // Sort order (asc or desc)
}

func GetQueryFilter(val url.Values) Filter {
	limit, limitErr := strconv.Atoi(val.Get("limit"))
	offset, offsetErr := strconv.Atoi(val.Get("offset"))
	sortBy := val.Get("sortBy")
	sortOrder := strings.ToUpper(val.Get("sortOrder"))

	filter := GetFilterDefault()

	filter.Q = val.Get("q")

	if limitErr == nil && int32(limit) > 0 {
		filter.Limit = int32(limit)
	}

	if offsetErr == nil && int32(offset) > 0 {
		filter.Offset = int32(offset)
	}

	if sortOrder == "ASC" {
		filter.SortOrder = sortOrder
	}

	if sortBy != "" {
		filter.SortBy = sortBy
	}

	return filter
}

func GetFilterDefault() Filter {
	return Filter{
		Q:         "",
		Limit:     10,
		Offset:    0,
		SortOrder: "DESC",
		SortBy:    "created_at",
	}
}

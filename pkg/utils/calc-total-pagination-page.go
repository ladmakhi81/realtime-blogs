package pkg_utils

import "math"

func CalcTotalPaginationPage(limit, rowsCount uint) uint {
	return uint(math.Ceil(float64(rowsCount) / float64(limit)))
}

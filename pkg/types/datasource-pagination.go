package pkg_types

type DatasourcePagination struct {
	Rows        any  `json:"rows"`
	TotalPage   uint `json:"totalPage"`
	TotalRows   uint `json:"totalRows"`
	CurrentPage uint `json:"currentPage"`
}

func NewDatasourcePagination[T any](
	rows []T,
	totalPage,
	totalRows,
	currentPage uint,
) *DatasourcePagination {
	return &DatasourcePagination{
		Rows:        rows,
		TotalPage:   totalPage,
		TotalRows:   totalRows,
		CurrentPage: currentPage,
	}
}

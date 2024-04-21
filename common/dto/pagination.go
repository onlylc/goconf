package dto

type Pagination struct {

	PageIndex int `from:"pageIndex"`
	PageSize int `from:"pageSize"`
}

func (m *Pagination) GetPageIndex() int {
	if m.PageIndex == 0 {
		m.PageIndex = 1
	}
	return m.PageIndex
}

func (m *Pagination) GetPageSize() int {
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	return m.PageSize
}


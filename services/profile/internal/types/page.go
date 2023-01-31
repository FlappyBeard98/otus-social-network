package types

// PageInfo represents requested page metadata
type PageInfo struct {
	From  int64 `json:"from"`  //page start position
	Count int   `json:"count"` //number of items per page
	Total int64 `json:"total"` //total elements
}

// PageRequest used page requesting
type PageRequest struct {
	Offset int64 `query:"offset"` //page start position
	Limit  int64 `query:"limit"`  //number of items per page
}

// PageResponse contains items and page metadata
type PageResponse[T any] struct {
	Items    []T      `json:"items"`    //page of items
	PageInfo PageInfo `json:"pageInfo"` //page metadata
}

func NewPageResponse[T any](r *PageRequest, items []T, total int64) *PageResponse[T] {
	return &PageResponse[T]{
		Items: items,
		PageInfo: PageInfo{
			From:  r.Offset,
			Count: len(items),
			Total: total,
		},
	}
}

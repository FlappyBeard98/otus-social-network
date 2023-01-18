package types

// PageInfo represents requested page metadata
type PageInfo struct {
	From  int `json:"from"`  //page start position
	Count int `json:"count"` //number of items per page
	Total int `json:"total"` //total elements
}

package dto

type Params struct {
	StatusCode int
	Message    string
	Pagination *Pagination
	Data       any
}

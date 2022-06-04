package data

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PostJsonRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

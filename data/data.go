package data

// type Tasks struct {
// 	Todo Task
// }

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

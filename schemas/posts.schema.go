package schemas

type Post struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

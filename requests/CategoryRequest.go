package requests

type CategoryRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

package requests

type CuisineRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

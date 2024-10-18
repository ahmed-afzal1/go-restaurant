package requests

type SubcategoryRequest struct {
	Name       string `json:"name" form:"name" binding:"required"`
	CategoryId string `json:"category_id" form:"category_id" binding:"required"`
}

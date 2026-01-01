package entity

type Person struct {
	Name  string `json:"name" xml:"name" form:"name" binding:"required,min=2,max=50"`
	Age   int    `json:"age" xml:"age" form:"age" binding:"gte=0,lte=120"`
	Email string `json:"email" xml:"email" form:"email" binding:"required,email"`
}

type Video struct {
	ID          string `json:"id" xml:"id" form:"id" binding:"required" validate:"is-idx"`
	Title       string `json:"title" xml:"title" form:"title" binding:"min=3,max=100"`
	Description string `json:"description" xml:"description" form:"description" binding:"max=500"`
	URL         string `json:"url" xml:"url" form:"url" binding:"required,url"`
	Author      Person `json:"author" xml:"author" form:"author" binding:"required"`
}

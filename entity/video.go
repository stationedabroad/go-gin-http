package entity

type Individual struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age"  binding:"gte=1,lte=110"`
	Email     string `json:"email" validate:"required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=2,max=10"`
	Description string `json:"desc" binding:"max=30"`
	URL         string `json:"url" binding:"required,url"`
	Author      string `json:"author" binding:"required"`
}

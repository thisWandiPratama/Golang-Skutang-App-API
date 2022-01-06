package dto

//RegisterDTO is used when client post from /register url
type RegisterDTO struct {
	Name       string `json:"name" form:"name" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email" `
	Password   string `json:"password" form:"password" binding:"required"`
	Nowhatsapp uint64 `json:"no_whatsapp" form:"no_whatsapp" binding:"required"`
	Alamat     string `json:"alamat" form:"alamat" binding:"required"`
}

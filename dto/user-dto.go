package dto

//UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	ID         uint64 `json:"id" form:"id"`
	Name       string `json:"name" form:"name" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	Password   string `json:"password,omitempty" form:"password,omitempty"`
	Nowhatsapp uint64 `json:"no_whatsapp" form:"no_whatsapp" binding:"required"`
	Alamat     string `json:"alamat" form:"alamat" binding:"required"`
}

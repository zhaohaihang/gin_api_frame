package serializer

type BasePage struct {
	PageNum  int `form:"page_num" json:"page_num"`
	PageSize int `form:"page_size" json:"page_size"`
}

type LoginUserInfo struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,emailorphone"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=20"`
	Type     int    `form:"type" json:"type" binding:"required,oneof=1 2"`
}

type Point struct {
	Lat float64 `form:"lat" json:"lat" binding:"latitude"`
	Lng float64 `form:"lng" json:"lng" binding:"longitude"`
}
type UpdateUserInfo struct {
	Biography string `form:"biography" json:"biography" binding:"omitempty,max=1000"`
	Address   string `form:"address" json:"address" binding:"omitempty,max=1000"`
	Email     string `form:"email" json:"email" binding:"omitempty,email"`
	Phone     string `form:"phone" json:"phone" binding:"omitempty,phone"`
	Location  Point  `form:"location" json:"location" binding:"omitempty,max=1000"`
	Extra     string `form:"extra" json:"extra" binding:"omitempty,max=1000"`
}


type NearInfo struct {
	Point
	Rad int `form:"rad" json:"rad" binding:"lte=100"`
}

type ChangePasswdInfo struct {
	OldPasswd string `form:"old_passwd" json:"old_passwd" binding:"required,min=8,max=20"`
	NewPasswd string `form:"new_passwd" json:"new_passwd" binding:"required,min=8,max=20,eqfield=RePasswd"`
	RePasswd  string `form:"re_passwd" json:"re_passwd" binding:"required,min=8,max=20,eqfield=NewPasswd"`
}

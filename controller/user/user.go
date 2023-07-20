package user

type Id struct {
	Id int64 `uri:"id" form:"id" json:"id" binding:"required"`
}

package request

type CommonRequest struct {
	Page   int    `form:"page"`
	Search string `form:"search"`
}

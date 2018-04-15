package controller

/**
 * 返回结果
 */
type APIJSON struct {
	code int
	msg  string
	data interface{}
}

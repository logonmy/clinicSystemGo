package controller

/**
 * 返回结果
 */
type APIJSON struct {
	code string
	msg  string
	data interface{}
}

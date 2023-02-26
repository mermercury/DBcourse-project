package response

import "Back_End/model"



func SetJSONResp(target *model.JSONResp, code int, msg string, data interface{}) {
	target.Code = code
	target.Message = msg
	target.Data = data
}
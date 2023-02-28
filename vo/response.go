package vo

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// A DataList is an response that is used when the request returned.
// swagger:response DataList
type DataList struct {
	// Required: true
	Item interface{} `json:"item"`
	// Required: true
	Total uint `json:"total"`
}

func BuildListResponse(item interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  item,
			Total: total,
		},
		Msg: "ok",
	}
}

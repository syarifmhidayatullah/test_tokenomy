package model

type Data struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

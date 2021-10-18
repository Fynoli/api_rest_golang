package main

type ResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (this *ResponseMessage) setStatus(data int) { //poner this no es necesario, esto no es objetos, podria llamarse tomate y da igual
	this.Status = data
}

func (this *ResponseMessage) setMessage(data string) {
	this.Message = data
}

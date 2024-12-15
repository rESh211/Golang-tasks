package patterns

import "fmt"

type RequestHandler interface {
	Handle(req *Request) bool
}

type Request struct {
	Data string
}

type HandlerA struct{}

func (h *HandlerA) Handle(req *Request) bool {
	if req.Data == "A" {
		fmt.Println("Обработан HandlerA")
		return true
	}
	return false
}

type HandlerB struct{}

func (h *HandlerB) Handle(req *Request) bool {
	if req.Data == "B" {
		fmt.Println("Обработан HandlerB")
		return true
	}
	return false
}

type HandlerC struct{}

func (h *HandlerC) Handle(req *Request) bool {
	if req.Data == "C" {
		fmt.Println("Обработан HandlerC")
		return true
	}
	return false
}

func RunArrayChain() {
	handlers := []RequestHandler{&HandlerA{}, &HandlerB{}, &HandlerC{}}
	req := &Request{Data: "B"}

	for _, handler := range handlers {
		if handler.Handle(req) {
			break // Запрос обработан
		}
	}
}

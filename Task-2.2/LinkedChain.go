package patterns

// Интерфейс возможностей цепочки.
type ILinkedHandler interface {
	// Добавить обработчик в цепочку.
	Use(handler ILinkedHandler) ILinkedHandler

	// Вызов цепочки обработчика.
	Handle()
}

type RequestLinkedHandler struct {
	ILinkedHandler
	next        ILinkedHandler
	handlerFunc func()
}

func (h *RequestLinkedHandler) Set(handlerFunc func()) ILinkedHandler {
	handler := new(RequestLinkedHandler)
	handler.handlerFunc = handlerFunc
	h.next = handler
	return h
}

func (h *RequestLinkedHandler) Handle() {
	defer func() {
		if rec := recover(); rec != nil {

			return
		}
	}()

	if h.handlerFunc != nil {
		h.handlerFunc()
	} else {
		panic("Handler Function is not determined")
	}
}

func (h *RequestLinkedHandler) Use(handler ILinkedHandler) ILinkedHandler {
	h.next = handler
	return h
}

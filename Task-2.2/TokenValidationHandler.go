package patterns

// Обработчик с проверкой валидности токена.
type TokenValidationHandler struct {
	ILinkedHandler
	next        ILinkedHandler
	handlerFunc func()
	sources     []ITokenSource
}

func (h *TokenValidationHandler) Source(sources []ITokenSource) *TokenValidationHandler {
	h.sources = sources
	return h
}

func (h *TokenValidationHandler) Set(handlerFunc func()) ILinkedHandler {
	handler := new(RequestLinkedHandler)
	handler.handlerFunc = handlerFunc
	h.next = handler
	return h
}

func (h *TokenValidationHandler) Handle() {
	defer func() {
		if rec := recover(); rec != nil {

			return
		}
	}()

	if h.sources == nil || len(h.sources) == 0 {
		panic("Token source from request not declared in handler")
	}

	// if tokenValid, err := TokenValid(w, r, h.sources); err == nil && tokenValid {
	// 	h.next.Handle(w, r)
	// } else {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }
}

func (h *TokenValidationHandler) Use(handler ILinkedHandler) ILinkedHandler {
	h.next = handler
	return h
}

// Интерфейс стратегии получения токена из запроса.
type ITokenSource interface {
	// Получение токена из запроса.
	get() (string, error)
}

// Получение токена из заголовка.
type HeaderToken struct {
	ITokenSource
}

// Получение токена из куки.
type CookieToken struct {
	ITokenSource
}

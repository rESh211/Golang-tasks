package chain

import "fmt"

type Request struct {
	cookie  map[string]string
	headers map[string]string
	body    any
}

type Response struct {
	body        any
	ok          bool
	description string
}

func Login(request *Request, response *Response) {
	response.ok = true
	response.description = "авторизация прошла"
}

func GetData(request *Request, response *Response) {
	response.ok = true
	response.description = "ffee"
}

type RequestHandler interface {
	Handle(request *Request, response *Response) bool
}

type LoginHandler struct{}

func (loginHandler *LoginHandler) Handle(request *Request, response *Response) bool {
	Login(request, response)
	return false
}

type GetDataHandler struct{}

func (GetDataHandler *GetDataHandler) Handle(request *Request, response *Response) bool {
	GetData(request, response)
	return false
}

type TokenValidationHandler struct{}

func CheckToken(token string) bool {
	gg := []string{"32", "f3", "44", "23f", "e2"}
	for _, i := range gg {
		if i == token {
			return true
		}
	}
	return false
}

func (TokenValidationHandler *TokenValidationHandler) Handle(request *Request, response *Response) bool {
	var headerSource = new(HeaderSource)
	var cookieSource = new(CookieSource)
	var defaultSources = []ITokenSource{headerSource, cookieSource}

	for _, sources := range defaultSources {
		token := sources.Get(request)
		if len(token) > 0 {
			TokenValid := CheckToken(token)
			if TokenValid {
				return true
			} else {
				response.description = "токен не действителен"
				response.ok = false
				return false
			}

		}
	}
	response.description = "запрос не авторизован, токен не найден "
	response.ok = false
	return false

}

type ITokenSource interface {
	Get(request *Request) string
}

type CookieSource struct{}

func (CookieSource *CookieSource) Get(request *Request) string {
	return request.cookie["token"]
}

type HeaderSource struct{}

func (HeaderSource *HeaderSource) Get(request *Request) string {
	return request.headers["authorization"]
}

func Run() {
	var cookie = make(map[string]string)
	cookie["token"] = "44"

	loginRequest := &Request{}
	getDataRequest := &Request{cookie: cookie}
	loginResponse := &Response{}
	getDataResponse := &Response{}
	//Login(loginRequest, loginResponse)
	loginChain := []RequestHandler{&LoginHandler{}}
	dataChain := []RequestHandler{&TokenValidationHandler{}, &GetDataHandler{}}

	for _, handler := range loginChain {
		if handler.Handle(loginRequest, loginResponse) == false {
			break
		}
	}
	for _, handler := range dataChain {
		if handler.Handle(getDataRequest, getDataResponse) == false {
			break
		}
	}
	//(getDataRequest, getDataResponse)
	fmt.Println(loginResponse, getDataResponse)
}

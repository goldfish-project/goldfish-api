package domain

type Header struct {
	Key string
	Value string
}

type Request struct {
	RequestId string
	Type      string
	Headers   []Header
	Body      interface{}
}
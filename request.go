package main

type Request interface {
	ApiName() string

	GetParams() map[string]string

	GetHttpMethod() string

	GetResponse() Response
}

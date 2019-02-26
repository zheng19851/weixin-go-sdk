package main

type Request interface {
	GetResponse() Response
}

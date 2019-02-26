package main

type Response interface {
	IsSuccess() bool
}

package interceptors

import "project/internal/observable"

var Interceptors []observable.Observer = make([]observable.Observer, 0)

func InterceptorRegister(intercept observable.Observer) {
	Interceptors = append(Interceptors, intercept)
}

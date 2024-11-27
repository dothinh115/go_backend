package interceptors

import "project/internal/observable"

var interceptors []observable.Observer = make([]observable.Observer, 0)

func InterceptorRegister(intercept observable.Observer) {
	interceptors = append(interceptors, intercept)
}

func GetInterceptors() []observable.Observer {
	return interceptors
}

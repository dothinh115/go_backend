package observable

import "sync"

type Observer func(interface{}) interface{}

type Observable struct {
	dataChannel chan interface{}
	transforms  []func(interface{}) interface{}
	w           sync.WaitGroup
}

func NewObservable() *Observable {
	return &Observable{
		dataChannel: make(chan interface{}),
		transforms:  make([]func(interface{}) interface{}, 0),
	}
}

func (o *Observable) Next(value interface{}) *Observable {
	o.w.Add(1)
	o.dataChannel <- value
	return o
}

func (o *Observable) Complete() {
	o.w.Wait()
	close(o.dataChannel)
}

func (o *Observable) Map(observer Observer) *Observable {
	o.transforms = append(o.transforms, observer)
	return o
}

func (o *Observable) Subscribe(callback func(interface{})) *Observable {
	go func() {
		for {
			data, ok := <-o.dataChannel
			if !ok {
				return
			}
			for _, transform := range o.transforms {
				data = transform(data)
			}
			callback(data)
			o.w.Done()
		}
	}()
	return o
}

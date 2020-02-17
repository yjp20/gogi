package main

type Handler func(interface{})

type Dispatcher struct {
	counter      int
	callbacks    map[string]map[int]Handler
	callbackName map[int]string
}

func (d *Dispatcher) Init() {
	d.callbacks = make(map[string]map[int]Handler)
	d.callbackName = make(map[int]string)
}

func (d *Dispatcher) Register(name string, callback Handler) int {
	if _, ok := d.callbacks[name]; !ok {
		d.callbacks[name] = make(map[int]Handler)
	}

	d.counter++
	id := d.counter
	d.callbacks[name][id] = callback
	d.callbackName[id] = name

	return id
}

func (d *Dispatcher) Dispatch(name string, data interface{}) {
	if m, ok := d.callbacks[name]; ok {
		for _, c := range m {
			c(data)
		}
	}
}

func (d *Dispatcher) Unregister(id int) {
	name := d.callbackName[id]
	delete(d.callbackName, id)
	delete(d.callbacks[name], id)
}

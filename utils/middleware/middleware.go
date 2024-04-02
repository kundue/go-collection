package middleware

type Handler func()

type app struct {
	index    int
	handlers []Handler
}

// New function to create instance
func New() app {
	return app{}
}

// Use function to collect and store handlers
func (a *app) Use(handlers ...Handler) {
	a.handlers = append(a.handlers, handlers...)
}

// Next function to run the next handler in the chain
func (a *app) Next() {
	if len(a.handlers) > a.index {
		a.handlers[a.index]()
		a.index++
	}
}

// Run function to execute all collected handlers sequentially
func (a *app) Run() {
	for a.index < len(a.handlers) {
		a.Next()
	}
}

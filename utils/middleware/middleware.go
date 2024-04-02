package middleware

type Handler func()

type App struct {
	index    int
	handlers []Handler
}

// Use function to collect and store handlers
func (a *App) Use(handlers ...Handler) {
	a.handlers = append(a.handlers, handlers...)
}

// Next function to run the next handler in the chain
func (a *App) Next() {
	if len(a.handlers) > a.index {
		a.handlers[a.index]()
		a.index++
	}
}

// Run function to execute all collected handlers sequentially
func (a *App) Run() {
	for a.index < len(a.handlers) {
		a.Next()
	}
}

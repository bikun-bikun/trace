package trace

type Tracer interface {
	Tracer(...interface{})
}
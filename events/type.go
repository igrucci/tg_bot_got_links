package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
	Meta interface{}
}

//type Fetcher interface {
//	Fetch(limit int) ([]Event, error)
//}
//
//type Processor interface {
//	Procces(e Event) error
//}
//
//type Type int
//
//const (
//	Unknown Type = iota
//	Message
//)
//
//type Event struct {
//	Type Type
//	Text string
//	Meta interface{}
//}

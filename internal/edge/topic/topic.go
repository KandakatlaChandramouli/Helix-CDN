package topic

import "sync"

type Registry struct {
	topics sync.Map
}

func New() *Registry {
	return &Registry{}
}

func (r *Registry) Subscribe(
	topic uint32,
	client uint64,
) {
	value, _ := r.topics.LoadOrStore(
		topic,
		&sync.Map{},
	)

	subs := value.(*sync.Map)

	subs.Store(client, struct{}{})
}

func (r *Registry) Unsubscribe(
	topic uint32,
	client uint64,
) {
	value, ok := r.topics.Load(topic)

	if !ok {
		return
	}

	subs := value.(*sync.Map)

	subs.Delete(client)
}

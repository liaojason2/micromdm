package queue

type QueueService struct {
	store Store
}

func New(store Store) *QueueService {
	return &QueueService{store: store}
}

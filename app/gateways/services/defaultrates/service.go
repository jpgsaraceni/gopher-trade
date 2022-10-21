package defaultrates

type Service struct {
	Client Client
	Cache  Cache
}

func NewService(client Client, cache Cache) Service {
	return Service{
		Client: client,
		Cache:  cache,
	}
}

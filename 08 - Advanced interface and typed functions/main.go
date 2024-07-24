package main

type Putter interface {
	Put(id int, val any) error
}

type Storage interface {
	Putter
	Get(id int) (any, error)
}

type FooStorage struct{}

func (s *FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *FooStorage) Put(id int, val any) error {
	return nil
}

type BarStorage struct{}

func (s *BarStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *BarStorage) Put(id int, val any) error {
	return nil
}

type Server struct {
	storage Storage
}

func main() {
	s := &Server{
		storage: &FooStorage{},
	}
	s.storage.Get(1)
	s.storage.Put(1, "foo")
}

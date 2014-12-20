package disk

import (
	"github.com/johnwilson/x"
)

type Store struct {
}

func (s *Store) Start(config string) error {
	return nil
}
func (s *Store) Add(i x.Image, path string) (int, error) {
	return 0, nil
}
func (s *Store) Get(p string) (x.Image, error) {
	return nil, nil
}

func NewStore() *Store {
	return &Store{}
}

func init() {
	x.RegisterStore("disk", NewStore())
}

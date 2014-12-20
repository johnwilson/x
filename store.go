package x

import (
	"fmt"
	"log"
)

var storePlugins = make(map[string]Store)

type Store interface {
	Start(config string) error
	Add(i Image, path string) (int, error)
	Get(p string) (Image, error)
}

func RegisterStore(name string, plugin Store) {
	if plugin == nil {
		log.Fatal("Store plugin registration: plugin is nil")
	}

	if _, exists := storePlugins[name]; exists {
		log.Printf("Store plugin registration: %q already registered\n", name)
		return
	}
	storePlugins[name] = plugin
}

func NewStore(name, config string) (plugin Store, err error) {
	plugin, ok := storePlugins[name]
	if !ok {
		err = fmt.Errorf("Store plugin Creation: unknown plugin name %q (forgot to import?)", name)
		return
	}
	err = plugin.Start(config)
	if err != nil {
		plugin = nil
	}
	return
}

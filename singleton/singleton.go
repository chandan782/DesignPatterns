// The Singleton pattern ensures that a class has only one instance and provides a global point of access to it.
package singleton

import "sync"

type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "single instance"}
	})

	return instance
}

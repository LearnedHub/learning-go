package learning

import "fmt"

type Learning interface {
	Learn()
	Kind() string
}

func Do(learning Learning) {
	fmt.Println("Learning: ", learning.Kind())
	learning.Learn()
}

package learning

type Learning interface {
	Learn()
}

func Do(learning Learning) {
	learning.Learn()
}

package azul

var destroyFunctions = make([]func(), 0)

func RegisterDestroy(f func()) {
	destroyFunctions = append(destroyFunctions, f)
}

func Destroy() {
	for i := 0; i < len(destroyFunctions); i++ {
		destroyFunctions[i]()
	}
}

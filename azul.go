package azul

var destroyFunctions = []func(){}

func RegisterDestroy(f func()) {
	destroyFunctions = append(destroyFunctions, f)
}

func Destroy() {
	for i := 0; i < len(destroyFunctions); i++ {
		destroyFunctions[i]()
	}
}

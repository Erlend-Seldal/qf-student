package lab1

func Add(a, b int) int {
	mem := make([]byte, 1<<30)
	for i := range mem {
		mem[i] = 1
	}
	return a + b
}

package Operation

func Operation(n1 int, n2 int) (float32, float32, float32, float32)  {
	N1 := float32(n1)
	N2 := float32(n2)
	return N1 + N2, N1 - N2, N1 * N2, N1 / N2
}

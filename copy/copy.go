package copy

func GenerateTestDataSrc() (src [2048][2048]int) {
	for i := 0; i < 2048; i++ {
		for j := 0; j < 2048; j++ {
			src[i][j] = i + j
		}
	}

	return
}

func Copyij(src *[2048][2048]int, dst *[2048][2048]int) {
	for i := 0; i < 2048; i++ {
		for j := 0; j < 2048; j++ {
			dst[i][j] = src[i][j]
		}
	}
}

func Copyji(src *[2048][2048]int, dst *[2048][2048]int) {
	for j := 0; j < 2048; j++ {
		for i := 0; i < 2048; i++ {
			dst[i][j] = src[i][j]
		}
	}
}

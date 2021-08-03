package iteration

func repeat(c string, r int) (rep string) {

	for i := 0; i < r; i++ {
		rep += c
	}

	return
}

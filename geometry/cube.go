package geometry

func CubeVolume(n int) int, err {

	if n != 0 {
		return n * n * n, nil

	} else {
		return 0, errors.New("Zero length edge is not allowed")
	}

}

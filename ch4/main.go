package ch4

func main() {
	switch v := 1; v {
	case 1:
		// x := 2
		fallthrough
	case 2:
		// fmt.Println("x:", x)
	default:
	}
}

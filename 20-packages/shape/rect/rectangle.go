package rect //must the name of immediate dir

// there are no access specifires ,

// only a concept exported or not exported
// any memeber/type start with uppercase exportedþ
// /any memeber/type start with lowercase is unexported(local only to package)

func Perimeter(l, b float32) float32 {
	return 2 * (l + b)
}

package sort

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	if len(a) == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return a
	}

	half := len(a) / 2
	l, r := a[:half], a[half:]
	return merge(MergeSort(l), MergeSort(r))
}

func merge(l []int, r []int) []int {
	a := make([]int, 0, len(l)+len(r))
	j := 0

	for i := 0; i < len(r); i++ {
		for j < len(l) && l[j] < r[i] {
			a = append(a, l[j])
			j++
		}
		a = append(a, r[i])
	}

	if j < len(l) {
		a = append(a, l[j:]...)
	}

	return a
}

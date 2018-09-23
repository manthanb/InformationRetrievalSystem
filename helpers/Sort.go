package helpers

func SortDocuments(A []Documents) []Documents {
	if len(A) <= 1 {
		return A
	}

	left, right := SplitDocuments(A)
	left = SortDocuments(left)
	right = SortDocuments(right)
	return MergeDocuments(left, right)
}

func SplitDocuments(A []Documents) ([]Documents, []Documents) {
	return A[0 : len(A)/2], A[len(A)/2:]
}

func MergeDocuments(A, B []Documents) []Documents {
	arr := make([]Documents, len(A)+len(B))

	j, k := 0, 0

	for i := 0; i < len(arr); i++ {

		if j >= len(A) {
			arr[i] = B[k]
			k++
			continue
		} else if k >= len(B) {
			arr[i] = A[j]
			j++
			continue
		}

		if A[j].TermRatio > B[k].TermRatio {
			arr[i] = B[k]
			k++
		} else {
			arr[i] = A[j]
			j++
		}
	}

	return arr

}

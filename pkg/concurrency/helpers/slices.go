package helpers

func SplitSlice[T any](data []T, chunk int) [][]T {
	var segments [][]T
	size := len(data)
	chunkSize := size / chunk
	if size%chunk != 0 {
		chunkSize++
	}

	for i := 0; i < size; i += chunkSize {
		end := i + chunkSize
		if end > size {
			end = size
		}
		segments = append(segments, data[i:end])
	}
	return segments
}

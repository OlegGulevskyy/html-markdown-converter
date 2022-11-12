package utils

func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) {

	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize])
	}

	return append(chunks, items)
}

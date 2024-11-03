package hashtable

// HashTable structure with custom hast function
type HashTable struct {
	buckets  [][]int
	size     int
	hashFunc func(int, int) int
}

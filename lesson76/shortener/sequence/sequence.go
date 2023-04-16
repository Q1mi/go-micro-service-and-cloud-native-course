package sequence

// Sequence 取号器接口
type Sequence interface {
	Next() (uint64, error)
}

package sequence

// 基于redis实现一个发号器

type Redis struct {
	// redis连接
}

func NewRedis(redisAddr string) Sequence {
	return &Redis{}
}

func (r *Redis) Next() (seq uint64, err error) {
	// 使用redis实现发号器的思路
	// incr
	return
}

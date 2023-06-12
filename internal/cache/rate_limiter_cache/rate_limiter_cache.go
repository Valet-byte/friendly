package cache

type RateLimiterCache interface {
	IncrementRequestCount(ip string) error
	GetRequestCount(ip string) (int, error)
}

package market

type Symbol interface {
	Code() string
}

type Interval interface {
	Seconds() int
	String() string
}

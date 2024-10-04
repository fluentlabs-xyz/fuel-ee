package helpers

type StartStop interface {
	Start() error
	Stop() error
}

type EntityWithUniqId interface {
	GetUniqueId() string
}

type EntityWithUniqIdProbably interface {
	TryGetUniqueId() (string, error)
}

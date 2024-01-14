package storage

type StorageError string

func (s StorageError) Error() string { return string(s) }

const (
	ParseError = StorageError("couldn't parse data to settings struct")
	SetError   = StorageError("couldn't save data to redis")
)

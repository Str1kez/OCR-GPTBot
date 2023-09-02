package storage

type (
	ParseError struct{}
	SetError   struct{}
)

func (p *ParseError) Error() string {
	return "couldn't parse data to settings struct"
}

func (s *SetError) Error() string {
	return "couldn't save data to redis"
}

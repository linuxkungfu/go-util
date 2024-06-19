package iorm

type NatsNCNilErr struct {
}

func (nerr *NatsNCNilErr) Error() string {
	return "nc is nil"
}

package fixture

type DiskReadError struct {
	Err error
}

func (e DiskReadError) Error() string {
	return "read settings file: " + e.Err.Error()
}

func (e DiskReadError) Unwrap() error {
	return e.Err
}

type DecodeError struct {
	Err error
}

func (e DecodeError) Error() string {
	return "decode settings: " + e.Err.Error()
}

func (e DecodeError) Unwrap() error {
	return e.Err
}

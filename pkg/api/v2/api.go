package apiv2

type API struct {
}

func NewAPI() (*API, error) {
	a := &API{}
	return a, nil
}
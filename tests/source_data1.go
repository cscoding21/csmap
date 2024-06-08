package tests

type TestSource struct {
	ID       string
	Name     string
	Age      int
	Location LocationSource
}

type LocationSource struct {
	Lat float64
	Lon float64
}

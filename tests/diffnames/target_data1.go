package diffnames

type TestTarget struct {
	ID       string
	Name     *string
	Age      int
	Location LocationTarget
}

type LocationTarget struct {
	Lat float64
	Lon float64
}

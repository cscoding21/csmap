package diffnames

// TestTarget used for testing
type TestTarget struct {
	ID       string
	Name     *string
	Age      int
	Location LocationTarget
}

// LocationTarget used for testing
type LocationTarget struct {
	Lat float64
	Lon float64
}

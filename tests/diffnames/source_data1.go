package diffnames

// TestSource used for testing
type TestSource struct {
	ID       string
	Name     string
	Age      *int
	Location LocationSource
}

// LocationSource used for testing
type LocationSource struct {
	Lat float64
	Lon float64
}

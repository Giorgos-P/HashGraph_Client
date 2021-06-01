package variables

var (
	// ID - This client's id.
	ID int

	// N - Number of processors
	N int

	// F - Number of faulty processors
	F int

	Clients int

	// Remote - If we are running locally or remotely
	Remote bool
)

// Initialize - Variables initializer method
func Initialize(id int, n int, clientsNum int, rem int) {
	ID = id
	N = n
	F = (N - 1) / 3
	Clients = clientsNum
	if rem == 1 {
		Remote = true
	} else {
		Remote = false
	}
}

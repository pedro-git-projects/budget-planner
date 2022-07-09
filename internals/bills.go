package internals

type Bill struct {
	title     string
	id        uint
	cost      Real
	frequency frequency
	status    status
}

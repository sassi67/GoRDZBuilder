package model

/*GridHandler is interface for all the data grids */
type GridHandler interface {
	Write(row uint, col uint, size uint, values []uint16) error
	Read(row uint, col uint, size uint) ([]uint16, error)
}

package model

import (
	"fmt"
	"sync"
)

const cGridModBusRows uint = 5000
const cGridModBusCols uint = 3

/*GridModbus holds the data of Modbus boards*/
type GridModbus struct {
	data  [cGridModBusRows * cGridModBusCols]uint16
	mutex sync.RWMutex
	rows  uint
	cols  uint
}

/*NewGridModBus returns a pointer to a newly created GridModbus*/
func NewGridModBus() *GridModbus {
	return &GridModbus{
		rows: cGridModBusRows,
		cols: cGridModBusCols,
	}
}

func (grid *GridModbus) Write(row uint, col uint, size uint, values []uint16) error {
	err := grid.checkGridBounds(row, col, size)
	if err != nil {
		return err
	}

	grid.mutex.Lock()
	defer grid.mutex.Unlock()

	var startFrom uint = grid.rows*(col-1) + row - 1
	for siz := uint(0); siz < size; siz++ {
		grid.data[startFrom+siz] = values[siz]
	}
	return nil
}

func (grid *GridModbus) Read(row uint, col uint, size uint) ([]uint16, error) {
	err := grid.checkGridBounds(row, col, size)
	if err != nil {
		return nil, err
	}

	grid.mutex.RLock()
	defer grid.mutex.RUnlock()

	var values []uint16
	var startFrom uint = grid.rows*(col-1) + row - 1
	for siz := uint(0); siz < size; siz++ {
		values = append(values, grid.data[startFrom+siz])
	}
	return values, nil
}

func (grid *GridModbus) checkGridBounds(row uint, col uint, size uint) error {
	if row == 0 {
		return fmt.Errorf("Row index cannot be zero")
	}
	if col == 0 {
		return fmt.Errorf("Column index cannot be zero")
	}
	if size == 0 {
		return fmt.Errorf("Size cannot be zero")
	}
	if row > grid.rows {
		return fmt.Errorf("Invalid row index: %d. Max allowed: %d", row, grid.rows)
	}
	if col > grid.cols {
		return fmt.Errorf("Invalid column index: %d. Max allowed: %d", col, grid.cols)
	}
	if len := row + size; len > grid.rows {
		return fmt.Errorf("Cannot write: %d elements. Max allowed: %d", len, grid.rows)
	}

	return nil
}

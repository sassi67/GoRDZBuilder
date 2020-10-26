package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGridModBus(t *testing.T) {
	grid := NewGridModBus()
	assert.NotNil(t, grid, "NewGridModBus should not return nil.")
	assert.Equal(t, cGridModBusCols, grid.cols, "Wrong number of cols")
	assert.Equal(t, cGridModBusRows, grid.rows, "Wrong number of rows")

}

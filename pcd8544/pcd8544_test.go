package pcd8544

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"tinygo.org/x/drivers/tester"
)

func TestDefaultPCD8544(t *testing.T) {
	c := qt.New(t)
	bus := tester.NewSPIBus(c)
	pin1 := tester.NewPin(c)
	pin2 := tester.NewPin(c)
	pin3 := tester.NewPin(c)
	dev := New(bus, pin1, pin2, pin3)
	c.Assert(dev, qt.Not(qt.IsNil))
}
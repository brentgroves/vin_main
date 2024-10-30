package main

import (
	"testing"

	vin "github.com/brentgroves/vin1"
)

const testVIN = "W09000051T2123456"

func TestVIN_Manufacturer(t *testing.T) {

	manufacturer := vin.Manufacturer(testVIN)
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
}

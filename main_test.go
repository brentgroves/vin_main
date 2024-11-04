package main_test

import (
	"testing"

	vin "github.com/brentgroves/vin1/vin-stages/4"
)

const (
	validVIN   = "W0L000051T2123456"
	invalidVIN = "W0"
	euSmallVIN = "W09000051T2123456"
)

func TestVIN_New(t *testing.T) {

	_, err := vin.NewVIN(validVIN)
	if err != nil {
		t.Errorf("creating valid VIN returned an error: %s", err.Error())
	}

	_, err = vin.NewVIN(invalidVIN)
	if err == nil {
		t.Error("creating invalid VIN did not return an error")
	}
}

func TestVIN_Manufacturer(t *testing.T) {

	testVIN, _ := vin.NewVIN(validVIN)
	manufacturer := testVIN.Manufacturer()
	if manufacturer != "W0L" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
}

func TestVIN_EU_SmallManufacturer_Polymorphism(t *testing.T) {

	// slice of vin or vinEU types that implement VIN interface
	var testVINs []vin.VIN
	testVIN, _ := vin.NewEUVIN(euSmallVIN)
	// now there is no need to cast!
	testVINs = append(testVINs, testVIN)

	for _, vin := range testVINs {
		manufacturer := vin.Manufacturer()
		if manufacturer != "W09123" {
			t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
		}
	}
}

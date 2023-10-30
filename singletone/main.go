package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	dm   *DeviceManager
)

type Device struct {
	ID     int
	Status string
}

type Printer struct {
	Device
}

type Camera struct {
	Device
}

type DeviceManager struct {
	Printers []*Printer
	Cameras  []*Camera
}

func NewDeviceManager() *DeviceManager {
	once.Do(func() {
		dm = &DeviceManager{
			Printers: []*Printer{},
			Cameras:  []*Camera{},
		}
	})
	return dm
}

func (dm *DeviceManager) AddPrinter(id int) {

	printer := &Printer{Device: Device{ID: id, Status: "Ready"}}
	dm.Printers = append(dm.Printers, printer)
}

func (dm *DeviceManager) AddCamera(id int) {

	camera := &Camera{Device: Device{ID: id, Status: "Ready"}}
	dm.Cameras = append(dm.Cameras, camera)
}

func (dm *DeviceManager) PrintStatus() {

	fmt.Println("Printers:")
	for _, printer := range dm.Printers {
		fmt.Printf("Printer ID %d: %s\n", printer.ID, printer.Status)
	}

	fmt.Println("Cameras:")
	for _, camera := range dm.Cameras {
		fmt.Printf("Camera ID %d: %s\n", camera.ID, camera.Status)
	}
}

func main() {
	dm := NewDeviceManager()

	dm.AddPrinter(1)
	dm.AddPrinter(2)
	dm.AddCamera(1)

	dm.PrintStatus()
}

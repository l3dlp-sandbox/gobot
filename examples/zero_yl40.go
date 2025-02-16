//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"fmt"
	"log"
	"time"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/drivers/i2c"
	"gobot.io/x/gobot/v2/platforms/radxa/zero"
)

func main() {
	// Wiring
	// PWR : 1, 17 (+3.3V, VCC), 6, 9, 14, 20, 25, 30, 34, 39 (GND)
	// I2C3: 3 (SDA), 5 (SCL), I2C1: 24/16 (SDA), 23/13 (SCL), I2C4: 7 (SDA), 11 (SCL)
	// YL-40 module: wire AOUT --> AIN2 for this example
	//
	// Note: temperature measurement is often buggy, because sensor is not properly grounded
	//       fix it by soldering a small bridge to the adjacent ground pin of brightness sensor
	board := zero.NewAdaptor()
	yl := i2c.NewYL40Driver(board, i2c.WithBus(3))

	work := func() {
		// the LED light is visible above ~1.7V
		writeVal, _ := yl.AOUT()

		gobot.Every(1000*time.Millisecond, func() {
			if err := yl.Write(writeVal); err != nil {
				fmt.Println(err)
			} else {
				log.Printf(" %.1f V written", writeVal)
				writeVal = writeVal + 0.1
				if writeVal > 3.3 {
					writeVal = 0
				}
			}

			if brightness, err := yl.ReadBrightness(); err != nil {
				fmt.Println(err)
			} else {
				log.Printf("Brightness: %.0f [0..1000]", brightness)
			}

			if temperature, err := yl.ReadTemperature(); err != nil {
				fmt.Println(err)
			} else {
				log.Printf("Temperature: %.1f °C", temperature)
			}

			if ain2, err := yl.ReadAIN2(); err != nil {
				fmt.Println(err)
			} else {
				log.Printf("Read back AOUT: %.1f [0..3.3]", ain2)
			}

			if potiState, err := yl.ReadPotentiometer(); err != nil {
				fmt.Println(err)
			} else {
				log.Printf("Resistor: %.0f %% [-100..+100]", potiState)
			}
		})
	}

	robot := gobot.NewRobot("yl40Bot",
		[]gobot.Connection{board},
		[]gobot.Device{yl},
		work,
	)

	if err := robot.Start(); err != nil {
		panic(err)
	}
}

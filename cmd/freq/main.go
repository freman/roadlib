package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/freman/roadlib/speed"
	"github.com/freman/roadlib/tire"
)

func must(f float64, e error) float64 {
	if e != nil {
		panic(e)
	}
	return f
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Syntax: freq {tiresize} {speed in km/hr}")
		os.Exit(1)
	}

	tireInfo, err := tire.Parse(os.Args[1])
	if err != nil {
		panic(err)
	}
	speed := speed.KMHR(must(strconv.ParseFloat(os.Args[2], 64)))

	frequency := float64(speed.MetresSecond()) / tireInfo.Revolutions.PerMetre()

	fmt.Printf("wheel imbalance should result vibrations of approximately %0.2f hz or %0.2f hz\n", frequency, frequency*2)
}

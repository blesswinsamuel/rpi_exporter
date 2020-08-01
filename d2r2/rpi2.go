package main

import (
	"fmt"
	"log"

	// "github.com/blesswinsamuel/rpi_exporter/d2r2/dht"

	"github.com/blesswinsamuel/rpi_exporter/dht2"
	"periph.io/x/periph/host/rpi"
)

func main() {
	// temperature, humidity, retried, err :=
	// 	dht.ReadDHTxxWithRetry(dht.DHT11, rpi.P1_7, false, 10)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	if err := dht2.HostInit(); err != nil {
		log.Fatal(err)
	}
	dht, err := dht2.NewDHT(rpi.P1_7, dht2.Celsius, "dht11")
	if err != nil {
		log.Fatal(err)
	}
	humidity, temperature, err := dht.ReadRetry(11)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Printf("humidity: %v\n", humidity)
	fmt.Printf("temperature: %v\n", temperature)

	// temperature, humidity, retried, err :=
	// 	dht2.ReadDHTxxWithRetry(dht.DHT11, rpi.P1_7, false, 10)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // Print temperature and humidity
	// fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
	// 	temperature, humidity, retried)
}

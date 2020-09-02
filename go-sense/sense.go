package main

import (
	"log"
	"fmt"

	"github.com/mdlayher/lmsensors"
)

func main() {
	scanner := lmsensors.New()
	devices,err := scanner.Scan()
	if  err != nil {
		log.Fatal(err)
	}
    	fmt.Printf("%s\t\t%s\t%s\t%s\n","Device","Sensor", "Label", "Temperature")
	for _ ,x := range devices {	
//		fmt.Printf("%s\n",x.Name)
		if x.Sensors != nil {
			for _, sensor := range x.Sensors {

				switch g := sensor.(type) {
  					case *lmsensors.TemperatureSensor:
    						fmt.Printf("%s\t%s\t%s\t%f\n",x.Name, g.Name, g.Label, g.Input)
//  					case int:
//    						fmt.Println("g is an integer, whose value is", g)
  					default:
    						fmt.Printf("I don't know what g is\n",g)
				}

//				temp := &sensor.(lmsensors.TemperatureSensor)
				//if ok == true {
					//fmt.Printf("%s \n",sensor)
				//}
			}
		}
	}
}

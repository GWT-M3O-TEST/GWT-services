package example

import (
	"fmt"
	"github.com/micro/services/clients/go/location"
	"os"
)

// Read an entity by its ID
func GetLocationById() {
	locationService := location.NewLocationService(os.Getenv("MICRO_API_TOKEN"))
	rsp, err := locationService.Read(&location.ReadRequest{
		Id: "1",
	})
	fmt.Println(rsp, err)
}
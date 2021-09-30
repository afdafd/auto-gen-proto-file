package main

import (
	_ "customPro/protoGen/database"
	"customPro/protoGen/routers"
	_ "customPro/protoGen/routers"
)

func main() {
	routers.Run()
}

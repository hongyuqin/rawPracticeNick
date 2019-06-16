package main

import (
	"fmt"
	"time"
)

func TestTimeLocal() {
	fmt.Println(*time.Local)
}

func main() {
	TestTimeLocal()
}

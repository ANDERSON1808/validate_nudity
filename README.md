# validate_nudity
library to validate nudes from any image


## Quick Start

```go
package main

import (
	"fmt"
	"github.com/ANDERSON1808/validate_nudity"
	"io/ioutil"
)

func main() {
	data, _ := ioutil.ReadFile("demo.png")
	isNude, err := validate_nudity.IsImageNude(&data)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	fmt.Printf("Img is nude %s",isNude)
}
```

## Installing

```
go get github.com/ANDERSON1808/validate_nudity
```

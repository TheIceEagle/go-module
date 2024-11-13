package hello

import (
	"fmt"

	"github.com/pkg/errors"
)

func Hello() {
	fmt.Println(errors.New("Testing Module!"))
}

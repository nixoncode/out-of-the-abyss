/*
Copyright © 2023 Nixon Kosgei <nkosgey6@gmail.com>
*/
package main

import (
	"github.com/nixoncode/out-of-the-abyss/internal/app"
)

func main() {
	application := app.New()

	err := application.Start()
	if err != nil {
		panic("Cannot start the application")
	}

}

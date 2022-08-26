package app

import "github.com/kai1/osx-activated/internal/app/graphical"

func init() {

}

func App() error {
	var err error

	graphical.UserInterface()

	return err
}

package smaug

import "fmt"

type Context struct {
	User Identity
}

type Identity struct {
	Id     string
	Agency string
}

func (i Identity) String() string {
	return fmt.Sprintf("%s@%s", i.Id, i.Agency)
}

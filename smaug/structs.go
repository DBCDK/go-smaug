package smaug

type Context struct {
	User Identity
}

type Identity struct {
	Id     string
	Agency string
}


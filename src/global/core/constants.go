package core

type UserRol string

const (
	RolUser      UserRol = "CTUSERR"
	RolOrganizer UserRol = "CTARGNZ"
	RolAdmin     UserRol = "CTADMIN"
	RolLogistic  UserRol = "CTLOGIR"
	RolSupport   UserRol = "CTSPPOT"
)

const PagLimit = 10

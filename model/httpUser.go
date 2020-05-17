package model

const (
	HTTPUserCategory_Hoster       int = 1
	HTTPUserCategory_Customer     int = 2
	HTTPUserCategory_SalesAdvisor int = 3
)

type HTTPUser struct {
	UserType int
	UserID   string
	UserName string
}

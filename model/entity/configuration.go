package entity

//Configuration default to application
type Configuration struct {
	DB db
}

type db struct {
	Engine   string
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

package helper

var connection string

func init() {
	connection = "MYSQL"
}

func GetConnection() string {
	return connection
}

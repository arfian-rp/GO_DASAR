package db

var connection string

func init() { //selalu dijalankan apabila package diakses (diimport pertama kali)
	println("Package Initialization dijalankan")
	connection = "HOST:PORT/MYDB/MYTABLE"

}

func GetConnection() string {
	return connection //"HOST:PORT/MYDB/MYTABLE"
}

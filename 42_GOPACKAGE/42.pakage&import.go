package main

import (
	_ "arfian-id.web.app/pkgs/db" //blank identifier => agar menjalankan init function tanpa memanggil function lain
	"arfian-id.web.app/pkgs/utils"
)

func main() {
	// println(db.GetConnection())
	utils.SayHello("Udeen")

}

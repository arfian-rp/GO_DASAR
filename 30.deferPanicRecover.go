package main

func main() {
	/*
		defer => func yg bisa dijadwalkan untuk dieksekusi setelah func fieksekusi.
		defer selalu dieksekusi meskipun terjadi error di funct sebelumnya.

		panic => function yg digunakan untuk menghentikan program.
		panic biasanya dipanggil ketika terjadi error.
		saat panic dipanggil program akan berhenti tetapi defer tetap dijalankan

		recover => function untuk menangkap data panic.
		dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
	*/
	runApp(false)
}

func logging() {
	message := recover()
	if message != nil {
		println("terjadi error", message)
	}
	println("selesai memanggil function")
}

func runApp(error bool) {
	defer logging() //dieksekusi setelah runApp() jalan walaupun terdapat error di runApp()
	if error {
		panic("App Error!") //defer tetap berjalan, code dibawah berhenti dijalankan
	}
	println("run app")
}

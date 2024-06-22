package helper

//bisa diakses dari luar package
func SayHello(name string) string {
	return "Hello " + name
}

// tidak bisa diakses dari luar package namun bisa diakses dari package yg sama
func sayHello(name string) string {
	return "Hello " + name
}

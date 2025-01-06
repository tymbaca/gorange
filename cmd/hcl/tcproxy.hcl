group {
	port = 8080
    strategy = "random"

    target { 
        addr = "localhost:8090" 
        primary = true 
    }
    target { addr = "localhost:8091" }
    target { addr = "localhost:8092" }
}

package server

func Start(config []byte) {
	router := setRouter(config)

	// Start listening and serving requests
	router.Run(":8080")
}

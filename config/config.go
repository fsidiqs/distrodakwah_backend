package config

var (
	PORT      = 0
	JWTSECRET []byte
)

func Load() {
	// var err error
	// err = godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	// if err != nil {
	// 	PORT = 9000
	// }

	// JWTSECRET = []byte(os.Getenv("API_SECRET"))

}

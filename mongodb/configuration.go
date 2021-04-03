package mongodb

type MongoDBConfiguration struct {
	Name     string `json:"name"`
	URI      string `json:"uri"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const LOCAL_DB_NAME = "localDb"
const LOCAL_DB_URI = "mongodb://127.0.0.1:27017"
const LOCAL_DB_USERNAME = "localuser"
const LOCAL_DB_PASSWORD = "1234abcd"

var MongoDBConfigurations = &[]MongoDBConfiguration{
	{
		Name:     LOCAL_DB_NAME,
		URI:      LOCAL_DB_URI,
		Username: LOCAL_DB_USERNAME,
		Password: LOCAL_DB_PASSWORD,
	},
}

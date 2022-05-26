package main

func main() {
	a := App{}
	a.Initialize("root", "pw", "db")
//    a.Initialize(
//        os.Getenv("APP_DB_USERNAME"),
//        os.Getenv("APP_DB_PASSWORD"),
//        os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}

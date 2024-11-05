package main

func main() {
	app := App{}
	app.Initialize(DbUser, DbPassword, DatabaseName)
	app.Run("localhost:10000")
}

package database

func getSourcePath() string {
	const sourcePath = "root:1q2w3e4r@tcp(127.0.0.1:3306)/plane_ticket?parseTime=true"
	return sourcePath
}

func getConnectionErrorMsg() string {
	const dbConnectionErrorMsg = "Database connection fail"
	return dbConnectionErrorMsg
}
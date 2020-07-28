package cmd


// avoid global variables

//var (
//	router = gin.Default()
//)

func StartApplication() error {
	router := MapUrls()
	if err := router.Run(":8081"); err != nil {
		return err
	}
	return nil
}

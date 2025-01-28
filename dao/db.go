package dao

func InitDB() error {
	dsn := "root:072231@tcp(127.0.0.1:3306)/first?charset=utf8&parseTime=True&loc=Local"
	if err := InitUsers(dsn); err != nil {
		return err
	}
	if err := InitProduct(dsn); err != nil {
		return err
	}
	if err := InitCart(dsn); err != nil {
		return err
	}
	return nil
}

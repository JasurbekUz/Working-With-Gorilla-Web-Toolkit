package main

import (

	"fmt"
)

const (

	Host = "localhost"
	User = "jasurbek"
	Password = "1001"
	Database = "classifieds"
	Port = 5432
)

var DB_CONFIG = fmt.Sprintf(
	"host=%s user=%s password=%s database=%s port=%d",
	Host, User, Password, Database, Port,
)
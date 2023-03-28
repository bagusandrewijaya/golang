package config
 
import (
    "database/sql"
)
 
func Connect() *sql.DB {
    dbDriver := "btzpwcfdumjrpj2zypou-mysql.services.clever-cloud.com"
    dbUser := "ucsfsfw7mdnec0ff"
    dbPass := "8vKxAEnExNyZxowgb6E3"
    dbName := "btzpwcfdumjrpj2zypou"
 
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}
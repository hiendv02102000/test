package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func init() {
	_ = flag.String("dirMigration", "/test/db/migrations", "DIR migration file")
	flag.Parse()
}
func main() {

	fmt.Println(GetStringFlag("dirMigration"))
	db, _ := sql.Open("mysql", "go_test:go_test@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local")
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///"+"/test/db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = m.Steps(2)
	if err != nil {
		currentVer, dirty, _ := m.Version()
		if dirty {
			if currentVer == 1 {
				_ = m.Force(-1)
			} else {
				_ = m.Force(int(currentVer) - 1)
			}
		}
		fmt.Println(err)
	}
}
func GetStringFlag(name string) string {
	return flag.Lookup(name).Value.(flag.Getter).Get().(string)
}

package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/snowman-mh/go-sample/src/config"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	if err := os.Chdir("/go/src/github.com/snowman-mh/go-sample/db/migrations"); err != nil {
		panic(err)
	}
	db = createDBConnection()
	clearDB(db)
	migrateDB(db)
	os.Exit(m.Run())
}

func createDBConnection() *gorm.DB {
	db, err := gorm.Open(config.Get().DB.Driver, config.Get().DB.DSN())
	if err != nil {
		panic(err)
	}
	return db
}

func clearDB(db *gorm.DB) {
	rows, err := db.Raw("SHOW TABLES").Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			panic(err)
		}
		if err := db.DropTable(tableName).Error; err != nil {
			panic(err)
		}
	}
}

func migrateDB(db *gorm.DB) {
	files, err := findMigrationFiles("./", regexp.MustCompile(`^\d.*\.sql$`))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		content := readFile(file)
		executeSQL(db, content)
	}
}

func findMigrationFiles(dir string, re *regexp.Regexp) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("could not find migrations directory `%s`", dir)
	}

	matches := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if !re.MatchString(name) {
			continue
		}

		matches = append(matches, name)
	}

	sort.Strings(matches)
	return matches, nil
}

func readFile(file string) string {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func executeSQL(db *gorm.DB, content string) {
	queries := strings.Split(strings.TrimSpace(content), ";")
	for _, query := range queries {
		if strings.Contains(query, "migrate:down") {
			break
		}
		if len(query) <= 0 {
			continue
		}
		if err := db.Exec(query).Error; err != nil {
			panic(err)
		}
	}
}

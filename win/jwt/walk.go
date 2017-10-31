package jwt

import (
	"database/sql"
	"flag"
	"log"

	_"github.com/go-sql-driver/mysql"
	"fmt"
)

func Get() error {
	return nil
}

func OpenSql() *sql.DB {

	dbConn := flag.String("db", "root:chenghai3c@tcp(127.0.0.1:3306)/server_group?charset=utf8&collation=utf8_general_ci&parseTime=true", "MySQL DB path")
	//dbConn := flag.String("db", "root:123456789@/ERP?charset=utf8", "MySQL DB path")

	flag.Parse()

	db, err := sql.Open("mysql", *dbConn)
	if err != nil {
		log.Panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
	log.Println("conn success.")
	return db
}

func FindDb(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM erp_tracks e WHERE e.status = 6")
	if err != nil {
		log.Println(err)
	}
	if rows.Next() {
		cols, err := rows.Columns()
		if err != nil {
			log.Println(err)
		}
		for i, col := range cols {
			log.Printf("col-%d :%v", i+1, col)
		}
		types, err := rows.ColumnTypes()
		if err != nil {
			log.Println(err)
		}
		for _, t := range types {
			null, _ := t.Nullable()
			l, _ := t.Length()
			d, e, _ := t.DecimalSize()
			log.Printf("type// 0--Name :%v, 1--dbTypeName :%v, 2--size :%v ~ %v, 3--len :%v, 4--isNull :%v, 5--scanType :%v", t.Name(), t.DatabaseTypeName(), d, e, l, null, t.ScanType())
		}
	}
	log.Println(liShi)
	return nil
}

var liShi = `
	     ____      _______
        /   /     /  _____\
	   /   /     /  /
	  /	  /     /  /______
	 /	 /      \_____   /
	/	/            /  /
	\	\____  _____/  /
	 \______/  \______/
`

func PrintLs() {
	as := [12][20]int{
		{0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
	}
	for i := 0; i < len(as); i++ {
		for k := 0; k < len(as[i]); k++ {
			if as[i][k] == 1 {
				fmt.Printf("%v", "*")
			} else {
				fmt.Printf("%v", " ")
			}
		}
		fmt.Println()
	}
}
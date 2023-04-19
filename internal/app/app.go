package app

import (
	"encoding/xml"
	"flag"
	"fmt"
	pssql "go-xml-parser/internal/database/postgresql"
	. "go-xml-parser/internal/models"
	"os"
	"time"
)

const addressObjectElementName = "ADDRESSOBJECTS"
const objectElementName = "OBJECT"

const paramsElementName = "PARAMS"
const paramElementName = "PARAM"
const dateConst = "2006-01-02"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Run(cfg Config) {

	fileName := flag.String("file", "somthing.xml", "")
	table := flag.String("table", "somthing", "")
	flag.Parse()
	fmt.Println(os.Args)

	var (
		object Object
		param  Param
	)
	currentTime := time.Now()
	fmt.Println("Start programm: ", currentTime.Format("2006-01-02 15:04:05.000"))

	//	xmlFile, err := os.Open("AS_HOUSES_PARAMS_20230306_0260d3dc-9456-4cde-b970-2a6446414dd7.xml")
	xmlFile, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	d := xml.NewDecoder(xmlFile)

	var sqldb pssql.Db
	err = sqldb.Init(cfg)
	if err != nil {
		os.Exit(1)
	}

	defer sqldb.Close()

	sqldb.ExecQuery("TRUNCATE TABLE " + *table)

	currentTime = time.Now()
	fmt.Println("Start parse: ", currentTime.Format("2006-01-02 15:04:05.000"))
	count := 0
	var execQuery string
	for t, _ := d.Token(); t != nil; t, _ = d.Token() {

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == addressObjectElementName {
				execQuery = pssql.AddressObjectQuery
			} else if se.Name.Local == paramsElementName {
				execQuery = pssql.ParamsQuery
			} else if se.Name.Local == objectElementName {
				d.DecodeElement(&object, &se)
				sqldb.QueryBatch(execQuery, object.ID, object.OBJECTGUID)
				if count < 50000 {
					count++
				} else {
					sqldb.ExecBatch()
					count = 0
				}
			} else if se.Name.Local == paramElementName {
				d.DecodeElement(&param, &se)
				udate, _ := time.Parse(dateConst, param.UPDATEDATE)
				sdate, _ := time.Parse(dateConst, param.STARTDATE)
				edate, _ := time.Parse(dateConst, param.ENDDATE)
				sqldb.QueryBatch(execQuery, param.ID, param.OBJECTID, udate, sdate, edate)
				if count < 10000 {
					count++
				} else {
					sqldb.ExecBatch()
					count = 0
				}
			}
		}
	}
	sqldb.ExecBatch()

	currentTime = time.Now()
	fmt.Println("End parse: ", currentTime.Format("2006-01-02 15:04:05.000"))
}

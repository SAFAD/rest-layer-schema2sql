package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/rs/rest-layer/schema"
)

/*
This package is intended to convert rest-layer/schema to an SQL Compliant (CREATE Statements)

first we map all the column types in SQL
second we map all field types in REST Layer
third we create command line arguments

*/

//SQLColumns maps all SQL column types
type SQLColumns struct {
	bigint          int8      //signed eight-byte integer
	bigserial       int8      //serial8	autoincrementing eight-byte integer
	bit             string    //[ (n) ]	 	fixed-length bit string
	bitvarying      string    //[ (n) ]	varbit	variable-length bit string
	boolean         bool      //logical Boolean (true/false)
	box             string    //rectangular box on a plane wtf?
	bytea           []byte    //binary data ("byte array")
	char            string    //[ (n) ]	char [ (n) ]	fixed-length character string
	varchar         string    // [ (n) ]	varchar [ (n) ]	variable-length character string
	cidr            string    // IPv4 or IPv6 network address TODO use net/ip type
	circle          string    //circle on a plane wtf??
	date            time.Time //calendar date (year, month, day)
	doubleprecision float64   //double precision floating-point number (8 bytes) TODO make it float8
	inet            string    //IPv4 or IPv6 host address TODO use net/ip type
	integer         int       //int4	signed four-byte integer
	interval        time.Time //[ fields ] [ (p) ]	 	time span
	json            string    //JSON data TODO use encoding/json type
	line            string    //infinite line on a plane wtf???
	lseg            string    //line segment on a plane wtf????
	macaddr         string    //MAC (Media Access Control) address TODO use hardwareAddr from net/mac
	money           int64     //currency amount
	numeric         int64     //[ (p, s) ]	decimal [ (p, s) ]	exact numeric of selectable precision
	path            string    // geometric path on a plane wtf?????
	point           string    // 	geometric point on a plane wtf??????
	polygon         string    // 	closed geometric path on a plane wtf???????
	real            float64   // single precision floating-point number (4 bytes)
	smallint        int       //signed two-byte integer
	smallserial     int       //serial2	autoincrementing two-byte integer
	serial          int       //serial4	autoincrementing four-byte integer
	text            string    // 	variable-length character string
	time            time.Time //[ (p) ] [ without time zone ]	 	time of day (no time zone)
	timestamp       time.Time //[ (p) ] with time zone	timestamptz	date and time, including time zone
	tsquery         string    // text search query
	tsvector        string    //text search document
	txid_snapshot   string    //user-level transaction ID snapshot TODO further research this
	uuid            string    // universally unique identifier
	xml             string    //XML data TODO use encoding/xml type
}

//RESTLayerFields maps all REST Layer Fields
type RESTLayerFields struct {
	String    schema.String    // Ensures the field is a string
	Integer   schema.Integer   // Ensures the field is an integer
	Float     schema.Float     // Ensures the field is a float
	Bool      schema.Bool      // Ensures the field is a Boolean
	Array     schema.Array     // Ensures the field is an array
	Dict      schema.Dict      // Ensures the field is a dict
	Object    schema.Object    // Ensures the field is an object validating against a sub-schema
	Time      schema.Time      // Ensures the field is a datetime
	URL       schema.URL       // Ensures the field is a valid URL
	IP        schema.IP        // Ensures the field is a valid IPv4 or IPv6
	Password  schema.Password  // Ensures the field is a valid password and bcrypt it
	Reference schema.Reference // Ensures the field contains a reference to another existing API item
	AnyOf     schema.AnyOf     // Ensures that at least one sub-validator is valid
	AllOf     schema.AllOf     // Ensures that at least all sub-validators are valid
}

func prepareQuery() {

}

func main() {
	fmt.Println("Welcome to schema2SQL CLI")

	flag.String("file", "schema.go", "Schema file location DEFAULT: schema.go")
	flag.Bool("drop", true, "Include DROP IF EXISTS in Query DEFAULT: True")

	flag.Parse()

	prepareQuery()
}

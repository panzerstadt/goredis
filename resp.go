package main

// We categorize every RESP data type as either simple, bulk or aggregate.
const (
	// resp 2
	STRING  = '+' // simple
	ERROR   = '-' // simple
	INTEGER = ':' // simple
	BULK    = '$' // aggregate (bulk strings)
	ARRAY   = '*' // aggregate

	// resp3
	NULL            = '_' // simple
	BOOLEAN         = '#' // simple
	DOUBLE          = ',' // simple
	BIG_NUMBER      = '(' // simple
	BULK_ERROR      = '!' // aggregate
	VERBATIM_STRING = '=' // aggregate
	MAP             = '%' // aggregate
	ATTR            = '`' // aggregate
	SET             = '~' // aggregate
	PUSH            = '>' // aggregate
)

type Value struct {
	typ   string
	str   string // simple strings
	num   int    // integers
	bulk  string // bulk strings
	array []Value
}

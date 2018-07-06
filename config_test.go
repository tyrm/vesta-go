package main

import "testing"

func TestDecodeEngine(t *testing.T) {
	conStrings := []struct {
		connectionString string
		dialect string
		args string
	}{
		{"postgresql://auser:Th31rPassw0rd!@myserver.test.com/thedatas","postgres", "host=myserver.test.com user=auser dbname=thedatas password=Th31rPassw0rd!"},
	}

	for _, conString := range conStrings {
		dialect, args := DecodeEngine(conString.connectionString)
		if dialect != conString.dialect {
			t.Errorf("Decoding dialect of (%s) was incorrect, got: %s, want: %s.", conString.connectionString, dialect, conString.dialect)
		}
		if args != conString.args {
			t.Errorf("Decoding dialect of (%s) was incorrect, got: %s, want: %s.", conString.connectionString, args, conString.args)
		}
	}
}
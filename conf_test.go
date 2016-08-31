package conf

import "testing"
func TestGetInt(t *testing.T){
	v, ok := GetInt("testInt")
	if !ok {
		t.Error("Expected GetInt OK, got NOK")
	}
	if v != 123 {
		t.Error("Expected GetInt = 123, got ", v)
	}
}

func TestGetFloat(t *testing.T){
	v, ok := GetFloat("testFloat")
	if !ok {
		t.Error("Expected GetFloat OK, got NOK")
	}
	if v != 123.123 {
		t.Error("Expected GetFloat = 123.123, got ", v)
	}
}

func TestGetBool(t *testing.T){
	v, ok := GetBool("testBool")
	if !ok {
		t.Error("Expected GetBool OK, got NOK")
	}
	if !v {
		t.Error("Expected GetBool = true, got ", v)
	}
}

func TestGetString(t *testing.T){
	v, ok := GetString("testString")
	if !ok {
		t.Error("Expected GetString OK, got NOK")
	}
	if v != "asdf" {
		t.Error("Expected GetString = asdf, got ", v)
	}
}

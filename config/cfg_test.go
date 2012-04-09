package config

import (
	"fmt"
	"testing"
)

func TestGetValue(t *testing.T){
	fmt.Println(GetValue("app.vendor"))
}

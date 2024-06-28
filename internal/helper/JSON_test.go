package helper_test

import (
	"applicant-backend/internal/helper"
	"testing"
)

type test struct {
	Name string
	Age  int
}

func TestStructToJsonToString(t *testing.T) {
	testInstance := test{Age: 10, Name: "TestName"}
	str, err := helper.StructToJsonToString(testInstance)
	if err != nil || str != "{\"Name\":\"TestName\",\"Age\":10}" {
		t.Fatalf("Expected: %s - Got: %s", "{\"Name\":\"TestName\",\"Age\":10}", str)
	}

}

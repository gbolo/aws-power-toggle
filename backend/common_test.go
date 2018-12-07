package backend

import "testing"

func TestComputeId(t *testing.T) {

	// test cases for case sensitivity
	testStringCase := map[string]string{
		"testString": "956265657d0b",
		"teststring": "b8473b86d4c2",
	}

	// test case for multiple strings
	testMultipleInputs1 := "teststring1"
	testMultipleInputs2 := "teststring2"
	testMultipleInputsExpectedId := "ab8675367360"

	for input, expectedResult := range testStringCase {
		if ComputeId(input) != expectedResult {
			t.Errorf("%s != %s - got: %s", input, expectedResult, ComputeId(input))
		}
	}

	if ComputeId(testMultipleInputs1, testMultipleInputs2) != testMultipleInputsExpectedId {
		t.Error("testMultipleInputsExpectedId did not match")
	}
}

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
	testMultipleInputsExpectedID := "ab8675367360"

	for input, expectedResult := range testStringCase {
		if ComputeID(input) != expectedResult {
			t.Errorf("%s != %s - got: %s", input, expectedResult, ComputeID(input))
		}
	}

	if ComputeID(testMultipleInputs1, testMultipleInputs2) != testMultipleInputsExpectedID {
		t.Error("testMultipleInputsExpectedID did not match")
	}
}

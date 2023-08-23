package entities

import (
	"testing"
)

func TestIsCategoryValid(t *testing.T) {
	testCases := []struct {
		title          string
		category       string
		expectedResult bool
	}{
		{title: "given_category_bebida_should_return_true", category: "BEBIDA", expectedResult: true},
		{title: "given_category_lanche_should_return_true", category: "LANCHE", expectedResult: true},
		{title: "given_category_sobremesa_should_return_true", category: "SOBREMESA", expectedResult: true},
		{title: "given_category_acompanhamento_should_return_true", category: "ACOMPANHAMENTO", expectedResult: true},
		{title: "given_invalid_category_should_return_false", category: "invalid_category", expectedResult: false},
	}

	for _, testCase := range testCases {
		item := &Item{Category: testCase.category}
		actualResult := item.IsCategoryValid()
		if actualResult != testCase.expectedResult {
			t.Fatalf("TestIsCategoryValid: %s\n\texpected: %t\n\tgot: %t", testCase.title, testCase.expectedResult, actualResult)
		}
	}
}

package powerdns_test

import (
	"fmt"
	"testing"

	"github.com/bishy999/go-powerdns/pkg/powerdns"
)

func TestCheckUserInput(t *testing.T) {

	tt := []struct {
		name      string
		flagValue string
		result    string
	}{
		{"Test 001 correct input", "test01", ""},
	}

	for _, tc := range tt {
		fmt.Println(tc.name)
		t.Run(tc.name, func(t *testing.T) {

			input, err := powerdns.CheckUserInput()
			fmt.Println(input)
			if err != nil {
				if err.Error() != tc.result {
					t.Errorf("Test %v result should be `%v`, got  `%v`", tc.name, tc.result, err)
				}
			}

		})
	}

}

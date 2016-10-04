package results

import (
	"testing"
)

const expectedWrong = "ожидалось %v, фактически %v"

func TestNumber(t *testing.T) {
	testValue := float64(42)
	var n Number = Number(testValue)

	// GetTotal
	if fact := n.GetTotal(); fact != testValue {
		t.Fatalf(expectedWrong, testValue, fact)
	}

	// GetDetails
	if n.GetDetails() != nil {
		t.Fatalf("%T не должен содержать деталей", n)
	}

	// проверка, что Number реализует MethodResult
	var _ MethodResult = n
}

func TestNumberSum(t *testing.T) {
	one, two, three := float64(100500), 3.5, float64(42)
	expectedSum := one + two + three
	n := NumberSum{"one":one, "two": two, "three": three}

	// GetTotal
	if fact := n.GetTotal(); fact != expectedSum {
		t.Fatalf(expectedWrong, expectedSum, fact)
	}

	// GetDetails
	for key, value := range n.GetDetails() {
		switch key {
		case "one":
			if value != one {
				t.Fatalf(expectedWrong, one, value)
			}
		case "two":
			if value != two {
				t.Fatalf(expectedWrong, two, value)
			}
		case "three":
			if value != three {
				t.Fatalf(expectedWrong, three, value)
			}
		default:
			t.Fatalf("неизвестный ключ %v со значением %v", key, value)
		}
	}

	// проверка, что NumberSum реализует MethodResult
	var _ MethodResult = n
}


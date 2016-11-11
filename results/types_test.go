package results

import (
	"testing"
)

const expectedWrong = "ожидалось %v, фактически %v"

func TestNumberList_GetSortedKeys(t *testing.T) {
	n := NumberList{"z":1, "a": 4, "c": 32, "d": 58, "b":2}

	sorted := ""
	for _, s := range *n.GetSortedKeys() {
		sorted += s
	}

	if sorted != "abcdz" {
		t.Fatal(expectedWrong, "abcdz", sorted)
	}
}

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

func TestNumberSum_GetSortedKeys(t *testing.T) {
	rightKeys := []string{"a", "b", "c", "d", "e"}
	n := NumberSum{rightKeys[4]: 5.4,
		rightKeys[0]: 15.0,
		rightKeys[2]: 23.2,
		rightKeys[1]: 1.1,
		rightKeys[3]: 57.3, }
	testKeys := *n.GetDetails().GetSortedKeys()
	for i := 0; i < 5; i++ {
		if rightKeys[i] != testKeys[i] {
			t.Errorf(expectedWrong, rightKeys[i], testKeys[i])
		}
	}
}

func TestNumberSum_GetSortedKeys_Empty(t *testing.T) {
	n := NumberSum{}
	testKeys := *n.GetDetails().GetSortedKeys()
	if count := len(testKeys); count != 0 {
		t.Errorf(expectedWrong, 0, count)
	}
}


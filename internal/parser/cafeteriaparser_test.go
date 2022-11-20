package parser

import (
	"os"
	"testing"
)

// TestParseCafeteria tests the ParseCafeteria function.
// The test data can be found in the meal_plan.html file.
// The test data is a sample of the HTML returned by mySu.
func TestParseCafeteria(t *testing.T) {
	var mealHTML, err = os.Open("testdata/cafeteria.html")
	if err != nil {
		t.Fatal(err)
	}

	cafeteria, err := ParseCafeteria(mealHTML)
	if err != nil {
		t.Fatal(err)
	}

	if len(cafeteria.Menu) != 17 {
		t.Errorf("len(cafeteria.Menu) = %d; want 17", len(cafeteria.Menu))
	}
	// Check the first item in the menu.
	if cafeteria.Menu[0].Name != "TARHANA (TRADITIONAL) SOUP" {
		t.Errorf("cafeteria.Menu[0].Name = %s; want TARHANA (TRADITIONAL) SOUP", cafeteria.Menu[0].Name)
	}
	if cafeteria.Menu[0].Calories != "123" {
		t.Errorf("cafeteria.Menu[0].Calories = %s; want 123", cafeteria.Menu[0].Calories)
	}
	// Check the last item in the menu.
	if cafeteria.Menu[16].Name != "AYRAN" {
		t.Errorf("cafeteria.Menu[16].Name = %s; want AYRAN", cafeteria.Menu[16].Name)
	}
	if cafeteria.Menu[16].Calories != "64" {
		t.Errorf("cafeteria.Menu[16].Calories = %s; want 64", cafeteria.Menu[16].Calories)
	}
}

package data

// Cafeteria represents the cafeteria menu.
// It contains a slice of Food structs.
type Cafeteria struct {
	Menu []Food
}

// Food represents a food item in the cafeteria menu.
type Food struct {
	Name     string
	Calories string
}

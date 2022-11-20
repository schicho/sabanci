package data

// Class represents one continous timeblock in the schedule.
type Class struct {
	Name      string
	ClassCode string
	Building  string
	TimeStart string
	TimeEnd   string
}

// Schedule wraps an array indexed by the day of the week.
// Each element contains the classes on that day.
// The 0th element is Monday, the 6th element is Sunday.
type Schedule struct {
	Classes [7][]Class
}

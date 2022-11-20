package parser

import (
	"os"
	"testing"
)

// TestParseSchedule tests the ParseSchedule function.
// It uses the testdata/schedule.html file as test data.
func TestParseSchedule(t *testing.T) {
	file, err := os.Open("testdata/schedule.html")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	schedule, err := ParseSchedule(file)
	if err != nil {
		t.Fatal(err)
	}

	// That's kind of obvious.
	if len(schedule.Classes) != 7 {
		t.Errorf("len(schedule.Classes) = %d; want 7", len(schedule.Classes))
	}

	// Check the first class on Monday.
	// It is my online IoT class. It starts at 09:40 and ends at 11:30.
	monday := t.Run("Monday's first class", func(t *testing.T) {

		if schedule.Classes[0][0].Name != "Special Topics in CS: Internet of Things Sensing System" {
			t.Errorf("schedule.Classes[0][0].Name = %s; want Special Topics in CS: Internet of Things Sensing System", schedule.Classes[0][0].Name)
		}
		if schedule.Classes[0][0].TimeStart != "09:40" {
			t.Errorf("schedule.Classes[0][0].Start = %s; want 09:40", schedule.Classes[0][0].TimeStart)
		}
		if schedule.Classes[0][0].TimeEnd != "11:30" {
			t.Errorf("schedule.Classes[0][0].End = %s; want 11:30", schedule.Classes[0][0].TimeEnd)
		}
		if schedule.Classes[0][0].ClassCode != "CS48007 - 0" {
			t.Errorf("schedule.Classes[0][0].ClassCode = %s; want CS48007 - 0", schedule.Classes[0][0].ClassCode)
		}
	})
	if !monday {
		t.Errorf("Parsing Monday's first class failed.")
	}

	// The second class on Monday is my Cybersecurity CS 437 class.
	// It starts at 14:40 and ends at 15:30. It's in classroom FENS L045.
	monday = t.Run("Monday's second class", func(t *testing.T) {
		if schedule.Classes[0][1].Name != "Cybersecurity Practices and Applications" {
			t.Errorf("schedule.Classes[0][1].Name = %s; want Cybersecurity Practices and Applications", schedule.Classes[0][1].Name)
		}
		if schedule.Classes[0][1].TimeStart != "14:40" {
			t.Errorf("schedule.Classes[0][1].Start = %s; want 14:40", schedule.Classes[0][1].TimeStart)
		}
		if schedule.Classes[0][1].TimeEnd != "15:30" {
			t.Errorf("schedule.Classes[0][1].End = %s; want 15:30", schedule.Classes[0][1].TimeEnd)
		}
		if schedule.Classes[0][1].ClassCode != "CS437 - 0" {
			t.Errorf("schedule.Classes[0][1].ClassCode = %s; want CS437 - 0", schedule.Classes[0][1].ClassCode)
		}
	})
	if !monday {
		t.Errorf("Parsing Monday's second class failed.")
	}
}

func TestCorrectNumberOfClasses(t *testing.T) {
	file, err := os.Open("testdata/schedule.html")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	schedule, err := ParseSchedule(file)
	if err != nil {
		t.Fatal(err)
	}

	// 3 classes on Monday, 1 on Tuesday, 2 on Wednesday, 2 on Thursday, 0 on Friday, 0 on Saturday, 0 on Sunday.
	if len(schedule.Classes[0]) != 3 {
		t.Errorf("len(schedule.Classes[0]) = %d; want 3", len(schedule.Classes[0]))
	}
	if len(schedule.Classes[1]) != 1 {
		t.Errorf("len(schedule.Classes[1]) = %d; want 1", len(schedule.Classes[1]))
	}
	if len(schedule.Classes[2]) != 2 {
		t.Errorf("len(schedule.Classes[2]) = %d; want 2", len(schedule.Classes[2]))
	}
	if len(schedule.Classes[3]) != 2 {
		t.Errorf("len(schedule.Classes[3]) = %d; want 2", len(schedule.Classes[3]))
	}
	if len(schedule.Classes[4]) != 0 {
		t.Errorf("len(schedule.Classes[4]) = %d; want 0", len(schedule.Classes[4]))
	}
	if len(schedule.Classes[5]) != 0 {
		t.Errorf("len(schedule.Classes[5]) = %d; want 0", len(schedule.Classes[5]))
	}
	if len(schedule.Classes[6]) != 0 {
		t.Errorf("len(schedule.Classes[6]) = %d; want 0", len(schedule.Classes[6]))
	}
}

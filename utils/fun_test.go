package utils

import (
	"testing"
)

func TestGetRandomActivity(t *testing.T) {
	// Test that GetRandomActivity returns a valid activity
	activity := GetRandomActivity()

	if activity == "" {
		t.Error("GetRandomActivity() returned empty string")
	}

	// Verify the returned activity is in the activities array
	found := false
	for _, validActivity := range activities {
		if activity == validActivity {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("GetRandomActivity() = %s; not found in activities array", activity)
	}
}

func TestGetRandomActivityVariety(t *testing.T) {
	// Test that GetRandomActivity returns different values over multiple calls
	// This is a probabilistic test, so we'll call it many times
	const iterations = 100
	results := make(map[string]bool)

	for i := 0; i < iterations; i++ {
		activity := GetRandomActivity()
		results[activity] = true
	}

	// We should get at least 2 different activities in 100 calls
	// (Probability of getting only 1 is extremely low with 14 activities)
	if len(results) < 2 {
		t.Errorf("GetRandomActivity() returned only %d unique value(s) in %d calls; expected variety", len(results), iterations)
	}
}

func TestActivitiesArrayNotEmpty(t *testing.T) {
	// Verify the activities array is not empty
	if len(activities) == 0 {
		t.Error("activities array is empty")
	}
}

func TestActivitiesArrayContent(t *testing.T) {
	// Verify all activities are non-empty strings
	for i, activity := range activities {
		if activity == "" {
			t.Errorf("activities[%d] is empty string", i)
		}
	}

	// Verify we have the expected number of activities
	expectedCount := 15
	if len(activities) != expectedCount {
		t.Errorf("activities array length = %d; want %d", len(activities), expectedCount)
	}
}

func TestActivitiesArrayValues(t *testing.T) {
	// Test that specific expected activities are present
	expectedActivities := []string{
		"Cooking",
		"Bamboozling",
		"Brewing",
		"Scheming",
		"Conjuring",
		"Crunching",
		"Calibrating",
		"Spinning",
		"Summoning",
		"Plotting",
		"Whirring",
		"Overthinking",
		"Vibing",
		"Meddling",
		"Fiddling",
	}

	for _, expected := range expectedActivities {
		found := false
		for _, activity := range activities {
			if activity == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected activity %s not found in activities array", expected)
		}
	}
}

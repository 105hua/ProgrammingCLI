package utils

import "math/rand/v2"

// activities
// An array of strings containing what the model can say it is doing.
// See GetRandomActivity.
var activities = [...]string{
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

// GetRandomActivity
// Fetches a random activity to print out instead of simply saying "Thinking..."
func GetRandomActivity() string {
	return activities[rand.IntN(len(activities))]
}

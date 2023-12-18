package onosclient

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestParseIntent(t *testing.T) {
	intent, err := os.ReadFile("testdata/intent.json")
	if err != nil {
		t.Fatal(err)
	}
	var tests = []struct {
		AppID       string
		ID          string
		Key         string
		State       string
		Type        string
		Resources   []string
		Selector    *Selector
		Treatment   *Treatment
		Priority    int
		Constraints []Constraints
		One         string
		Two         string
	}{
		{
			AppID:     "org.onosproject.cli",
			ID:        "0x300154",
			Key:       "0x100005",
			State:     "FAILED",
			Type:      "HostToHostIntent",
			Resources: []string{"00:00:00:00:00:01/None", "00:00:00:00:00:99/None"},
			//Selector:    &Selector
			//Treatment:   *Treatment
			Priority: 100,
			//Constraints: []Constraints
			One: "00:00:00:00:00:01/None",
			Two: "00:00:00:00:00:99/None",
		}, /*
			{
				AppID:     "shouldfail",
				ID:        "shouldfail",
				Key:       "shouldfail",
				State:     "shouldfail",
				Type:      "shouldfail",
				Resources: []string{"shouldfail", "shouldfail"},
				//Selector:    &Selector
				//Treatment:   *Treatment
				Priority: 99,
				//Constraints: []Constraints
				One: "shouldfail",
				Two: "shouldfail",
			}, */ // add more structs here to introduce further tests.
		// currently not testing treatment, selector, or constraints as the're not used anywhere
	}

	got, err := ParseIntent(intent)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.Key)
		t.Run(testname, func(t *testing.T) {
			if got.AppID != tt.AppID {
				t.Errorf("Got AppID: %s,, wanted: %s", got.AppID, tt.AppID)
			} else {
				t.Logf("AppID Passed with Value: %s", tt.AppID)
			}
			if got.ID != tt.ID {
				t.Errorf("Got ID: %s,, wanted: %s", got.ID, tt.ID)
			} else {
				t.Logf("ID Passed with Value: %s", tt.ID)
			}
			if got.Key != tt.Key {
				t.Errorf("Got Key: %s,, wanted: %s", got.Key, tt.Key)
			} else {
				t.Logf("Key Passed with Value: %s", tt.Key)
			}
			if got.State != tt.State {
				t.Errorf("Got State: %s,, wanted: %s", got.State, tt.State)
			} else {
				t.Logf("State Passed with Value: %s", tt.State)
			}
			if got.Type != tt.Type {
				t.Errorf("Got Type: %s,, wanted: %s", got.Type, tt.Type)
			} else {
				t.Logf("Type Passed with Value: %s", tt.Type)
			}
			if !reflect.DeepEqual(got.Resources, tt.Resources) {
				t.Errorf("Got Resources: %q,, wanted: %q", got.Resources, tt.Resources)
			} else {
				t.Logf("Resources Passed with Value: %q", tt.Resources)
			}
			if got.Priority != tt.Priority {
				t.Errorf("Got Priority: %d,, wanted: %d", got.Priority, tt.Priority)
			} else {
				t.Logf("AppID Passed with Value: %d", tt.Priority)
			}
			if got.One != tt.One {
				t.Errorf("Got One: %s,, wanted: %s", got.One, tt.One)
			} else {
				t.Logf("One Passed with Value: %s", tt.One)
			}
			if got.Two != tt.Two {
				t.Errorf("Got Two: %s,, wanted: %s", got.Two, tt.Two)
			} else {
				t.Logf("Two Passed with Value: %s", tt.Two)
			}
		})
	}
}

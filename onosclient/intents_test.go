package onosclient

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

// Reference: https://bitfieldconsulting.com/golang/api-client

func TestParseIntent_CorrectJSON(t *testing.T) {
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

func TestParseIntent_ErrOnEmpty(t *testing.T) {
	_, err := ParseIntent([]byte{})
	if err == nil {
		t.Fatal("want error parsing empty response, got nil")
	}
}

func TestParseIntents_CorrectJSON(t *testing.T) {
	intents, err := os.ReadFile("testdata/intents.json")
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		Intents []Intent
	}{
		{
			Intents: []Intent{{
				AppID:     "org.onosproject.cli",
				ID:        "0x40004f",
				Key:       "0x100005",
				State:     "INSTALLED",
				Type:      "HostToHostIntent",
				Resources: []string{"00:00:00:00:00:01/None", "00:00:00:00:00:02/None"},
			},
				{
					AppID:     "org.onosproject.cli",
					ID:        "0x40004d",
					Key:       "0x300009",
					State:     "FAILED",
					Type:      "HostToHostIntent",
					Resources: []string{"00:00:00:00:00:01/None", "00:00:00:00:00:99/None"},
				},
				{
					AppID:     "org.onosproject.cli",
					ID:        "0x40004e",
					Key:       "0x100006",
					State:     "FAILED",
					Type:      "HostToHostIntent",
					Resources: []string{"00:00:00:00:00:02/None", "00:00:00:00:00:88/None"},
				},
			},
		},
	}

	got, err := ParseIntents(intents)
	if err != nil {
		t.Fatal(err)
	}

	for _, subtest := range tests {
		if len(got.Intents) != len(subtest.Intents) {
			t.Fatal(errors.New("Incorrect number of intents in response"))
		}
		for i, tt := range subtest.Intents {
			testname := fmt.Sprintf("%s", tt.Key)
			t.Run(testname, func(t *testing.T) {
				if got.Intents[i].AppID != tt.AppID {
					t.Errorf("Got AppID: %s,, wanted: %s", got.Intents[i].AppID, tt.AppID)
				} else {
					t.Logf("AppID Passed with Value: %s", tt.AppID)
				}
				if got.Intents[i].ID != tt.ID {
					t.Errorf("Got ID: %s,, wanted: %s", got.Intents[i].ID, tt.ID)
				} else {
					t.Logf("ID Passed with Value: %s", tt.ID)
				}
				if got.Intents[i].Key != tt.Key {
					t.Errorf("Got Key: %s,, wanted: %s", got.Intents[i].Key, tt.Key)
				} else {
					t.Logf("Key Passed with Value: %s", tt.Key)
				}
				if got.Intents[i].State != tt.State {
					t.Errorf("Got State: %s,, wanted: %s", got.Intents[i].State, tt.State)
				} else {
					t.Logf("State Passed with Value: %s", tt.State)
				}
				if got.Intents[i].Type != tt.Type {
					t.Errorf("Got Type: %s,, wanted: %s", got.Intents[i].Type, tt.Type)
				} else {
					t.Logf("Type Passed with Value: %s", tt.Type)
				}
				if !reflect.DeepEqual(got.Intents[i].Resources, tt.Resources) {
					t.Errorf("Got Resources: %q,, wanted: %q", got.Intents[i].Resources, tt.Resources)
				} else {
					t.Logf("Resources Passed with Value: %q", tt.Resources)
				}

			})
		}
	}
}

func TestParseIntents_ErrOnEmpty(t *testing.T) {
	_, err := ParseIntents([]byte{})
	if err == nil {
		t.Fatal("want error parsing empty response, got nil")
	}
}

func TestGetIntent_InvalidIntent(t *testing.T) {
	c, err := NewClient("host", "username", "password")
	if err != nil {
		t.Fatal("Error creating client")
	}
	intent := Intent{
		Key: "0x300009",
	}
	_, err = c.GetIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing appid ")
	}

	intent = Intent{
		AppID: "org.onosproject.cli",
	}
	_, err = c.GetIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing key ")
	}
}

func TestCreateIntent_InvalidIntent(t *testing.T) {
	c, err := NewClient("host", "username", "password")
	if err != nil {
		t.Fatal("Error creating client")
	}

	intent := Intent{
		//Type:  "HostToHostIntent",
		AppID: "org.onosproject.cli",
		One:   "00:00:00:00:00:01/None",
		Two:   "00:00:00:00:00:99/None",
	}

	_, err = c.CreateIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing type")
	}

	intent = Intent{
		Type: "HostToHostIntent",
		//AppID: "org.onosproject.cli",
		One: "00:00:00:00:00:01/None",
		Two: "00:00:00:00:00:99/None",
	}

	_, err = c.CreateIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing appid")
	}

	intent = Intent{
		Type:  "HostToHostIntent",
		AppID: "org.onosproject.cli",
		//One:   "00:00:00:00:00:01/None",
		Two: "00:00:00:00:00:99/None",
	}

	_, err = c.CreateIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing one")
	}

	intent = Intent{
		Type:  "HostToHostIntent",
		AppID: "org.onosproject.cli",
		One:   "00:00:00:00:00:01/None",
		//Two:   "00:00:00:00:00:99/None",
	}

	_, err = c.CreateIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing two")
	}
}

func TestUpdateIntent_InvalidIntent(t *testing.T) {
	c, err := NewClient("host", "username", "password")
	if err != nil {
		t.Fatal("Error creating client")
	}

	intent := Intent{
		//Key:   "0x300009",
		Type:  "HostToHostIntent",
		AppID: "org.onosproject.cli",
		One:   "00:00:00:00:00:01/None",
		Two:   "00:00:00:00:00:99/None",
	}

	_, err = c.UpdateIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing key")
	}

	intent = Intent{
		Key: "0x300009",
		//Type:  "HostToHostIntent",
		AppID: "org.onosproject.cli",
		One:   "00:00:00:00:00:01/None",
		Two:   "00:00:00:00:00:99/None",
	}

	_, err = c.UpdateIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing type")
	}

	intent = Intent{
		Key:  "0x300009",
		Type: "HostToHostIntent",
		//AppID: "org.onosproject.cli",
		One: "00:00:00:00:00:01/None",
		Two: "00:00:00:00:00:99/None",
	}

	_, err = c.UpdateIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing appid")
	}

	intent = Intent{
		Key:   "0x300009",
		Type:  "HostToHostIntent",
		AppID: "org.onosproject.cli",
		//One:   "00:00:00:00:00:01/None",
		Two: "00:00:00:00:00:99/None",
	}

	_, err = c.UpdateIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing one")
	}

	intent = Intent{
		Key:   "0x300009",
		Type:  "HostToHostIntent",
		AppID: "org.onosproject.cli",
		One:   "00:00:00:00:00:01/None",
		//Two:   "00:00:00:00:00:99/None",
	}

	_, err = c.UpdateIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing two")
	}
}

func TestDeleteIntent_InvalidIntent(t *testing.T) {
	c, err := NewClient("host", "username", "password")
	if err != nil {
		t.Fatal("Error creating client")
	}
	intent := Intent{
		Key: "0x300009",
	}
	err = c.DeleteIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing appid")
	}

	intent = Intent{
		AppID: "org.onosproject.cli",
	}
	err = c.DeleteIntent(intent)
	if err == nil {
		t.Fatal("want invalid intent error: missing key")
	}
}

func TestGetIntent_ReturnExpectedJSON(t *testing.T) {
	want := Intent{
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
	}

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "testdata/intent.json")
		}))
	defer ts.Close()
	c, err := NewClient("host", "username", "password")
	if err != nil {
		t.Fatal("Error creating client")
	}
	c.HTTPClient = ts.Client()
	c.HostURL = ts.URL

	intent := Intent{
		AppID: "org.onosproject.cli",
		Key:   "0x100005",
	}
	got, err := c.GetIntent(intent)
	if err != nil {
		t.Fatal(err)
	}
	if got.AppID != want.AppID {
		t.Errorf("Got AppID: %s,, wanted: %s", got.AppID, want.AppID)
	} else {
		t.Logf("AppID Passed with Value: %s", want.AppID)
	}
	if got.ID != want.ID {
		t.Errorf("Got ID: %s,, wanted: %s", got.ID, want.ID)
	} else {
		t.Logf("ID Passed with Value: %s", want.ID)
	}
	if got.Key != want.Key {
		t.Errorf("Got Key: %s,, wanted: %s", got.Key, want.Key)
	} else {
		t.Logf("Key Passed with Value: %s", want.Key)
	}
	if got.State != want.State {
		t.Errorf("Got State: %s,, wanted: %s", got.State, want.State)
	} else {
		t.Logf("State Passed with Value: %s", want.State)
	}
	if got.Type != want.Type {
		t.Errorf("Got Type: %s,, wanted: %s", got.Type, want.Type)
	} else {
		t.Logf("Type Passed with Value: %s", want.Type)
	}
	if !reflect.DeepEqual(got.Resources, want.Resources) {
		t.Errorf("Got Resources: %q,, wanted: %q", got.Resources, want.Resources)
	} else {
		t.Logf("Resources Passed with Value: %q", want.Resources)
	}
	if got.Priority != want.Priority {
		t.Errorf("Got Priority: %d,, wanted: %d", got.Priority, want.Priority)
	} else {
		t.Logf("AppID Passed with Value: %d", want.Priority)
	}
	if got.One != want.One {
		t.Errorf("Got One: %s,, wanted: %s", got.One, want.One)
	} else {
		t.Logf("One Passed with Value: %s", want.One)
	}
	if got.Two != want.Two {
		t.Errorf("Got Two: %s,, wanted: %s", got.Two, want.Two)
	} else {
		t.Logf("Two Passed with Value: %s", want.Two)
	}
}

func TestGetIntents_ReturnExpectedJSON(t *testing.T) {
	want := Intents{
		[]Intent{
			{
				AppID:     "org.onosproject.cli",
				ID:        "0x40004f",
				Key:       "0x100005",
				State:     "INSTALLED",
				Type:      "HostToHostIntent",
				Resources: []string{"00:00:00:00:00:01/None", "00:00:00:00:00:02/None"},
			},
			{
				AppID:     "org.onosproject.cli",
				ID:        "0x40004d",
				Key:       "0x300009",
				State:     "FAILED",
				Type:      "HostToHostIntent",
				Resources: []string{"00:00:00:00:00:01/None", "00:00:00:00:00:99/None"},
			},
			{
				AppID:     "org.onosproject.cli",
				ID:        "0x40004e",
				Key:       "0x100006",
				State:     "FAILED",
				Type:      "HostToHostIntent",
				Resources: []string{"00:00:00:00:00:02/None", "00:00:00:00:00:88/None"},
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "testdata/intents.json")
		}))
	defer ts.Close()
	c, err := NewClient("host", "username", "password")
	if err != nil {
		t.Fatal("Error creating client")
	}
	c.HTTPClient = ts.Client()
	c.HostURL = ts.URL

	got, err := c.GetIntents()
	if err != nil {
		t.Fatal(err)
	}

	if len(got.Intents) != len(want.Intents) {
		t.Fatal(errors.New("Incorrect number of intents in response"))
	}
	for i, intent := range want.Intents {
		testname := fmt.Sprintf("%s", intent.Key)
		t.Run(testname, func(t *testing.T) {
			if got.Intents[i].AppID != intent.AppID {
				t.Errorf("Got AppID: %s,, wanted: %s", got.Intents[i].AppID, intent.AppID)
			} else {
				t.Logf("AppID Passed with Value: %s", intent.AppID)
			}
			if got.Intents[i].ID != intent.ID {
				t.Errorf("Got ID: %s,, wanted: %s", got.Intents[i].ID, intent.ID)
			} else {
				t.Logf("ID Passed with Value: %s", intent.ID)
			}
			if got.Intents[i].Key != intent.Key {
				t.Errorf("Got Key: %s,, wanted: %s", got.Intents[i].Key, intent.Key)
			} else {
				t.Logf("Key Passed with Value: %s", intent.Key)
			}
			if got.Intents[i].State != intent.State {
				t.Errorf("Got State: %s,, wanted: %s", got.Intents[i].State, intent.State)
			} else {
				t.Logf("State Passed with Value: %s", intent.State)
			}
			if got.Intents[i].Type != intent.Type {
				t.Errorf("Got Type: %s,, wanted: %s", got.Intents[i].Type, intent.Type)
			} else {
				t.Logf("Type Passed with Value: %s", intent.Type)
			}
			if !reflect.DeepEqual(got.Intents[i].Resources, intent.Resources) {
				t.Errorf("Got Resources: %q,, wanted: %q", got.Intents[i].Resources, intent.Resources)
			} else {
				t.Logf("Resources Passed with Value: %q", intent.Resources)
			}
		})
	}
}

//add http test tests for create and update, including delay

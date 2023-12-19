package onosclient

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestParseHosts_CorrectJSON(t *testing.T) {
	hosts, err := os.ReadFile("testdata/hosts.json")
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		Hosts []Host
	}{
		{
			Hosts: []Host{{
				ID:          "00:00:00:00:00:03/None",
				Mac:         "00:00:00:00:00:03",
				Vlan:        "None",
				InnerVlan:   "None",
				OuterTpid:   "0x0000",
				Configured:  false,
				Suspended:   false,
				IPAddresses: []string{"10.0.0.3"},
				Locations: []Location{{
					ElementID: "of:0000000000000003",
					Port:      "1",
				},
				},
			},
				{
					ID:          "00:00:00:00:00:04/None",
					Mac:         "00:00:00:00:00:04",
					Vlan:        "None",
					InnerVlan:   "None",
					OuterTpid:   "0x0000",
					Configured:  false,
					Suspended:   false,
					IPAddresses: []string{"10.0.0.4"},
					Locations: []Location{{
						ElementID: "of:0000000000000003",
						Port:      "2",
					},
					},
				},
				{
					ID:          "00:00:00:00:00:01/None",
					Mac:         "00:00:00:00:00:01",
					Vlan:        "None",
					InnerVlan:   "None",
					OuterTpid:   "0x0000",
					Configured:  false,
					Suspended:   false,
					IPAddresses: []string{"10.0.0.1"},
					Locations: []Location{{
						ElementID: "of:0000000000000002",
						Port:      "1",
					},
					},
				},
				{
					ID:          "00:00:00:00:00:02/None",
					Mac:         "00:00:00:00:00:02",
					Vlan:        "None",
					InnerVlan:   "None",
					OuterTpid:   "0x0000",
					Configured:  false,
					Suspended:   false,
					IPAddresses: []string{"10.0.0.2"},
					Locations: []Location{{
						ElementID: "of:0000000000000002",
						Port:      "2",
					},
					},
				},
			},
		},
	}

	got, err := ParseHosts(hosts)

	if err != nil {
		t.Fatal(err)
	}

	for _, subtest := range tests {
		if len(got.Hosts) != len(subtest.Hosts) {
			t.Fatal(errors.New("Incorrect number of hosts in response"))
		}
		for i, tt := range subtest.Hosts {
			testname := tt.ID
			t.Run(testname, func(t *testing.T) {
				if got.Hosts[i].ID != tt.ID {
					t.Errorf("Got ID: %s,, wanted: %s", got.Hosts[i].ID, tt.ID)
				} else {
					t.Logf("ID Passed with Value: %s", tt.ID)
				}
				if got.Hosts[i].Mac != tt.Mac {
					t.Errorf("Got Mac: %s,, wanted: %s", got.Hosts[i].Mac, tt.Mac)
				} else {
					t.Logf("Mac Passed with Value: %s", tt.Mac)
				}
				if got.Hosts[i].Vlan != tt.Vlan {
					t.Errorf("Got Vlan: %s,, wanted: %s", got.Hosts[i].Vlan, tt.Vlan)
				} else {
					t.Logf("Vlan Passed with Value: %s", tt.Vlan)
				}
				if got.Hosts[i].InnerVlan != tt.InnerVlan {
					t.Errorf("Got InnerVlan: %s,, wanted: %s", got.Hosts[i].InnerVlan, tt.InnerVlan)
				} else {
					t.Logf("InnerVlan Passed with Value: %s", tt.InnerVlan)
				}
				if got.Hosts[i].OuterTpid != tt.OuterTpid {
					t.Errorf("Got OuterTpid: %s,, wanted: %s", got.Hosts[i].OuterTpid, tt.OuterTpid)
				} else {
					t.Logf("OuterTpid Passed with Value: %s", tt.OuterTpid)
				}
				if got.Hosts[i].Configured != tt.Configured {
					t.Errorf("Got Configured: %t,, wanted: %t", got.Hosts[i].Configured, tt.Configured)
				} else {
					t.Logf("Configured Passed with Value: %t", tt.Configured)
				}
				if got.Hosts[i].Suspended != tt.Suspended {
					t.Errorf("Got Suspended: %t,, wanted: %t", got.Hosts[i].Suspended, tt.Suspended)
				} else {
					t.Logf("Suspended Passed with Value: %t", tt.Suspended)
				}
				if !reflect.DeepEqual(got.Hosts[i].IPAddresses, tt.IPAddresses) {
					t.Errorf("Got IPAddresses: %q,, wanted: %q", got.Hosts[i].IPAddresses, tt.IPAddresses)
				} else {
					t.Logf("IPAddresses Passed with Value: %q", tt.IPAddresses)
				}
				if !reflect.DeepEqual(got.Hosts[i].Locations, tt.Locations) {
					t.Errorf("Got Locations: %q,, wanted: %q", got.Hosts[i].Locations, tt.Locations)
				} else {
					t.Logf("Locations Passed with Value: %q", tt.Locations)
				}

			})
		}
	}
}

func TestParseHosts_ErrOnEmpty(t *testing.T) {
	_, err := ParseHosts([]byte{})
	if err == nil {
		t.Fatal("want error parsing empty response, got nil")
	}
}

func TestGetHosts_ReturnExpectedJSON(t *testing.T) {
	want := Hosts{
		[]Host{
			{
				ID:          "00:00:00:00:00:03/None",
				Mac:         "00:00:00:00:00:03",
				Vlan:        "None",
				InnerVlan:   "None",
				OuterTpid:   "0x0000",
				Configured:  false,
				Suspended:   false,
				IPAddresses: []string{"10.0.0.3"},
				Locations: []Location{{
					ElementID: "of:0000000000000003",
					Port:      "1",
				},
				},
			},
			{
				ID:          "00:00:00:00:00:04/None",
				Mac:         "00:00:00:00:00:04",
				Vlan:        "None",
				InnerVlan:   "None",
				OuterTpid:   "0x0000",
				Configured:  false,
				Suspended:   false,
				IPAddresses: []string{"10.0.0.4"},
				Locations: []Location{{
					ElementID: "of:0000000000000003",
					Port:      "2",
				},
				},
			},
			{
				ID:          "00:00:00:00:00:01/None",
				Mac:         "00:00:00:00:00:01",
				Vlan:        "None",
				InnerVlan:   "None",
				OuterTpid:   "0x0000",
				Configured:  false,
				Suspended:   false,
				IPAddresses: []string{"10.0.0.1"},
				Locations: []Location{{
					ElementID: "of:0000000000000002",
					Port:      "1",
				},
				},
			},
			{
				ID:          "00:00:00:00:00:02/None",
				Mac:         "00:00:00:00:00:02",
				Vlan:        "None",
				InnerVlan:   "None",
				OuterTpid:   "0x0000",
				Configured:  false,
				Suspended:   false,
				IPAddresses: []string{"10.0.0.2"},
				Locations: []Location{{
					ElementID: "of:0000000000000002",
					Port:      "2",
				},
				},
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "testdata/hosts.json")
		}))
	defer ts.Close()
	c, err := NewClient("host", "username", "password")
	if err != nil {
		t.Fatal("Error creating client")
	}
	c.HTTPClient = ts.Client()
	c.HostURL = ts.URL

	got, err := c.GetHosts()
	if err != nil {
		t.Fatal(err)
	}

	if len(got.Hosts) != len(want.Hosts) {
		t.Fatal(errors.New("Incorrect number of hosts in response"))
	}
	for i, host := range want.Hosts {
		testname := host.ID
		t.Run(testname, func(t *testing.T) {
			if got.Hosts[i].ID != host.ID {
				t.Errorf("Got ID: %s,, wanted: %s", got.Hosts[i].ID, host.ID)
			} else {
				t.Logf("ID Passed with Value: %s", host.ID)
			}
			if got.Hosts[i].Mac != host.Mac {
				t.Errorf("Got Mac: %s,, wanted: %s", got.Hosts[i].Mac, host.Mac)
			} else {
				t.Logf("Mac Passed with Value: %s", host.Mac)
			}
			if got.Hosts[i].Vlan != host.Vlan {
				t.Errorf("Got Vlan: %s,, wanted: %s", got.Hosts[i].Vlan, host.Vlan)
			} else {
				t.Logf("Vlan Passed with Value: %s", host.Vlan)
			}
			if got.Hosts[i].InnerVlan != host.InnerVlan {
				t.Errorf("Got InnerVlan: %s,, wanted: %s", got.Hosts[i].InnerVlan, host.InnerVlan)
			} else {
				t.Logf("InnerVlan Passed with Value: %s", host.InnerVlan)
			}
			if got.Hosts[i].OuterTpid != host.OuterTpid {
				t.Errorf("Got OuterTpid: %s,, wanted: %s", got.Hosts[i].OuterTpid, host.OuterTpid)
			} else {
				t.Logf("OuterTpid Passed with Value: %s", host.OuterTpid)
			}
			if got.Hosts[i].Configured != host.Configured {
				t.Errorf("Got Configured: %t,, wanted: %t", got.Hosts[i].Configured, host.Configured)
			} else {
				t.Logf("Configured Passed with Value: %t", host.Configured)
			}
			if got.Hosts[i].Suspended != host.Suspended {
				t.Errorf("Got Suspended: %t,, wanted: %t", got.Hosts[i].Suspended, host.Suspended)
			} else {
				t.Logf("Suspended Passed with Value: %t", host.Suspended)
			}
			if !reflect.DeepEqual(got.Hosts[i].IPAddresses, host.IPAddresses) {
				t.Errorf("Got IPAddresses: %q,, wanted: %q", got.Hosts[i].IPAddresses, host.IPAddresses)
			} else {
				t.Logf("IPAddresses Passed with Value: %q", host.IPAddresses)
			}
			if !reflect.DeepEqual(got.Hosts[i].Locations, host.Locations) {
				t.Errorf("Got Locations: %q,, wanted: %q", got.Hosts[i].Locations, host.Locations)
			} else {
				t.Logf("Locations Passed with Value: %q", host.Locations)
			}
		})
	}

}

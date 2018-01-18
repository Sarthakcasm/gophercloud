package testing

import (
	"net/http"
	"fmt"
	"testing"

	fake "github.com/Sarthakcasm/gophercloud/openstack/networking/v2/evs/common"
	"github.com/Sarthakcasm/gophercloud/openstack/networking/v2/evs"
	th "github.com/gophercloud/gophercloud/testhelper"
	//"github.com/Sarthakcasm/gophercloud/openstack/networking/v1/evs"
)

func TestListEvs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2/85636478b0bd8e67e89469c7749d4127/cloudvolumes/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "volumes": [
        {
            "id": "dae79325-ce3a-4905-8112-c6be961933d6",
            "name": "ecs-a334",
            "status": "in-use",
            "shareable": "false",
            "availability_zone": "eu-de-02"
        },
        {
            "id": "e1163763-3f54-46b7-a0a5-f24fe344a0f1",
            "name": "ecs-nordea-test",
            "status": "in-use",
            "shareable": "false",
            "availability_zone": "eu-de-02"
        },
        {
           "id": "9601e0db-9455-4ee1-9047-14dc98cbbaee",
            "name": "test_ecs",
            "status": "in-use",
            "shareable": "false",
            "availability_zone": "eu-de-02"
        }
    ]
}
			`)
	})

	//count := 0

	actual, err := evs.List(fake.ServiceClient(), evs.ListOpts{"","","","","","",""})
	if err != nil {
		t.Errorf("Failed to extract evs: %v", err)
	}

	expected := []evs.EVS{
		{
			ID: "dae79325-ce3a-4905-8112-c6be961933d6",
			Name: "ecs-a334",
			Status: "in-use",
			Availability_zone: "eu-de-02",
			Shareable: "false",

		},
		{
			ID: "e1163763-3f54-46b7-a0a5-f24fe344a0f1",
			Name: "ecs-nordea-test",
			Status: "in-use",
			Availability_zone: "eu-de-02",
			Shareable: "false",
		},
		{
			ID: "9601e0db-9455-4ee1-9047-14dc98cbbaee",
			Name: "test_ecs",
			Status: "in-use",
			Availability_zone: "eu-de-02",
			Shareable: "false",
		},
	}

	th.AssertDeepEquals(t, expected, actual)
}

func TestGetEvs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2/85636478b0bd8e67e89469c7749d4127/volumes/e1163763-3f54-46b7-a0a5-f24fe344a0f1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "volume": {
        "id": "e1163763-3f54-46b7-a0a5-f24fe344a0f1",
            "name": "ecs-nordea-test",
            "status": "in-use",
            "availability_zone": "eu-de-02"
    }
}
		`)
	})

	n, err:= evs.Get(fake.ServiceClient(), "e1163763-3f54-46b7-a0a5-f24fe344a0f1").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "e1163763-3f54-46b7-a0a5-f24fe344a0f1", n.ID)
	th.AssertEquals(t, "ecs-nordea-test", n.Name)
	th.AssertEquals(t, "in-use", n.Status)
	//th.AssertEquals(t, "OK", n.Status)
	//th.AssertDeepEquals(t, "false", n.Shareable)
	th.AssertEquals(t, "eu-de-02", n.Availability_zone)

}

//  "shareable": "false",
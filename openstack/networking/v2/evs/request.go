package evs

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"reflect"
)

type ListOpts struct {
	// ID is the unique identifier for the evs.
	ID string `json:"id"`

	// Name is the human readable name for the evs. It does not have to be
	// unique.
	Name string `json:"name"`

	Status string `json:"status"`

	Availability_zone string `json:"availability_zone"`

	Created_at string `json:"created_at"`

	Volume_type string `json:"volume_type"`


	Shareable string `json:"shareable"`

}

func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(EVSresourceURL(c,id), &r.Body, nil)

	return
}
func List(c *gophercloud.ServiceClient,opts ListOpts)([]EVS, error) {
	u := EvsListURL(c)


	pages, err := pagination.NewPager(c, u, func(r pagination.PageResult) pagination.Page {
		return EVSPage{pagination.LinkedPageBase{PageResult: r}}
	}).AllPages()


	allEvs, err := ExtractEVS(pages)
	if err != nil {
		panic(err)
	}
	return FilterEVSs(allEvs, opts)
}


func FilterEVSs(evs []EVS, opts ListOpts) ([]EVS, error) {

	var refinedEVSs []EVS
	var matched bool
	m := map[string]interface{}{}

	if opts.ID != "" {
		m["ID"] = opts.ID
	}
	if opts.Name != "" {
		m["Name"] = opts.Name
	}
	if opts.Status != "" {
		m["Status"] = opts.Status
	}
	if opts.Availability_zone != "" {
		m["Availability_zone"] = opts.Availability_zone
	}
	if opts.Created_at != "" {
		m["Created_at"] = opts.Created_at
	}
	if opts.Volume_type != "" {
		m["Volume_type"] = opts.Volume_type
	}

	if opts.Shareable != "" {
		m["Shareable"] = opts.Shareable
	}
	if len(m) > 0 && len(evs) > 0 {
		for _, evs := range evs {
			matched = true

			for key, value := range m {
				if sVal := getStructField(&evs, key); !(sVal == value) {
					matched = false
				}
			}

			if matched {
				refinedEVSs = append(refinedEVSs, evs)
			}
		}

	} else {
		refinedEVSs = evs
	}

	return refinedEVSs, nil
}


func getStructField(v *EVS, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

// Delete will permanently delete a particular evs based on its unique ID.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(EVSresourceURL(c, id), nil)
	return
}


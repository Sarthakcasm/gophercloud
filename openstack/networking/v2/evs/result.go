package evs

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud"

)

/*type Links struct {
//	href string `json:"href"`
//	rel string `json:"rel"`
href string `json:"href"`
rel string `json:"rel"`

}*/


type EVS struct{

	//ID string `json:"id"`

	ID string `json:"id"`

	Link interface{} `json:"links"`

	Name string `json:"name"`

	Status string `json:"status"`

	Availability_zone string `json:"availability_zone"`

	Created_at string `json:"created_at"`

	Volume_type string `json:"volume_type"`

	Size int `json:"size"`

	Shareable string `json:"shareable"`

	Source_volid string `json:"source_volid"`

	Snapshot_id string `json:"snapshot_id"`


	Description string `json:"description"`

	Os_vol_tenant_attr string `json:"os-vol-tenant-attr:tenant_id"`

	Bootable string `json:"bootable"`

	Message string `json:"message"`

	Code string `json:"code"`

	Os_vol_host_attr string `json:"os-vol-host-attr:host"`

	Metadata interface{} `json:"metadata"`

	Volume_image_metadata interface{} `json:"volume_image_metadata"`




}

func (r EVSPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"volumes_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

func (r EVSPage) IsEmpty() (bool, error) {
	is, err := ExtractEVS(r)
	return len(is) == 0, err
}

func (r commonResult) Extract() (*EVS, error) {
	var s struct {
		Evs *EVS `json:"volume"`
	}
	err := r.ExtractInto(&s)
	return s.Evs, err
}

/*func (r commonResult)ExtractSingleEVS() (*EVS, error){
	var s struct {
		Evs []EVS `json:"volume"`
	}
	err := (r.(EVSPage)).ExtractInto(&s)
	fmt.Println("Error",err)
	return s.Evs, err
}
*/



func ExtractEVS(r pagination.Page) ([]EVS, error) {

	var s struct {
		Evs []EVS `json:"volumes"`
	}
	err := (r.(EVSPage)).ExtractInto(&s)
	return s.Evs, err
}

type GetResult struct {
	commonResult
}

type EVSPage struct {
	pagination.LinkedPageBase
}

type commonResult struct {
	gophercloud.Result
}

type DeleteResult struct {
	gophercloud.ErrResult
}
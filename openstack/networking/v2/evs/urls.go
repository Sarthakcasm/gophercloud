package evs


import "github.com/gophercloud/gophercloud"

//const resourcePath = "/cloudvolumes/detail"

func EvsListURL(c *gophercloud.ServiceClient) string {

		resourcePath := "cloudvolumes/detail"
		return c.ServiceURL(resourcePath)

}

func EVSresourceURL(c *gophercloud.ServiceClient, id string) string {
	resourcePath := "volumes"
	return c.ServiceURL(resourcePath, id)
}

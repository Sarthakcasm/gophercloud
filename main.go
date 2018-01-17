package main

import(
//"github.com/Savasw/gophercloud/openstack"
"./openstack"
"github.com/gophercloud/gophercloud"
"fmt"
//"reflect"
//	"github.com/Savasw/gophercloud/openstack/networking/v1/vpcs"
//"github.com/rackspace/gophercloud/pagination"
//"github.com/rackspace/gophercloud/openstack/compute/v2/flavors"
//"os"

//"github.com/Savasw/gophercloud/openstack/networking/v2/networks"
//"github.com/rackspace/gophercloud/rackspace/compute/v2/networks"
//"os"
	//"github.com/Sarthakcasm/gophercloud/openstack/networking/v2/evs"
	"github.com/Sarthakcasm/gophercloud/openstack/networking/v2/evs"
	//"github.com/Sarthakcasm/gophercloud/openstack/networking/v1/vpcs"
)

func main() {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "https://iam.eu-de.otc.t-systems.com/v3",
		Username:         "lizhonghua",
		Password:         "slob@123",
		TenantID:         "87a56a48977e42068f70ad3280c50f0e",
		DomainName:       "OTC00000000001000010501",
	}

	provider, err := openstack.AuthenticatedClient(opts) 			//to generate token

	fmt.Println(err)


	opp := gophercloud.EndpointOpts{}


	client, err1 := openstack.NewListServiceV2(provider,opp)

	fmt.Println(err1)
	if err != nil {
		panic(err)
	}

	//List EVS
	opts1 := evs.ListOpts{"","","","","","",""}
	evss , err2 := evs.List(client,opts1) 			//to list existing evs
	for _, evss := range evss {
		fmt.Printf("%+v\n", evss)
	}
	if err2 != nil {
		panic(err)
	}

	//List EVS using ID
	evsid := "dae79325-ce3a-4905-8112-c6be961933d6"

	singleevs, err :=  evs.Get(client,evsid).Extract()



	fmt.Println("EVS ID",singleevs)










}








package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"
	"os"

	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
)

func main() {
	// create a VirtualNetworks client
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	vnetClient := network.NewVirtualNetworksClient(subscriptionID)

	// create an authorizer from env vars or Azure Managed Service Idenity
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err == nil {
		vnetClient.Authorizer = authorizer
	}

	// call the VirtualNetworks CreateOrUpdate API

	vm, err := vnetClient.CreateOrUpdate(context.Background(),
		"test1",
		"vnet-ansh",
		network.VirtualNetwork{
			Location: to.StringPtr("eastus"),
			VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: &[]string{"10.0.0.0/8"},
				},
				Subnets: &[]network.Subnet{
					{
						Name: to.StringPtr("snetone"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefix: to.StringPtr("10.0.0.0/16"),
						},
					},
					{
						Name: to.StringPtr("snettwo"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefix: to.StringPtr("10.1.0.0/16"),
						},
					},
				},
			},
		})

	if err == nil {
		fmt.Println("Success!")
	} else {
		fmt.Println(err)
	}

	fmt.Println(vm)

}


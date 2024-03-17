package main

import (
	"fmt"
	"sort"
)

func main() {
	instances := []*NATInstance{
		&NATInstance{
			Id:   "1",
			Zone: "us-west1-a",
		},
		&NATInstance{
			Id:   "2",
			Zone: "us-west1-b",
		},
		&NATInstance{
			Id:   "3",
			Zone: "us-west1-b",
		},
	}

	subnets := []*Subnet{
		&Subnet{
			Id:   "1",
			Zone: "us-west1-a",
		},
		&Subnet{
			Id:   "2",
			Zone: "us-west1-b",
		},
		&Subnet{
			Id:   "3",
			Zone: "us-west1-b",
		},
		&Subnet{
			Id:   "4",
			Zone: "us-west1-c",
		},
	}

	allocate(instances, subnets)
	printInstances(instances)
}

type Subnet struct {
	Id   string
	Zone string
}

type NATInstance struct {
	Id      string
	Zone    string
	Subnets []*Subnet
}

func printInstances(instances []*NATInstance) {
	for _, i := range instances {
		fmt.Printf("Instance (%v - %v):\n", i.Id, i.Zone)
		for _, s := range i.Subnets {
			fmt.Printf("\tsubnet (%v - %v)\n", s.Id, s.Zone)
		}
	}
}

func sortNatInstancesByNumberOfSubnet(instances []*NATInstance) {
	sort.SliceStable(instances, func(i, j int) bool {
		return len(instances[i].Subnets) < len(instances[j].Subnets)
	})
	/*
		For the bonus question, this function can be modifed like this:
			return sumWeightOf(instances[i].Subnets) < sumWeightOf(instances[j].Subnets)
		
		sumWeightOf: return total weight of all Subnets

		With good inputs, it may have an output that close to what you need.
	*/
}

// allocate Subnets to Instances
func allocate(instances []*NATInstance, subnets []*Subnet) {
	if len(instances) == 0 || len(subnets) == 0 {
		return
	}
	// create a map to query if we have NatInstance in a specific zone
	m := make(map[string]bool)
	for _, instance := range instances {
		m[instance.Zone] = true
	}

	for _, subnet := range subnets {
		// sort to make sure a instance with less subnet will be prioritized
		sortNatInstancesByNumberOfSubnet(instances)
		_, ok := m[subnet.Zone]
		if ok {
			for _, instance := range instances {
				if instance.Zone == subnet.Zone {
					// add to the first instance we found, 
					// if there are many instances in same zone, the first will be the one with least #subnet
					instance.Subnets = append(instance.Subnets, subnet)
					break
				}				
			}
		} else {
			// just add to an instance that have least #subnet
			instances[0].Subnets = append(instances[0].Subnets, subnet)
		}
	}

}

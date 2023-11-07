package main

type Resource struct {
	Name   string
	Amount int
}

type StructureRecipe struct {
	Structure         string
	RequiredResources []ResourceRequirement
}

type ResourceRequirement struct {
	ResourceName string
	Amount       int
}

func initializeResources() []Resource {
	return []Resource{
		{"Silica", 100},
		{"Metal", 50},
		{"Energy", 200},
	}
}

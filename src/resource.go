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
		{"Silica", 0},
		{"Steel", 0},
		{"Steel Frame", 0},
		{"Mining Drill", 0},
		{"Control Panel", 0},
		{"Conveyor Belt", 0},
		{"Power Generator", 0},
		{"Copper Wire", 0},
		{"Electronic Circuit", 0},
		{"Hydrolic Pump", 0},
	}
}

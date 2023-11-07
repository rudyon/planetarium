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

package main

func initializeStructureRecipes() []StructureRecipe {
	return []StructureRecipe{
		{
			Structure: "Miner",
			RequiredResources: []ResourceRequirement{
				{ResourceName: "Metal", Amount: 10},
			},
		},
		{
			Structure: "Furnace",
			RequiredResources: []ResourceRequirement{
				{ResourceName: "Metal", Amount: 10},
			},
		},
	}
}

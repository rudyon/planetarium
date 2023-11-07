package main

func initializeStructureRecipes() []StructureRecipe {
	return []StructureRecipe{
		{
			Structure: "Miner",
			RequiredResources: []ResourceRequirement{
				{ResourceName: "Metal", Amount: 10},
			},
		},
	}
}

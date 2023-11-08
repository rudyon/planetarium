package main

func initializeStructureRecipes() []StructureRecipe {
	return []StructureRecipe{
		{
			Structure: "Miner",
			RequiredResources: []ResourceRequirement{
				{ResourceName: "Iron", Amount: 1},
			},
		},
		{
			Structure: "Furnace",
			RequiredResources: []ResourceRequirement{
				{ResourceName: "Iron", Amount: 2},
			},
		},
	}
}

// TODO resource crafting
// func initializeResourceRecipes() []ResourceRecipe {
// 	return []StructureRecipe{
// 		{
// 			Structure: "Miner",
// 			RequiredResources: []ResourceRequirement{
// 				{ResourceName: "Iron", Amount: 1},
// 			},
// 		},
// 		{
// 			Structure: "Furnace",
// 			RequiredResources: []ResourceRequirement{
// 				{ResourceName: "Iron", Amount: 2},
// 			},
// 		},
// 	}
// }

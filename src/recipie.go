package main

func initializeStructureRecipes() []StructureRecipe {
	return []StructureRecipe{
		{
			Structure: "Miner",
			RequiredResources: []ResourceRequirement{
				{ResourceName: "Steel Frame", Amount: 1},
				{ResourceName: "Mining Drill", Amount: 1},
				{ResourceName: "Control Panel", Amount: 1},
				{ResourceName: "Conveyor Belt", Amount: 10},
				{ResourceName: "Power Generator", Amount: 1},
				{ResourceName: "Copper Wire", Amount: 20},
				{ResourceName: "Electronic Circuit", Amount: 5},
				{ResourceName: "Hydraulic Pump", Amount: 2},
			},
		},
		{
			Structure: "Furnace",
			RequiredResources: []ResourceRequirement{
				{ResourceName: "Steel", Amount: 10},
			},
		},
	}
}

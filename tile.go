package main

type Tile struct {
	ResourceType   string
	ResourceAmount int
	TerrainType    string
	Structure      string
}

func createTiles() []Tile {
	var tiles []Tile

	for i := 0; i < 25; i++ {
		resourceType := "Silica"
		resourceAmount := i * 10
		terrainType := "Plains"

		tile := Tile{
			ResourceType:   resourceType,
			ResourceAmount: resourceAmount,
			TerrainType:    terrainType,
			Structure:      "None",
		}

		tiles = append(tiles, tile)
	}

	return tiles
}

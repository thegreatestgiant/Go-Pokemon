package main

func main() {
	pokeApiClient := pokeapi.newClient()

	locationAreas, err := pokeApiClient.getLocationAreas()

	// startRepl()
}

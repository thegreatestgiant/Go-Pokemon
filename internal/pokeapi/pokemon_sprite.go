package pokeapi

import (
	"fmt"
	"image"
	"net/http"

	"github.com/qeesung/image2ascii/convert"
)

func (p *Pokemon) PrintPokemonSprite() error {
	fetchAndRender(p.Sprites.FrontDefault)
	return nil
}

func fetchAndRender(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching image: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Bad status: %s\n", resp.Status)
		return
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		return
	}

	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 30
	convertOptions.FixedHeight = 20
	convertOptions.Ratio = 0.45

	converter := convert.NewImageConverter()
	ascii := converter.Image2ASCIIString(img, &convertOptions)

	fmt.Println(ascii)
}

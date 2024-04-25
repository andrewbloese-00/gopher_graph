package reader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/andrewbloese-00/gopher_graph/pkg/strains"
)

func ReadCannabisData(filepath string) []strains.StrainEntry {
	file, ferr := os.Open(filepath)
	if ferr != nil {
		fmt.Printf("[Error] %v\n", ferr.Error())
	}

	defer file.Close()
	r := csv.NewReader(file)
	sList := make([]strains.StrainEntry, 0)

	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("[Error] %v\n", err.Error())
	}
	for i, row := range records {
		if i == 0 {
			continue
		} //skip headers

		//get kind into enum type
		var kind strains.StrainType
		switch row[1] {
		case "sativa":
			kind = strains.SATIVA
			break
		case "indica":
			kind = strains.INDICA
			break
		default:
			kind = strains.HYBRID
			break
		}

		effects := make([]strains.StrainEffect, 0)
		//get effects into array of enum type
		effectstrs := strings.Split(row[3], ",")
		for _, eff := range effectstrs {
			low := strings.ToLower(eff)

			switch low {
			case "relaxed":
				effects = append(effects, strains.RELAXED)
				break
			case "uplifted":
				effects = append(effects, strains.UPLIFTED)
				break
			case "tingly":
				effects = append(effects, strains.TINGLY)
				break
			case "happy":
				effects = append(effects, strains.HAPPY)
				break
			case "energetic":
				effects = append(effects, strains.ENERGETIC)
				break
			case "euphoric":
				effects = append(effects, strains.EUPHORIC)
				break
			case "talkative":
				effects = append(effects, strains.TALKATIVE)
				break
			case "focused":
				effects = append(effects, strains.FOCUSED)
				break
			case "sleepy":
				effects = append(effects, strains.SLEEPY)
				break
			case "hungry":
				effects = append(effects, strains.HUNGRY)
				break
			case "aroused":
				effects = append(effects, strains.AROUSED)
				break
			case "giggly":
				effects = append(effects, strains.GIGGLY)
				break
			default:
				break

			}

		}

		//get flavors into array of enum type
		flavors := make([]strains.StrainFlavor, 0)
		flavorstrs := strings.Split(row[4], ",")

		for _, f := range flavorstrs {
			low := strings.ToLower(f)
			switch low {
			case "earthy":
				flavors = append(flavors, strains.EARTHY)
				break
			case "flowery":
				flavors = append(flavors, strains.FLOWERY)
				break
			case "spicy/herbal":
				flavors = append(flavors, strains.SPICY_HERBAL)
				break
			case "apricot":
				flavors = append(flavors, strains.APRICOT)
				break
			case "citrus":
				flavors = append(flavors, strains.CITRUS)
				break
			case "pine":
				flavors = append(flavors, strains.PINE)
				break
			case "diesel":
				flavors = append(flavors, strains.DIESEL)
				break
			case "menthol":
				flavors = append(flavors, strains.MENTHOL)
				break
			case "lemon":
				flavors = append(flavors, strains.LEMON)
				break
			case "sweet":
				flavors = append(flavors, strains.SWEET)
				break
			case "pungent":
				flavors = append(flavors, strains.PUNGENT)
				break
			case "nutty":
				flavors = append(flavors, strains.NUTTY)
				break
			case "skunk":
				flavors = append(flavors, strains.SKUNK)
				break
			case "woody":
				flavors = append(flavors, strains.WOODY)
				break
			case "tobacco":
				flavors = append(flavors, strains.TOBACCO)
				break
			case "pepper":
				flavors = append(flavors, strains.PEPPER)
				break
			case "cheese":
				flavors = append(flavors, strains.CHEESE)
				break
			case "grape":
				flavors = append(flavors, strains.GRAPE)
				break
			case "tropical":
				flavors = append(flavors, strains.TROPICAL)
				break
			case "blueberry":
				flavors = append(flavors, strains.BLUEBERRY)
				break
			case "strawberry":
				flavors = append(flavors, strains.STRAWBERRY)
				break
			case "apple":
				flavors = append(flavors, strains.APPLE)
				break
			case "ammonia":
				flavors = append(flavors, strains.AMMONIA)
				break
			case "lime":
				flavors = append(flavors, strains.LIME)
				break
			case "mango":
				flavors = append(flavors, strains.MANGO)
				break
			case "peach":
				flavors = append(flavors, strains.PEACH)
				break
			case "sage":
				flavors = append(flavors, strains.SAGE)
				break
			case "butter":
				flavors = append(flavors, strains.BUTTER)
				break
			case "vanilla":
				flavors = append(flavors, strains.VANILLA)
				break
			case "coffee":
				flavors = append(flavors, strains.COFFEE)
				break
			case "lavender":
				flavors = append(flavors, strains.LAVENDER)
				break
			case "chemical":
				flavors = append(flavors, strains.CHEMICAL)
				break
			case "honey":
				flavors = append(flavors, strains.HONEY)
				break
			case "pinapple":
				flavors = append(flavors, strains.PINEAPPLE)
				break
			case "plum":
				flavors = append(flavors, strains.PLUM)
				break
			case "tar":
				flavors = append(flavors, strains.TAR)
				break
			case "grapefruit":
				flavors = append(flavors, strains.GRAPEFRUIT)
				break
			case "orange":
				flavors = append(flavors, strains.ORANGE)
				break
			case "rose":
				flavors = append(flavors, strains.ROSE)
				break
			case "tea":
				flavors = append(flavors, strains.TEA)
				break
			case "pear":
				flavors = append(flavors, strains.PEAR)
				break
			case "mint":
				flavors = append(flavors, strains.MINTY)
				break
			case "minty":
				flavors = append(flavors, strains.MINTY)
				break
			default:
				break
			} //end switch
		}

		sList = append(sList, strains.StrainEntry{
			Name:    row[0],
			Kind:    kind,
			Effects: effects,
			Flavors: flavors,
		})

	}

	return sList
}

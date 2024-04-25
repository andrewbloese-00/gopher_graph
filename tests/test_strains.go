package tests

import (
	"fmt"

	"github.com/andrewbloese-00/gopher_graph/pkg/strains"
)

func TestStrains() {

	fmt.Println("\n\n == Test Creating And Comparing Strains ==")
	blueDreamEffects := []strains.StrainEffect{
		strains.UPLIFTED,
		strains.ENERGETIC,
		strains.CREATIVE,
	}

	blueDreamFlavors := []strains.StrainFlavor{strains.EARTHY, strains.PINE}

	BlueDream := &strains.StrainEntry{
		Name:    "Blue Dream",
		Kind:    strains.SATIVA,
		Effects: blueDreamEffects,
		Flavors: blueDreamFlavors,
	}

	otherE := []strains.StrainEffect{strains.UPLIFTED, strains.AROUSED, strains.CREATIVE}
	otherF := []strains.StrainFlavor{strains.PINE}
	AnotherStrain := &strains.StrainEntry{
		Name:    "Other Strain",
		Kind:    strains.HYBRID,
		Effects: otherE,
		Flavors: otherF,
	}
	fmt.Println("Effects Blue Dream: ", BlueDream.GetEffectEmbedding())
	fmt.Println("Effects 'Another':  ", AnotherStrain.GetEffectEmbedding())
	fmt.Println("\nFlavors Blue Dream:  ", BlueDream.GetFlavorEmbedding())
	fmt.Println("Flavors 'Another':  ", AnotherStrain.GetFlavorEmbedding())

	fmt.Println("\nCommon Effects: ", BlueDream.CommonEffectsWith(AnotherStrain))
	fmt.Println("Common Flavors: ", BlueDream.CommonFlavorsWith(AnotherStrain))
}

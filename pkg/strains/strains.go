package strains

import "fmt"

// Strain Enum
type StrainType int

const (
	INDICA StrainType = iota
	HYBRID
	SATIVA
)

// Effects Enum
type StrainEffect int

const (
	CREATIVE StrainEffect = iota
	RELAXED
	UPLIFTED
	HAPPY
	TINGLY
	ENERGETIC
	EUPHORIC
	TALKATIVE
	FOCUSED
	SLEEPY
	HUNGRY
	AROUSED
	GIGGLY
)

// Flavor Enum
type StrainFlavor int

const (
	EARTHY StrainFlavor = iota
	FLOWERY
	SPICY_HERBAL
	APRICOT
	CITRUS
	PINE
	DIESEL
	MENTHOL
	LEMON
	SWEET
	PUNGENT
	NUTTY
	SKUNK
	WOODY
	TOBACCO
	PEPPER
	CHEESE
	GRAPE
	TROPICAL
	BLUEBERRY
	BERRY
	APPLE
	AMMONIA
	STRAWBERRY
	MINTY
	TREE
	LIME
	MANGO
	PEACH
	SAGE
	BUTTER
	VANILLA
	COFFEE
	LAVENDER
	CHEMICAL
	HONEY
	PINEAPPLE
	PLUM
	TAR
	GRAPEFRUIT
	ORANGE
	ROSE
	TEA
	PEAR
)

type StrainEntry struct {
	Kind    StrainType
	Name    string
	Effects []StrainEffect
	Flavors []StrainFlavor
}

var effects_array []string = []string{
	"creative",
	"relaxed",
	"uplifted",
	"tingly",
	"happy",
	"energetic",
	"euphoric",
	"talkative",
	"focused",
	"sleepy",
	"hungry",
	"aroused",
	"giggly",
	"dry mouth",
}
var flavors_array []string = []string{
	"earthy",
	"flowery",
	"spicy/herbal",
	"apricot",
	"citrus",
	"pine",
	"diesel",
	"menthol",
	"lemon",
	"sweet",
	"pungent",
	"nutty",
	"skunk",
	"woody",
	"tobacco",
	"pepper",
	"cheese",
	"grape",
	"tropical",
	"blueberry",
	"berry",
	"apple",
	"ammonia",
	"strawberry",
	"minty",
	"tree",
	"lime",
	"mango",
	"peach",
	"sage",
	"butter",
	"vanilla",
	"coffee",
	"lavender",
	"chemical",
	"honey",
	"pineapple",
	"plum",
	"tar",
	"grapefruit",
	"orange",
	"rose",
	"tea",
	"pear",
}

// helpers grab effects / flavors from effects and flavors arrays
func GetEffectName(i int) string {
	if i > len(effects_array) {
		return "UNDEFINED"
	}
	return effects_array[i]
}

func GetFlavorName(i int) string {
	if i > len(flavors_array) {
		return "UNDEFINED"
	}
	return flavors_array[i]
}

func (s *StrainEntry) GetEffectEmbedding() []byte {
	bits := make([]byte, len(effects_array))
	for _, se := range s.Effects {
		bits[int(se)] = 1
	}
	return bits
}

func (s *StrainEntry) GetFlavorEmbedding() []byte {
	bits := make([]byte, len(flavors_array))
	for _, sf := range s.Flavors {
		bits[int(sf)] = 1
	}
	return bits
}

func (s1 *StrainEntry) CommonEffectsWith(s2 *StrainEntry) []string {
	similar_bits := Similar1Bits(s1.GetEffectEmbedding(), s2.GetEffectEmbedding())
	res := make([]string, 0)
	for _, bitIndex := range similar_bits {
		res = append(res, GetEffectName(bitIndex))
	}
	return res
}

func (s1 *StrainEntry) CommonFlavorsWith(s2 *StrainEntry) []string {
	similar_bits := Similar1Bits(s1.GetFlavorEmbedding(), s2.GetFlavorEmbedding())
	res := make([]string, 0)
	for _, bitIndex := range similar_bits {
		res = append(res, GetFlavorName(bitIndex))
	}
	return res
}

// helper compares two byte arrays and returns the indices of shared '1' bits
func Similar1Bits(a, b []byte) []int {
	sims := make([]int, 0)
	if len(a) != len(b) {
		fmt.Println("Warning: Cannot get similar bits for != length embeddings...")
		return sims
	}

	for i := 0; i < len(a); i += 1 {
		if a[i] == b[i] && a[i] == 1 {
			sims = append(sims, i)
		}
	}
	return sims
}

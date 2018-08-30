package npcgenerator

func (npc *NpcData) setPsyche() {
	violencePsyche := selectWeightedPsycheTrait(weightedViolencePsyche, 13)

}

func selectWeightedPsycheTrait(traitMap map[string][]int, max int) string {
	weightedSelector := r.Intn(max)

	for trait, intRange := range traitMap {
		min := intRange[0]
		max := intRange[len(intRange)-1]
		if (weightedSelector > min && weightedSelector < max) || weightedSelector == min || weightedSelector == max {
			return trait
		}
	}
	return ""
}

var weightedViolencePsyche = map[string][]int{
	"Actively seeks out excuses to perform violence":          generateIntRange(0, 1),
	"Seeks out violence to prove prowess in combat":           generateIntRange(2, 3),
	"Willing to participate in a fight":                       generateIntRange(4, 5),
	"Takes a neutral stance on violence":                      generateIntRange(6, 7),
	"Reluctant to participate in a fight":                     generateIntRange(8, 9),
	"Seeks peaceful resolution over combat":                   generateIntRange(10, 11),
	"Actively avoids situations which could lead to violence": generateIntRange(12, 13),
}
var weightedAnimalPsyche = map[string][]int{
	"More of a cat person":        generateIntRange(0, 1),
	"More of a dog person":        generateIntRange(2, 3),
	"More of a bear person":       generateIntRange(4, 5),
	"More of a reptile person":    generateIntRange(6, 7),
	"More of a amphibian person":  generateIntRange(8, 9),
	"More of a fish person":       generateIntRange(10, 11),
	"More of a bird person":       generateIntRange(12, 13),
	"More of a insect person":     generateIntRange(16, 17),
	"Loves animals of every kind": generateIntRange(18, 19),
}
var weightedRelationPsyche = map[string][]int{
	"Avoids forming relationships with other people": generateIntRange(0, 1),
	"Seeks out relationships with other people":      generateIntRange(0, 1),
	"Cares most about their family":                  generateIntRange(0, 1),
	"Bad relationship with their family":             generateIntRange(0, 1),
	"Makes friends with everyone they meet":          generateIntRange(0, 1),
	"Loner": generateIntRange(0, 1),
	"Keeps a circle of close friends": generateIntRange(0, 1),
}
var weightedRomancePsyche = map[string][]int{
	"Completely aromantic": generateIntRange(0, 1),
	"Cold Fish":            generateIntRange(0, 1),
	"Clumsy at love":       generateIntRange(0, 1),
	"Attentive partner":    generateIntRange(0, 1),
	"True romantic":        generateIntRange(0, 1),
	"Insatiable":           generateIntRange(0, 1),
}
var weightedRacialPsyche = map[string][]int{
	"Completely xenophobic":    generateIntRange(0, 1),
	"Racist against dwarves":   generateIntRange(0, 1),
	"Racist against elves":     generateIntRange(0, 1),
	"Racist against orcs":      generateIntRange(0, 1),
	"Racist against halflings": generateIntRange(0, 1),
	"Racist against goblins":   generateIntRange(0, 1),
	"Racist against gnomes":    generateIntRange(0, 1),
	"Racist against humans":    generateIntRange(0, 1),
}
var weightedMagicPscyhe = map[string][]int{
	"Magic makes everything better":                                   generateIntRange(0, 1),
	"Magic is a tool like any other. It may be used for good or evil": generateIntRange(0, 1),
	"Magic ruins everything it touches":                               generateIntRange(0, 1),
}
var weightedTechPsyche = map[string][]int{
	"Technology makes everything better":                                   generateIntRange(0, 1),
	"Technology is a tool like any other. It may be used for good or evil": generateIntRange(0, 1),
	"Technology ruins everything it touches":                               generateIntRange(0, 1),
}
var miscPscyheList = []string{
	"",
}

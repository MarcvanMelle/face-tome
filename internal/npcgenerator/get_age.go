package npcgenerator

import api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"

type npcAge int32

func (npc *NpcData) setAge() {
	race := npc.request.GetRace()
	ageGroup := npc.request.GetRelativeAge()

	if ageGroup == api.AgeGroup_AGE_UNKNOWN {
		ageGroup = selectWeightedAge()
	}
	if race == api.RaceName_RACE_UNKNOWN {
		race = selectWeightedRace()
	}

	ageGroupMap := selectRacialAgeMap(race)
	ageRange := ageGroupMap[ageGroup]
	selectedAge := npcAge(ageRange[r.Intn(len(ageRange))])

	npc.npcRace = &npcRace{raceName: race}
	npc.npcAge = int32(selectedAge)
}

func selectRacialAgeMap(race api.RaceName) map[api.AgeGroup][]int {
	switch race {
	case api.RaceName_RACE_DRAGONBORN:
		return dragonbornAgeMap
	case api.RaceName_RACE_DWARF_HILL:
		return dwarfHillAgeMap
	case api.RaceName_RACE_DWARF_MOUNTAIN:
		return dwarfMountainAgeMap
	case api.RaceName_RACE_ELF_DROW:
		return elfDrowAgeMap
	case api.RaceName_RACE_ELF_HIGH:
		return elfHighAgeMap
	case api.RaceName_RACE_ELF_WOOD:
		return elfWoodAgeMap
	case api.RaceName_RACE_GNOME_FOREST:
		return gnomeForestAgeMap
	case api.RaceName_RACE_GNOME_ROCK:
		return gnomeRockAgeMap
	case api.RaceName_RACE_HALF_ELF:
		return halfElfAgeMap
	case api.RaceName_RACE_HALF_ORC:
		return halfOrcAgeMap
	case api.RaceName_RACE_HALFING_LIGHTFOOT:
		return halfingLightfootAgeMap
	case api.RaceName_RACE_HALFLING_STOUT:
		return halflingStoutAgeMap
	case api.RaceName_RACE_HUMAN:
		return humanAgeMap
	case api.RaceName_RACE_TIEFLING:
		return tieflingAgeMap
	default:
		return humanAgeMap
	}
}

func selectWeightedRace() api.RaceName {
	weightedSelector := r.Intn(99)

	for raceName, intRange := range weightedRaces {
		min := intRange[0]
		max := intRange[len(intRange)-1]
		if (weightedSelector > min && weightedSelector < max) || weightedSelector == min || weightedSelector == max {
			return raceName
		}
	}
	return api.RaceName_RACE_HUMAN
}

func selectWeightedAge() api.AgeGroup {
	weightedSelector := r.Intn(99)

	for ageGroup, intRange := range weightedAgeGroups {
		min := intRange[0]
		max := intRange[len(intRange)-1]
		if (weightedSelector > min && weightedSelector < max) || weightedSelector == min || weightedSelector == max {
			return ageGroup
		}
	}
	return api.AgeGroup_AGE_ADULT
}

func generateIntRange(min, max int) []int {
	intRange := make([]int, max-min+1)
	for i := range intRange {
		intRange[i] = min + i
	}
	return intRange
}

var ageGroups = []api.AgeGroup{
	api.AgeGroup_AGE_INFANT,
	api.AgeGroup_AGE_TODDLER,
	api.AgeGroup_AGE_CHILD,
	api.AgeGroup_AGE_ADOLESCANT,
	api.AgeGroup_AGE_TEENAGER,
	api.AgeGroup_AGE_YOUNG_ADULT,
	api.AgeGroup_AGE_ADULT,
	api.AgeGroup_AGE_MIDDLE_AGE,
	api.AgeGroup_AGE_OLD_AGE,
	api.AgeGroup_AGE_CENTIGENARIAN,
	api.AgeGroup_AGE_ANCIENT,
	api.AgeGroup_AGE_TIMELESS,
}

var weightedAgeGroups = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 1),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(2, 4),
	api.AgeGroup_AGE_CHILD:         generateIntRange(5, 8),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(9, 12),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(13, 18),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(19, 39),
	api.AgeGroup_AGE_ADULT:         generateIntRange(40, 69),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(70, 89),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(90, 94),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(95, 97),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(98, 98),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(99, 99),
}

var races = []api.RaceName{
	api.RaceName_RACE_DWARF_HILL,
	api.RaceName_RACE_DWARF_MOUNTAIN,
	api.RaceName_RACE_ELF_HIGH,
	api.RaceName_RACE_ELF_WOOD,
	api.RaceName_RACE_ELF_DROW,
	api.RaceName_RACE_HALFING_LIGHTFOOT,
	api.RaceName_RACE_HALFLING_STOUT,
	api.RaceName_RACE_HUMAN,
	api.RaceName_RACE_DRAGONBORN,
	api.RaceName_RACE_GNOME_FOREST,
	api.RaceName_RACE_GNOME_ROCK,
	api.RaceName_RACE_HALF_ELF,
	api.RaceName_RACE_HALF_ORC,
	api.RaceName_RACE_TIEFLING,
}

var weightedRaces = map[api.RaceName][]int{
	api.RaceName_RACE_DWARF_HILL:        generateIntRange(36, 40),
	api.RaceName_RACE_DWARF_MOUNTAIN:    generateIntRange(41, 45),
	api.RaceName_RACE_ELF_HIGH:          generateIntRange(46, 50),
	api.RaceName_RACE_ELF_WOOD:          generateIntRange(51, 55),
	api.RaceName_RACE_ELF_DROW:          generateIntRange(56, 60),
	api.RaceName_RACE_HALFING_LIGHTFOOT: generateIntRange(61, 65),
	api.RaceName_RACE_HALFLING_STOUT:    generateIntRange(66, 70),
	api.RaceName_RACE_HUMAN:             generateIntRange(0, 35),
	api.RaceName_RACE_DRAGONBORN:        generateIntRange(71, 75),
	api.RaceName_RACE_GNOME_FOREST:      generateIntRange(76, 80),
	api.RaceName_RACE_GNOME_ROCK:        generateIntRange(81, 85),
	api.RaceName_RACE_HALF_ELF:          generateIntRange(86, 90),
	api.RaceName_RACE_HALF_ORC:          generateIntRange(91, 95),
	api.RaceName_RACE_TIEFLING:          generateIntRange(96, 99),
}

var dragonbornAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 1),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(2, 4),
	api.AgeGroup_AGE_CHILD:         generateIntRange(5, 7),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(8, 12),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(13, 18),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(19, 26),
	api.AgeGroup_AGE_ADULT:         generateIntRange(27, 39),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(40, 65),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(66, 99),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(100, 199),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(200, 999),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(999, 1000000),
}
var dwarfHillAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 3),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(4, 9),
	api.AgeGroup_AGE_CHILD:         generateIntRange(10, 19),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(19, 29),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(30, 39),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(40, 79),
	api.AgeGroup_AGE_ADULT:         generateIntRange(80, 139),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(140, 199),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(200, 299),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(300, 499),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(500, 9999),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(10000, 10000000),
}
var dwarfMountainAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 3),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(4, 9),
	api.AgeGroup_AGE_CHILD:         generateIntRange(10, 19),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(19, 29),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(30, 39),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(40, 79),
	api.AgeGroup_AGE_ADULT:         generateIntRange(80, 139),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(140, 199),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(200, 299),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(300, 499),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(500, 9999),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(10000, 10000000),
}
var elfDrowAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 4),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(5, 15),
	api.AgeGroup_AGE_CHILD:         generateIntRange(16, 34),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(35, 69),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(70, 109),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(110, 159),
	api.AgeGroup_AGE_ADULT:         generateIntRange(160, 399),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(400, 799),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(800, 1499),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(1500, 2999),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(3000, 19999),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(20000, 1000000),
}
var elfHighAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 4),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(5, 15),
	api.AgeGroup_AGE_CHILD:         generateIntRange(16, 34),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(35, 69),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(70, 109),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(110, 159),
	api.AgeGroup_AGE_ADULT:         generateIntRange(160, 399),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(400, 799),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(800, 1499),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(1500, 2999),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(3000, 19999),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(20000, 1000000),
}
var elfWoodAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 4),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(5, 15),
	api.AgeGroup_AGE_CHILD:         generateIntRange(16, 34),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(35, 69),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(70, 109),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(110, 159),
	api.AgeGroup_AGE_ADULT:         generateIntRange(160, 399),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(400, 799),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(800, 1499),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(1500, 2999),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(3000, 19999),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(20000, 1000000),
}
var gnomeForestAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 3),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(4, 9),
	api.AgeGroup_AGE_CHILD:         generateIntRange(10, 19),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(19, 29),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(30, 39),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(40, 79),
	api.AgeGroup_AGE_ADULT:         generateIntRange(80, 139),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(140, 199),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(200, 299),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(300, 499),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(500, 9999),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(10000, 10000000),
}
var gnomeRockAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 3),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(4, 9),
	api.AgeGroup_AGE_CHILD:         generateIntRange(10, 19),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(19, 29),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(30, 39),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(40, 79),
	api.AgeGroup_AGE_ADULT:         generateIntRange(80, 139),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(140, 199),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(200, 299),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(300, 499),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(500, 9999),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(10000, 10000000),
}
var halfElfAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 1),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(2, 4),
	api.AgeGroup_AGE_CHILD:         generateIntRange(5, 7),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(8, 12),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(13, 18),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(19, 26),
	api.AgeGroup_AGE_ADULT:         generateIntRange(27, 39),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(40, 65),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(66, 99),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(100, 199),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(200, 999),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(999, 1000000),
}
var halfOrcAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 1),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(2, 4),
	api.AgeGroup_AGE_CHILD:         generateIntRange(5, 7),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(8, 12),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(13, 18),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(19, 26),
	api.AgeGroup_AGE_ADULT:         generateIntRange(27, 39),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(40, 65),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(66, 99),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(100, 199),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(200, 999),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(999, 1000000),
}
var halfingLightfootAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 2),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(3, 6),
	api.AgeGroup_AGE_CHILD:         generateIntRange(7, 14),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(15, 20),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(21, 29),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(30, 39),
	api.AgeGroup_AGE_ADULT:         generateIntRange(40, 64),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(65, 79),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(80, 110),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(110, 139),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(140, 199),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(200, 1000000),
}
var halflingStoutAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 2),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(3, 6),
	api.AgeGroup_AGE_CHILD:         generateIntRange(7, 14),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(15, 20),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(21, 29),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(30, 39),
	api.AgeGroup_AGE_ADULT:         generateIntRange(40, 64),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(65, 79),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(80, 110),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(110, 139),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(140, 199),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(200, 1000000),
}
var humanAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 1),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(2, 4),
	api.AgeGroup_AGE_CHILD:         generateIntRange(5, 7),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(8, 12),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(13, 18),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(19, 26),
	api.AgeGroup_AGE_ADULT:         generateIntRange(27, 39),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(40, 65),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(66, 99),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(100, 199),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(200, 999),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(999, 1000000),
}
var tieflingAgeMap = map[api.AgeGroup][]int{
	api.AgeGroup_AGE_INFANT:        generateIntRange(0, 1),
	api.AgeGroup_AGE_TODDLER:       generateIntRange(2, 4),
	api.AgeGroup_AGE_CHILD:         generateIntRange(5, 7),
	api.AgeGroup_AGE_ADOLESCANT:    generateIntRange(8, 12),
	api.AgeGroup_AGE_TEENAGER:      generateIntRange(13, 18),
	api.AgeGroup_AGE_YOUNG_ADULT:   generateIntRange(19, 26),
	api.AgeGroup_AGE_ADULT:         generateIntRange(27, 39),
	api.AgeGroup_AGE_MIDDLE_AGE:    generateIntRange(40, 65),
	api.AgeGroup_AGE_OLD_AGE:       generateIntRange(66, 99),
	api.AgeGroup_AGE_CENTIGENARIAN: generateIntRange(100, 199),
	api.AgeGroup_AGE_ANCIENT:       generateIntRange(200, 999),
	api.AgeGroup_AGE_TIMELESS:      generateIntRange(999, 1000000),
}

package npcgenerator

import (
	"fmt"
	"sort"

	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"
)

type npcStats struct {
	PrimaryStat    string
	SecondaryStat  string
	TertiaryStat   string
	QuartenaryStat string
	QuinaryStat    string
	SenaryStat     string
	Stats          map[string]int
}

var defaultBlock = map[string]int{
	"str": 9,
	"con": 9,
	"dex": 9,
	"int": 9,
	"wis": 9,
	"cha": 9,
}

func (npc *NpcData) setStats() {
	stats := make([]int, 6, 6)
	if npc.npcClass[0].className == api.ClassName_CLASSNAME_COMMONER {
		stats = []int{9, 9, 9, 9, 9, 9}
	} else {
		for i := 0; i < 6; i++ {
			stats[i] = generateStatValue()
		}

		sort.Sort(sort.Reverse(sort.IntSlice(stats)))
	}

	fmt.Println("Stat Imps: ", npc.numStatImps, " Feats: ", npc.numFeats)
	fmt.Println("Base stats: ", stats)
	npc.statsForClass(stats)
	fmt.Println("Class stats: ", npc.npcStats.Stats)
	npc.statsForRace()
	fmt.Println("Race stats: ", npc.npcStats.Stats)
	npc.statsForLevel()
	fmt.Println("Level stats: ", npc.npcStats.Stats)

	// npc.npcStats = raceBlock
}

func (npc *NpcData) statsForClass(statList []int) {
	var stats npcStats

	switch npc.npcClass[0].className {
	case api.ClassName_CLASSNAME_BARBARIAN:
		stats = npcStats{PrimaryStat: "str", SecondaryStat: "con", TertiaryStat: "dex", QuartenaryStat: "wis", QuinaryStat: "cha", SenaryStat: "int"}
	case api.ClassName_CLASSNAME_BARD:
		stats = npcStats{PrimaryStat: "cha", SecondaryStat: "dex", TertiaryStat: "con", QuartenaryStat: "wis", QuinaryStat: "int", SenaryStat: "str"}
	case api.ClassName_CLASSNAME_CLERIC:
		stats = npcStats{PrimaryStat: "wis", SecondaryStat: "con", TertiaryStat: "str", QuartenaryStat: "dex", QuinaryStat: "cha", SenaryStat: "int"}
	case api.ClassName_CLASSNAME_DRUID:
		stats = npcStats{PrimaryStat: "wis", SecondaryStat: "con", TertiaryStat: "dex", QuartenaryStat: "str", QuinaryStat: "int", SenaryStat: "cha"}
	case api.ClassName_CLASSNAME_FIGHTER:
		dexChance := r.Intn(6)
		intChance := r.Intn(10)

		switch {
		case dexChance == 5 && intChance == 9:
			stats = npcStats{PrimaryStat: "dex", SecondaryStat: "int", TertiaryStat: "con", QuartenaryStat: "str", QuinaryStat: "wis", SenaryStat: "cha"}
		case dexChance == 5:
			stats = npcStats{PrimaryStat: "dex", SecondaryStat: "con", TertiaryStat: "str", QuartenaryStat: "wis", QuinaryStat: "cha", SenaryStat: "int"}
		case intChance == 9:
			stats = npcStats{PrimaryStat: "str", SecondaryStat: "int", TertiaryStat: "con", QuartenaryStat: "dex", QuinaryStat: "wis", SenaryStat: "cha"}
		default:
			stats = npcStats{PrimaryStat: "str", SecondaryStat: "con", TertiaryStat: "dex", QuartenaryStat: "wis", QuinaryStat: "cha", SenaryStat: "int"}
		}
	case api.ClassName_CLASSNAME_MONK:
		stats = npcStats{PrimaryStat: "dex", SecondaryStat: "wis", TertiaryStat: "con", QuartenaryStat: "str", QuinaryStat: "cha", SenaryStat: "int"}
	case api.ClassName_CLASSNAME_PALADIN:
		stats = npcStats{PrimaryStat: "str", SecondaryStat: "cha", TertiaryStat: "con", QuartenaryStat: "wis", QuinaryStat: "dex", SenaryStat: "int"}
	case api.ClassName_CLASSNAME_RANGER:
		stats = npcStats{PrimaryStat: "dex", SecondaryStat: "wis", TertiaryStat: "con", QuartenaryStat: "str", QuinaryStat: "cha", SenaryStat: "int"}
	case api.ClassName_CLASSNAME_ROGUE:
		intChance := r.Intn(4)
		if intChance == 3 {
			stats = npcStats{PrimaryStat: "dex", SecondaryStat: "int", TertiaryStat: "con", QuartenaryStat: "cha", QuinaryStat: "wis", SenaryStat: "str"}
		} else {
			stats = npcStats{PrimaryStat: "dex", SecondaryStat: "cha", TertiaryStat: "con", QuartenaryStat: "int", QuinaryStat: "wis", SenaryStat: "str"}
		}
	case api.ClassName_CLASSNAME_SORCEROR:
		stats = npcStats{PrimaryStat: "cha", SecondaryStat: "con", TertiaryStat: "dex", QuartenaryStat: "int", QuinaryStat: "wis", SenaryStat: "str"}
	case api.ClassName_CLASSNAME_WARLOCK:
		stats = npcStats{PrimaryStat: "cha", SecondaryStat: "con", TertiaryStat: "dex", QuartenaryStat: "wis", QuinaryStat: "int", SenaryStat: "str"}
	case api.ClassName_CLASSNAME_WIZARD:
		stats = npcStats{PrimaryStat: "int", SecondaryStat: "con", TertiaryStat: "dex", QuartenaryStat: "wis", QuinaryStat: "cha", SenaryStat: "str"}
	default:
		stats = npcStats{PrimaryStat: "str", SecondaryStat: "con", TertiaryStat: "dex", QuartenaryStat: "int", QuinaryStat: "wis", SenaryStat: "cha"}
	}

	stats.Stats = map[string]int{
		stats.PrimaryStat:    statList[0],
		stats.SecondaryStat:  statList[1],
		stats.TertiaryStat:   statList[2],
		stats.QuartenaryStat: statList[3],
		stats.QuinaryStat:    statList[4],
		stats.SenaryStat:     statList[5],
	}
	npc.npcStats = stats
}

func (npc *NpcData) statsForRace() {
	classBlock := npc.npcStats.Stats
	if npc.npcRace.raceName == api.RaceName_RACE_HUMAN {
		if variantChance := r.Intn(2); variantChance == 1 {
			classBlock[npc.npcStats.PrimaryStat] += 1
			classBlock[npc.npcStats.SecondaryStat] += 1
		} else {
			racialMod := statModMap[npc.npcRace.raceName]
			classBlock["str"] += racialMod["str"]
			classBlock["con"] += racialMod["con"]
			classBlock["dex"] += racialMod["dex"]
			classBlock["int"] += racialMod["int"]
			classBlock["wis"] += racialMod["wis"]
			classBlock["cha"] += racialMod["cha"]
		}
	} else {
		racialMod := statModMap[npc.npcRace.raceName]
		classBlock["str"] += racialMod["str"]
		classBlock["con"] += racialMod["con"]
		classBlock["dex"] += racialMod["dex"]
		classBlock["int"] += racialMod["int"]
		classBlock["wis"] += racialMod["wis"]
		classBlock["cha"] += racialMod["cha"]
	}

	npc.npcStats.Stats = classBlock
}

func (npc *NpcData) statsForLevel() {
	if npc.numStatImps == 0 {
		return
	}
	primary := npc.npcStats.PrimaryStat
	secondary := npc.npcStats.SecondaryStat
	tertiary := npc.npcStats.TertiaryStat
	quartenary := npc.npcStats.QuartenaryStat
	quinary := npc.npcStats.QuinaryStat
	senary := npc.npcStats.SenaryStat

	for i := 0; i < npc.numStatImps; i++ {
		npc.improveStats()
	}

	npc.npcStats.PrimaryStat = primary
	npc.npcStats.SecondaryStat = secondary
	npc.npcStats.TertiaryStat = tertiary
	npc.npcStats.QuartenaryStat = quartenary
	npc.npcStats.QuinaryStat = quinary
	npc.npcStats.SenaryStat = senary
}

func (npc *NpcData) improveStats() {
	stats := npc.npcStats.Stats
	if stats[npc.npcStats.PrimaryStat] == 20 {
		npc.npcStats.PrimaryStat = npc.npcStats.SecondaryStat
		npc.npcStats.SecondaryStat = npc.npcStats.TertiaryStat
		npc.npcStats.TertiaryStat = npc.npcStats.QuartenaryStat
		npc.npcStats.QuartenaryStat = npc.npcStats.QuinaryStat
		npc.npcStats.QuinaryStat = npc.npcStats.SenaryStat
	}

	if stats[npc.npcStats.PrimaryStat] == 19 {
		stats[npc.npcStats.PrimaryStat]++
		stats[npc.npcStats.SecondaryStat]++
	} else {
		stats[npc.npcStats.PrimaryStat] += 2
	}

	npc.npcStats.Stats = stats
}

func generateStatValue() int {
	var sum int
	rolls := make([]int, 4, 4)
	for i := 0; i < 4; i++ {
		rolls[i] = r.Intn(6) + 1
	}
	sort.Ints(rolls)
	for _, val := range rolls[1:] {
		sum += val
	}
	return sum
}

var statNames = []string{
	"str",
	"con",
	"dex",
	"int",
	"wis",
	"cha",
}

var statModMap = map[api.RaceName]map[string]int{
	api.RaceName_RACE_DWARF_HILL: {
		"str": 0,
		"con": 2,
		"dex": 0,
		"int": 0,
		"wis": 1,
		"cha": 0,
	},
	api.RaceName_RACE_DWARF_MOUNTAIN: {
		"str": 2,
		"con": 2,
		"dex": 0,
		"int": 0,
		"wis": 0,
		"cha": 0,
	},
	api.RaceName_RACE_ELF_HIGH: {
		"str": 0,
		"con": 0,
		"dex": 2,
		"int": 1,
		"wis": 0,
		"cha": 0,
	},
	api.RaceName_RACE_ELF_WOOD: {
		"str": 0,
		"con": 0,
		"dex": 2,
		"int": 0,
		"wis": 1,
		"cha": 0,
	},
	api.RaceName_RACE_ELF_DROW: {
		"str": 0,
		"con": 0,
		"dex": 2,
		"int": 0,
		"wis": 0,
		"cha": 1,
	},
	api.RaceName_RACE_HALFING_LIGHTFOOT: {
		"str": 0,
		"con": 0,
		"dex": 2,
		"int": 0,
		"wis": 0,
		"cha": 1,
	},
	api.RaceName_RACE_HALFLING_STOUT: {
		"str": 0,
		"con": 1,
		"dex": 2,
		"int": 0,
		"wis": 0,
		"cha": 0,
	},
	api.RaceName_RACE_HUMAN: {
		"str": 1,
		"con": 1,
		"dex": 1,
		"int": 1,
		"wis": 1,
		"cha": 1,
	},
	api.RaceName_RACE_DRAGONBORN: {
		"str": 2,
		"con": 0,
		"dex": 0,
		"int": 0,
		"wis": 0,
		"cha": 1,
	},
	api.RaceName_RACE_GNOME_FOREST: {
		"str": 0,
		"con": 0,
		"dex": 1,
		"int": 2,
		"wis": 0,
		"cha": 0,
	},
	api.RaceName_RACE_GNOME_ROCK: {
		"str": 0,
		"con": 1,
		"dex": 0,
		"int": 2,
		"wis": 0,
		"cha": 0,
	},
	api.RaceName_RACE_HALF_ELF: {
		"str": 0,
		"con": 0,
		"dex": 0,
		"int": 0,
		"wis": 0,
		"cha": 2,
	},
	api.RaceName_RACE_HALF_ORC: {
		"str": 2,
		"con": 1,
		"dex": 0,
		"int": 0,
		"wis": 0,
		"cha": 0,
	},
	api.RaceName_RACE_TIEFLING: {
		"str": 0,
		"con": 0,
		"dex": 0,
		"int": 1,
		"wis": 0,
		"cha": 2,
	},
}

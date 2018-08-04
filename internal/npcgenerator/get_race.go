package npcgenerator

import api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"

type npcRace struct {
	raceName  api.RaceName
	raceSpeed int32
}

func (npc *NpcData) setRace() {
	race := npc.npcRace.raceName
	npc.npcRace.raceSpeed = raceSpeedMap[race]

	secondLanguage := raceLangMap[race]
	slicedLangList := langList[1:]
	if secondLanguage == api.Language_LANG_COMMON {
		secondLanguage = slicedLangList[r.Intn(len(slicedLangList))]
	}
	npc.npcLang = append(npc.npcLang, secondLanguage)

	if race == api.RaceName_RACE_ELF_DROW {
		npc.npcLang = append(npc.npcLang, api.Language_LANG_UNDERCOMMON)
	}
	if race == api.RaceName_RACE_HALF_ELF {
		npc.npcLang = append(npc.npcLang, slicedLangList[r.Intn(len(slicedLangList))])
	}
}

var raceSpeedMap = map[api.RaceName]int32{
	api.RaceName_RACE_DWARF_HILL:        25,
	api.RaceName_RACE_DWARF_MOUNTAIN:    25,
	api.RaceName_RACE_ELF_HIGH:          30,
	api.RaceName_RACE_ELF_WOOD:          30,
	api.RaceName_RACE_ELF_DROW:          30,
	api.RaceName_RACE_HALFING_LIGHTFOOT: 25,
	api.RaceName_RACE_HALFLING_STOUT:    25,
	api.RaceName_RACE_HUMAN:             30,
	api.RaceName_RACE_DRAGONBORN:        30,
	api.RaceName_RACE_GNOME_FOREST:      25,
	api.RaceName_RACE_GNOME_ROCK:        25,
	api.RaceName_RACE_HALF_ELF:          30,
	api.RaceName_RACE_HALF_ORC:          30,
	api.RaceName_RACE_TIEFLING:          30,
}

var raceLangMap = map[api.RaceName]api.Language{
	api.RaceName_RACE_DWARF_HILL:        api.Language_LANG_DWARVISH,
	api.RaceName_RACE_DWARF_MOUNTAIN:    api.Language_LANG_DWARVISH,
	api.RaceName_RACE_ELF_HIGH:          api.Language_LANG_ELVISH,
	api.RaceName_RACE_ELF_WOOD:          api.Language_LANG_ELVISH,
	api.RaceName_RACE_ELF_DROW:          api.Language_LANG_ELVISH,
	api.RaceName_RACE_HALFING_LIGHTFOOT: api.Language_LANG_HALFLING,
	api.RaceName_RACE_HALFLING_STOUT:    api.Language_LANG_HALFLING,
	api.RaceName_RACE_HUMAN:             api.Language_LANG_COMMON,
	api.RaceName_RACE_DRAGONBORN:        api.Language_LANG_DRACONIC,
	api.RaceName_RACE_GNOME_FOREST:      api.Language_LANG_GNOMISH,
	api.RaceName_RACE_GNOME_ROCK:        api.Language_LANG_GNOMISH,
	api.RaceName_RACE_HALF_ELF:          api.Language_LANG_ELVISH,
	api.RaceName_RACE_HALF_ORC:          api.Language_LANG_ORCISH,
	api.RaceName_RACE_TIEFLING:          api.Language_LANG_INFERNAL,
}

var langList = []api.Language{
	api.Language_LANG_COMMON,
	api.Language_LANG_DWARVISH,
	api.Language_LANG_ELVISH,
	api.Language_LANG_GIANT,
	api.Language_LANG_GNOMISH,
	api.Language_LANG_GOBLIN,
	api.Language_LANG_HALFLING,
	api.Language_LANG_ORCISH,
	api.Language_LANG_ABYSSAL,
	api.Language_LANG_CELESTIAL,
	api.Language_LANG_DRACONIC,
	api.Language_LANG_DEEP,
	api.Language_LANG_INFERNAL,
	api.Language_LANG_PRIMORDIAL,
	api.Language_LANG_SYLVAN,
	api.Language_LANG_UNDERCOMMON,
}

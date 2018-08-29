package npcgenerator

import (
	"fmt"

	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"
)

type npcRace struct {
	raceName         api.RaceName
	raceSpeed        int32
	draconicAncestry api.DraconicAncestry
	racialTraits     map[string]bool
	variantHuman     bool
	halfElf          bool
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
	if race == api.RaceName_RACE_DRAGONBORN {
		npc.npcRace.draconicAncestry = draconicAncestryList[r.Intn(len(draconicAncestryList))]
	}

	npc.setRacialTraits()
}

func (npc *NpcData) setRacialTraits() {
	baseTraits := getBaseTraits()

	switch npc.npcRace.raceName {
	case api.RaceName_RACE_DWARF_HILL:
		baseTraits["darkvision"] = true
		baseTraits["dwarven_combat_training"] = true
		baseTraits["dwarven_resilience"] = true
		baseTraits["dwarven_tool_proficiency"] = true
		baseTraits["dwarven_toughness"] = true
		baseTraits["stone_cunning"] = true
	case api.RaceName_RACE_DWARF_MOUNTAIN:
		baseTraits["darkvision"] = true
		baseTraits["dwarven_armor_training"] = true
		baseTraits["dwarven_combat_training"] = true
		baseTraits["dwarven_resilience"] = true
		baseTraits["dwarven_tool_proficiency"] = true
		baseTraits["stone_cunning"] = true
	case api.RaceName_RACE_ELF_HIGH:
		baseTraits["cantrip"] = true
		baseTraits["darkvision"] = true
		baseTraits["elf_weapon_training"] = true
		baseTraits["extra_language"] = true
		baseTraits["fey_ancestry"] = true
		baseTraits["keen_senses"] = true
		baseTraits["trance"] = true
	case api.RaceName_RACE_ELF_WOOD:
		baseTraits["darkvision"] = true
		baseTraits["elf_weapon_training"] = true
		baseTraits["fey_ancestry"] = true
		baseTraits["fleet_of_foot"] = true
		baseTraits["keen_senses"] = true
		baseTraits["mask_of_the_wild"] = true
		baseTraits["trance"] = true
	case api.RaceName_RACE_ELF_DROW:
		baseTraits["darkvision"] = true
		baseTraits["drow_magic"] = true
		baseTraits["drow_weapon_training"] = true
		baseTraits["keen_senses"] = true
		baseTraits["fey_ancestry"] = true
		baseTraits["sunlight_sensitive"] = true
		baseTraits["superior_darkvision"] = true
		baseTraits["trance"] = true
	case api.RaceName_RACE_HALFING_LIGHTFOOT:
		baseTraits["brave"] = true
		baseTraits["halfling_nimble"] = true
		baseTraits["lucky"] = true
		baseTraits["natural_stealth"] = true
	case api.RaceName_RACE_HALFLING_STOUT:
		baseTraits["brave"] = true
		baseTraits["halfling_nimble"] = true
		baseTraits["lucky"] = true
		baseTraits["stout_resilience"] = true
	case api.RaceName_RACE_HUMAN:
	case api.RaceName_RACE_DRAGONBORN:
		baseTraits["breath_weapon"] = true
		baseTraits["damage_resistance"] = true
	case api.RaceName_RACE_GNOME_FOREST:
		baseTraits["darkvision"] = true
		baseTraits["gnome_cunning"] = true
		baseTraits["natural_illusion"] = true
		baseTraits["speak_with_beasts"] = true
	case api.RaceName_RACE_GNOME_ROCK:
		baseTraits["artifice_lore"] = true
		baseTraits["darkvision"] = true
		baseTraits["gnome_cunning"] = true
		baseTraits["tinker"] = true
	case api.RaceName_RACE_HALF_ELF:
		baseTraits["darkvision"] = true
		baseTraits["fey_ancestry"] = true
		baseTraits["skill_vesatility"] = true
	case api.RaceName_RACE_HALF_ORC:
		baseTraits["darkvision"] = true
		baseTraits["menacing"] = true
		baseTraits["relentless_endure"] = true
		baseTraits["savage_attacks"] = true
	case api.RaceName_RACE_TIEFLING:
		baseTraits["darkvision"] = true
		baseTraits["hellish_resistance"] = true
		baseTraits["infernal_legacy"] = true
	}

	fmt.Println(baseTraits)

	for key, keep := range baseTraits {
		if !keep {
			delete(baseTraits, key)
		}
	}

	npc.npcRace.racialTraits = baseTraits
}

func getBaseTraits() map[string]bool {
	return map[string]bool{
		"artifice_lore":            false,
		"brave":                    false,
		"breath_weapon":            false,
		"cantrip":                  false,
		"damage_resistance":        false,
		"darkvision":               false,
		"drow_magic":               false,
		"drow_weapon_training":     false,
		"dwarven_combat_training":  false,
		"dwarven_resilience":       false,
		"dwarven_tool_proficiency": false,
		"dwarven_toughness":        false,
		"dwarven_armor_training":   false,
		"elven_weapon_trainging":   false,
		"extra_language":           false,
		"fey_ancestry":             false,
		"gnome_cunning":            false,
		"fleet_of_foot":            false,
		"halfling_nimble":          false,
		"hellish_resistance":       false,
		"infernal_legacy":          false,
		"keen_senses":              false,
		"lucky":                    false,
		"mask_of_the_wild":         false,
		"menacing":                 false,
		"natural_illusion":         false,
		"natural_stealth":          false,
		"relentless_endurance":     false,
		"savage_attacks":           false,
		"skill_versatility":        false,
		"speak_with_breaks":        false,
		"stone_cunning":            false,
		"stout_resilience":         false,
		"superior_dark_vision":     false,
		"sunlight_sensitive":       false,
		"tinker":                   false,
		"trance":                   false,
	}
}

var raceSpeedMap = map[api.RaceName]int32{
	api.RaceName_RACE_DWARF_HILL:        25,
	api.RaceName_RACE_DWARF_MOUNTAIN:    25,
	api.RaceName_RACE_ELF_HIGH:          30,
	api.RaceName_RACE_ELF_WOOD:          35,
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

var draconicAncestryList = []api.DraconicAncestry{
	api.DraconicAncestry_DRAC_ANCS_BLACK,
	api.DraconicAncestry_DRAC_ANCS_BLUE,
	api.DraconicAncestry_DRAC_ANCS_BRASS,
	api.DraconicAncestry_DRAC_ANCS_BRONZE,
	api.DraconicAncestry_DRAC_ANCS_COPPER,
	api.DraconicAncestry_DRAC_ANCS_GOLD,
	api.DraconicAncestry_DRAC_ANCS_GREEN,
	api.DraconicAncestry_DRAC_ANCS_RED,
	api.DraconicAncestry_DRAC_ANCS_SILVER,
	api.DraconicAncestry_DRAC_ANCS_WHITE,
}

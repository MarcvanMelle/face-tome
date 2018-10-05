package npcgenerator

import api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"

func (npc *NpcData) setBackground() {
	background := npc.request.GetBackground()

	if background == api.Background_BACK_UNKNOWN {
		baseClass := npc.npcClass[0].className

		if baseClass == api.ClassName_CLASSNAME_COMMONER || r.Intn(10) == 0 {
			background = backgroundList[r.Intn(len(backgroundList))]
			npc.background = background
			return
		}

		switch baseClass {
		case api.ClassName_CLASSNAME_BARBARIAN:
			background = api.Background_BACK_OUTLANDER
		case api.ClassName_CLASSNAME_BARD:
			background = api.Background_BACK_ENTERTAINER
		case api.ClassName_CLASSNAME_CLERIC:
			background = api.Background_BACK_ACOLYTE
		case api.ClassName_CLASSNAME_DRUID:
			background = api.Background_BACK_HERMIT
		case api.ClassName_CLASSNAME_FIGHTER:
			background = api.Background_BACK_SOLDIER
		case api.ClassName_CLASSNAME_MONK:
			background = api.Background_BACK_HERMIT
		case api.ClassName_CLASSNAME_PALADIN:
			background = api.Background_BACK_NOBLE
		case api.ClassName_CLASSNAME_RANGER:
			background = api.Background_BACK_OUTLANDER
		case api.ClassName_CLASSNAME_ROGUE:
			background = api.Background_BACK_CHARLATAN
		case api.ClassName_CLASSNAME_SORCEROR:
			background = api.Background_BACK_HERMIT
		case api.ClassName_CLASSNAME_WARLOCK:
			background = api.Background_BACK_CHARLATAN
		case api.ClassName_CLASSNAME_WIZARD:
			background = api.Background_BACK_SAGE
		}
	}

	npc.background = background
}

var backgroundList = []api.Background{
	api.Background_BACK_ACOLYTE,
	api.Background_BACK_CHARLATAN,
	api.Background_BACK_CRIMINAL,
	api.Background_BACK_ENTERTAINER,
	api.Background_BACK_FOLK_HERO,
	api.Background_BACK_GUILD_ARTISAN,
	api.Background_BACK_HERMIT,
	api.Background_BACK_NOBLE,
	api.Background_BACK_OUTLANDER,
	api.Background_BACK_SAGE,
	api.Background_BACK_SAILOR,
	api.Background_BACK_SOLDIER,
	api.Background_BACK_URCHIN,
}

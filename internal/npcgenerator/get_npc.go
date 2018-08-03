package npcgenerator

import (
	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"
)

// GetNPC returns the fully generated NPC response based on provided protocol buffer parameters
func GetNPC(request *api.GetNPCRequest) (*api.GetNPCResponse, error) {
	npcResponse := &api.GetNPCResponse{
		NpcData: &api.NPC{
			FirstName: "Dougie",
			LastName:  "Mooper",
			Age:       32,
			Alignment: api.NPC_ALIGN_LG,
			Speed:     30,
			Language:  []api.NPC_Language{api.NPC_LANG_COMMON, api.NPC_LANG_DWARVISH},
			Class:     []*api.Class{&api.Class{Name: api.Class_CLASSNAME_BARD, Level: api.Class_LEVEL_ELEVEN}},
			Race: &api.Race{
				Race: api.Race_RACE_DWARF_MOUNTAIN,
				RacialTraits: &api.Race_MountainDwarfTraits{
					MountainDwarfTraits: &api.MountainDwarfTraits{
						StatMod: &api.Stats{
							Str: 2,
							Con: 2,
							Dex: 0,
							Int: 0,
							Wis: 0,
							Cha: 0,
						},
						Darkvision:             true,
						DwarvenResilience:      true,
						DwarvenCombatTraining:  true,
						DwarvenToolProficiency: true,
						StoneCunning:           true,
						DwarvenArmorTraining:   true,
					},
				},
			},
			Stats: &api.Stats{
				Str: 12,
				Con: 13,
				Dex: 18,
				Int: 14,
				Wis: 10,
				Cha: 20,
			},
			Skill: []*api.Skill{
				&api.Skill{
					SkillName:   api.Skill_SKILL_ACROBATICS,
					Proficiency: true,
					StatMod:     "Dex", // Should probably make this an enum
				},
				&api.Skill{
					SkillName:   api.Skill_SKILL_DECEPTION,
					Proficiency: true,
					StatMod:     "Cha",
				},
				&api.Skill{
					SkillName:   api.Skill_SKILL_SLEIGHT,
					Proficiency: true,
					StatMod:     "Dex",
				},
			},
			PhysicalTraits: &api.PhysicalTraits{
				HeightFeet: 4,
				HeightInch: 2,
				Weight:     160,
				SkinTone:   api.PhysicalTraits_SKIN_DUSKY,
				Traits:     []string{},
			},
			PsychologicalTraits: &api.PsychologicalTraits{
				Traits: []string{},
			},
			Gender: api.NPC_GEN_MALE,
		},
	}

	return npcResponse, nil
}

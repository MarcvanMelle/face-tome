package npcgenerator

import (
	"math/rand"
	"time"

	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"
)

type NpcData struct {
	request      *api.GetNPCRequest
	npcName      *npcName
	npcGender    api.Gender
	npcAge       int32
	npcRace      *npcRace
	npcAlign     api.Alignment
	npcClass     []*npcClass
	npcLang      []api.Language
	levelSum     api.Level
	fighterLevel api.Level
	numStatImps  int
	numFeats     int
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// GetNPC returns the fully generated NPC response based on provided protocol buffer parameters
func GetNPC(request *api.GetNPCRequest) (*api.GetNPCResponse, error) {
	npc := &NpcData{request: request, npcLang: []api.Language{api.Language_LANG_COMMON}}
	npc.setName()
	npc.setAge()
	npc.setAlignment()
	npc.setClass()
	npc.setRace()

	var apiClasses []*api.Class
	for _, class := range npc.npcClass {
		apiClass := &api.Class{Name: class.className, Level: class.classLevel}
		apiClasses = append(apiClasses, apiClass)
	}

	npcResponse := &api.GetNPCResponse{
		NpcData: &api.NPC{
			FirstName: npc.npcName.firstName,
			LastName:  npc.npcName.lastName,
			Age:       npc.npcAge,
			Alignment: npc.npcAlign,
			Speed:     npc.npcRace.raceSpeed,
			Language:  npc.npcLang,
			Class:     apiClasses,
			Race: &api.Race{
				Race: npc.npcRace.raceName,
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
			Gender: npc.npcGender,
		},
	}

	return npcResponse, nil
}

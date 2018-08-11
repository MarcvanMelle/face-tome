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
	npcStats     npcStats
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
	npc.setStats()

	apiClasses := make([]*api.Class, len(npc.npcClass), len(npc.npcClass))
	for i, class := range npc.npcClass {
		apiClasses[i] = &api.Class{Name: class.className, Level: class.classLevel}
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
				Str: int32(npc.npcStats.Stats["str"]),
				Con: int32(npc.npcStats.Stats["con"]),
				Dex: int32(npc.npcStats.Stats["dex"]),
				Int: int32(npc.npcStats.Stats["int"]),
				Wis: int32(npc.npcStats.Stats["wis"]),
				Cha: int32(npc.npcStats.Stats["cha"]),
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

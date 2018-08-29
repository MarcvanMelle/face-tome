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
	skills       []*api.Skill
	background   api.Background
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
	npc.setBackground()
	npc.setSkills()

	apiClasses := make([]*api.Class, len(npc.npcClass), len(npc.npcClass))
	for i, class := range npc.npcClass {
		apiClasses[i] = &api.Class{Name: class.className, Level: class.classLevel}
	}

	npcResponse := &api.GetNPCResponse{
		NpcData: &api.NPC{
			FirstName: npc.npcName.firstName,
			LastName:  npc.npcName.lastName,
			Gender:    npc.npcGender,
			Age:       npc.npcAge,
			Alignment: npc.npcAlign,
			Speed:     npc.npcRace.raceSpeed,
			Language:  npc.npcLang,
			Class:     apiClasses,
			Race: &api.Race{
				Race:         npc.npcRace.raceName,
				RacialTraits: npc.npcRace.racialTraits,
			},
			PhysicalTraits: &api.PhysicalTraits{
				HeightFeet: 4,
				HeightInch: 2,
				Weight:     160,
				SkinTone:   api.PhysicalTraits_SKIN_DUSKY,
				Traits:     []string{},
			},
			Stats: &api.Stats{
				Str: int32(npc.npcStats.Stats["str"]),
				Con: int32(npc.npcStats.Stats["con"]),
				Dex: int32(npc.npcStats.Stats["dex"]),
				Int: int32(npc.npcStats.Stats["int"]),
				Wis: int32(npc.npcStats.Stats["wis"]),
				Cha: int32(npc.npcStats.Stats["cha"]),
			},
			Skill: npc.skills,
			PsychologicalTraits: &api.PsychologicalTraits{
				Traits: []string{},
			},
		},
	}

	return npcResponse, nil
}

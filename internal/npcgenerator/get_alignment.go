package npcgenerator

import api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"

func (npc *NpcData) setAlignment() {
	alignment := npc.request.GetAlignment()
	class := npc.request.GetClass()

	if class == api.ClassName_CLASSNAME_UNKNOWN {
		class = classList[r.Intn(len(classList))]
	}

	if alignment == api.Alignment_ALIGN_UNKNOWN {
		if class == api.ClassName_CLASSNAME_PALADIN {
			chance := r.Intn(99)
			if chance < 10 {
				selectWeightedAlignment()
			} else {
				alignment = api.Alignment_ALIGN_LG
			}
		} else {
			alignment = selectWeightedAlignment()
		}
	}

	npc.npcAlign = alignment
	npc.npcClass = []*npcClass{&npcClass{className: class}}
}

func selectWeightedAlignment() api.Alignment {
	weightedSelector := r.Intn(99)

	for alignment, intRange := range weightedAlignments {
		min := intRange[0]
		max := intRange[len(intRange)-1]
		if (weightedSelector > min && weightedSelector < max) || weightedSelector == min || weightedSelector == max {
			return alignment
		}
	}
	return api.Alignment_ALIGN_NN
}

var weightedAlignments = map[api.Alignment][]int{
	api.Alignment_ALIGN_CE: generateIntRange(94, 99),
	api.Alignment_ALIGN_NE: generateIntRange(88, 93),
	api.Alignment_ALIGN_LE: generateIntRange(82, 87),
	api.Alignment_ALIGN_CG: generateIntRange(76, 81),
	api.Alignment_ALIGN_NG: generateIntRange(70, 75),
	api.Alignment_ALIGN_LG: generateIntRange(64, 69),
	api.Alignment_ALIGN_CN: generateIntRange(43, 63),
	api.Alignment_ALIGN_NN: generateIntRange(0, 21),
	api.Alignment_ALIGN_LN: generateIntRange(22, 42),
}

var classList = []api.ClassName{
	api.ClassName_CLASSNAME_COMMONER,
	api.ClassName_CLASSNAME_BARBARIAN,
	api.ClassName_CLASSNAME_BARD,
	api.ClassName_CLASSNAME_CLERIC,
	api.ClassName_CLASSNAME_DRUID,
	api.ClassName_CLASSNAME_FIGHTER,
	api.ClassName_CLASSNAME_MONK,
	api.ClassName_CLASSNAME_PALADIN,
	api.ClassName_CLASSNAME_RANGER,
	api.ClassName_CLASSNAME_SORCEROR,
	api.ClassName_CLASSNAME_WARLOCK,
	api.ClassName_CLASSNAME_WARRIOR,
	api.ClassName_CLASSNAME_WIZARD,
}

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
			if chance == 0 {
				alignment = alignmentList[r.Intn(len(alignmentList))]
			} else {
				nonEvilList := alignmentList[3:]
				alignment = nonEvilList[r.Intn(len(nonEvilList))]
			}
		} else {
			alignment = alignmentList[r.Intn(len(alignmentList))]
		}
	}

	npc.npcAlign = alignment
	npc.npcClass = []*npcClass{&npcClass{className: class}}
}

var alignmentList = []api.Alignment{
	api.Alignment_ALIGN_CE,
	api.Alignment_ALIGN_NE,
	api.Alignment_ALIGN_LE,
	api.Alignment_ALIGN_CG,
	api.Alignment_ALIGN_NG,
	api.Alignment_ALIGN_LG,
	api.Alignment_ALIGN_CN,
	api.Alignment_ALIGN_NN,
	api.Alignment_ALIGN_LN,
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

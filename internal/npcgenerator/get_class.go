package npcgenerator

import (
	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"
)

type npcClass struct {
	className  api.ClassName
	classLevel api.Level
	statImps   int
	feats      int
}

func (npc *NpcData) setClass() {
	firstClass := npc.npcClass[0]
	if firstClass.className == api.ClassName_CLASSNAME_COMMONER {
		firstClass.classLevel = api.Level_LEVEL_ONE
		return
	}
	level := selectWeightedLevel()
	firstClass.classLevel = level
	npc.levelSum = level
	if firstClass.className == api.ClassName_CLASSNAME_FIGHTER {
		npc.fighterLevel = level
	}

	chance := r.Intn(100)
	if chance == 0 && npc.levelSum < 15 {
		generateNewClass(npc)
	}

	npc.calculateAbilityScoreImps()
}

func generateNewClass(npc *NpcData) {
	var newClassName api.ClassName
	var newClassLevel api.Level
	var existingClasses []api.ClassName

	for _, class := range npc.npcClass {
		existingClasses = append(existingClasses, class.className)
	}

	safeClassList := classList[1:]
	for newClassName == api.ClassName_CLASSNAME_UNKNOWN || containClass(existingClasses, newClassName) {
		newClassName = safeClassList[r.Intn(len(safeClassList))]
	}

	for newClassLevel == api.Level_LEVEL_UNKNOWN || levelSumExceeded(npc.levelSum, newClassLevel) {
		newClassLevel = selectWeightedLevel()
	}

	newClass := &npcClass{className: newClassName, classLevel: newClassLevel}
	npc.npcClass = append(npc.npcClass, newClass)
	npc.levelSum += newClassLevel
	if newClassName == api.ClassName_CLASSNAME_FIGHTER {
		npc.fighterLevel = newClassLevel
	}

	chance := r.Intn(100)
	if chance == 0 && npc.levelSum < 15 {
		generateNewClass(npc)
	}
}

func selectWeightedLevel() api.Level {
	weightedSelector := r.Intn(100)

	for level, intRange := range weightedLevels {
		min := intRange[0]
		max := intRange[len(intRange)-1]
		if (weightedSelector > min && weightedSelector < max) || weightedSelector == min || weightedSelector == max {
			return level
		}
	}
	return api.Level_LEVEL_ONE
}

func levelSumExceeded(sum api.Level, level api.Level) bool {
	if (sum + level) > 20 {
		return true
	}
	return false
}

func containClass(s []api.ClassName, val api.ClassName) bool {
	for _, el := range s {
		if el == val {
			return true
		}
	}
	return false
}

func (npc *NpcData) calculateAbilityScoreImps() {
	var totalImps int
	switch {
	case npc.levelSum >= 19:
		totalImps = 5
	case npc.levelSum >= 16:
		totalImps = 4
	case npc.levelSum >= 12:
		totalImps = 3
	case npc.levelSum >= 8:
		totalImps = 2
	case npc.levelSum >= 4:
		totalImps = 1
	default:
		totalImps = 0
	}

	if npc.fighterLevel >= 6 {
		switch {
		case npc.fighterLevel >= 14:
			totalImps++
		case npc.fighterLevel >= 6:
			totalImps++
		}
	}

	for i := 1; i <= totalImps; i++ {
		chance := r.Intn(100)
		if chance < 50 {
			npc.numStatImps++
		} else {
			npc.numFeats++
		}
	}
}

var weightedLevels = map[api.Level][]int{
	api.Level_LEVEL_ONE:       generateIntRange(0, 26),
	api.Level_LEVEL_TWO:       generateIntRange(27, 32),
	api.Level_LEVEL_THREE:     generateIntRange(33, 38),
	api.Level_LEVEL_FOUR:      generateIntRange(39, 44),
	api.Level_LEVEL_FIVE:      generateIntRange(45, 49),
	api.Level_LEVEL_SIX:       generateIntRange(50, 54),
	api.Level_LEVEL_SEVEN:     generateIntRange(55, 59),
	api.Level_LEVEL_EIGHT:     generateIntRange(60, 64),
	api.Level_LEVEL_NINE:      generateIntRange(65, 68),
	api.Level_LEVEL_TEN:       generateIntRange(69, 72),
	api.Level_LEVEL_ELEVEN:    generateIntRange(73, 76),
	api.Level_LEVEL_TWELVE:    generateIntRange(77, 80),
	api.Level_LEVEL_THIRTEEN:  generateIntRange(81, 83),
	api.Level_LEVEL_FOURTEEN:  generateIntRange(84, 86),
	api.Level_LEVEL_FIFTEEN:   generateIntRange(87, 89),
	api.Level_LEVEL_SIXTEEN:   generateIntRange(90, 92),
	api.Level_LEVEL_SEVENTEEN: generateIntRange(93, 94),
	api.Level_LEVEL_EIGHTEEN:  generateIntRange(95, 96),
	api.Level_LEVEL_NINETEEN:  generateIntRange(97, 98),
	api.Level_LEVEL_TWENTY:    generateIntRange(99, 99),
}

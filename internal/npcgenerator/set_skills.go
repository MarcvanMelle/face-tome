package npcgenerator

import api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"

func (npc *NpcData) setSkills() {
	skills := make([]api.SkillName, 0, 10)
	npcSkills := make([]*api.Skill, 0, 10)

	skills = backgroundSkills(npc.background, skills)
	skills = classSkills(npc.npcClass[0].className, skills)
	skills = racialSkills(npc.npcRace.raceName, npc.npcRace.variantHuman, skills)

	for _, skill := range skills {
		npcSkills = append(npcSkills, &api.Skill{SkillName: skill})
	}

	npc.skills = npcSkills
}

func backgroundSkills(background api.Background, skills []api.SkillName) []api.SkillName {
	switch background {
	case api.Background_BACK_ACOLYTE:
		skills = append(skills, api.SkillName_SKILL_INSIGHT)
		skills = append(skills, api.SkillName_SKILL_RELIGION)
	case api.Background_BACK_CHARLATAN:
		skills = append(skills, api.SkillName_SKILL_DECEPTION)
		skills = append(skills, api.SkillName_SKILL_SLEIGHT)
	case api.Background_BACK_CRIMINAL:
		skills = append(skills, api.SkillName_SKILL_DECEPTION)
		skills = append(skills, api.SkillName_SKILL_STEALTH)
	case api.Background_BACK_ENTERTAINER:
		skills = append(skills, api.SkillName_SKILL_ACROBATICS)
		skills = append(skills, api.SkillName_SKILL_PERFORMANCE)
	case api.Background_BACK_FOLK_HERO:
		skills = append(skills, api.SkillName_SKILL_ANIMAL_HANDLING)
		skills = append(skills, api.SkillName_SKILL_SURVIVAL)
	case api.Background_BACK_GUILD_ARTISAN:
		skills = append(skills, api.SkillName_SKILL_INSIGHT)
		skills = append(skills, api.SkillName_SKILL_PERSUASION)
	case api.Background_BACK_HERMIT:
		skills = append(skills, api.SkillName_SKILL_MEDICINE)
		skills = append(skills, api.SkillName_SKILL_RELIGION)
	case api.Background_BACK_NOBLE:
		skills = append(skills, api.SkillName_SKILL_HISTORY)
		skills = append(skills, api.SkillName_SKILL_PERSUASION)
	case api.Background_BACK_OUTLANDER:
		skills = append(skills, api.SkillName_SKILL_ATHLETICS)
		skills = append(skills, api.SkillName_SKILL_SURVIVAL)
	case api.Background_BACK_SAGE:
		skills = append(skills, api.SkillName_SKILL_ARCANA)
		skills = append(skills, api.SkillName_SKILL_HISTORY)
	case api.Background_BACK_SAILOR:
		skills = append(skills, api.SkillName_SKILL_ATHLETICS)
		skills = append(skills, api.SkillName_SKILL_PERCEPTION)
	case api.Background_BACK_SOLDIER:
		skills = append(skills, api.SkillName_SKILL_ATHLETICS)
		skills = append(skills, api.SkillName_SKILL_INTIMIDATION)
	case api.Background_BACK_URCHIN:
		skills = append(skills, api.SkillName_SKILL_SLEIGHT)
		skills = append(skills, api.SkillName_SKILL_STEALTH)
	}

	return skills
}

func classSkills(class api.ClassName, skills []api.SkillName) []api.SkillName {
	switch class {
	case api.ClassName_CLASSNAME_COMMONER:
		newSkill := randomSkill(skillList, skills)
		skills = append(skills, newSkill)
	case api.ClassName_CLASSNAME_BARBARIAN:
		for i := 0; i < 2; i++ {
			newSkill := randomSkill(barbarianSkills, skills)
			skills = append(skills, newSkill)
		}
	case api.ClassName_CLASSNAME_BARD:
		for i := 0; i < 3; i++ {
			newSkill := randomSkill(bardSkills, skills)
			skills = append(skills, newSkill)
		}
	case api.ClassName_CLASSNAME_CLERIC:
		for i := 0; i < 2; i++ {
			newSkill := randomSkill(clericSkills, skills)
			skills = append(skills, newSkill)
		}
	case api.ClassName_CLASSNAME_DRUID:
		for i := 0; i < 2; i++ {
			newSkill := randomSkill(druidSkills, skills)
			skills = append(skills, newSkill)
		}
	case api.ClassName_CLASSNAME_FIGHTER:
		for i := 0; i < 2; i++ {
			newSkill := randomSkill(fighterSkills, skills)
			skills = append(skills, newSkill)
		}
	case api.ClassName_CLASSNAME_MONK:
		for i := 0; i < 2; i++ {
			newSkill := randomSkill(monkSkills, skills)
			skills = append(skills, newSkill)
		}
	case api.ClassName_CLASSNAME_PALADIN:
		for i := 0; i < 2; i++ {
			newSkill := randomSkill(paladinSkills, skills)
			skills = append(skills, newSkill)
		}
	case api.ClassName_CLASSNAME_RANGER:
		for i := 0; i < 3; i++ {
			newSkill := randomSkill(rangerSkills, skills)
			skills = append(skills, newSkill)
		}
	case api.ClassName_CLASSNAME_ROGUE:
		for i := 0; i < 4; i++ {
			newSkill := randomSkill(rougeSkills, skills)
			skills = append(skills, newSkill)
		}
	case api.ClassName_CLASSNAME_SORCEROR:
		for i := 0; i < 2; i++ {
			newSkill := randomSkill(sorcerorSkills, skills)
			skills = append(skills, newSkill)
		}
	case api.ClassName_CLASSNAME_WARLOCK:
		for i := 0; i < 2; i++ {
			newSkill := randomSkill(warlockSkills, skills)
			skills = append(skills, newSkill)
		}
	case api.ClassName_CLASSNAME_WIZARD:
		for i := 0; i < 2; i++ {
			newSkill := randomSkill(wizardSkills, skills)
			skills = append(skills, newSkill)
		}
	}

	return skills
}

func racialSkills(race api.RaceName, variantHuman bool, skills []api.SkillName) []api.SkillName {
	if variantHuman {
		newSkill := randomSkill(skillList, skills)
		skills = append(skills, newSkill)
		return skills
	}

	if race == api.RaceName_RACE_HALF_ELF {
		for i := 0; i < 2; i++ {
			newSkill := randomSkill(skillList, skills)
			skills = append(skills, newSkill)
		}
		return skills
	}

	if race == api.RaceName_RACE_HALF_ORC && !containSkill(skills, api.SkillName_SKILL_INTIMIDATION) {
		skills = append(skills, api.SkillName_SKILL_INTIMIDATION)
		return skills
	}

	return skills
}

func randomSkill(availableSkills, existingSkills []api.SkillName) api.SkillName {
	var newSkill api.SkillName
	for newSkill == api.SkillName_SKILL_UNKNOWN || containSkill(existingSkills, newSkill) {
		newSkill = availableSkills[r.Intn(len(availableSkills))]
	}
	return newSkill
}

func containSkill(s []api.SkillName, val api.SkillName) bool {
	for _, el := range s {
		if el == val {
			return true
		}
	}
	return false
}

var skillList = []api.SkillName{
	api.SkillName_SKILL_ATHLETICS,
	api.SkillName_SKILL_ACROBATICS,
	api.SkillName_SKILL_SLEIGHT,
	api.SkillName_SKILL_STEALTH,
	api.SkillName_SKILL_ARCANA,
	api.SkillName_SKILL_HISTORY,
	api.SkillName_SKILL_INVESTIGATION,
	api.SkillName_SKILL_NATURE,
	api.SkillName_SKILL_RELIGION,
	api.SkillName_SKILL_ANIMAL_HANDLING,
	api.SkillName_SKILL_INSIGHT,
	api.SkillName_SKILL_MEDICINE,
	api.SkillName_SKILL_PERCEPTION,
	api.SkillName_SKILL_SURVIVAL,
	api.SkillName_SKILL_DECEPTION,
	api.SkillName_SKILL_INTIMIDATION,
	api.SkillName_SKILL_PERFORMANCE,
	api.SkillName_SKILL_PERSUASION,
}

var barbarianSkills = []api.SkillName{
	api.SkillName_SKILL_ANIMAL_HANDLING,
	api.SkillName_SKILL_ATHLETICS,
	api.SkillName_SKILL_INTIMIDATION,
	api.SkillName_SKILL_NATURE,
	api.SkillName_SKILL_PERCEPTION,
	api.SkillName_SKILL_SURVIVAL,
}

var bardSkills = skillList

var clericSkills = []api.SkillName{
	api.SkillName_SKILL_HISTORY,
	api.SkillName_SKILL_INSIGHT,
	api.SkillName_SKILL_MEDICINE,
	api.SkillName_SKILL_PERSUASION,
	api.SkillName_SKILL_RELIGION,
}

var druidSkills = []api.SkillName{
	api.SkillName_SKILL_ARCANA,
	api.SkillName_SKILL_ANIMAL_HANDLING,
	api.SkillName_SKILL_INSIGHT,
	api.SkillName_SKILL_MEDICINE,
	api.SkillName_SKILL_NATURE,
	api.SkillName_SKILL_PERCEPTION,
	api.SkillName_SKILL_RELIGION,
	api.SkillName_SKILL_SURVIVAL,
}

var fighterSkills = []api.SkillName{
	api.SkillName_SKILL_ACROBATICS,
	api.SkillName_SKILL_ANIMAL_HANDLING,
	api.SkillName_SKILL_ATHLETICS,
	api.SkillName_SKILL_HISTORY,
	api.SkillName_SKILL_INSIGHT,
	api.SkillName_SKILL_INTIMIDATION,
	api.SkillName_SKILL_PERCEPTION,
	api.SkillName_SKILL_SURVIVAL,
}

var monkSkills = []api.SkillName{
	api.SkillName_SKILL_ACROBATICS,
	api.SkillName_SKILL_ATHLETICS,
	api.SkillName_SKILL_HISTORY,
	api.SkillName_SKILL_INSIGHT,
	api.SkillName_SKILL_RELIGION,
	api.SkillName_SKILL_STEALTH,
}

var paladinSkills = []api.SkillName{
	api.SkillName_SKILL_ATHLETICS,
	api.SkillName_SKILL_INSIGHT,
	api.SkillName_SKILL_INTIMIDATION,
	api.SkillName_SKILL_MEDICINE,
	api.SkillName_SKILL_PERSUASION,
	api.SkillName_SKILL_RELIGION,
}

var rangerSkills = []api.SkillName{
	api.SkillName_SKILL_ANIMAL_HANDLING,
	api.SkillName_SKILL_ATHLETICS,
	api.SkillName_SKILL_INSIGHT,
	api.SkillName_SKILL_INVESTIGATION,
	api.SkillName_SKILL_NATURE,
	api.SkillName_SKILL_PERCEPTION,
	api.SkillName_SKILL_STEALTH,
	api.SkillName_SKILL_SURVIVAL,
}

var rougeSkills = []api.SkillName{
	api.SkillName_SKILL_ACROBATICS,
	api.SkillName_SKILL_ATHLETICS,
	api.SkillName_SKILL_DECEPTION,
	api.SkillName_SKILL_INSIGHT,
	api.SkillName_SKILL_INTIMIDATION,
	api.SkillName_SKILL_INVESTIGATION,
	api.SkillName_SKILL_PERCEPTION,
	api.SkillName_SKILL_PERFORMANCE,
	api.SkillName_SKILL_PERSUASION,
	api.SkillName_SKILL_SLEIGHT,
	api.SkillName_SKILL_STEALTH,
}

var sorcerorSkills = []api.SkillName{
	api.SkillName_SKILL_ARCANA,
	api.SkillName_SKILL_DECEPTION,
	api.SkillName_SKILL_INSIGHT,
	api.SkillName_SKILL_INTIMIDATION,
	api.SkillName_SKILL_PERSUASION,
	api.SkillName_SKILL_RELIGION,
}

var warlockSkills = []api.SkillName{
	api.SkillName_SKILL_ARCANA,
	api.SkillName_SKILL_DECEPTION,
	api.SkillName_SKILL_HISTORY,
	api.SkillName_SKILL_INTIMIDATION,
	api.SkillName_SKILL_INVESTIGATION,
	api.SkillName_SKILL_NATURE,
	api.SkillName_SKILL_RELIGION,
}

var wizardSkills = []api.SkillName{
	api.SkillName_SKILL_ARCANA,
	api.SkillName_SKILL_HISTORY,
	api.SkillName_SKILL_INSIGHT,
	api.SkillName_SKILL_INVESTIGATION,
	api.SkillName_SKILL_MEDICINE,
	api.SkillName_SKILL_RELIGION,
}

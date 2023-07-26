package configs

import "errors"

var ErrUnknownStage = errors.New("unknown stage")

type Stage uint8

const (
	UNKNOWN Stage = iota
	DEV           // Stage of Development
	SIT           // Stage of System Integration Testing
	STG           // Stage of Staging
	UAT           // Stage of User Acceptance Testing
	PRD           // Stage of Production
)

func (s Stage) String() string {
	switch s {
	case DEV:
		return "Stage[DEV]"
	case SIT:
		return "Stage[SIT]"
	case STG:
		return "Stage[STG]"
	case UAT:
		return "Stage[UAT]"
	case PRD:
		return "Stage[PRD]"
	default:
		return "Stage[Unknown]"
	}
}

// StageParsing parse a string as Stage throw error if unknown stage comes up
func StageParsing(s string) (Stage, error) {
	switch s {
	case "DEV":
		return DEV, nil
	case "SIT":
		return SIT, nil
	case "STG":
		return STG, nil
	case "UAT":
		return UAT, nil
	case "PRD":
		return PRD, nil
	default:
		return UNKNOWN, ErrUnknownStage
	}
}

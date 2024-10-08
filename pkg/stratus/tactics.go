package stratus

import (
	"database/sql"
	"github.com/datadog/stratus-red-team/v2/pkg/stratus/mitreattack"
	"github.com/sirupsen/logrus"
	"log"
	"math/rand"
	"time"
)

// TacticToString gets the tactic name as string from mitreattack.Tactic
func TacticToString(tactic mitreattack.Tactic) string {
	switch tactic {
	case mitreattack.CredentialAccess:
		return "Credential Access"
	case mitreattack.DefenseEvasion:
		return "Defense Evasion"
	case mitreattack.Discovery:
		return "Discovery"
	case mitreattack.Execution:
		return "Execution"
	case mitreattack.Exfiltration:
		return "Exfiltration"
	case mitreattack.Impact:
		return "Impact"
	case mitreattack.InitialAccess:
		return "Initial Access"
	case mitreattack.LateralMovement:
		return "Lateral Movement"
	case mitreattack.Persistence:
		return "Persistence"
	default:
		return "Unknown"
	}
}

func GetUnusedTactic(db *sql.DB) mitreattack.Tactic {
	tactics := []mitreattack.Tactic{
		mitreattack.CredentialAccess,
		mitreattack.DefenseEvasion,
		mitreattack.Discovery,
		mitreattack.Execution,
		mitreattack.Exfiltration,
		mitreattack.Impact,
		mitreattack.InitialAccess,
		mitreattack.LateralMovement,
		mitreattack.Persistence,
	}

	var unusedTactics []mitreattack.Tactic
	for _, tactic := range tactics {
		used, err := IsTacticUsed(db, TacticToString(tactic))
		if err != nil {
			logrus.Fatalf("Error checking if tactic is used: %v", err)
			continue
		}
		if !used {
			unusedTactics = append(unusedTactics, tactic)
		}
	}

	if len(unusedTactics) == 0 {
		err := ResetTactics(db)
		if err != nil {
			logrus.Fatalf("Error resetting tactics: %v\n", err)
		}
		return GetUnusedTactic(db)
	}

	rand.Seed(time.Now().UnixNano())
	selectedTactic := unusedTactics[rand.Intn(len(unusedTactics))]
	log.Printf("Unused tactic: %s\n", TacticToString(selectedTactic))
	return selectedTactic
}

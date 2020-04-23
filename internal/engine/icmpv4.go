package engine

import (
	"gitlab.com/Alvoras/pinknoise/internal/events"
	"gitlab.com/Alvoras/pinknoise/internal/rules"
)

func qualifyICMPv4Event(ev *events.ICMPv4Event) {
	var matches []rules.Rule

	for _, rules := range rules.GlobalRules {
		for _, rule := range rules {
			if rule.Layer != ev.Kind {
				continue
			}
			if rule.MatchICMPv4Event(*ev) {
				matches = append(matches, rule)
			}
		}
	}
	if len(matches) > 0 {
		for _, match := range matches {
			ev.AddTags(match.Tags)
			ev.AddMeta(match.Metadata)
			ev.AddRefs(match.References)
			ev.AddStatements(match.Statements)
		}
	}
}
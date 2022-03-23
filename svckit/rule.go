package svckit

func NewRule(triggerConditionOperator Operator, comparisonValue Event) map[Operator]Event {
	return map[Operator]Event{triggerConditionOperator: comparisonValue}
}
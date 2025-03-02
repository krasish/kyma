package beb

import (
	"reflect"

	eventingv1alpha1 "github.com/kyma-project/kyma/components/eventing-controller/api/v1alpha1"
)

// conditionsEquals checks if two list of conditions are equal.
func conditionsEquals(existing, expected []eventingv1alpha1.Condition) bool {
	// not equal if length is different
	if len(existing) != len(expected) {
		return false
	}

	// compile map of Conditions per ConditionType
	existingMap := make(map[eventingv1alpha1.ConditionType]eventingv1alpha1.Condition, len(existing))
	for _, value := range existing {
		existingMap[value.Type] = value
	}

	for _, value := range expected {
		if !conditionEquals(existingMap[value.Type], value) {
			return false
		}
	}

	return true
}

// conditionsEquals checks if two conditions are equal.
func conditionEquals(existing, expected eventingv1alpha1.Condition) bool {
	isTypeEqual := existing.Type == expected.Type
	isStatusEqual := existing.Status == expected.Status
	isReasonEqual := existing.Reason == expected.Reason
	isMessageEqual := existing.Message == expected.Message

	if !isStatusEqual || !isReasonEqual || !isMessageEqual || !isTypeEqual {
		return false
	}

	return true
}

func isSubscriptionStatusEqual(oldStatus, newStatus eventingv1alpha1.SubscriptionStatus) bool {
	oldStatusWithoutCond := oldStatus.DeepCopy()
	newStatusWithoutCond := newStatus.DeepCopy()

	// remove conditions, so that we don't compare them
	oldStatusWithoutCond.Conditions = []eventingv1alpha1.Condition{}
	newStatusWithoutCond.Conditions = []eventingv1alpha1.Condition{}

	return reflect.DeepEqual(oldStatusWithoutCond, newStatusWithoutCond) && conditionsEquals(oldStatus.Conditions, newStatus.Conditions)
}

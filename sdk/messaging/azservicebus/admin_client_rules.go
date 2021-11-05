// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/atom"
)

// RuleProperties represent the properties of a Subscription Rule
type RuleProperties struct {
	// RuleName of the rule
	RuleName string

	// TopicName is the topic for the subscription for this rule.
	TopicName string

	// SubscriptionName is the name of the subscription for this rule.
	SubscriptionName string

	// CreatedAt is the time this rule was created
	CreatedAt time.Time

	// Filter is the filter for this rule.
	Filter RuleFilter

	// Action is the action for this rule.
	Action RuleAction
}

// RuleFilters and implementations.
type (
	// RuleFilter can be used to set a filter condition for a rule.
	// More information on types of rule filters can be found here: https://docs.microsoft.com/en-us/azure/service-bus-messaging/topic-filters#filters
	RuleFilter interface {
		ruleFilter()
	}

	// CorrelationFilter holds a set of conditions that are matched against one or more of an arriving
	// message's user and system properties.
	CorrelationRuleFilter struct {
		ApplicationProperties map[string]string
		ContentType           *string
		CorrelationID         *string
		MessageID             *string
		ReplyTo               *string
		ReplyToSessionID      *string
		SessionID             *string
		Subject               *string
		To                    *string
	}

	// FalseFilter causes all arriving messages to be ignored.
	FalseRuleFilter struct{}

	// TrueFilter causes all arriving messages to be selected.
	TrueRuleFilter struct{}

	// SQLRuleFilter holds a SQL-like conditional expression that is evaluated in the broker against the arriving message's
	// user-defined properties and system properties.
	SQLRuleFilter struct {
		// Expression is the SQL expression for this rule.
		Expression string

		// Parameters are the parameters for the Expression.
		Parameters map[string]interface{}
	}
)

// RuleActions and implementations.
type (
	// RuleAction is an action when a rule is activated.
	RuleAction interface {
		ruleAction()
	}

	// SQLRuleAction can annotate the message by adding, removing, or replacing properties and their values.
	// The action uses a SQL-like expression that loosely leans on the SQL UPDATE statement syntax.
	// More information can be found here: https://docs.microsoft.com/azure/service-bus-messaging/topic-filters
	SQLRuleAction struct {
		// Expression is the SQL expression for this rule.
		Expression string

		// Parameters are the parameters for the Expression.
		Parameters map[string]interface{}
	}
)

func (a *SQLRuleAction) ruleAction()         {}
func (a *SQLRuleFilter) ruleFilter()         {}
func (a *TrueRuleFilter) ruleFilter()        {}
func (a *FalseRuleFilter) ruleFilter()       {}
func (a *CorrelationRuleFilter) ruleFilter() {}

type AddRuleResponse struct {
	Rule        *RuleProperties
	RawResponse *http.Response
}

type GetRuleResponse struct {
	Rule        *RuleProperties
	RawResponse *http.Response
}

type RuleExistsResponse struct {
	Exists      bool
	RawResponse *http.Response
}

type UpdateRuleResponse struct {
	Rule        *RuleProperties
	RawResponse *http.Response
}

type DeleteRuleResponse struct {
	RawResponse *http.Response
}

type ListRulesResponse struct {
	Rules       []*RuleProperties
	RawResponse *http.Response
}

// RulePropertiesPager provides iteration over ListRulesProperties pages.
type RulePropertiesPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// PageResponse returns the current RuleProperties.
	PageResponse() *ListRulesResponse

	// Err returns the last error encountered while paging.
	Err() error
}

// AddRule adds a new rule to a subscription.
func (ac *AdminClient) AddRule(ctx context.Context, properties *RuleProperties) (*AddRuleResponse, error) {
	resp, rule, err := ac.addOrUpdateRule(ctx, properties, true)

	if err != nil {
		return nil, err
	}

	return &AddRuleResponse{
		RawResponse: resp,
		Rule:        rule,
	}, nil
}

// GetRule gets an existing rule by name.
func (ac *AdminClient) GetRule(ctx context.Context, topicName string, subscriptionName string, ruleName string) (*GetRuleResponse, error) {
	var atomResp *atom.RuleEnvelope
	resp, err := ac.em.Get(ctx, fmtRuleURL(topicName, subscriptionName, ruleName), &atomResp)

	if err != nil {
		return nil, err
	}

	rule, err := newRuleProperties(topicName, subscriptionName, ruleName, &atomResp.Content.RuleDescription)

	if err != nil {
		return nil, err
	}

	return &GetRuleResponse{
		RawResponse: resp,
		Rule:        rule,
	}, nil
}

type ListRulesOptions struct {
	MaxPage
}

// ListRules returns a pager that can be used to list all the rules in a subscription.
func (ac *AdminClient) ListRules(ctx context.Context, topicName string, subscriptionName string) (*RulePropertiesPager, error) {
	ac.getRulePager(topicName, subscriptionName)
}

func (ac *AdminClient) getRulePager(topicName string, subscriptionName string, top int) ruleFeedPagerFunc {
	skip := 0

	return func(ctx context.Context) (*atom.RuleFeed, *http.Response, error) {
		url := fmtRuleURL(topicName, subscriptionName, "") + "?"

		if top > 0 {
			url += fmt.Sprintf("&$top=%d", top)
		}

		var atomResp *atom.RuleFeed
		resp, err := ac.em.Get(ctx, url, &atomResp)

		if err != nil {
			return nil, nil, err
		}

		if len(atomResp.Entries) == 0 {
			return nil, nil, nil
		}

		skip += len(atomResp.Entries)
		return atomResp, resp, nil
	}
}

// RuleExists checks if a subscription rule exists.
// Returns true if the rule is found
// (false, nil) if the rule is not found
// (false, err) if an error occurred while trying to check if the rule exists.
func (ac *AdminClient) RuleExists(ctx context.Context, topicName string, subscriptionName string, ruleName string) (*RuleExistsResponse, error) {
	resp, err := ac.em.Get(ctx, fmtRuleURL(topicName, subscriptionName, ruleName), &struct{}{})

	if err == nil {
		return &RuleExistsResponse{
			RawResponse: resp,
			Exists:      true,
		}, nil
	}

	if atom.NotFound(err) {
		return &RuleExistsResponse{
			RawResponse: resp,
			Exists:      false,
		}, nil
	}

	return nil, err
}

// UpdateRule updates an existing rule.
func (ac *AdminClient) UpdateRule(ctx context.Context, properties *RuleProperties) (*UpdateRuleResponse, error) {
	resp, rule, err := ac.addOrUpdateRule(ctx, properties, false)

	if err != nil {
		return nil, err
	}

	return &UpdateRuleResponse{
		RawResponse: resp,
		Rule:        rule,
	}, nil
}

// DeleteRule deletes a rule from a subscription.
func (ac *AdminClient) DeleteRule(ctx context.Context, topicName string, subscriptionName string, ruleName string) (*DeleteRuleResponse, error) {
	resp, err := ac.em.Delete(ctx, fmtRuleURL(topicName, subscriptionName, ruleName))

	if err != nil {
		return nil, err
	}

	return &DeleteRuleResponse{
		RawResponse: resp,
	}, nil
}

func fmtRuleURL(topicName, subscriptionName, optionalRuleName string) string {
	if optionalRuleName != "" {
		return fmt.Sprintf("/%s/Subscriptions/%s/Rules/%s", topicName, subscriptionName, optionalRuleName)
	}

	return fmt.Sprintf("/%s/Subscriptions/%s/Rules", topicName, subscriptionName)
}

// convert from ATOM entities

func newRuleProperties(topicName, subscriptionName, ruleName string, desc *atom.RuleDescription) (*RuleProperties, error) {
	props := &RuleProperties{
		RuleName:         ruleName,
		TopicName:        topicName,
		SubscriptionName: subscriptionName,
		CreatedAt:        desc.CreatedAt.Time,
	}

	switch desc.Action.Type {
	case "EmptyRuleAction":
		break
	case "SqlRuleAction":
		props.Action = &SQLRuleAction{
			Expression: desc.Action.SQLExpression,
			// TODO: Parameters
			// Parameters:  ,
		}
	default:
		return nil, fmt.Errorf("unknown Action.Type %s", desc.Action.Type)
	}

	switch desc.Filter.Type {
	case "TrueFilter":
		props.Filter = &TrueRuleFilter{}
	case "FalseFilter":
		props.Filter = &FalseRuleFilter{}
	case "SqlFilter":
		props.Filter = &SQLRuleFilter{
			Expression: *desc.Filter.SQLExpression,
			// TODO: SQLParameters
			// Parameters: ,
		}
	case "CorrelationFilter":
		props.Filter = &CorrelationRuleFilter{
			ContentType:      desc.Filter.ContentType,
			CorrelationID:    desc.Filter.CorrelationID,
			MessageID:        desc.Filter.MessageID,
			ReplyTo:          desc.Filter.ReplyTo,
			ReplyToSessionID: desc.Filter.ReplyToSessionID,
			SessionID:        desc.Filter.SessionID,
			Subject:          desc.Filter.Label,
			To:               desc.Filter.To,
		}
	default:
		return nil, fmt.Errorf("unknown Filter.Type %s", desc.Filter.Type)
	}

	return props, nil
}

// convert to ATOM entities

func newRuleEnvelope(props *RuleProperties) (*atom.RuleEnvelope, error) {
	filterDesc, err := newFilterDescription(props.Filter)

	if err != nil {
		return nil, err
	}

	actionDesc, err := newActionDescription(props.Action)

	if err != nil {
		return nil, err
	}

	desc := &atom.RuleDescription{
		Filter: *filterDesc,
		Action: actionDesc,
	}

	return atom.WrapWithRuleEnvelope(desc), nil
}

func newFilterDescription(filter RuleFilter) (*atom.FilterDescription, error) {
	var desc atom.FilterDescription

	switch asType := filter.(type) {
	case *SQLRuleFilter:
		desc = atom.SQLFilter{
			Expression: asType.Expression,
			// TODO: need to handle parameters
		}.ToFilterDescription()
	case *CorrelationRuleFilter:
		desc = atom.CorrelationFilter{
			CorrelationID:    asType.CorrelationID,
			MessageID:        asType.MessageID,
			To:               asType.To,
			ReplyTo:          asType.ReplyTo,
			Label:            asType.Subject,
			SessionID:        asType.SessionID,
			ReplyToSessionID: asType.ReplyToSessionID,
			ContentType:      asType.ContentType,
		}.ToFilterDescription()
	case *FalseRuleFilter:
		desc = atom.FalseFilter{}.ToFilterDescription()
	case *TrueRuleFilter:
		desc = atom.TrueFilter{}.ToFilterDescription()
	default:
		return nil, fmt.Errorf("unknown filter type %T", filter)
	}

	return &desc, nil
}

func newActionDescription(action RuleAction) (*atom.ActionDescription, error) {
	var desc atom.ActionDescription

	if action == nil {
		return nil, nil
	}

	switch asType := action.(type) {
	case *SQLRuleAction:
		desc = atom.SQLAction{
			Expression: asType.Expression,
			// TODO: parameters
		}.ToActionDescription()
	default:
		return nil, fmt.Errorf("unknown action type %T", action)
	}

	return &desc, nil
}

func (ac *AdminClient) addOrUpdateRule(ctx context.Context, properties *RuleProperties, creating bool) (*http.Response, *RuleProperties, error) {
	putBody, err := newRuleEnvelope(properties)

	if err != nil {
		return nil, nil, err
	}

	var atomResp *atom.RuleEnvelope
	resp, err := ac.em.Put(ctx,
		fmtRuleURL(properties.TopicName, properties.SubscriptionName, properties.RuleName),
		putBody,
		&atomResp)

	if err != nil {
		return nil, nil, err
	}

	rule, err := newRuleProperties(
		properties.TopicName,
		properties.SubscriptionName,
		properties.RuleName,
		&atomResp.Content.RuleDescription)

	if err != nil {
		return nil, nil, err
	}

	return resp, rule, nil
}

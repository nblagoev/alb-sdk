// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// WafWhitelistLog waf whitelist log
// swagger:model WafWhitelistLog
type WafWhitelistLog struct {

	// Actions generated by this rule. Enum options - WAF_POLICY_WHITELIST_ACTION_ALLOW, WAF_POLICY_WHITELIST_ACTION_DETECTION_MODE, WAF_POLICY_WHITELIST_ACTION_CONTINUE. Field deprecated in 20.1.3. Field introduced in 18.2.3.
	Actions []string `json:"actions,omitempty"`

	// Name of the matched rule. Field deprecated in 20.1.3. Field introduced in 18.2.3.
	RuleName *string `json:"rule_name,omitempty"`
}

// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// MicroServiceContainer micro service container
// swagger:model MicroServiceContainer
type MicroServiceContainer struct {

	// ID of the container.
	ContainerID *string `json:"container_id,omitempty"`

	// ID or name of the host where the container is.
	Host *string `json:"host,omitempty"`

	// IP Address of the container.
	// Required: true
	IP *IPAddr `json:"ip"`

	// Port nunber of the instance.
	Port *int32 `json:"port,omitempty"`

	// Marathon Task ID of the instance.
	TaskID *string `json:"task_id,omitempty"`
}

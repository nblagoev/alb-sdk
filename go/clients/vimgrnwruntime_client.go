/***************************************************************************
 *
 * AVI CONFIDENTIAL
 * __________________
 *
 * [2013] - [2018] Avi Networks Incorporated
 * All Rights Reserved.
 *
 * NOTICE: All information contained herein is, and remains the property
 * of Avi Networks Incorporated and its suppliers, if any. The intellectual
 * and technical concepts contained herein are proprietary to Avi Networks
 * Incorporated, and its suppliers and are covered by U.S. and Foreign
 * Patents, patents in process, and are protected by trade secret or
 * copyright law, and other laws. Dissemination of this information or
 * reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Avi Networks Incorporated.
 */

package clients

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

import (
	"github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/sdk/go/session"
)

// VIMgrNWRuntimeClient is a client for avi VIMgrNWRuntime resource
type VIMgrNWRuntimeClient struct {
	aviSession *session.AviSession
}

// NewVIMgrNWRuntimeClient creates a new client for VIMgrNWRuntime resource
func NewVIMgrNWRuntimeClient(aviSession *session.AviSession) *VIMgrNWRuntimeClient {
	return &VIMgrNWRuntimeClient{aviSession: aviSession}
}

func (client *VIMgrNWRuntimeClient) getAPIPath(uuid string, options ...session.ApiOptionsParams) (string, error) {
	path := "api/vimgrnwruntime"
	var err error
	if uuid != "" {
		path += "/" + uuid
	} else {
		path, err = session.SetApiFilter(path, options...)
		if err != nil {
			return "", err
		}
	}
	return path, nil
}

// GetAll is a collection API to get a list of VIMgrNWRuntime objects
func (client *VIMgrNWRuntimeClient) GetAll(options ...session.ApiOptionsParams) ([]*models.VIMgrNWRuntime, error) {
	var plist []*models.VIMgrNWRuntime
	path, err := client.getAPIPath("", options...)
	if err == nil {
		err = client.aviSession.GetCollection(path, &plist, options...)
	}
	return plist, err
}

// Get an existing VIMgrNWRuntime by uuid
func (client *VIMgrNWRuntimeClient) Get(uuid string, options ...session.ApiOptionsParams) (*models.VIMgrNWRuntime, error) {
	var obj *models.VIMgrNWRuntime
	path, _ := client.getAPIPath(uuid)
	err := client.aviSession.Get(path, &obj, options...)
	return obj, err
}

// GetByName - Get an existing VIMgrNWRuntime by name
func (client *VIMgrNWRuntimeClient) GetByName(name string, options ...session.ApiOptionsParams) (*models.VIMgrNWRuntime, error) {
	var obj *models.VIMgrNWRuntime
	err := client.aviSession.GetObjectByName("vimgrnwruntime", name, &obj, options...)
	return obj, err
}

// GetObject - Get an existing VIMgrNWRuntime by filters like name, cloud, tenant
// Api creates VIMgrNWRuntime object with every call.
func (client *VIMgrNWRuntimeClient) GetObject(options ...session.ApiOptionsParams) (*models.VIMgrNWRuntime, error) {
	var obj *models.VIMgrNWRuntime
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("vimgrnwruntime", newOptions...)
	return obj, err
}

// Create a new VIMgrNWRuntime object
func (client *VIMgrNWRuntimeClient) Create(obj *models.VIMgrNWRuntime, options ...session.ApiOptionsParams) (*models.VIMgrNWRuntime, error) {
	var robj *models.VIMgrNWRuntime
	path, _ := client.getAPIPath("")
	err := client.aviSession.Post(path, obj, &robj, options...)
	return robj, err
}

// Update an existing VIMgrNWRuntime object
func (client *VIMgrNWRuntimeClient) Update(obj *models.VIMgrNWRuntime, options ...session.ApiOptionsParams) (*models.VIMgrNWRuntime, error) {
	var robj *models.VIMgrNWRuntime
	path, _ := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, options...)
	return robj, err
}

// Patch an existing VIMgrNWRuntime object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.VIMgrNWRuntime
// or it should be json compatible of form map[string]interface{}
func (client *VIMgrNWRuntimeClient) Patch(uuid string, patch interface{}, patchOp string, options ...session.ApiOptionsParams) (*models.VIMgrNWRuntime, error) {
	var robj *models.VIMgrNWRuntime
	path, _ := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, options...)
	return robj, err
}

// Delete an existing VIMgrNWRuntime object with a given UUID
func (client *VIMgrNWRuntimeClient) Delete(uuid string, options ...session.ApiOptionsParams) error {
	path, _ := client.getAPIPath(uuid)
	if len(options) == 0 {
		return client.aviSession.Delete(path)
	} else {
		return client.aviSession.DeleteObject(path, options...)
	}
}

// DeleteByName - Delete an existing VIMgrNWRuntime object with a given name
func (client *VIMgrNWRuntimeClient) DeleteByName(name string, options ...session.ApiOptionsParams) error {
	res, err := client.GetByName(name, options...)
	if err != nil {
		return err
	}
	return client.Delete(*res.UUID, options...)
}

// GetAviSession
func (client *VIMgrNWRuntimeClient) GetAviSession() *session.AviSession {
	return client.aviSession
}

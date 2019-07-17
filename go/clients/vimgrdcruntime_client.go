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

// VIMgrDCRuntimeClient is a client for avi VIMgrDCRuntime resource
type VIMgrDCRuntimeClient struct {
	aviSession *session.AviSession
}

// NewVIMgrDCRuntimeClient creates a new client for VIMgrDCRuntime resource
func NewVIMgrDCRuntimeClient(aviSession *session.AviSession) *VIMgrDCRuntimeClient {
	return &VIMgrDCRuntimeClient{aviSession: aviSession}
}

func (client *VIMgrDCRuntimeClient) getAPIPath(uuid string, options ...session.ApiOptionsParams) (string, error) {
	path := "api/vimgrdcruntime"
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

// GetAll is a collection API to get a list of VIMgrDCRuntime objects
func (client *VIMgrDCRuntimeClient) GetAll(options ...session.ApiOptionsParams) ([]*models.VIMgrDCRuntime, error) {
	var plist []*models.VIMgrDCRuntime
	path, err := client.getAPIPath("", options...)
	if err == nil {
		err = client.aviSession.GetCollection(path, &plist, options...)
	}
	return plist, err
}

// Get an existing VIMgrDCRuntime by uuid
func (client *VIMgrDCRuntimeClient) Get(uuid string, options ...session.ApiOptionsParams) (*models.VIMgrDCRuntime, error) {
	var obj *models.VIMgrDCRuntime
	path, _ := client.getAPIPath(uuid)
	err := client.aviSession.Get(path, &obj, options...)
	return obj, err
}

// GetByName - Get an existing VIMgrDCRuntime by name
func (client *VIMgrDCRuntimeClient) GetByName(name string, options ...session.ApiOptionsParams) (*models.VIMgrDCRuntime, error) {
	var obj *models.VIMgrDCRuntime
	err := client.aviSession.GetObjectByName("vimgrdcruntime", name, &obj, options...)
	return obj, err
}

// GetObject - Get an existing VIMgrDCRuntime by filters like name, cloud, tenant
// Api creates VIMgrDCRuntime object with every call.
func (client *VIMgrDCRuntimeClient) GetObject(options ...session.ApiOptionsParams) (*models.VIMgrDCRuntime, error) {
	var obj *models.VIMgrDCRuntime
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("vimgrdcruntime", newOptions...)
	return obj, err
}

// Create a new VIMgrDCRuntime object
func (client *VIMgrDCRuntimeClient) Create(obj *models.VIMgrDCRuntime, options ...session.ApiOptionsParams) (*models.VIMgrDCRuntime, error) {
	var robj *models.VIMgrDCRuntime
	path, _ := client.getAPIPath("")
	err := client.aviSession.Post(path, obj, &robj, options...)
	return robj, err
}

// Update an existing VIMgrDCRuntime object
func (client *VIMgrDCRuntimeClient) Update(obj *models.VIMgrDCRuntime, options ...session.ApiOptionsParams) (*models.VIMgrDCRuntime, error) {
	var robj *models.VIMgrDCRuntime
	path, _ := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, options...)
	return robj, err
}

// Patch an existing VIMgrDCRuntime object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.VIMgrDCRuntime
// or it should be json compatible of form map[string]interface{}
func (client *VIMgrDCRuntimeClient) Patch(uuid string, patch interface{}, patchOp string, options ...session.ApiOptionsParams) (*models.VIMgrDCRuntime, error) {
	var robj *models.VIMgrDCRuntime
	path, _ := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, options...)
	return robj, err
}

// Delete an existing VIMgrDCRuntime object with a given UUID
func (client *VIMgrDCRuntimeClient) Delete(uuid string, options ...session.ApiOptionsParams) error {
	path, _ := client.getAPIPath(uuid)
	if len(options) == 0 {
		return client.aviSession.Delete(path)
	} else {
		return client.aviSession.DeleteObject(path, options...)
	}
}

// DeleteByName - Delete an existing VIMgrDCRuntime object with a given name
func (client *VIMgrDCRuntimeClient) DeleteByName(name string, options ...session.ApiOptionsParams) error {
	res, err := client.GetByName(name, options...)
	if err != nil {
		return err
	}
	return client.Delete(*res.UUID, options...)
}

// GetAviSession
func (client *VIMgrDCRuntimeClient) GetAviSession() *session.AviSession {
	return client.aviSession
}

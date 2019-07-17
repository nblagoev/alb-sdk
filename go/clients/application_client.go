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

// ApplicationClient is a client for avi Application resource
type ApplicationClient struct {
	aviSession *session.AviSession
}

// NewApplicationClient creates a new client for Application resource
func NewApplicationClient(aviSession *session.AviSession) *ApplicationClient {
	return &ApplicationClient{aviSession: aviSession}
}

func (client *ApplicationClient) getAPIPath(uuid string, options ...session.ApiOptionsParams) (string, error) {
	path := "api/application"
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

// GetAll is a collection API to get a list of Application objects
func (client *ApplicationClient) GetAll(options ...session.ApiOptionsParams) ([]*models.Application, error) {
	var plist []*models.Application
	path, err := client.getAPIPath("", options...)
	if err == nil {
		err = client.aviSession.GetCollection(path, &plist, options...)
	}
	return plist, err
}

// Get an existing Application by uuid
func (client *ApplicationClient) Get(uuid string, options ...session.ApiOptionsParams) (*models.Application, error) {
	var obj *models.Application
	path, _ := client.getAPIPath(uuid)
	err := client.aviSession.Get(path, &obj, options...)
	return obj, err
}

// GetByName - Get an existing Application by name
func (client *ApplicationClient) GetByName(name string, options ...session.ApiOptionsParams) (*models.Application, error) {
	var obj *models.Application
	err := client.aviSession.GetObjectByName("application", name, &obj, options...)
	return obj, err
}

// GetObject - Get an existing Application by filters like name, cloud, tenant
// Api creates Application object with every call.
func (client *ApplicationClient) GetObject(options ...session.ApiOptionsParams) (*models.Application, error) {
	var obj *models.Application
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("application", newOptions...)
	return obj, err
}

// Create a new Application object
func (client *ApplicationClient) Create(obj *models.Application, options ...session.ApiOptionsParams) (*models.Application, error) {
	var robj *models.Application
	path, _ := client.getAPIPath("")
	err := client.aviSession.Post(path, obj, &robj, options...)
	return robj, err
}

// Update an existing Application object
func (client *ApplicationClient) Update(obj *models.Application, options ...session.ApiOptionsParams) (*models.Application, error) {
	var robj *models.Application
	path, _ := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, options...)
	return robj, err
}

// Patch an existing Application object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.Application
// or it should be json compatible of form map[string]interface{}
func (client *ApplicationClient) Patch(uuid string, patch interface{}, patchOp string, options ...session.ApiOptionsParams) (*models.Application, error) {
	var robj *models.Application
	path, _ := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, options...)
	return robj, err
}

// Delete an existing Application object with a given UUID
func (client *ApplicationClient) Delete(uuid string, options ...session.ApiOptionsParams) error {
	path, _ := client.getAPIPath(uuid)
	if len(options) == 0 {
		return client.aviSession.Delete(path)
	} else {
		return client.aviSession.DeleteObject(path, options...)
	}
}

// DeleteByName - Delete an existing Application object with a given name
func (client *ApplicationClient) DeleteByName(name string, options ...session.ApiOptionsParams) error {
	res, err := client.GetByName(name, options...)
	if err != nil {
		return err
	}
	return client.Delete(*res.UUID, options...)
}

// GetAviSession
func (client *ApplicationClient) GetAviSession() *session.AviSession {
	return client.aviSession
}

// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// SourceApplicationInfo The information about the application.
type SourceApplicationInfo struct {

	// The OCID of the workspace containing the application. This allows cross workspace deployment to publish one application in one workspace into the current workspace specified in thi operation.
	WorkspaceId *string `mandatory:"false" json:"workspaceId"`

	// The source application key to use when creating the application.
	ApplicationKey *string `mandatory:"false" json:"applicationKey"`

	// The source application version of the application.
	ApplicationVersion *string `mandatory:"false" json:"applicationVersion"`

	// The last patch key for the application.
	LastPatchKey *string `mandatory:"false" json:"lastPatchKey"`
}

func (m SourceApplicationInfo) String() string {
	return common.PointerString(m)
}
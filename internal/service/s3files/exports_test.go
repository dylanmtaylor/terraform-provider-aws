// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package s3files

var (
	ResourceFileSystem                   = newFileSystemResource
	ResourceMountTarget                  = newMountTargetResource
	ResourceFileSystemPolicy             = newFileSystemPolicyResource
	ResourceAccessPoint                  = newAccessPointResource
	ResourceSynchronizationConfiguration = newSynchronizationConfigurationResource

	FindFileSystemByID       = findFileSystemByID
	FindMountTargetByID      = findMountTargetByID
	FindFileSystemPolicyByID = findFileSystemPolicyByID
	FindAccessPointByID      = findAccessPointByID
	FindSyncConfigByID       = findSyncConfigByID
)

// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	release "github.com/jsdidierlaurent/azure-devops-go-api/azuredevops/release"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateFolder provides a mock function with given fields: _a0, _a1
func (_m *Client) CreateFolder(_a0 context.Context, _a1 release.CreateFolderArgs) (*release.Folder, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateFolder")
	}

	var r0 *release.Folder
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.CreateFolderArgs) (*release.Folder, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.CreateFolderArgs) *release.Folder); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.Folder)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.CreateFolderArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRelease provides a mock function with given fields: _a0, _a1
func (_m *Client) CreateRelease(_a0 context.Context, _a1 release.CreateReleaseArgs) (*release.Release, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateRelease")
	}

	var r0 *release.Release
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.CreateReleaseArgs) (*release.Release, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.CreateReleaseArgs) *release.Release); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.Release)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.CreateReleaseArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReleaseDefinition provides a mock function with given fields: _a0, _a1
func (_m *Client) CreateReleaseDefinition(_a0 context.Context, _a1 release.CreateReleaseDefinitionArgs) (*release.ReleaseDefinition, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateReleaseDefinition")
	}

	var r0 *release.ReleaseDefinition
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.CreateReleaseDefinitionArgs) (*release.ReleaseDefinition, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.CreateReleaseDefinitionArgs) *release.ReleaseDefinition); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.ReleaseDefinition)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.CreateReleaseDefinitionArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFolder provides a mock function with given fields: _a0, _a1
func (_m *Client) DeleteFolder(_a0 context.Context, _a1 release.DeleteFolderArgs) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFolder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, release.DeleteFolderArgs) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReleaseDefinition provides a mock function with given fields: _a0, _a1
func (_m *Client) DeleteReleaseDefinition(_a0 context.Context, _a1 release.DeleteReleaseDefinitionArgs) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteReleaseDefinition")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, release.DeleteReleaseDefinitionArgs) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetApprovals provides a mock function with given fields: _a0, _a1
func (_m *Client) GetApprovals(_a0 context.Context, _a1 release.GetApprovalsArgs) (*release.GetApprovalsResponseValue, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetApprovals")
	}

	var r0 *release.GetApprovalsResponseValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetApprovalsArgs) (*release.GetApprovalsResponseValue, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetApprovalsArgs) *release.GetApprovalsResponseValue); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.GetApprovalsResponseValue)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetApprovalsArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefinitionRevision provides a mock function with given fields: _a0, _a1
func (_m *Client) GetDefinitionRevision(_a0 context.Context, _a1 release.GetDefinitionRevisionArgs) (io.ReadCloser, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetDefinitionRevision")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetDefinitionRevisionArgs) (io.ReadCloser, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetDefinitionRevisionArgs) io.ReadCloser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetDefinitionRevisionArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeployments provides a mock function with given fields: _a0, _a1
func (_m *Client) GetDeployments(_a0 context.Context, _a1 release.GetDeploymentsArgs) (*release.GetDeploymentsResponseValue, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetDeployments")
	}

	var r0 *release.GetDeploymentsResponseValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetDeploymentsArgs) (*release.GetDeploymentsResponseValue, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetDeploymentsArgs) *release.GetDeploymentsResponseValue); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.GetDeploymentsResponseValue)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetDeploymentsArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFolders provides a mock function with given fields: _a0, _a1
func (_m *Client) GetFolders(_a0 context.Context, _a1 release.GetFoldersArgs) (*[]release.Folder, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetFolders")
	}

	var r0 *[]release.Folder
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetFoldersArgs) (*[]release.Folder, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetFoldersArgs) *[]release.Folder); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]release.Folder)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetFoldersArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogs provides a mock function with given fields: _a0, _a1
func (_m *Client) GetLogs(_a0 context.Context, _a1 release.GetLogsArgs) (io.ReadCloser, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetLogs")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetLogsArgs) (io.ReadCloser, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetLogsArgs) io.ReadCloser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetLogsArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetManualIntervention provides a mock function with given fields: _a0, _a1
func (_m *Client) GetManualIntervention(_a0 context.Context, _a1 release.GetManualInterventionArgs) (*release.ManualIntervention, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetManualIntervention")
	}

	var r0 *release.ManualIntervention
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetManualInterventionArgs) (*release.ManualIntervention, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetManualInterventionArgs) *release.ManualIntervention); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.ManualIntervention)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetManualInterventionArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetManualInterventions provides a mock function with given fields: _a0, _a1
func (_m *Client) GetManualInterventions(_a0 context.Context, _a1 release.GetManualInterventionsArgs) (*[]release.ManualIntervention, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetManualInterventions")
	}

	var r0 *[]release.ManualIntervention
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetManualInterventionsArgs) (*[]release.ManualIntervention, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetManualInterventionsArgs) *[]release.ManualIntervention); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]release.ManualIntervention)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetManualInterventionsArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRelease provides a mock function with given fields: _a0, _a1
func (_m *Client) GetRelease(_a0 context.Context, _a1 release.GetReleaseArgs) (*release.Release, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetRelease")
	}

	var r0 *release.Release
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseArgs) (*release.Release, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseArgs) *release.Release); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.Release)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetReleaseArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReleaseDefinition provides a mock function with given fields: _a0, _a1
func (_m *Client) GetReleaseDefinition(_a0 context.Context, _a1 release.GetReleaseDefinitionArgs) (*release.ReleaseDefinition, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetReleaseDefinition")
	}

	var r0 *release.ReleaseDefinition
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseDefinitionArgs) (*release.ReleaseDefinition, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseDefinitionArgs) *release.ReleaseDefinition); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.ReleaseDefinition)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetReleaseDefinitionArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReleaseDefinitionHistory provides a mock function with given fields: _a0, _a1
func (_m *Client) GetReleaseDefinitionHistory(_a0 context.Context, _a1 release.GetReleaseDefinitionHistoryArgs) (*[]release.ReleaseDefinitionRevision, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetReleaseDefinitionHistory")
	}

	var r0 *[]release.ReleaseDefinitionRevision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseDefinitionHistoryArgs) (*[]release.ReleaseDefinitionRevision, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseDefinitionHistoryArgs) *[]release.ReleaseDefinitionRevision); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]release.ReleaseDefinitionRevision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetReleaseDefinitionHistoryArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReleaseDefinitions provides a mock function with given fields: _a0, _a1
func (_m *Client) GetReleaseDefinitions(_a0 context.Context, _a1 release.GetReleaseDefinitionsArgs) (*release.GetReleaseDefinitionsResponseValue, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetReleaseDefinitions")
	}

	var r0 *release.GetReleaseDefinitionsResponseValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseDefinitionsArgs) (*release.GetReleaseDefinitionsResponseValue, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseDefinitionsArgs) *release.GetReleaseDefinitionsResponseValue); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.GetReleaseDefinitionsResponseValue)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetReleaseDefinitionsArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReleaseEnvironment provides a mock function with given fields: _a0, _a1
func (_m *Client) GetReleaseEnvironment(_a0 context.Context, _a1 release.GetReleaseEnvironmentArgs) (*release.ReleaseEnvironment, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetReleaseEnvironment")
	}

	var r0 *release.ReleaseEnvironment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseEnvironmentArgs) (*release.ReleaseEnvironment, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseEnvironmentArgs) *release.ReleaseEnvironment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.ReleaseEnvironment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetReleaseEnvironmentArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReleaseRevision provides a mock function with given fields: _a0, _a1
func (_m *Client) GetReleaseRevision(_a0 context.Context, _a1 release.GetReleaseRevisionArgs) (io.ReadCloser, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetReleaseRevision")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseRevisionArgs) (io.ReadCloser, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseRevisionArgs) io.ReadCloser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetReleaseRevisionArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReleaseTaskAttachmentContent provides a mock function with given fields: _a0, _a1
func (_m *Client) GetReleaseTaskAttachmentContent(_a0 context.Context, _a1 release.GetReleaseTaskAttachmentContentArgs) (io.ReadCloser, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetReleaseTaskAttachmentContent")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseTaskAttachmentContentArgs) (io.ReadCloser, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseTaskAttachmentContentArgs) io.ReadCloser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetReleaseTaskAttachmentContentArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReleaseTaskAttachments provides a mock function with given fields: _a0, _a1
func (_m *Client) GetReleaseTaskAttachments(_a0 context.Context, _a1 release.GetReleaseTaskAttachmentsArgs) (*[]release.ReleaseTaskAttachment, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetReleaseTaskAttachments")
	}

	var r0 *[]release.ReleaseTaskAttachment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseTaskAttachmentsArgs) (*[]release.ReleaseTaskAttachment, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleaseTaskAttachmentsArgs) *[]release.ReleaseTaskAttachment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]release.ReleaseTaskAttachment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetReleaseTaskAttachmentsArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReleases provides a mock function with given fields: _a0, _a1
func (_m *Client) GetReleases(_a0 context.Context, _a1 release.GetReleasesArgs) (*release.GetReleasesResponseValue, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetReleases")
	}

	var r0 *release.GetReleasesResponseValue
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleasesArgs) (*release.GetReleasesResponseValue, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetReleasesArgs) *release.GetReleasesResponseValue); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.GetReleasesResponseValue)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetReleasesArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaskLog provides a mock function with given fields: _a0, _a1
func (_m *Client) GetTaskLog(_a0 context.Context, _a1 release.GetTaskLogArgs) (io.ReadCloser, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskLog")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.GetTaskLogArgs) (io.ReadCloser, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.GetTaskLogArgs) io.ReadCloser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.GetTaskLogArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFolder provides a mock function with given fields: _a0, _a1
func (_m *Client) UpdateFolder(_a0 context.Context, _a1 release.UpdateFolderArgs) (*release.Folder, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFolder")
	}

	var r0 *release.Folder
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateFolderArgs) (*release.Folder, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateFolderArgs) *release.Folder); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.Folder)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.UpdateFolderArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGates provides a mock function with given fields: _a0, _a1
func (_m *Client) UpdateGates(_a0 context.Context, _a1 release.UpdateGatesArgs) (*release.ReleaseGates, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateGates")
	}

	var r0 *release.ReleaseGates
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateGatesArgs) (*release.ReleaseGates, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateGatesArgs) *release.ReleaseGates); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.ReleaseGates)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.UpdateGatesArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateManualIntervention provides a mock function with given fields: _a0, _a1
func (_m *Client) UpdateManualIntervention(_a0 context.Context, _a1 release.UpdateManualInterventionArgs) (*release.ManualIntervention, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateManualIntervention")
	}

	var r0 *release.ManualIntervention
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateManualInterventionArgs) (*release.ManualIntervention, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateManualInterventionArgs) *release.ManualIntervention); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.ManualIntervention)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.UpdateManualInterventionArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRelease provides a mock function with given fields: _a0, _a1
func (_m *Client) UpdateRelease(_a0 context.Context, _a1 release.UpdateReleaseArgs) (*release.Release, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRelease")
	}

	var r0 *release.Release
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateReleaseArgs) (*release.Release, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateReleaseArgs) *release.Release); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.Release)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.UpdateReleaseArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReleaseApproval provides a mock function with given fields: _a0, _a1
func (_m *Client) UpdateReleaseApproval(_a0 context.Context, _a1 release.UpdateReleaseApprovalArgs) (*release.ReleaseApproval, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateReleaseApproval")
	}

	var r0 *release.ReleaseApproval
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateReleaseApprovalArgs) (*release.ReleaseApproval, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateReleaseApprovalArgs) *release.ReleaseApproval); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.ReleaseApproval)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.UpdateReleaseApprovalArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReleaseDefinition provides a mock function with given fields: _a0, _a1
func (_m *Client) UpdateReleaseDefinition(_a0 context.Context, _a1 release.UpdateReleaseDefinitionArgs) (*release.ReleaseDefinition, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateReleaseDefinition")
	}

	var r0 *release.ReleaseDefinition
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateReleaseDefinitionArgs) (*release.ReleaseDefinition, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateReleaseDefinitionArgs) *release.ReleaseDefinition); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.ReleaseDefinition)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.UpdateReleaseDefinitionArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReleaseEnvironment provides a mock function with given fields: _a0, _a1
func (_m *Client) UpdateReleaseEnvironment(_a0 context.Context, _a1 release.UpdateReleaseEnvironmentArgs) (*release.ReleaseEnvironment, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateReleaseEnvironment")
	}

	var r0 *release.ReleaseEnvironment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateReleaseEnvironmentArgs) (*release.ReleaseEnvironment, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateReleaseEnvironmentArgs) *release.ReleaseEnvironment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.ReleaseEnvironment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.UpdateReleaseEnvironmentArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReleaseResource provides a mock function with given fields: _a0, _a1
func (_m *Client) UpdateReleaseResource(_a0 context.Context, _a1 release.UpdateReleaseResourceArgs) (*release.Release, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateReleaseResource")
	}

	var r0 *release.Release
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateReleaseResourceArgs) (*release.Release, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, release.UpdateReleaseResourceArgs) *release.Release); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*release.Release)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, release.UpdateReleaseResourceArgs) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

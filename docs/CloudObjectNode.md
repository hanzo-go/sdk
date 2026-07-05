# CloudObjectNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoQuery** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CpuSize** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EnableRemoteApp** | Pointer to **bool** |  | [optional] 
**IsPermanent** | Pointer to **bool** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**MachineName** | Pointer to **string** |  | [optional] 
**MemSize** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Patches** | Pointer to [**[]CloudObjectPatch**](CloudObjectPatch.md) |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**RemoteApps** | Pointer to [**[]CloudObjectRemoteApp**](CloudObjectRemoteApp.md) |  | [optional] 
**RemotePassword** | Pointer to **string** |  | [optional] 
**RemotePort** | Pointer to **int64** |  | [optional] 
**RemoteProtocol** | Pointer to **string** |  | [optional] 
**RemoteUsername** | Pointer to **string** |  | [optional] 
**Services** | Pointer to [**[]CloudObjectService**](CloudObjectService.md) |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudObjectNode

`func NewCloudObjectNode() *CloudObjectNode`

NewCloudObjectNode instantiates a new CloudObjectNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectNodeWithDefaults

`func NewCloudObjectNodeWithDefaults() *CloudObjectNode`

NewCloudObjectNodeWithDefaults instantiates a new CloudObjectNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoQuery

`func (o *CloudObjectNode) GetAutoQuery() bool`

GetAutoQuery returns the AutoQuery field if non-nil, zero value otherwise.

### GetAutoQueryOk

`func (o *CloudObjectNode) GetAutoQueryOk() (*bool, bool)`

GetAutoQueryOk returns a tuple with the AutoQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoQuery

`func (o *CloudObjectNode) SetAutoQuery(v bool)`

SetAutoQuery sets AutoQuery field to given value.

### HasAutoQuery

`func (o *CloudObjectNode) HasAutoQuery() bool`

HasAutoQuery returns a boolean if a field has been set.

### GetCategory

`func (o *CloudObjectNode) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudObjectNode) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudObjectNode) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudObjectNode) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCpuSize

`func (o *CloudObjectNode) GetCpuSize() string`

GetCpuSize returns the CpuSize field if non-nil, zero value otherwise.

### GetCpuSizeOk

`func (o *CloudObjectNode) GetCpuSizeOk() (*string, bool)`

GetCpuSizeOk returns a tuple with the CpuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSize

`func (o *CloudObjectNode) SetCpuSize(v string)`

SetCpuSize sets CpuSize field to given value.

### HasCpuSize

`func (o *CloudObjectNode) HasCpuSize() bool`

HasCpuSize returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CloudObjectNode) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudObjectNode) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudObjectNode) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudObjectNode) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *CloudObjectNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudObjectNode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudObjectNode) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudObjectNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudObjectNode) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudObjectNode) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudObjectNode) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudObjectNode) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableRemoteApp

`func (o *CloudObjectNode) GetEnableRemoteApp() bool`

GetEnableRemoteApp returns the EnableRemoteApp field if non-nil, zero value otherwise.

### GetEnableRemoteAppOk

`func (o *CloudObjectNode) GetEnableRemoteAppOk() (*bool, bool)`

GetEnableRemoteAppOk returns a tuple with the EnableRemoteApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteApp

`func (o *CloudObjectNode) SetEnableRemoteApp(v bool)`

SetEnableRemoteApp sets EnableRemoteApp field to given value.

### HasEnableRemoteApp

`func (o *CloudObjectNode) HasEnableRemoteApp() bool`

HasEnableRemoteApp returns a boolean if a field has been set.

### GetIsPermanent

`func (o *CloudObjectNode) GetIsPermanent() bool`

GetIsPermanent returns the IsPermanent field if non-nil, zero value otherwise.

### GetIsPermanentOk

`func (o *CloudObjectNode) GetIsPermanentOk() (*bool, bool)`

GetIsPermanentOk returns a tuple with the IsPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPermanent

`func (o *CloudObjectNode) SetIsPermanent(v bool)`

SetIsPermanent sets IsPermanent field to given value.

### HasIsPermanent

`func (o *CloudObjectNode) HasIsPermanent() bool`

HasIsPermanent returns a boolean if a field has been set.

### GetLanguage

`func (o *CloudObjectNode) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CloudObjectNode) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CloudObjectNode) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CloudObjectNode) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMachineName

`func (o *CloudObjectNode) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *CloudObjectNode) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *CloudObjectNode) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *CloudObjectNode) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### GetMemSize

`func (o *CloudObjectNode) GetMemSize() string`

GetMemSize returns the MemSize field if non-nil, zero value otherwise.

### GetMemSizeOk

`func (o *CloudObjectNode) GetMemSizeOk() (*string, bool)`

GetMemSizeOk returns a tuple with the MemSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemSize

`func (o *CloudObjectNode) SetMemSize(v string)`

SetMemSize sets MemSize field to given value.

### HasMemSize

`func (o *CloudObjectNode) HasMemSize() bool`

HasMemSize returns a boolean if a field has been set.

### GetName

`func (o *CloudObjectNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudObjectNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudObjectNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudObjectNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *CloudObjectNode) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *CloudObjectNode) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *CloudObjectNode) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *CloudObjectNode) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOwner

`func (o *CloudObjectNode) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudObjectNode) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudObjectNode) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudObjectNode) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPatches

`func (o *CloudObjectNode) GetPatches() []CloudObjectPatch`

GetPatches returns the Patches field if non-nil, zero value otherwise.

### GetPatchesOk

`func (o *CloudObjectNode) GetPatchesOk() (*[]CloudObjectPatch, bool)`

GetPatchesOk returns a tuple with the Patches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatches

`func (o *CloudObjectNode) SetPatches(v []CloudObjectPatch)`

SetPatches sets Patches field to given value.

### HasPatches

`func (o *CloudObjectNode) HasPatches() bool`

HasPatches returns a boolean if a field has been set.

### GetPrivateIp

`func (o *CloudObjectNode) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *CloudObjectNode) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *CloudObjectNode) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *CloudObjectNode) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *CloudObjectNode) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *CloudObjectNode) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *CloudObjectNode) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *CloudObjectNode) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRemoteApps

`func (o *CloudObjectNode) GetRemoteApps() []CloudObjectRemoteApp`

GetRemoteApps returns the RemoteApps field if non-nil, zero value otherwise.

### GetRemoteAppsOk

`func (o *CloudObjectNode) GetRemoteAppsOk() (*[]CloudObjectRemoteApp, bool)`

GetRemoteAppsOk returns a tuple with the RemoteApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteApps

`func (o *CloudObjectNode) SetRemoteApps(v []CloudObjectRemoteApp)`

SetRemoteApps sets RemoteApps field to given value.

### HasRemoteApps

`func (o *CloudObjectNode) HasRemoteApps() bool`

HasRemoteApps returns a boolean if a field has been set.

### GetRemotePassword

`func (o *CloudObjectNode) GetRemotePassword() string`

GetRemotePassword returns the RemotePassword field if non-nil, zero value otherwise.

### GetRemotePasswordOk

`func (o *CloudObjectNode) GetRemotePasswordOk() (*string, bool)`

GetRemotePasswordOk returns a tuple with the RemotePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePassword

`func (o *CloudObjectNode) SetRemotePassword(v string)`

SetRemotePassword sets RemotePassword field to given value.

### HasRemotePassword

`func (o *CloudObjectNode) HasRemotePassword() bool`

HasRemotePassword returns a boolean if a field has been set.

### GetRemotePort

`func (o *CloudObjectNode) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *CloudObjectNode) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *CloudObjectNode) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *CloudObjectNode) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetRemoteProtocol

`func (o *CloudObjectNode) GetRemoteProtocol() string`

GetRemoteProtocol returns the RemoteProtocol field if non-nil, zero value otherwise.

### GetRemoteProtocolOk

`func (o *CloudObjectNode) GetRemoteProtocolOk() (*string, bool)`

GetRemoteProtocolOk returns a tuple with the RemoteProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProtocol

`func (o *CloudObjectNode) SetRemoteProtocol(v string)`

SetRemoteProtocol sets RemoteProtocol field to given value.

### HasRemoteProtocol

`func (o *CloudObjectNode) HasRemoteProtocol() bool`

HasRemoteProtocol returns a boolean if a field has been set.

### GetRemoteUsername

`func (o *CloudObjectNode) GetRemoteUsername() string`

GetRemoteUsername returns the RemoteUsername field if non-nil, zero value otherwise.

### GetRemoteUsernameOk

`func (o *CloudObjectNode) GetRemoteUsernameOk() (*string, bool)`

GetRemoteUsernameOk returns a tuple with the RemoteUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUsername

`func (o *CloudObjectNode) SetRemoteUsername(v string)`

SetRemoteUsername sets RemoteUsername field to given value.

### HasRemoteUsername

`func (o *CloudObjectNode) HasRemoteUsername() bool`

HasRemoteUsername returns a boolean if a field has been set.

### GetServices

`func (o *CloudObjectNode) GetServices() []CloudObjectService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CloudObjectNode) GetServicesOk() (*[]CloudObjectService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CloudObjectNode) SetServices(v []CloudObjectService)`

SetServices sets Services field to given value.

### HasServices

`func (o *CloudObjectNode) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSize

`func (o *CloudObjectNode) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CloudObjectNode) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CloudObjectNode) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CloudObjectNode) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTag

`func (o *CloudObjectNode) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CloudObjectNode) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CloudObjectNode) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CloudObjectNode) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *CloudObjectNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudObjectNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudObjectNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudObjectNode) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *CloudObjectNode) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *CloudObjectNode) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *CloudObjectNode) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *CloudObjectNode) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



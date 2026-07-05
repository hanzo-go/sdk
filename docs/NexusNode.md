# NexusNode

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
**Patches** | Pointer to [**[]NexusPatch**](NexusPatch.md) |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**RemoteApps** | Pointer to [**[]NexusRemoteApp**](NexusRemoteApp.md) |  | [optional] 
**RemotePassword** | Pointer to **string** |  | [optional] 
**RemotePort** | Pointer to **int64** |  | [optional] 
**RemoteProtocol** | Pointer to **string** |  | [optional] 
**RemoteUsername** | Pointer to **string** |  | [optional] 
**Services** | Pointer to [**[]NexusService**](NexusService.md) |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 

## Methods

### NewNexusNode

`func NewNexusNode() *NexusNode`

NewNexusNode instantiates a new NexusNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusNodeWithDefaults

`func NewNexusNodeWithDefaults() *NexusNode`

NewNexusNodeWithDefaults instantiates a new NexusNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoQuery

`func (o *NexusNode) GetAutoQuery() bool`

GetAutoQuery returns the AutoQuery field if non-nil, zero value otherwise.

### GetAutoQueryOk

`func (o *NexusNode) GetAutoQueryOk() (*bool, bool)`

GetAutoQueryOk returns a tuple with the AutoQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoQuery

`func (o *NexusNode) SetAutoQuery(v bool)`

SetAutoQuery sets AutoQuery field to given value.

### HasAutoQuery

`func (o *NexusNode) HasAutoQuery() bool`

HasAutoQuery returns a boolean if a field has been set.

### GetCategory

`func (o *NexusNode) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NexusNode) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NexusNode) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NexusNode) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCpuSize

`func (o *NexusNode) GetCpuSize() string`

GetCpuSize returns the CpuSize field if non-nil, zero value otherwise.

### GetCpuSizeOk

`func (o *NexusNode) GetCpuSizeOk() (*string, bool)`

GetCpuSizeOk returns a tuple with the CpuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSize

`func (o *NexusNode) SetCpuSize(v string)`

SetCpuSize sets CpuSize field to given value.

### HasCpuSize

`func (o *NexusNode) HasCpuSize() bool`

HasCpuSize returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NexusNode) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NexusNode) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NexusNode) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NexusNode) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *NexusNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NexusNode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NexusNode) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NexusNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *NexusNode) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NexusNode) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NexusNode) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NexusNode) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableRemoteApp

`func (o *NexusNode) GetEnableRemoteApp() bool`

GetEnableRemoteApp returns the EnableRemoteApp field if non-nil, zero value otherwise.

### GetEnableRemoteAppOk

`func (o *NexusNode) GetEnableRemoteAppOk() (*bool, bool)`

GetEnableRemoteAppOk returns a tuple with the EnableRemoteApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteApp

`func (o *NexusNode) SetEnableRemoteApp(v bool)`

SetEnableRemoteApp sets EnableRemoteApp field to given value.

### HasEnableRemoteApp

`func (o *NexusNode) HasEnableRemoteApp() bool`

HasEnableRemoteApp returns a boolean if a field has been set.

### GetIsPermanent

`func (o *NexusNode) GetIsPermanent() bool`

GetIsPermanent returns the IsPermanent field if non-nil, zero value otherwise.

### GetIsPermanentOk

`func (o *NexusNode) GetIsPermanentOk() (*bool, bool)`

GetIsPermanentOk returns a tuple with the IsPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPermanent

`func (o *NexusNode) SetIsPermanent(v bool)`

SetIsPermanent sets IsPermanent field to given value.

### HasIsPermanent

`func (o *NexusNode) HasIsPermanent() bool`

HasIsPermanent returns a boolean if a field has been set.

### GetLanguage

`func (o *NexusNode) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *NexusNode) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *NexusNode) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *NexusNode) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMachineName

`func (o *NexusNode) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *NexusNode) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *NexusNode) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *NexusNode) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### GetMemSize

`func (o *NexusNode) GetMemSize() string`

GetMemSize returns the MemSize field if non-nil, zero value otherwise.

### GetMemSizeOk

`func (o *NexusNode) GetMemSizeOk() (*string, bool)`

GetMemSizeOk returns a tuple with the MemSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemSize

`func (o *NexusNode) SetMemSize(v string)`

SetMemSize sets MemSize field to given value.

### HasMemSize

`func (o *NexusNode) HasMemSize() bool`

HasMemSize returns a boolean if a field has been set.

### GetName

`func (o *NexusNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NexusNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NexusNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NexusNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *NexusNode) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *NexusNode) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *NexusNode) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *NexusNode) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOwner

`func (o *NexusNode) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NexusNode) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NexusNode) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NexusNode) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPatches

`func (o *NexusNode) GetPatches() []NexusPatch`

GetPatches returns the Patches field if non-nil, zero value otherwise.

### GetPatchesOk

`func (o *NexusNode) GetPatchesOk() (*[]NexusPatch, bool)`

GetPatchesOk returns a tuple with the Patches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatches

`func (o *NexusNode) SetPatches(v []NexusPatch)`

SetPatches sets Patches field to given value.

### HasPatches

`func (o *NexusNode) HasPatches() bool`

HasPatches returns a boolean if a field has been set.

### GetPrivateIp

`func (o *NexusNode) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *NexusNode) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *NexusNode) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *NexusNode) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *NexusNode) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *NexusNode) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *NexusNode) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *NexusNode) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRemoteApps

`func (o *NexusNode) GetRemoteApps() []NexusRemoteApp`

GetRemoteApps returns the RemoteApps field if non-nil, zero value otherwise.

### GetRemoteAppsOk

`func (o *NexusNode) GetRemoteAppsOk() (*[]NexusRemoteApp, bool)`

GetRemoteAppsOk returns a tuple with the RemoteApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteApps

`func (o *NexusNode) SetRemoteApps(v []NexusRemoteApp)`

SetRemoteApps sets RemoteApps field to given value.

### HasRemoteApps

`func (o *NexusNode) HasRemoteApps() bool`

HasRemoteApps returns a boolean if a field has been set.

### GetRemotePassword

`func (o *NexusNode) GetRemotePassword() string`

GetRemotePassword returns the RemotePassword field if non-nil, zero value otherwise.

### GetRemotePasswordOk

`func (o *NexusNode) GetRemotePasswordOk() (*string, bool)`

GetRemotePasswordOk returns a tuple with the RemotePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePassword

`func (o *NexusNode) SetRemotePassword(v string)`

SetRemotePassword sets RemotePassword field to given value.

### HasRemotePassword

`func (o *NexusNode) HasRemotePassword() bool`

HasRemotePassword returns a boolean if a field has been set.

### GetRemotePort

`func (o *NexusNode) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *NexusNode) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *NexusNode) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *NexusNode) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetRemoteProtocol

`func (o *NexusNode) GetRemoteProtocol() string`

GetRemoteProtocol returns the RemoteProtocol field if non-nil, zero value otherwise.

### GetRemoteProtocolOk

`func (o *NexusNode) GetRemoteProtocolOk() (*string, bool)`

GetRemoteProtocolOk returns a tuple with the RemoteProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProtocol

`func (o *NexusNode) SetRemoteProtocol(v string)`

SetRemoteProtocol sets RemoteProtocol field to given value.

### HasRemoteProtocol

`func (o *NexusNode) HasRemoteProtocol() bool`

HasRemoteProtocol returns a boolean if a field has been set.

### GetRemoteUsername

`func (o *NexusNode) GetRemoteUsername() string`

GetRemoteUsername returns the RemoteUsername field if non-nil, zero value otherwise.

### GetRemoteUsernameOk

`func (o *NexusNode) GetRemoteUsernameOk() (*string, bool)`

GetRemoteUsernameOk returns a tuple with the RemoteUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUsername

`func (o *NexusNode) SetRemoteUsername(v string)`

SetRemoteUsername sets RemoteUsername field to given value.

### HasRemoteUsername

`func (o *NexusNode) HasRemoteUsername() bool`

HasRemoteUsername returns a boolean if a field has been set.

### GetServices

`func (o *NexusNode) GetServices() []NexusService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *NexusNode) GetServicesOk() (*[]NexusService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *NexusNode) SetServices(v []NexusService)`

SetServices sets Services field to given value.

### HasServices

`func (o *NexusNode) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSize

`func (o *NexusNode) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NexusNode) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NexusNode) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *NexusNode) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTag

`func (o *NexusNode) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *NexusNode) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *NexusNode) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *NexusNode) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *NexusNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NexusNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NexusNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NexusNode) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *NexusNode) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *NexusNode) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *NexusNode) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *NexusNode) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



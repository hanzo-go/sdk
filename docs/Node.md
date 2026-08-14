# Node

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
**Patches** | Pointer to [**[]Patch**](Patch.md) |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**RemoteApps** | Pointer to [**[]RemoteApp**](RemoteApp.md) |  | [optional] 
**RemotePassword** | Pointer to **string** |  | [optional] 
**RemotePort** | Pointer to **int32** |  | [optional] 
**RemoteProtocol** | Pointer to **string** |  | [optional] 
**RemoteUsername** | Pointer to **string** |  | [optional] 
**Services** | Pointer to [**[]Service**](Service.md) |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 

## Methods

### NewNode

`func NewNode() *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoQuery

`func (o *Node) GetAutoQuery() bool`

GetAutoQuery returns the AutoQuery field if non-nil, zero value otherwise.

### GetAutoQueryOk

`func (o *Node) GetAutoQueryOk() (*bool, bool)`

GetAutoQueryOk returns a tuple with the AutoQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoQuery

`func (o *Node) SetAutoQuery(v bool)`

SetAutoQuery sets AutoQuery field to given value.

### HasAutoQuery

`func (o *Node) HasAutoQuery() bool`

HasAutoQuery returns a boolean if a field has been set.

### GetCategory

`func (o *Node) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Node) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Node) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Node) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCpuSize

`func (o *Node) GetCpuSize() string`

GetCpuSize returns the CpuSize field if non-nil, zero value otherwise.

### GetCpuSizeOk

`func (o *Node) GetCpuSizeOk() (*string, bool)`

GetCpuSizeOk returns a tuple with the CpuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSize

`func (o *Node) SetCpuSize(v string)`

SetCpuSize sets CpuSize field to given value.

### HasCpuSize

`func (o *Node) HasCpuSize() bool`

HasCpuSize returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Node) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Node) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Node) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Node) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *Node) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Node) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Node) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Node) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Node) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Node) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Node) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Node) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableRemoteApp

`func (o *Node) GetEnableRemoteApp() bool`

GetEnableRemoteApp returns the EnableRemoteApp field if non-nil, zero value otherwise.

### GetEnableRemoteAppOk

`func (o *Node) GetEnableRemoteAppOk() (*bool, bool)`

GetEnableRemoteAppOk returns a tuple with the EnableRemoteApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteApp

`func (o *Node) SetEnableRemoteApp(v bool)`

SetEnableRemoteApp sets EnableRemoteApp field to given value.

### HasEnableRemoteApp

`func (o *Node) HasEnableRemoteApp() bool`

HasEnableRemoteApp returns a boolean if a field has been set.

### GetIsPermanent

`func (o *Node) GetIsPermanent() bool`

GetIsPermanent returns the IsPermanent field if non-nil, zero value otherwise.

### GetIsPermanentOk

`func (o *Node) GetIsPermanentOk() (*bool, bool)`

GetIsPermanentOk returns a tuple with the IsPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPermanent

`func (o *Node) SetIsPermanent(v bool)`

SetIsPermanent sets IsPermanent field to given value.

### HasIsPermanent

`func (o *Node) HasIsPermanent() bool`

HasIsPermanent returns a boolean if a field has been set.

### GetLanguage

`func (o *Node) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Node) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Node) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Node) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMachineName

`func (o *Node) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *Node) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *Node) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *Node) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### GetMemSize

`func (o *Node) GetMemSize() string`

GetMemSize returns the MemSize field if non-nil, zero value otherwise.

### GetMemSizeOk

`func (o *Node) GetMemSizeOk() (*string, bool)`

GetMemSizeOk returns a tuple with the MemSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemSize

`func (o *Node) SetMemSize(v string)`

SetMemSize sets MemSize field to given value.

### HasMemSize

`func (o *Node) HasMemSize() bool`

HasMemSize returns a boolean if a field has been set.

### GetName

`func (o *Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Node) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Node) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *Node) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Node) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Node) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *Node) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOwner

`func (o *Node) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Node) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Node) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Node) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPatches

`func (o *Node) GetPatches() []Patch`

GetPatches returns the Patches field if non-nil, zero value otherwise.

### GetPatchesOk

`func (o *Node) GetPatchesOk() (*[]Patch, bool)`

GetPatchesOk returns a tuple with the Patches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatches

`func (o *Node) SetPatches(v []Patch)`

SetPatches sets Patches field to given value.

### HasPatches

`func (o *Node) HasPatches() bool`

HasPatches returns a boolean if a field has been set.

### GetPrivateIp

`func (o *Node) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *Node) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *Node) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *Node) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicIp

`func (o *Node) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *Node) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *Node) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *Node) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRemoteApps

`func (o *Node) GetRemoteApps() []RemoteApp`

GetRemoteApps returns the RemoteApps field if non-nil, zero value otherwise.

### GetRemoteAppsOk

`func (o *Node) GetRemoteAppsOk() (*[]RemoteApp, bool)`

GetRemoteAppsOk returns a tuple with the RemoteApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteApps

`func (o *Node) SetRemoteApps(v []RemoteApp)`

SetRemoteApps sets RemoteApps field to given value.

### HasRemoteApps

`func (o *Node) HasRemoteApps() bool`

HasRemoteApps returns a boolean if a field has been set.

### GetRemotePassword

`func (o *Node) GetRemotePassword() string`

GetRemotePassword returns the RemotePassword field if non-nil, zero value otherwise.

### GetRemotePasswordOk

`func (o *Node) GetRemotePasswordOk() (*string, bool)`

GetRemotePasswordOk returns a tuple with the RemotePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePassword

`func (o *Node) SetRemotePassword(v string)`

SetRemotePassword sets RemotePassword field to given value.

### HasRemotePassword

`func (o *Node) HasRemotePassword() bool`

HasRemotePassword returns a boolean if a field has been set.

### GetRemotePort

`func (o *Node) GetRemotePort() int32`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *Node) GetRemotePortOk() (*int32, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *Node) SetRemotePort(v int32)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *Node) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetRemoteProtocol

`func (o *Node) GetRemoteProtocol() string`

GetRemoteProtocol returns the RemoteProtocol field if non-nil, zero value otherwise.

### GetRemoteProtocolOk

`func (o *Node) GetRemoteProtocolOk() (*string, bool)`

GetRemoteProtocolOk returns a tuple with the RemoteProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProtocol

`func (o *Node) SetRemoteProtocol(v string)`

SetRemoteProtocol sets RemoteProtocol field to given value.

### HasRemoteProtocol

`func (o *Node) HasRemoteProtocol() bool`

HasRemoteProtocol returns a boolean if a field has been set.

### GetRemoteUsername

`func (o *Node) GetRemoteUsername() string`

GetRemoteUsername returns the RemoteUsername field if non-nil, zero value otherwise.

### GetRemoteUsernameOk

`func (o *Node) GetRemoteUsernameOk() (*string, bool)`

GetRemoteUsernameOk returns a tuple with the RemoteUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUsername

`func (o *Node) SetRemoteUsername(v string)`

SetRemoteUsername sets RemoteUsername field to given value.

### HasRemoteUsername

`func (o *Node) HasRemoteUsername() bool`

HasRemoteUsername returns a boolean if a field has been set.

### GetServices

`func (o *Node) GetServices() []Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Node) GetServicesOk() (*[]Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Node) SetServices(v []Service)`

SetServices sets Services field to given value.

### HasServices

`func (o *Node) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSize

`func (o *Node) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Node) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Node) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Node) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTag

`func (o *Node) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Node) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Node) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Node) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *Node) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Node) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Node) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Node) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *Node) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *Node) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *Node) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *Node) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



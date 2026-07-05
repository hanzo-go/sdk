# CloudObjectMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**CpuSize** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ExpireTime** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**MemSize** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RemotePassword** | Pointer to **string** |  | [optional] 
**RemotePort** | Pointer to **int64** |  | [optional] 
**RemoteProtocol** | Pointer to **string** |  | [optional] 
**RemoteUsername** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudObjectMachine

`func NewCloudObjectMachine() *CloudObjectMachine`

NewCloudObjectMachine instantiates a new CloudObjectMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectMachineWithDefaults

`func NewCloudObjectMachineWithDefaults() *CloudObjectMachine`

NewCloudObjectMachineWithDefaults instantiates a new CloudObjectMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CloudObjectMachine) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudObjectMachine) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudObjectMachine) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudObjectMachine) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCpuSize

`func (o *CloudObjectMachine) GetCpuSize() string`

GetCpuSize returns the CpuSize field if non-nil, zero value otherwise.

### GetCpuSizeOk

`func (o *CloudObjectMachine) GetCpuSizeOk() (*string, bool)`

GetCpuSizeOk returns a tuple with the CpuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSize

`func (o *CloudObjectMachine) SetCpuSize(v string)`

SetCpuSize sets CpuSize field to given value.

### HasCpuSize

`func (o *CloudObjectMachine) HasCpuSize() bool`

HasCpuSize returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CloudObjectMachine) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudObjectMachine) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudObjectMachine) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudObjectMachine) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudObjectMachine) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudObjectMachine) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudObjectMachine) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudObjectMachine) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExpireTime

`func (o *CloudObjectMachine) GetExpireTime() string`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *CloudObjectMachine) GetExpireTimeOk() (*string, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *CloudObjectMachine) SetExpireTime(v string)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *CloudObjectMachine) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetId

`func (o *CloudObjectMachine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudObjectMachine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudObjectMachine) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudObjectMachine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *CloudObjectMachine) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CloudObjectMachine) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CloudObjectMachine) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CloudObjectMachine) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMemSize

`func (o *CloudObjectMachine) GetMemSize() string`

GetMemSize returns the MemSize field if non-nil, zero value otherwise.

### GetMemSizeOk

`func (o *CloudObjectMachine) GetMemSizeOk() (*string, bool)`

GetMemSizeOk returns a tuple with the MemSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemSize

`func (o *CloudObjectMachine) SetMemSize(v string)`

SetMemSize sets MemSize field to given value.

### HasMemSize

`func (o *CloudObjectMachine) HasMemSize() bool`

HasMemSize returns a boolean if a field has been set.

### GetName

`func (o *CloudObjectMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudObjectMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudObjectMachine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudObjectMachine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *CloudObjectMachine) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *CloudObjectMachine) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *CloudObjectMachine) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *CloudObjectMachine) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOwner

`func (o *CloudObjectMachine) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudObjectMachine) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudObjectMachine) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudObjectMachine) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrivateIp

`func (o *CloudObjectMachine) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *CloudObjectMachine) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *CloudObjectMachine) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *CloudObjectMachine) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetProvider

`func (o *CloudObjectMachine) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudObjectMachine) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudObjectMachine) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudObjectMachine) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublicIp

`func (o *CloudObjectMachine) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *CloudObjectMachine) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *CloudObjectMachine) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *CloudObjectMachine) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRegion

`func (o *CloudObjectMachine) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudObjectMachine) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudObjectMachine) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudObjectMachine) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRemotePassword

`func (o *CloudObjectMachine) GetRemotePassword() string`

GetRemotePassword returns the RemotePassword field if non-nil, zero value otherwise.

### GetRemotePasswordOk

`func (o *CloudObjectMachine) GetRemotePasswordOk() (*string, bool)`

GetRemotePasswordOk returns a tuple with the RemotePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePassword

`func (o *CloudObjectMachine) SetRemotePassword(v string)`

SetRemotePassword sets RemotePassword field to given value.

### HasRemotePassword

`func (o *CloudObjectMachine) HasRemotePassword() bool`

HasRemotePassword returns a boolean if a field has been set.

### GetRemotePort

`func (o *CloudObjectMachine) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *CloudObjectMachine) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *CloudObjectMachine) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *CloudObjectMachine) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetRemoteProtocol

`func (o *CloudObjectMachine) GetRemoteProtocol() string`

GetRemoteProtocol returns the RemoteProtocol field if non-nil, zero value otherwise.

### GetRemoteProtocolOk

`func (o *CloudObjectMachine) GetRemoteProtocolOk() (*string, bool)`

GetRemoteProtocolOk returns a tuple with the RemoteProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProtocol

`func (o *CloudObjectMachine) SetRemoteProtocol(v string)`

SetRemoteProtocol sets RemoteProtocol field to given value.

### HasRemoteProtocol

`func (o *CloudObjectMachine) HasRemoteProtocol() bool`

HasRemoteProtocol returns a boolean if a field has been set.

### GetRemoteUsername

`func (o *CloudObjectMachine) GetRemoteUsername() string`

GetRemoteUsername returns the RemoteUsername field if non-nil, zero value otherwise.

### GetRemoteUsernameOk

`func (o *CloudObjectMachine) GetRemoteUsernameOk() (*string, bool)`

GetRemoteUsernameOk returns a tuple with the RemoteUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUsername

`func (o *CloudObjectMachine) SetRemoteUsername(v string)`

SetRemoteUsername sets RemoteUsername field to given value.

### HasRemoteUsername

`func (o *CloudObjectMachine) HasRemoteUsername() bool`

HasRemoteUsername returns a boolean if a field has been set.

### GetSize

`func (o *CloudObjectMachine) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CloudObjectMachine) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CloudObjectMachine) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CloudObjectMachine) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *CloudObjectMachine) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudObjectMachine) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudObjectMachine) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudObjectMachine) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *CloudObjectMachine) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CloudObjectMachine) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CloudObjectMachine) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CloudObjectMachine) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *CloudObjectMachine) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudObjectMachine) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudObjectMachine) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudObjectMachine) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *CloudObjectMachine) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *CloudObjectMachine) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *CloudObjectMachine) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *CloudObjectMachine) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetZone

`func (o *CloudObjectMachine) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CloudObjectMachine) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CloudObjectMachine) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CloudObjectMachine) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudObjectContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**SizeRootFs** | Pointer to **int64** |  | [optional] 
**SizeRw** | Pointer to **int64** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudObjectContainer

`func NewCloudObjectContainer() *CloudObjectContainer`

NewCloudObjectContainer instantiates a new CloudObjectContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectContainerWithDefaults

`func NewCloudObjectContainerWithDefaults() *CloudObjectContainer`

NewCloudObjectContainerWithDefaults instantiates a new CloudObjectContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *CloudObjectContainer) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CloudObjectContainer) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CloudObjectContainer) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CloudObjectContainer) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CloudObjectContainer) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudObjectContainer) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudObjectContainer) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudObjectContainer) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudObjectContainer) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudObjectContainer) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudObjectContainer) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudObjectContainer) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetImage

`func (o *CloudObjectContainer) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CloudObjectContainer) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CloudObjectContainer) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CloudObjectContainer) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageId

`func (o *CloudObjectContainer) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CloudObjectContainer) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CloudObjectContainer) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CloudObjectContainer) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetName

`func (o *CloudObjectContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudObjectContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudObjectContainer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudObjectContainer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *CloudObjectContainer) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudObjectContainer) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudObjectContainer) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudObjectContainer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPorts

`func (o *CloudObjectContainer) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *CloudObjectContainer) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *CloudObjectContainer) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *CloudObjectContainer) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProvider

`func (o *CloudObjectContainer) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudObjectContainer) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudObjectContainer) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudObjectContainer) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSizeRootFs

`func (o *CloudObjectContainer) GetSizeRootFs() int64`

GetSizeRootFs returns the SizeRootFs field if non-nil, zero value otherwise.

### GetSizeRootFsOk

`func (o *CloudObjectContainer) GetSizeRootFsOk() (*int64, bool)`

GetSizeRootFsOk returns a tuple with the SizeRootFs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeRootFs

`func (o *CloudObjectContainer) SetSizeRootFs(v int64)`

SetSizeRootFs sets SizeRootFs field to given value.

### HasSizeRootFs

`func (o *CloudObjectContainer) HasSizeRootFs() bool`

HasSizeRootFs returns a boolean if a field has been set.

### GetSizeRw

`func (o *CloudObjectContainer) GetSizeRw() int64`

GetSizeRw returns the SizeRw field if non-nil, zero value otherwise.

### GetSizeRwOk

`func (o *CloudObjectContainer) GetSizeRwOk() (*int64, bool)`

GetSizeRwOk returns a tuple with the SizeRw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeRw

`func (o *CloudObjectContainer) SetSizeRw(v int64)`

SetSizeRw sets SizeRw field to given value.

### HasSizeRw

`func (o *CloudObjectContainer) HasSizeRw() bool`

HasSizeRw returns a boolean if a field has been set.

### GetState

`func (o *CloudObjectContainer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudObjectContainer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudObjectContainer) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudObjectContainer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *CloudObjectContainer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudObjectContainer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudObjectContainer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudObjectContainer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ArgoCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionState** | Pointer to [**ArgoConnectionState**](ArgoConnectionState.md) |  | [optional] 
**Info** | Pointer to [**ArgoClusterInfo**](ArgoClusterInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 

## Methods

### NewArgoCluster

`func NewArgoCluster() *ArgoCluster`

NewArgoCluster instantiates a new ArgoCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoClusterWithDefaults

`func NewArgoClusterWithDefaults() *ArgoCluster`

NewArgoClusterWithDefaults instantiates a new ArgoCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionState

`func (o *ArgoCluster) GetConnectionState() ArgoConnectionState`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *ArgoCluster) GetConnectionStateOk() (*ArgoConnectionState, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *ArgoCluster) SetConnectionState(v ArgoConnectionState)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *ArgoCluster) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetInfo

`func (o *ArgoCluster) GetInfo() ArgoClusterInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ArgoCluster) GetInfoOk() (*ArgoClusterInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ArgoCluster) SetInfo(v ArgoClusterInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ArgoCluster) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetName

`func (o *ArgoCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArgoCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArgoCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArgoCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServer

`func (o *ArgoCluster) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ArgoCluster) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ArgoCluster) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *ArgoCluster) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



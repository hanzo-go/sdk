# CloudArgoCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionState** | Pointer to [**CloudArgoConnectionState**](CloudArgoConnectionState.md) |  | [optional] 
**Info** | Pointer to [**CloudArgoClusterInfo**](CloudArgoClusterInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudArgoCluster

`func NewCloudArgoCluster() *CloudArgoCluster`

NewCloudArgoCluster instantiates a new CloudArgoCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoClusterWithDefaults

`func NewCloudArgoClusterWithDefaults() *CloudArgoCluster`

NewCloudArgoClusterWithDefaults instantiates a new CloudArgoCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionState

`func (o *CloudArgoCluster) GetConnectionState() CloudArgoConnectionState`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *CloudArgoCluster) GetConnectionStateOk() (*CloudArgoConnectionState, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *CloudArgoCluster) SetConnectionState(v CloudArgoConnectionState)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *CloudArgoCluster) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetInfo

`func (o *CloudArgoCluster) GetInfo() CloudArgoClusterInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CloudArgoCluster) GetInfoOk() (*CloudArgoClusterInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CloudArgoCluster) SetInfo(v CloudArgoClusterInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CloudArgoCluster) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetName

`func (o *CloudArgoCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudArgoCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudArgoCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudArgoCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServer

`func (o *CloudArgoCluster) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CloudArgoCluster) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CloudArgoCluster) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *CloudArgoCluster) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



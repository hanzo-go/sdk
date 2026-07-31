# CloudArgoClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationsCount** | Pointer to **int32** |  | [optional] 
**ConnectionState** | Pointer to [**CloudArgoConnectionState**](CloudArgoConnectionState.md) |  | [optional] 
**ServerVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudArgoClusterInfo

`func NewCloudArgoClusterInfo() *CloudArgoClusterInfo`

NewCloudArgoClusterInfo instantiates a new CloudArgoClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoClusterInfoWithDefaults

`func NewCloudArgoClusterInfoWithDefaults() *CloudArgoClusterInfo`

NewCloudArgoClusterInfoWithDefaults instantiates a new CloudArgoClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationsCount

`func (o *CloudArgoClusterInfo) GetApplicationsCount() int32`

GetApplicationsCount returns the ApplicationsCount field if non-nil, zero value otherwise.

### GetApplicationsCountOk

`func (o *CloudArgoClusterInfo) GetApplicationsCountOk() (*int32, bool)`

GetApplicationsCountOk returns a tuple with the ApplicationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsCount

`func (o *CloudArgoClusterInfo) SetApplicationsCount(v int32)`

SetApplicationsCount sets ApplicationsCount field to given value.

### HasApplicationsCount

`func (o *CloudArgoClusterInfo) HasApplicationsCount() bool`

HasApplicationsCount returns a boolean if a field has been set.

### GetConnectionState

`func (o *CloudArgoClusterInfo) GetConnectionState() CloudArgoConnectionState`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *CloudArgoClusterInfo) GetConnectionStateOk() (*CloudArgoConnectionState, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *CloudArgoClusterInfo) SetConnectionState(v CloudArgoConnectionState)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *CloudArgoClusterInfo) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetServerVersion

`func (o *CloudArgoClusterInfo) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *CloudArgoClusterInfo) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *CloudArgoClusterInfo) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *CloudArgoClusterInfo) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



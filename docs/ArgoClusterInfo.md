# ArgoClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationsCount** | Pointer to **int64** | ApplicationsCount is how many of THE CALLER&#39;S applications reconcile into this cluster, so a tenant sees its own count and a SuperAdmin the fleet&#39;s. It is zero for the in-cluster destination when the caller owns nothing, since that destination is listed whether or not anything targets it. | [optional] 
**ConnectionState** | Pointer to [**ArgoConnectionState**](ArgoConnectionState.md) | ConnectionState repeats the cluster&#39;s own connection state, which is where ArgoCD&#39;s UI reads it from on this object. | [optional] 
**ServerVersion** | Pointer to **string** | ServerVersion is the kubernetes version of the destination. Always absent: nothing here queries the API server for it. | [optional] 

## Methods

### NewArgoClusterInfo

`func NewArgoClusterInfo() *ArgoClusterInfo`

NewArgoClusterInfo instantiates a new ArgoClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoClusterInfoWithDefaults

`func NewArgoClusterInfoWithDefaults() *ArgoClusterInfo`

NewArgoClusterInfoWithDefaults instantiates a new ArgoClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationsCount

`func (o *ArgoClusterInfo) GetApplicationsCount() int64`

GetApplicationsCount returns the ApplicationsCount field if non-nil, zero value otherwise.

### GetApplicationsCountOk

`func (o *ArgoClusterInfo) GetApplicationsCountOk() (*int64, bool)`

GetApplicationsCountOk returns a tuple with the ApplicationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationsCount

`func (o *ArgoClusterInfo) SetApplicationsCount(v int64)`

SetApplicationsCount sets ApplicationsCount field to given value.

### HasApplicationsCount

`func (o *ArgoClusterInfo) HasApplicationsCount() bool`

HasApplicationsCount returns a boolean if a field has been set.

### GetConnectionState

`func (o *ArgoClusterInfo) GetConnectionState() ArgoConnectionState`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *ArgoClusterInfo) GetConnectionStateOk() (*ArgoConnectionState, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *ArgoClusterInfo) SetConnectionState(v ArgoConnectionState)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *ArgoClusterInfo) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetServerVersion

`func (o *ArgoClusterInfo) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *ArgoClusterInfo) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *ArgoClusterInfo) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *ArgoClusterInfo) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



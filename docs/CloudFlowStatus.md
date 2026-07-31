# CloudFlowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reachable** | Pointer to **bool** | Reachable is true when the flow service answered its health probe. | [optional] 
**Version** | Pointer to **string** | Version is the flow service&#39;s own version, present only when reachable. | [optional] 

## Methods

### NewCloudFlowStatus

`func NewCloudFlowStatus() *CloudFlowStatus`

NewCloudFlowStatus instantiates a new CloudFlowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFlowStatusWithDefaults

`func NewCloudFlowStatusWithDefaults() *CloudFlowStatus`

NewCloudFlowStatusWithDefaults instantiates a new CloudFlowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReachable

`func (o *CloudFlowStatus) GetReachable() bool`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *CloudFlowStatus) GetReachableOk() (*bool, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *CloudFlowStatus) SetReachable(v bool)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *CloudFlowStatus) HasReachable() bool`

HasReachable returns a boolean if a field has been set.

### GetVersion

`func (o *CloudFlowStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudFlowStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudFlowStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudFlowStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



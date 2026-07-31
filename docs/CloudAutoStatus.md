# CloudAutoStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reachable** | Pointer to **bool** | Reachable is true when the auto service answered its health probe. | [optional] 

## Methods

### NewCloudAutoStatus

`func NewCloudAutoStatus() *CloudAutoStatus`

NewCloudAutoStatus instantiates a new CloudAutoStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAutoStatusWithDefaults

`func NewCloudAutoStatusWithDefaults() *CloudAutoStatus`

NewCloudAutoStatusWithDefaults instantiates a new CloudAutoStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReachable

`func (o *CloudAutoStatus) GetReachable() bool`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *CloudAutoStatus) GetReachableOk() (*bool, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *CloudAutoStatus) SetReachable(v bool)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *CloudAutoStatus) HasReachable() bool`

HasReachable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



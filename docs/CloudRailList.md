# CloudRailList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rails** | Pointer to [**[]CloudRailView**](CloudRailView.md) | Rails is every (chain, token, treasury) triple this deployment accepts. | [optional] 

## Methods

### NewCloudRailList

`func NewCloudRailList() *CloudRailList`

NewCloudRailList instantiates a new CloudRailList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRailListWithDefaults

`func NewCloudRailListWithDefaults() *CloudRailList`

NewCloudRailListWithDefaults instantiates a new CloudRailList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRails

`func (o *CloudRailList) GetRails() []CloudRailView`

GetRails returns the Rails field if non-nil, zero value otherwise.

### GetRailsOk

`func (o *CloudRailList) GetRailsOk() (*[]CloudRailView, bool)`

GetRailsOk returns a tuple with the Rails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRails

`func (o *CloudRailList) SetRails(v []CloudRailView)`

SetRails sets Rails field to given value.

### HasRails

`func (o *CloudRailList) HasRails() bool`

HasRails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



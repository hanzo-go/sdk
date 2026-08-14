# ValidatorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SlotView**](SlotView.md) | Data is one entry per slot this org has claimed. | [optional] 
**Network** | Pointer to **string** | Network is the luxd network slug new nodes join on this deployment. | [optional] 

## Methods

### NewValidatorList

`func NewValidatorList() *ValidatorList`

NewValidatorList instantiates a new ValidatorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorListWithDefaults

`func NewValidatorListWithDefaults() *ValidatorList`

NewValidatorListWithDefaults instantiates a new ValidatorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ValidatorList) GetData() []SlotView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ValidatorList) GetDataOk() (*[]SlotView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ValidatorList) SetData(v []SlotView)`

SetData sets Data field to given value.

### HasData

`func (o *ValidatorList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNetwork

`func (o *ValidatorList) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ValidatorList) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ValidatorList) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ValidatorList) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



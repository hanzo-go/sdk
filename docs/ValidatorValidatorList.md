# ValidatorValidatorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ValidatorSlotView**](ValidatorSlotView.md) | Data is one entry per slot this org has claimed. | [optional] 
**Network** | Pointer to **string** | Network is the luxd network slug new nodes join on this deployment. | [optional] 

## Methods

### NewValidatorValidatorList

`func NewValidatorValidatorList() *ValidatorValidatorList`

NewValidatorValidatorList instantiates a new ValidatorValidatorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorValidatorListWithDefaults

`func NewValidatorValidatorListWithDefaults() *ValidatorValidatorList`

NewValidatorValidatorListWithDefaults instantiates a new ValidatorValidatorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ValidatorValidatorList) GetData() []ValidatorSlotView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ValidatorValidatorList) GetDataOk() (*[]ValidatorSlotView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ValidatorValidatorList) SetData(v []ValidatorSlotView)`

SetData sets Data field to given value.

### HasData

`func (o *ValidatorValidatorList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNetwork

`func (o *ValidatorValidatorList) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ValidatorValidatorList) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ValidatorValidatorList) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ValidatorValidatorList) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



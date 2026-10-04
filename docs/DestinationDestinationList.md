# DestinationDestinationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | Pointer to [**[]DestinationDestinationStatus**](DestinationDestinationStatus.md) | Destinations is one card per registered platform, in slug order. | [optional] 

## Methods

### NewDestinationDestinationList

`func NewDestinationDestinationList() *DestinationDestinationList`

NewDestinationDestinationList instantiates a new DestinationDestinationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationDestinationListWithDefaults

`func NewDestinationDestinationListWithDefaults() *DestinationDestinationList`

NewDestinationDestinationListWithDefaults instantiates a new DestinationDestinationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *DestinationDestinationList) GetDestinations() []DestinationDestinationStatus`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *DestinationDestinationList) GetDestinationsOk() (*[]DestinationDestinationStatus, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *DestinationDestinationList) SetDestinations(v []DestinationDestinationStatus)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *DestinationDestinationList) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudDestinationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | Pointer to [**[]CloudDestinationStatus**](CloudDestinationStatus.md) | Destinations is one card per registered platform, in slug order. | [optional] 

## Methods

### NewCloudDestinationList

`func NewCloudDestinationList() *CloudDestinationList`

NewCloudDestinationList instantiates a new CloudDestinationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDestinationListWithDefaults

`func NewCloudDestinationListWithDefaults() *CloudDestinationList`

NewCloudDestinationListWithDefaults instantiates a new CloudDestinationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *CloudDestinationList) GetDestinations() []CloudDestinationStatus`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *CloudDestinationList) GetDestinationsOk() (*[]CloudDestinationStatus, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *CloudDestinationList) SetDestinations(v []CloudDestinationStatus)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *CloudDestinationList) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



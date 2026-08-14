# AccList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AccView**](AccView.md) | Data is the org&#39;s tracked accreditation records, newest first. | [optional] 
**Disclaimer** | Pointer to **string** | Disclaimer states that statuses are tracked or provider-reported, never a platform assertion of legal or regulatory compliance. | [optional] 

## Methods

### NewAccList

`func NewAccList() *AccList`

NewAccList instantiates a new AccList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccListWithDefaults

`func NewAccListWithDefaults() *AccList`

NewAccListWithDefaults instantiates a new AccList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AccList) GetData() []AccView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccList) GetDataOk() (*[]AccView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccList) SetData(v []AccView)`

SetData sets Data field to given value.

### HasData

`func (o *AccList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisclaimer

`func (o *AccList) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *AccList) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *AccList) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *AccList) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



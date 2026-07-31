# CloudAccList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudAccView**](CloudAccView.md) | Data is the org&#39;s tracked accreditation records, newest first. | [optional] 
**Disclaimer** | Pointer to **string** | Disclaimer states that statuses are tracked or provider-reported, never a platform assertion of legal or regulatory compliance. | [optional] 

## Methods

### NewCloudAccList

`func NewCloudAccList() *CloudAccList`

NewCloudAccList instantiates a new CloudAccList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccListWithDefaults

`func NewCloudAccListWithDefaults() *CloudAccList`

NewCloudAccListWithDefaults instantiates a new CloudAccList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudAccList) GetData() []CloudAccView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudAccList) GetDataOk() (*[]CloudAccView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudAccList) SetData(v []CloudAccView)`

SetData sets Data field to given value.

### HasData

`func (o *CloudAccList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisclaimer

`func (o *CloudAccList) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *CloudAccList) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *CloudAccList) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *CloudAccList) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



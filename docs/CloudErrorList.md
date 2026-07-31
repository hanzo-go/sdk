# CloudErrorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudCapturedError**](CloudCapturedError.md) | Data is the errors, newest first. Empty rather than absent when there are none. | [optional] 

## Methods

### NewCloudErrorList

`func NewCloudErrorList() *CloudErrorList`

NewCloudErrorList instantiates a new CloudErrorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudErrorListWithDefaults

`func NewCloudErrorListWithDefaults() *CloudErrorList`

NewCloudErrorListWithDefaults instantiates a new CloudErrorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudErrorList) GetData() []CloudCapturedError`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudErrorList) GetDataOk() (*[]CloudCapturedError, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudErrorList) SetData(v []CloudCapturedError)`

SetData sets Data field to given value.

### HasData

`func (o *CloudErrorList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



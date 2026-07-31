# CloudApplicationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudApplication**](CloudApplication.md) | Data is the page of applications, newest first. | [optional] 

## Methods

### NewCloudApplicationList

`func NewCloudApplicationList() *CloudApplicationList`

NewCloudApplicationList instantiates a new CloudApplicationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudApplicationListWithDefaults

`func NewCloudApplicationListWithDefaults() *CloudApplicationList`

NewCloudApplicationListWithDefaults instantiates a new CloudApplicationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudApplicationList) GetData() []CloudApplication`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudApplicationList) GetDataOk() (*[]CloudApplication, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudApplicationList) SetData(v []CloudApplication)`

SetData sets Data field to given value.

### HasData

`func (o *CloudApplicationList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudKeyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudKeyView**](CloudKeyView.md) | Data holds the org&#39;s keys. | [optional] 

## Methods

### NewCloudKeyList

`func NewCloudKeyList() *CloudKeyList`

NewCloudKeyList instantiates a new CloudKeyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudKeyListWithDefaults

`func NewCloudKeyListWithDefaults() *CloudKeyList`

NewCloudKeyListWithDefaults instantiates a new CloudKeyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudKeyList) GetData() []CloudKeyView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudKeyList) GetDataOk() (*[]CloudKeyView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudKeyList) SetData(v []CloudKeyView)`

SetData sets Data field to given value.

### HasData

`func (o *CloudKeyList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



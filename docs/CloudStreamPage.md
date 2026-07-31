# CloudStreamPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudStreamRecord**](CloudStreamRecord.md) | Data are the org&#39;s streams. | [optional] 

## Methods

### NewCloudStreamPage

`func NewCloudStreamPage() *CloudStreamPage`

NewCloudStreamPage instantiates a new CloudStreamPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStreamPageWithDefaults

`func NewCloudStreamPageWithDefaults() *CloudStreamPage`

NewCloudStreamPageWithDefaults instantiates a new CloudStreamPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudStreamPage) GetData() []CloudStreamRecord`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudStreamPage) GetDataOk() (*[]CloudStreamRecord, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudStreamPage) SetData(v []CloudStreamRecord)`

SetData sets Data field to given value.

### HasData

`func (o *CloudStreamPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



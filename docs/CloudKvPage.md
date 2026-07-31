# CloudKvPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudKvEntry**](CloudKvEntry.md) | Data are the key&#39;s retained revisions. | [optional] 

## Methods

### NewCloudKvPage

`func NewCloudKvPage() *CloudKvPage`

NewCloudKvPage instantiates a new CloudKvPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudKvPageWithDefaults

`func NewCloudKvPageWithDefaults() *CloudKvPage`

NewCloudKvPageWithDefaults instantiates a new CloudKvPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudKvPage) GetData() []CloudKvEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudKvPage) GetDataOk() (*[]CloudKvEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudKvPage) SetData(v []CloudKvEntry)`

SetData sets Data field to given value.

### HasData

`func (o *CloudKvPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



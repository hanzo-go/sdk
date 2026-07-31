# CloudOppList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudOpportunity**](CloudOpportunity.md) | Data is the page of opportunities, most recently updated first. | [optional] 

## Methods

### NewCloudOppList

`func NewCloudOppList() *CloudOppList`

NewCloudOppList instantiates a new CloudOppList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudOppListWithDefaults

`func NewCloudOppListWithDefaults() *CloudOppList`

NewCloudOppListWithDefaults instantiates a new CloudOppList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudOppList) GetData() []CloudOpportunity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudOppList) GetDataOk() (*[]CloudOpportunity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudOppList) SetData(v []CloudOpportunity)`

SetData sets Data field to given value.

### HasData

`func (o *CloudOppList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudFlowPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudFlow**](CloudFlow.md) | Data is the page of flows, newest-updated first. | [optional] 

## Methods

### NewCloudFlowPage

`func NewCloudFlowPage() *CloudFlowPage`

NewCloudFlowPage instantiates a new CloudFlowPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFlowPageWithDefaults

`func NewCloudFlowPageWithDefaults() *CloudFlowPage`

NewCloudFlowPageWithDefaults instantiates a new CloudFlowPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudFlowPage) GetData() []CloudFlow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudFlowPage) GetDataOk() (*[]CloudFlow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudFlowPage) SetData(v []CloudFlow)`

SetData sets Data field to given value.

### HasData

`func (o *CloudFlowPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



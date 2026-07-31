# CloudRunPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudFlowRun**](CloudFlowRun.md) | Data is the page of runs, newest first. | [optional] 

## Methods

### NewCloudRunPage

`func NewCloudRunPage() *CloudRunPage`

NewCloudRunPage instantiates a new CloudRunPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRunPageWithDefaults

`func NewCloudRunPageWithDefaults() *CloudRunPage`

NewCloudRunPageWithDefaults instantiates a new CloudRunPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudRunPage) GetData() []CloudFlowRun`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudRunPage) GetDataOk() (*[]CloudFlowRun, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudRunPage) SetData(v []CloudFlowRun)`

SetData sets Data field to given value.

### HasData

`func (o *CloudRunPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



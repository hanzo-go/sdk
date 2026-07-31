# CloudEndpointList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudEndpoint**](CloudEndpoint.md) | Data is the org&#39;s endpoints, newest first, each with its signing secret REDACTED — the secret leaves the server only on create and on rotate. | [optional] 

## Methods

### NewCloudEndpointList

`func NewCloudEndpointList() *CloudEndpointList`

NewCloudEndpointList instantiates a new CloudEndpointList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEndpointListWithDefaults

`func NewCloudEndpointListWithDefaults() *CloudEndpointList`

NewCloudEndpointListWithDefaults instantiates a new CloudEndpointList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudEndpointList) GetData() []CloudEndpoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudEndpointList) GetDataOk() (*[]CloudEndpoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudEndpointList) SetData(v []CloudEndpoint)`

SetData sets Data field to given value.

### HasData

`func (o *CloudEndpointList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



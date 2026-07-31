# CloudCollabResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **map[string]string** | Content maps each document field to its value for the verb: the new blob ref after a createContent, the stored markup after a getContent. | [optional] 
**Error** | Pointer to **string** | Error carries a SEMANTIC refusal, which this RPC reports under 200 because the client throws on result.error — auth and tenancy failures are HTTP statuses instead. | [optional] 

## Methods

### NewCloudCollabResult

`func NewCloudCollabResult() *CloudCollabResult`

NewCloudCollabResult instantiates a new CloudCollabResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCollabResultWithDefaults

`func NewCloudCollabResultWithDefaults() *CloudCollabResult`

NewCloudCollabResultWithDefaults instantiates a new CloudCollabResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CloudCollabResult) GetContent() map[string]string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CloudCollabResult) GetContentOk() (*map[string]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CloudCollabResult) SetContent(v map[string]string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CloudCollabResult) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetError

`func (o *CloudCollabResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudCollabResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudCollabResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CloudCollabResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



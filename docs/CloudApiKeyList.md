# CloudApiKeyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to [**[]CloudApiKey**](CloudApiKey.md) | Keys is every key the caller holds, at most one per type. | [optional] 

## Methods

### NewCloudApiKeyList

`func NewCloudApiKeyList() *CloudApiKeyList`

NewCloudApiKeyList instantiates a new CloudApiKeyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudApiKeyListWithDefaults

`func NewCloudApiKeyListWithDefaults() *CloudApiKeyList`

NewCloudApiKeyListWithDefaults instantiates a new CloudApiKeyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *CloudApiKeyList) GetKeys() []CloudApiKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *CloudApiKeyList) GetKeysOk() (*[]CloudApiKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *CloudApiKeyList) SetKeys(v []CloudApiKey)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *CloudApiKeyList) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



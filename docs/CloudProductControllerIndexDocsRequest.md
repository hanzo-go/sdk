# CloudProductControllerIndexDocsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** |  | [optional] 
**Documents** | **[]map[string]interface{}** |  | 
**Replace** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudProductControllerIndexDocsRequest

`func NewCloudProductControllerIndexDocsRequest(documents []map[string]interface{}, ) *CloudProductControllerIndexDocsRequest`

NewCloudProductControllerIndexDocsRequest instantiates a new CloudProductControllerIndexDocsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProductControllerIndexDocsRequestWithDefaults

`func NewCloudProductControllerIndexDocsRequestWithDefaults() *CloudProductControllerIndexDocsRequest`

NewCloudProductControllerIndexDocsRequestWithDefaults instantiates a new CloudProductControllerIndexDocsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *CloudProductControllerIndexDocsRequest) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CloudProductControllerIndexDocsRequest) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CloudProductControllerIndexDocsRequest) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *CloudProductControllerIndexDocsRequest) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetDocuments

`func (o *CloudProductControllerIndexDocsRequest) GetDocuments() []map[string]interface{}`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *CloudProductControllerIndexDocsRequest) GetDocumentsOk() (*[]map[string]interface{}, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *CloudProductControllerIndexDocsRequest) SetDocuments(v []map[string]interface{})`

SetDocuments sets Documents field to given value.


### GetReplace

`func (o *CloudProductControllerIndexDocsRequest) GetReplace() bool`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *CloudProductControllerIndexDocsRequest) GetReplaceOk() (*bool, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *CloudProductControllerIndexDocsRequest) SetReplace(v bool)`

SetReplace sets Replace field to given value.

### HasReplace

`func (o *CloudProductControllerIndexDocsRequest) HasReplace() bool`

HasReplace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



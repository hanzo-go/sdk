# FormImportIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to **bool** | File also files each new response in the CRM, at the time it was made. Nothing is mailed and no webhook is told. | [optional] 
**Id** | Pointer to **string** | ID is the form&#39;s id or slug, from the path. | [optional] 
**Responses** | Pointer to [**[]FormImported**](FormImported.md) | Responses are the responses, at most 500. | [optional] 

## Methods

### NewFormImportIn

`func NewFormImportIn() *FormImportIn`

NewFormImportIn instantiates a new FormImportIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormImportInWithDefaults

`func NewFormImportInWithDefaults() *FormImportIn`

NewFormImportInWithDefaults instantiates a new FormImportIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *FormImportIn) GetFile() bool`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *FormImportIn) GetFileOk() (*bool, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *FormImportIn) SetFile(v bool)`

SetFile sets File field to given value.

### HasFile

`func (o *FormImportIn) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetId

`func (o *FormImportIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormImportIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormImportIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormImportIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResponses

`func (o *FormImportIn) GetResponses() []FormImported`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *FormImportIn) GetResponsesOk() (*[]FormImported, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *FormImportIn) SetResponses(v []FormImported)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *FormImportIn) HasResponses() bool`

HasResponses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



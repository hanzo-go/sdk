# EsignState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the document id. | [optional] 
**Status** | Pointer to **string** | Status is PENDING while it is out for signature. | [optional] 
**Title** | Pointer to **string** | Title is the document&#39;s name. | [optional] 

## Methods

### NewEsignState

`func NewEsignState() *EsignState`

NewEsignState instantiates a new EsignState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignStateWithDefaults

`func NewEsignStateWithDefaults() *EsignState`

NewEsignStateWithDefaults instantiates a new EsignState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EsignState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsignState) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsignState) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsignState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *EsignState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EsignState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EsignState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EsignState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *EsignState) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EsignState) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EsignState) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EsignState) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



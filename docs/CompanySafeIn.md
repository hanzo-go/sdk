# CompanySafeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentIds** | Pointer to **[]string** | DocumentIDs are data room document ids to raise a signature request over. Required. | [optional] 
**Signers** | Pointer to [**[]CompanySigner**](CompanySigner.md) | Signers are the recipients, each a name and an email. Required. | [optional] 

## Methods

### NewCompanySafeIn

`func NewCompanySafeIn() *CompanySafeIn`

NewCompanySafeIn instantiates a new CompanySafeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanySafeInWithDefaults

`func NewCompanySafeInWithDefaults() *CompanySafeIn`

NewCompanySafeInWithDefaults instantiates a new CompanySafeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentIds

`func (o *CompanySafeIn) GetDocumentIds() []string`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *CompanySafeIn) GetDocumentIdsOk() (*[]string, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *CompanySafeIn) SetDocumentIds(v []string)`

SetDocumentIds sets DocumentIds field to given value.

### HasDocumentIds

`func (o *CompanySafeIn) HasDocumentIds() bool`

HasDocumentIds returns a boolean if a field has been set.

### GetSigners

`func (o *CompanySafeIn) GetSigners() []CompanySigner`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *CompanySafeIn) GetSignersOk() (*[]CompanySigner, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *CompanySafeIn) SetSigners(v []CompanySigner)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *CompanySafeIn) HasSigners() bool`

HasSigners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



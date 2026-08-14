# Approval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code overrides the minted code; else the requested vanity code, else a derived slug. | [optional] 
**Id** | Pointer to **string** | ID is the affiliate to approve, from the path. | [optional] 

## Methods

### NewApproval

`func NewApproval() *Approval`

NewApproval instantiates a new Approval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWithDefaults

`func NewApprovalWithDefaults() *Approval`

NewApprovalWithDefaults instantiates a new Approval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Approval) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Approval) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Approval) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Approval) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetId

`func (o *Approval) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Approval) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Approval) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Approval) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



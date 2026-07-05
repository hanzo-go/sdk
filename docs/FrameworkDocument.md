# FrameworkDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Doctype** | Pointer to **string** |  | [optional] 
**Docstatus** | Pointer to **int32** | &#39;0&#x3D;draft, 1&#x3D;submitted, 2&#x3D;cancelled&#39; | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewFrameworkDocument

`func NewFrameworkDocument() *FrameworkDocument`

NewFrameworkDocument instantiates a new FrameworkDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameworkDocumentWithDefaults

`func NewFrameworkDocumentWithDefaults() *FrameworkDocument`

NewFrameworkDocumentWithDefaults instantiates a new FrameworkDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FrameworkDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FrameworkDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FrameworkDocument) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FrameworkDocument) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDoctype

`func (o *FrameworkDocument) GetDoctype() string`

GetDoctype returns the Doctype field if non-nil, zero value otherwise.

### GetDoctypeOk

`func (o *FrameworkDocument) GetDoctypeOk() (*string, bool)`

GetDoctypeOk returns a tuple with the Doctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctype

`func (o *FrameworkDocument) SetDoctype(v string)`

SetDoctype sets Doctype field to given value.

### HasDoctype

`func (o *FrameworkDocument) HasDoctype() bool`

HasDoctype returns a boolean if a field has been set.

### GetDocstatus

`func (o *FrameworkDocument) GetDocstatus() int32`

GetDocstatus returns the Docstatus field if non-nil, zero value otherwise.

### GetDocstatusOk

`func (o *FrameworkDocument) GetDocstatusOk() (*int32, bool)`

GetDocstatusOk returns a tuple with the Docstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocstatus

`func (o *FrameworkDocument) SetDocstatus(v int32)`

SetDocstatus sets Docstatus field to given value.

### HasDocstatus

`func (o *FrameworkDocument) HasDocstatus() bool`

HasDocstatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FrameworkDocument) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FrameworkDocument) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FrameworkDocument) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FrameworkDocument) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FrameworkDocument) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FrameworkDocument) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FrameworkDocument) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FrameworkDocument) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# O11yO11yTransactionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorized** | Pointer to **bool** | Authorized says whether the caller may do it. | [optional] 
**Object** | Pointer to [**O11yO11yObject**](O11yO11yObject.md) | Object is the resource it would act on. | [optional] 
**Relation** | Pointer to **string** | Relation is the verb that was asked about. | [optional] 

## Methods

### NewO11yO11yTransactionResult

`func NewO11yO11yTransactionResult() *O11yO11yTransactionResult`

NewO11yO11yTransactionResult instantiates a new O11yO11yTransactionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTransactionResultWithDefaults

`func NewO11yO11yTransactionResultWithDefaults() *O11yO11yTransactionResult`

NewO11yO11yTransactionResultWithDefaults instantiates a new O11yO11yTransactionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorized

`func (o *O11yO11yTransactionResult) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *O11yO11yTransactionResult) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *O11yO11yTransactionResult) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *O11yO11yTransactionResult) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### GetObject

`func (o *O11yO11yTransactionResult) GetObject() O11yO11yObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *O11yO11yTransactionResult) GetObjectOk() (*O11yO11yObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *O11yO11yTransactionResult) SetObject(v O11yO11yObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *O11yO11yTransactionResult) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRelation

`func (o *O11yO11yTransactionResult) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *O11yO11yTransactionResult) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *O11yO11yTransactionResult) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *O11yO11yTransactionResult) HasRelation() bool`

HasRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



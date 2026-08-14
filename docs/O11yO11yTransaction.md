# O11yO11yTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to [**O11yO11yObject**](O11yO11yObject.md) | Object is the resource the verb would act on. | [optional] 
**Relation** | Pointer to **string** | Relation is the verb being asked about, e.g. read, create, update, delete. | [optional] 

## Methods

### NewO11yO11yTransaction

`func NewO11yO11yTransaction() *O11yO11yTransaction`

NewO11yO11yTransaction instantiates a new O11yO11yTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTransactionWithDefaults

`func NewO11yO11yTransactionWithDefaults() *O11yO11yTransaction`

NewO11yO11yTransactionWithDefaults instantiates a new O11yO11yTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *O11yO11yTransaction) GetObject() O11yO11yObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *O11yO11yTransaction) GetObjectOk() (*O11yO11yObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *O11yO11yTransaction) SetObject(v O11yO11yObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *O11yO11yTransaction) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRelation

`func (o *O11yO11yTransaction) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *O11yO11yTransaction) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *O11yO11yTransaction) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *O11yO11yTransaction) HasRelation() bool`

HasRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



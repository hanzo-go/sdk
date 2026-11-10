# PatentImplBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Implementations** | Pointer to [**[]PatentPatentImplementation**](PatentPatentImplementation.md) | Implementations are the records to append, at most 2000. | [optional] 

## Methods

### NewPatentImplBatch

`func NewPatentImplBatch() *PatentImplBatch`

NewPatentImplBatch instantiates a new PatentImplBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentImplBatchWithDefaults

`func NewPatentImplBatchWithDefaults() *PatentImplBatch`

NewPatentImplBatchWithDefaults instantiates a new PatentImplBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImplementations

`func (o *PatentImplBatch) GetImplementations() []PatentPatentImplementation`

GetImplementations returns the Implementations field if non-nil, zero value otherwise.

### GetImplementationsOk

`func (o *PatentImplBatch) GetImplementationsOk() (*[]PatentPatentImplementation, bool)`

GetImplementationsOk returns a tuple with the Implementations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementations

`func (o *PatentImplBatch) SetImplementations(v []PatentPatentImplementation)`

SetImplementations sets Implementations field to given value.

### HasImplementations

`func (o *PatentImplBatch) HasImplementations() bool`

HasImplementations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PatentDisclosureBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclosures** | Pointer to [**[]PatentPatentDisclosure**](PatentPatentDisclosure.md) | Disclosures are the documents to append, at most 2000. | [optional] 

## Methods

### NewPatentDisclosureBatch

`func NewPatentDisclosureBatch() *PatentDisclosureBatch`

NewPatentDisclosureBatch instantiates a new PatentDisclosureBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentDisclosureBatchWithDefaults

`func NewPatentDisclosureBatchWithDefaults() *PatentDisclosureBatch`

NewPatentDisclosureBatchWithDefaults instantiates a new PatentDisclosureBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclosures

`func (o *PatentDisclosureBatch) GetDisclosures() []PatentPatentDisclosure`

GetDisclosures returns the Disclosures field if non-nil, zero value otherwise.

### GetDisclosuresOk

`func (o *PatentDisclosureBatch) GetDisclosuresOk() (*[]PatentPatentDisclosure, bool)`

GetDisclosuresOk returns a tuple with the Disclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosures

`func (o *PatentDisclosureBatch) SetDisclosures(v []PatentPatentDisclosure)`

SetDisclosures sets Disclosures field to given value.

### HasDisclosures

`func (o *PatentDisclosureBatch) HasDisclosures() bool`

HasDisclosures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



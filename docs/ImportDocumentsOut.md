# ImportDocumentsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formation** | Pointer to [**Formation**](Formation.md) | Formation is the org&#39;s incorporation record with the imported document ids. | [optional] 
**Ingested** | Pointer to **int32** | Ingested is how many files this call put in the data room. | [optional] 

## Methods

### NewImportDocumentsOut

`func NewImportDocumentsOut() *ImportDocumentsOut`

NewImportDocumentsOut instantiates a new ImportDocumentsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportDocumentsOutWithDefaults

`func NewImportDocumentsOutWithDefaults() *ImportDocumentsOut`

NewImportDocumentsOutWithDefaults instantiates a new ImportDocumentsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormation

`func (o *ImportDocumentsOut) GetFormation() Formation`

GetFormation returns the Formation field if non-nil, zero value otherwise.

### GetFormationOk

`func (o *ImportDocumentsOut) GetFormationOk() (*Formation, bool)`

GetFormationOk returns a tuple with the Formation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormation

`func (o *ImportDocumentsOut) SetFormation(v Formation)`

SetFormation sets Formation field to given value.

### HasFormation

`func (o *ImportDocumentsOut) HasFormation() bool`

HasFormation returns a boolean if a field has been set.

### GetIngested

`func (o *ImportDocumentsOut) GetIngested() int32`

GetIngested returns the Ingested field if non-nil, zero value otherwise.

### GetIngestedOk

`func (o *ImportDocumentsOut) GetIngestedOk() (*int32, bool)`

GetIngestedOk returns a tuple with the Ingested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngested

`func (o *ImportDocumentsOut) SetIngested(v int32)`

SetIngested sets Ingested field to given value.

### HasIngested

`func (o *ImportDocumentsOut) HasIngested() bool`

HasIngested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudImportDocumentsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formation** | Pointer to [**CloudFormation**](CloudFormation.md) | Formation is the org&#39;s incorporation record with the imported document ids. | [optional] 
**Ingested** | Pointer to **int32** | Ingested is how many files this call put in the data room. | [optional] 

## Methods

### NewCloudImportDocumentsOut

`func NewCloudImportDocumentsOut() *CloudImportDocumentsOut`

NewCloudImportDocumentsOut instantiates a new CloudImportDocumentsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudImportDocumentsOutWithDefaults

`func NewCloudImportDocumentsOutWithDefaults() *CloudImportDocumentsOut`

NewCloudImportDocumentsOutWithDefaults instantiates a new CloudImportDocumentsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormation

`func (o *CloudImportDocumentsOut) GetFormation() CloudFormation`

GetFormation returns the Formation field if non-nil, zero value otherwise.

### GetFormationOk

`func (o *CloudImportDocumentsOut) GetFormationOk() (*CloudFormation, bool)`

GetFormationOk returns a tuple with the Formation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormation

`func (o *CloudImportDocumentsOut) SetFormation(v CloudFormation)`

SetFormation sets Formation field to given value.

### HasFormation

`func (o *CloudImportDocumentsOut) HasFormation() bool`

HasFormation returns a boolean if a field has been set.

### GetIngested

`func (o *CloudImportDocumentsOut) GetIngested() int32`

GetIngested returns the Ingested field if non-nil, zero value otherwise.

### GetIngestedOk

`func (o *CloudImportDocumentsOut) GetIngestedOk() (*int32, bool)`

GetIngestedOk returns a tuple with the Ingested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngested

`func (o *CloudImportDocumentsOut) SetIngested(v int32)`

SetIngested sets Ingested field to given value.

### HasIngested

`func (o *CloudImportDocumentsOut) HasIngested() bool`

HasIngested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



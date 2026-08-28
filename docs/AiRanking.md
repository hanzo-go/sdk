# AiRanking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]AiRankedDoc**](AiRankedDoc.md) |  | [optional] 
**Usage** | Pointer to [**AiRankUsage**](AiRankUsage.md) |  | [optional] 

## Methods

### NewAiRanking

`func NewAiRanking() *AiRanking`

NewAiRanking instantiates a new AiRanking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRankingWithDefaults

`func NewAiRankingWithDefaults() *AiRanking`

NewAiRankingWithDefaults instantiates a new AiRanking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiRanking) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiRanking) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiRanking) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiRanking) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetObject

`func (o *AiRanking) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiRanking) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiRanking) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiRanking) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetResults

`func (o *AiRanking) GetResults() []AiRankedDoc`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AiRanking) GetResultsOk() (*[]AiRankedDoc, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AiRanking) SetResults(v []AiRankedDoc)`

SetResults sets Results field to given value.

### HasResults

`func (o *AiRanking) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetUsage

`func (o *AiRanking) GetUsage() AiRankUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiRanking) GetUsageOk() (*AiRankUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiRanking) SetUsage(v AiRankUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiRanking) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



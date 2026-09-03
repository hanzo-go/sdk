# WebEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the engine, matching the &#x60;engine&#x60; stamped on each result. | [optional] 
**Outcome** | Pointer to **string** | Outcome is \&quot;answered\&quot;, \&quot;blind\&quot; or \&quot;failed\&quot; — see outcome.go. \&quot;blind\&quot; means the page came back and no results could be read out of it. | [optional] 
**Results** | Pointer to **int64** | Results is how many hits this engine contributed, before the merge deduplicated them against the others. | [optional] 

## Methods

### NewWebEngine

`func NewWebEngine() *WebEngine`

NewWebEngine instantiates a new WebEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebEngineWithDefaults

`func NewWebEngineWithDefaults() *WebEngine`

NewWebEngineWithDefaults instantiates a new WebEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebEngine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebEngine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebEngine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebEngine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutcome

`func (o *WebEngine) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *WebEngine) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *WebEngine) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *WebEngine) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetResults

`func (o *WebEngine) GetResults() int64`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *WebEngine) GetResultsOk() (*int64, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *WebEngine) SetResults(v int64)`

SetResults sets Results field to given value.

### HasResults

`func (o *WebEngine) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PatentPatentWriteOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Canonical** | Pointer to **int64** | Canonical is how many entries the org&#39;s log now holds for this kind. | [optional] 
**Project** | Pointer to **string** | Project is the scope the batch was filed under, as the server decided it. | [optional] 
**Recorded** | Pointer to **int64** | Recorded is how many new versions this call appended. | [optional] 

## Methods

### NewPatentPatentWriteOut

`func NewPatentPatentWriteOut() *PatentPatentWriteOut`

NewPatentPatentWriteOut instantiates a new PatentPatentWriteOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentWriteOutWithDefaults

`func NewPatentPatentWriteOutWithDefaults() *PatentPatentWriteOut`

NewPatentPatentWriteOutWithDefaults instantiates a new PatentPatentWriteOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonical

`func (o *PatentPatentWriteOut) GetCanonical() int64`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *PatentPatentWriteOut) GetCanonicalOk() (*int64, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *PatentPatentWriteOut) SetCanonical(v int64)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *PatentPatentWriteOut) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetProject

`func (o *PatentPatentWriteOut) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatentPatentWriteOut) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatentPatentWriteOut) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatentPatentWriteOut) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRecorded

`func (o *PatentPatentWriteOut) GetRecorded() int64`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *PatentPatentWriteOut) GetRecordedOk() (*int64, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *PatentPatentWriteOut) SetRecorded(v int64)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *PatentPatentWriteOut) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



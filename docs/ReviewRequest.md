# ReviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Glossary** | Pointer to **map[string]string** | Glossary is the terminology the entry was translated under. Its VERSION — the digest of the sorted terms — is part of the entry&#39;s identity, so editing a term yields a new entry rather than overwriting the old rendering. | [optional] 
**Source** | Pointer to **string** | Source is the ORIGINAL string this entry translates. Required; part of the entry&#39;s identity, so a different source is a different entry. | [optional] 
**State** | Pointer to **string** | State is the entry&#39;s new position on the review ladder: suggested, approved or published. &#x60;machine&#x60; is engine-only and is refused here — a human may not demote a string back into the churn. | [optional] 
**Target** | Pointer to **string** | Target is the target language tag (BCP-47, e.g. \&quot;es\&quot; or \&quot;pt-BR\&quot;). Required; part of the entry&#39;s identity. | [optional] 
**Text** | Pointer to **string** | Text is the reviewed translation to store. A human write always wins over the stored value. | [optional] 
**Tier** | Pointer to **string** | Tier is the engine tier the entry belongs to, quality (the default) or bulk. Part of the entry&#39;s identity: the two tiers keep separate renderings. | [optional] 

## Methods

### NewReviewRequest

`func NewReviewRequest() *ReviewRequest`

NewReviewRequest instantiates a new ReviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewRequestWithDefaults

`func NewReviewRequestWithDefaults() *ReviewRequest`

NewReviewRequestWithDefaults instantiates a new ReviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlossary

`func (o *ReviewRequest) GetGlossary() map[string]string`

GetGlossary returns the Glossary field if non-nil, zero value otherwise.

### GetGlossaryOk

`func (o *ReviewRequest) GetGlossaryOk() (*map[string]string, bool)`

GetGlossaryOk returns a tuple with the Glossary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlossary

`func (o *ReviewRequest) SetGlossary(v map[string]string)`

SetGlossary sets Glossary field to given value.

### HasGlossary

`func (o *ReviewRequest) HasGlossary() bool`

HasGlossary returns a boolean if a field has been set.

### GetSource

`func (o *ReviewRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReviewRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReviewRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ReviewRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *ReviewRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReviewRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReviewRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ReviewRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTarget

`func (o *ReviewRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ReviewRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ReviewRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ReviewRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetText

`func (o *ReviewRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ReviewRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ReviewRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ReviewRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTier

`func (o *ReviewRequest) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *ReviewRequest) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *ReviewRequest) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *ReviewRequest) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



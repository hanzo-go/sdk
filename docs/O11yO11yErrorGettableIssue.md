# O11yO11yErrorGettableIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issue** | Pointer to [**O11yO11yErrorIssue**](O11yO11yErrorIssue.md) | Issue is the lifecycle row. | [optional] 
**LatestEvent** | Pointer to [**O11yO11yOccurrence**](O11yO11yOccurrence.md) | LatestEvent is the most recent occurrence that landed on the issue. | [optional] 

## Methods

### NewO11yO11yErrorGettableIssue

`func NewO11yO11yErrorGettableIssue() *O11yO11yErrorGettableIssue`

NewO11yO11yErrorGettableIssue instantiates a new O11yO11yErrorGettableIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yErrorGettableIssueWithDefaults

`func NewO11yO11yErrorGettableIssueWithDefaults() *O11yO11yErrorGettableIssue`

NewO11yO11yErrorGettableIssueWithDefaults instantiates a new O11yO11yErrorGettableIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssue

`func (o *O11yO11yErrorGettableIssue) GetIssue() O11yO11yErrorIssue`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *O11yO11yErrorGettableIssue) GetIssueOk() (*O11yO11yErrorIssue, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *O11yO11yErrorGettableIssue) SetIssue(v O11yO11yErrorIssue)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *O11yO11yErrorGettableIssue) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetLatestEvent

`func (o *O11yO11yErrorGettableIssue) GetLatestEvent() O11yO11yOccurrence`

GetLatestEvent returns the LatestEvent field if non-nil, zero value otherwise.

### GetLatestEventOk

`func (o *O11yO11yErrorGettableIssue) GetLatestEventOk() (*O11yO11yOccurrence, bool)`

GetLatestEventOk returns a tuple with the LatestEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestEvent

`func (o *O11yO11yErrorGettableIssue) SetLatestEvent(v O11yO11yOccurrence)`

SetLatestEvent sets LatestEvent field to given value.

### HasLatestEvent

`func (o *O11yO11yErrorGettableIssue) HasLatestEvent() bool`

HasLatestEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



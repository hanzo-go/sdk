# EmailPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the provider (ses, sendgrid) or the campaign. | [optional] 
**Tally** | Pointer to [**EmailTally**](EmailTally.md) | Tally is its mail counted over the range. | [optional] 

## Methods

### NewEmailPart

`func NewEmailPart() *EmailPart`

NewEmailPart instantiates a new EmailPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailPartWithDefaults

`func NewEmailPartWithDefaults() *EmailPart`

NewEmailPartWithDefaults instantiates a new EmailPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailPart) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailPart) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailPart) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailPart) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTally

`func (o *EmailPart) GetTally() EmailTally`

GetTally returns the Tally field if non-nil, zero value otherwise.

### GetTallyOk

`func (o *EmailPart) GetTallyOk() (*EmailTally, bool)`

GetTallyOk returns a tuple with the Tally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTally

`func (o *EmailPart) SetTally(v EmailTally)`

SetTally sets Tally field to given value.

### HasTally

`func (o *EmailPart) HasTally() bool`

HasTally returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# O11yO11yQueryWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message is the warning. | [optional] 
**Url** | Pointer to **string** | URL points at the relevant documentation. | [optional] 
**Warnings** | Pointer to [**[]O11yO11yQueryWarningNote**](O11yO11yQueryWarningNote.md) | Warnings carries additional notes. | [optional] 

## Methods

### NewO11yO11yQueryWarning

`func NewO11yO11yQueryWarning() *O11yO11yQueryWarning`

NewO11yO11yQueryWarning instantiates a new O11yO11yQueryWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueryWarningWithDefaults

`func NewO11yO11yQueryWarningWithDefaults() *O11yO11yQueryWarning`

NewO11yO11yQueryWarningWithDefaults instantiates a new O11yO11yQueryWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *O11yO11yQueryWarning) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yO11yQueryWarning) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yO11yQueryWarning) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yO11yQueryWarning) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetUrl

`func (o *O11yO11yQueryWarning) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *O11yO11yQueryWarning) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *O11yO11yQueryWarning) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *O11yO11yQueryWarning) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWarnings

`func (o *O11yO11yQueryWarning) GetWarnings() []O11yO11yQueryWarningNote`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *O11yO11yQueryWarning) GetWarningsOk() (*[]O11yO11yQueryWarningNote, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *O11yO11yQueryWarning) SetWarnings(v []O11yO11yQueryWarningNote)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *O11yO11yQueryWarning) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



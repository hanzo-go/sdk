# O11yQueryWarnData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Warnings** | Pointer to [**[]O11yQueryWarnDataAdditional**](O11yQueryWarnDataAdditional.md) |  | [optional] 

## Methods

### NewO11yQueryWarnData

`func NewO11yQueryWarnData() *O11yQueryWarnData`

NewO11yQueryWarnData instantiates a new O11yQueryWarnData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yQueryWarnDataWithDefaults

`func NewO11yQueryWarnDataWithDefaults() *O11yQueryWarnData`

NewO11yQueryWarnDataWithDefaults instantiates a new O11yQueryWarnData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *O11yQueryWarnData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yQueryWarnData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yQueryWarnData) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yQueryWarnData) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetUrl

`func (o *O11yQueryWarnData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *O11yQueryWarnData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *O11yQueryWarnData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *O11yQueryWarnData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWarnings

`func (o *O11yQueryWarnData) GetWarnings() []O11yQueryWarnDataAdditional`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *O11yQueryWarnData) GetWarningsOk() (*[]O11yQueryWarnDataAdditional, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *O11yQueryWarnData) SetWarnings(v []O11yQueryWarnDataAdditional)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *O11yQueryWarnData) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



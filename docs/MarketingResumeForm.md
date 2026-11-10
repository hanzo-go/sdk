# MarketingResumeForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **int64** | Expiry is how long the form is good for, in seconds. | [optional] 
**Fields** | Pointer to **map[string]string** | Fields go in the form ahead of the file, exactly as given. | [optional] 
**Key** | Pointer to **string** | Key is the object the file becomes, to name as the application&#39;s resume_file. | [optional] 
**Url** | Pointer to **string** | URL is where the form posts. | [optional] 

## Methods

### NewMarketingResumeForm

`func NewMarketingResumeForm() *MarketingResumeForm`

NewMarketingResumeForm instantiates a new MarketingResumeForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingResumeFormWithDefaults

`func NewMarketingResumeFormWithDefaults() *MarketingResumeForm`

NewMarketingResumeFormWithDefaults instantiates a new MarketingResumeForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *MarketingResumeForm) GetExpiry() int64`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *MarketingResumeForm) GetExpiryOk() (*int64, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *MarketingResumeForm) SetExpiry(v int64)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *MarketingResumeForm) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetFields

`func (o *MarketingResumeForm) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MarketingResumeForm) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MarketingResumeForm) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MarketingResumeForm) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetKey

`func (o *MarketingResumeForm) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MarketingResumeForm) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MarketingResumeForm) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MarketingResumeForm) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUrl

`func (o *MarketingResumeForm) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MarketingResumeForm) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MarketingResumeForm) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MarketingResumeForm) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



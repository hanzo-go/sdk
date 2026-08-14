# Number

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capable** | Pointer to **[]string** | voice | sms | mms | fax | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**E164** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Monthly** | Pointer to **int32** | minor units, as the carrier quoted it | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewNumber

`func NewNumber() *Number`

NewNumber instantiates a new Number object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberWithDefaults

`func NewNumberWithDefaults() *Number`

NewNumberWithDefaults instantiates a new Number object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapable

`func (o *Number) GetCapable() []string`

GetCapable returns the Capable field if non-nil, zero value otherwise.

### GetCapableOk

`func (o *Number) GetCapableOk() (*[]string, bool)`

GetCapableOk returns a tuple with the Capable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapable

`func (o *Number) SetCapable(v []string)`

SetCapable sets Capable field to given value.

### HasCapable

`func (o *Number) HasCapable() bool`

HasCapable returns a boolean if a field has been set.

### GetCountry

`func (o *Number) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Number) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Number) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Number) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *Number) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Number) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Number) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Number) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetE164

`func (o *Number) GetE164() string`

GetE164 returns the E164 field if non-nil, zero value otherwise.

### GetE164Ok

`func (o *Number) GetE164Ok() (*string, bool)`

GetE164Ok returns a tuple with the E164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE164

`func (o *Number) SetE164(v string)`

SetE164 sets E164 field to given value.

### HasE164

`func (o *Number) HasE164() bool`

HasE164 returns a boolean if a field has been set.

### GetId

`func (o *Number) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Number) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Number) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Number) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMonthly

`func (o *Number) GetMonthly() int32`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *Number) GetMonthlyOk() (*int32, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *Number) SetMonthly(v int32)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *Number) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.

### GetOrg

`func (o *Number) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Number) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Number) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Number) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetType

`func (o *Number) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Number) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Number) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Number) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



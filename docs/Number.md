# Number

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capable** | Pointer to **[]string** | Capable is what the number can carry: any of \&quot;voice\&quot;, \&quot;sms\&quot;, \&quot;mms\&quot;, \&quot;fax\&quot;. A number missing \&quot;sms\&quot; cannot send one no matter what this platform does. | [optional] 
**Country** | Pointer to **string** | Country is the ISO 3166-1 alpha-2 code the number is issued under. Numbering is national, so this is what makes a search answerable at all. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO 4217 code Monthly is denominated in. Without it the number beside it means nothing, so the two are always read together. | [optional] 
**E164** | Pointer to **string** | E164 is the number in E.164: a leading + and digits only, no spaces or dashes. That is what a carrier accepts and what a search result must be bought by. | [optional] 
**Id** | Pointer to **string** | ID is the carrier&#39;s handle for the number, and the id every route here addresses it by. It is not the number itself — see E164. | [optional] 
**Monthly** | Pointer to **int64** | Monthly is the recurring rental in the MINOR unit of Currency (cents for USD), exactly as the carrier quoted it. It is a price, not a charge: nothing is billed by this field. | [optional] 
**Org** | Pointer to **string** | Org is the tenant holding the number. A search result carries none — nobody holds it yet — which is how an available number is told from a held one. | [optional] 
**Type** | Pointer to **string** | Type is what kind of number it is: \&quot;local\&quot;, \&quot;national\&quot;, \&quot;tollfree\&quot; or \&quot;mobile\&quot;. It decides both price and what a carrier will let it originate. | [optional] 

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

`func (o *Number) GetMonthly() int64`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *Number) GetMonthlyOk() (*int64, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *Number) SetMonthly(v int64)`

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



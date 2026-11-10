# MarketingImportOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basis** | Pointer to **string** | Basis is the consent basis recorded for every row: explicit, legitimate_interest or unverified (the default). | [optional] 
**DryRun** | Pointer to **bool** | DryRun reads and judges every row and writes nothing. | [optional] 
**Format** | Pointer to **string** | Format is csv or json; ignored for a source. | [optional] 
**From** | Pointer to **string** | From is a source to read instead of a file, as the host offers them: commerce (the org&#39;s commerce form subscribers) or crm (the confirmations the retired newsletter filed in the org&#39;s CRM). Empty is a file. | [optional] 
**Key** | Pointer to **string** | Key makes creating the import idempotent: a second create with the same key answers the first import. | [optional] 
**Map** | Pointer to [**MarketingMapping**](MarketingMapping.md) | Map says which column holds what; a column left out is found by its header. | [optional] 
**Name** | Pointer to **string** | Name is the file&#39;s name, for the record. | [optional] 
**Size** | Pointer to **int64** | Size is the file&#39;s size in bytes, when known; the upload is refused past it. | [optional] 
**Status** | Pointer to **string** | Status is what a new row becomes: subscribed (the default), unsubscribed, bounced, complained or cleaned, as an export from another provider records. A row whose file carries its own status (map.status) takes that one. | [optional] 
**Tags** | Pointer to **[]string** | Tags are put on every row the import writes. | [optional] 
**Verify** | Pointer to **bool** | Verify checks each address deliverable before writing it. | [optional] 

## Methods

### NewMarketingImportOptions

`func NewMarketingImportOptions() *MarketingImportOptions`

NewMarketingImportOptions instantiates a new MarketingImportOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingImportOptionsWithDefaults

`func NewMarketingImportOptionsWithDefaults() *MarketingImportOptions`

NewMarketingImportOptionsWithDefaults instantiates a new MarketingImportOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *MarketingImportOptions) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *MarketingImportOptions) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *MarketingImportOptions) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *MarketingImportOptions) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetDryRun

`func (o *MarketingImportOptions) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *MarketingImportOptions) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *MarketingImportOptions) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *MarketingImportOptions) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFormat

`func (o *MarketingImportOptions) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MarketingImportOptions) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MarketingImportOptions) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MarketingImportOptions) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFrom

`func (o *MarketingImportOptions) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MarketingImportOptions) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MarketingImportOptions) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MarketingImportOptions) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetKey

`func (o *MarketingImportOptions) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MarketingImportOptions) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MarketingImportOptions) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MarketingImportOptions) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMap

`func (o *MarketingImportOptions) GetMap() MarketingMapping`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *MarketingImportOptions) GetMapOk() (*MarketingMapping, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *MarketingImportOptions) SetMap(v MarketingMapping)`

SetMap sets Map field to given value.

### HasMap

`func (o *MarketingImportOptions) HasMap() bool`

HasMap returns a boolean if a field has been set.

### GetName

`func (o *MarketingImportOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingImportOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingImportOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingImportOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *MarketingImportOptions) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MarketingImportOptions) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MarketingImportOptions) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MarketingImportOptions) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *MarketingImportOptions) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketingImportOptions) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketingImportOptions) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketingImportOptions) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *MarketingImportOptions) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MarketingImportOptions) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MarketingImportOptions) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MarketingImportOptions) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVerify

`func (o *MarketingImportOptions) GetVerify() bool`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *MarketingImportOptions) GetVerifyOk() (*bool, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *MarketingImportOptions) SetVerify(v bool)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *MarketingImportOptions) HasVerify() bool`

HasVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



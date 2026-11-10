# FormExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Csv** | Pointer to **string** | CSV is the responses as CSV, for csv: one row each, one column per field, then the hidden values, the respondent and the times. | [optional] 
**Filename** | Pointer to **string** | Filename is what to save it as. | [optional] 
**Format** | Pointer to **string** | Format is json or csv. | [optional] 
**Items** | Pointer to [**[]FormResponse**](FormResponse.md) | Items are the responses, for json. | [optional] 
**Next** | Pointer to **int64** |  | [optional] 
**Page** | Pointer to **int64** | Page is this page; Next the page after it, 0 when this is the last. | [optional] 
**Total** | Pointer to **int64** | Total is how many responses answer the filter. | [optional] 

## Methods

### NewFormExport

`func NewFormExport() *FormExport`

NewFormExport instantiates a new FormExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormExportWithDefaults

`func NewFormExportWithDefaults() *FormExport`

NewFormExportWithDefaults instantiates a new FormExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsv

`func (o *FormExport) GetCsv() string`

GetCsv returns the Csv field if non-nil, zero value otherwise.

### GetCsvOk

`func (o *FormExport) GetCsvOk() (*string, bool)`

GetCsvOk returns a tuple with the Csv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsv

`func (o *FormExport) SetCsv(v string)`

SetCsv sets Csv field to given value.

### HasCsv

`func (o *FormExport) HasCsv() bool`

HasCsv returns a boolean if a field has been set.

### GetFilename

`func (o *FormExport) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FormExport) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FormExport) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *FormExport) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFormat

`func (o *FormExport) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *FormExport) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *FormExport) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *FormExport) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetItems

`func (o *FormExport) GetItems() []FormResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FormExport) GetItemsOk() (*[]FormResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FormExport) SetItems(v []FormResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *FormExport) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNext

`func (o *FormExport) GetNext() int64`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FormExport) GetNextOk() (*int64, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FormExport) SetNext(v int64)`

SetNext sets Next field to given value.

### HasNext

`func (o *FormExport) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPage

`func (o *FormExport) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *FormExport) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *FormExport) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *FormExport) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotal

`func (o *FormExport) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FormExport) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FormExport) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FormExport) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



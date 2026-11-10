# MarketingImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to **string** | By is who created it. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt, UpdatedAt and FinishedAt are unix seconds. | [optional] 
**Error** | Pointer to **string** | Error is why a failed import stopped. | [optional] 
**FinishedAt** | Pointer to **int64** | FinishedAt is when it ended; 0 while it runs. | [optional] 
**Id** | Pointer to **string** | ID is the import&#39;s id (\&quot;imp_\&quot; and 96 random bits in hex): the batch every subscriber it adds records as its source ref. | [optional] 
**List** | Pointer to **string** | List is the list it fills. | [optional] 
**Options** | Pointer to [**MarketingImportOptions**](MarketingImportOptions.md) | Options are what the import was created with. | [optional] 
**Parsed** | Pointer to **int64** | Parsed is how many bytes of the file the worker has read and judged. | [optional] 
**Received** | Pointer to **int64** | Received is how many bytes of the file have arrived; the next piece starts here. | [optional] 
**Report** | Pointer to [**MarketingReport**](MarketingReport.md) | Report is what became of the rows read so far. | [optional] 
**State** | Pointer to **string** | State is receiving (waiting for its bytes), queued, processing, done, failed or cancelled. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is the last progress. | [optional] 

## Methods

### NewMarketingImport

`func NewMarketingImport() *MarketingImport`

NewMarketingImport instantiates a new MarketingImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingImportWithDefaults

`func NewMarketingImportWithDefaults() *MarketingImport`

NewMarketingImportWithDefaults instantiates a new MarketingImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *MarketingImport) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *MarketingImport) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *MarketingImport) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *MarketingImport) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketingImport) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketingImport) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketingImport) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketingImport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetError

`func (o *MarketingImport) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MarketingImport) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MarketingImport) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MarketingImport) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFinishedAt

`func (o *MarketingImport) GetFinishedAt() int64`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *MarketingImport) GetFinishedAtOk() (*int64, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *MarketingImport) SetFinishedAt(v int64)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *MarketingImport) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *MarketingImport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingImport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingImport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingImport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetList

`func (o *MarketingImport) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *MarketingImport) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *MarketingImport) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *MarketingImport) HasList() bool`

HasList returns a boolean if a field has been set.

### GetOptions

`func (o *MarketingImport) GetOptions() MarketingImportOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MarketingImport) GetOptionsOk() (*MarketingImportOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MarketingImport) SetOptions(v MarketingImportOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MarketingImport) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetParsed

`func (o *MarketingImport) GetParsed() int64`

GetParsed returns the Parsed field if non-nil, zero value otherwise.

### GetParsedOk

`func (o *MarketingImport) GetParsedOk() (*int64, bool)`

GetParsedOk returns a tuple with the Parsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsed

`func (o *MarketingImport) SetParsed(v int64)`

SetParsed sets Parsed field to given value.

### HasParsed

`func (o *MarketingImport) HasParsed() bool`

HasParsed returns a boolean if a field has been set.

### GetReceived

`func (o *MarketingImport) GetReceived() int64`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *MarketingImport) GetReceivedOk() (*int64, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *MarketingImport) SetReceived(v int64)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *MarketingImport) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetReport

`func (o *MarketingImport) GetReport() MarketingReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *MarketingImport) GetReportOk() (*MarketingReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *MarketingImport) SetReport(v MarketingReport)`

SetReport sets Report field to given value.

### HasReport

`func (o *MarketingImport) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetState

`func (o *MarketingImport) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MarketingImport) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MarketingImport) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MarketingImport) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MarketingImport) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketingImport) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketingImport) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MarketingImport) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



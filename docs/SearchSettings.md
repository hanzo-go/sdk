# SearchSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayedAttributes** | Pointer to **[]string** |  | [optional] 
**SearchableAttributes** | Pointer to **[]string** |  | [optional] 
**FilterableAttributes** | Pointer to [**[]AuthorsErrorError**](AuthorsErrorError.md) |  | [optional] 
**SortableAttributes** | Pointer to **[]string** |  | [optional] 
**RankingRules** | Pointer to **[]string** |  | [optional] 
**StopWords** | Pointer to **[]string** |  | [optional] 
**NonSeparatorTokens** | Pointer to **[]string** |  | [optional] 
**SeparatorTokens** | Pointer to **[]string** |  | [optional] 
**Dictionary** | Pointer to **[]string** |  | [optional] 
**Synonyms** | Pointer to **map[string][]string** |  | [optional] 
**DistinctAttribute** | Pointer to **string** |  | [optional] 
**ProximityPrecision** | Pointer to **string** |  | [optional] 
**TypoTolerance** | Pointer to [**SearchSettingsTypoTolerance**](SearchSettingsTypoTolerance.md) |  | [optional] 
**Faceting** | Pointer to [**SearchSettingsFaceting**](SearchSettingsFaceting.md) |  | [optional] 
**Pagination** | Pointer to [**SearchSettingsPagination**](SearchSettingsPagination.md) |  | [optional] 
**Embedders** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**SearchCutoffMs** | Pointer to **int32** |  | [optional] 
**LocalizedAttributes** | Pointer to [**[]SearchSettingsLocalizedAttributesInner**](SearchSettingsLocalizedAttributesInner.md) |  | [optional] 
**PrefixSearch** | Pointer to **string** |  | [optional] 
**FacetSearch** | Pointer to **bool** |  | [optional] 

## Methods

### NewSearchSettings

`func NewSearchSettings() *SearchSettings`

NewSearchSettings instantiates a new SearchSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSettingsWithDefaults

`func NewSearchSettingsWithDefaults() *SearchSettings`

NewSearchSettingsWithDefaults instantiates a new SearchSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayedAttributes

`func (o *SearchSettings) GetDisplayedAttributes() []string`

GetDisplayedAttributes returns the DisplayedAttributes field if non-nil, zero value otherwise.

### GetDisplayedAttributesOk

`func (o *SearchSettings) GetDisplayedAttributesOk() (*[]string, bool)`

GetDisplayedAttributesOk returns a tuple with the DisplayedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayedAttributes

`func (o *SearchSettings) SetDisplayedAttributes(v []string)`

SetDisplayedAttributes sets DisplayedAttributes field to given value.

### HasDisplayedAttributes

`func (o *SearchSettings) HasDisplayedAttributes() bool`

HasDisplayedAttributes returns a boolean if a field has been set.

### GetSearchableAttributes

`func (o *SearchSettings) GetSearchableAttributes() []string`

GetSearchableAttributes returns the SearchableAttributes field if non-nil, zero value otherwise.

### GetSearchableAttributesOk

`func (o *SearchSettings) GetSearchableAttributesOk() (*[]string, bool)`

GetSearchableAttributesOk returns a tuple with the SearchableAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableAttributes

`func (o *SearchSettings) SetSearchableAttributes(v []string)`

SetSearchableAttributes sets SearchableAttributes field to given value.

### HasSearchableAttributes

`func (o *SearchSettings) HasSearchableAttributes() bool`

HasSearchableAttributes returns a boolean if a field has been set.

### GetFilterableAttributes

`func (o *SearchSettings) GetFilterableAttributes() []AuthorsErrorError`

GetFilterableAttributes returns the FilterableAttributes field if non-nil, zero value otherwise.

### GetFilterableAttributesOk

`func (o *SearchSettings) GetFilterableAttributesOk() (*[]AuthorsErrorError, bool)`

GetFilterableAttributesOk returns a tuple with the FilterableAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterableAttributes

`func (o *SearchSettings) SetFilterableAttributes(v []AuthorsErrorError)`

SetFilterableAttributes sets FilterableAttributes field to given value.

### HasFilterableAttributes

`func (o *SearchSettings) HasFilterableAttributes() bool`

HasFilterableAttributes returns a boolean if a field has been set.

### GetSortableAttributes

`func (o *SearchSettings) GetSortableAttributes() []string`

GetSortableAttributes returns the SortableAttributes field if non-nil, zero value otherwise.

### GetSortableAttributesOk

`func (o *SearchSettings) GetSortableAttributesOk() (*[]string, bool)`

GetSortableAttributesOk returns a tuple with the SortableAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortableAttributes

`func (o *SearchSettings) SetSortableAttributes(v []string)`

SetSortableAttributes sets SortableAttributes field to given value.

### HasSortableAttributes

`func (o *SearchSettings) HasSortableAttributes() bool`

HasSortableAttributes returns a boolean if a field has been set.

### GetRankingRules

`func (o *SearchSettings) GetRankingRules() []string`

GetRankingRules returns the RankingRules field if non-nil, zero value otherwise.

### GetRankingRulesOk

`func (o *SearchSettings) GetRankingRulesOk() (*[]string, bool)`

GetRankingRulesOk returns a tuple with the RankingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankingRules

`func (o *SearchSettings) SetRankingRules(v []string)`

SetRankingRules sets RankingRules field to given value.

### HasRankingRules

`func (o *SearchSettings) HasRankingRules() bool`

HasRankingRules returns a boolean if a field has been set.

### GetStopWords

`func (o *SearchSettings) GetStopWords() []string`

GetStopWords returns the StopWords field if non-nil, zero value otherwise.

### GetStopWordsOk

`func (o *SearchSettings) GetStopWordsOk() (*[]string, bool)`

GetStopWordsOk returns a tuple with the StopWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopWords

`func (o *SearchSettings) SetStopWords(v []string)`

SetStopWords sets StopWords field to given value.

### HasStopWords

`func (o *SearchSettings) HasStopWords() bool`

HasStopWords returns a boolean if a field has been set.

### GetNonSeparatorTokens

`func (o *SearchSettings) GetNonSeparatorTokens() []string`

GetNonSeparatorTokens returns the NonSeparatorTokens field if non-nil, zero value otherwise.

### GetNonSeparatorTokensOk

`func (o *SearchSettings) GetNonSeparatorTokensOk() (*[]string, bool)`

GetNonSeparatorTokensOk returns a tuple with the NonSeparatorTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSeparatorTokens

`func (o *SearchSettings) SetNonSeparatorTokens(v []string)`

SetNonSeparatorTokens sets NonSeparatorTokens field to given value.

### HasNonSeparatorTokens

`func (o *SearchSettings) HasNonSeparatorTokens() bool`

HasNonSeparatorTokens returns a boolean if a field has been set.

### GetSeparatorTokens

`func (o *SearchSettings) GetSeparatorTokens() []string`

GetSeparatorTokens returns the SeparatorTokens field if non-nil, zero value otherwise.

### GetSeparatorTokensOk

`func (o *SearchSettings) GetSeparatorTokensOk() (*[]string, bool)`

GetSeparatorTokensOk returns a tuple with the SeparatorTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparatorTokens

`func (o *SearchSettings) SetSeparatorTokens(v []string)`

SetSeparatorTokens sets SeparatorTokens field to given value.

### HasSeparatorTokens

`func (o *SearchSettings) HasSeparatorTokens() bool`

HasSeparatorTokens returns a boolean if a field has been set.

### GetDictionary

`func (o *SearchSettings) GetDictionary() []string`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *SearchSettings) GetDictionaryOk() (*[]string, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *SearchSettings) SetDictionary(v []string)`

SetDictionary sets Dictionary field to given value.

### HasDictionary

`func (o *SearchSettings) HasDictionary() bool`

HasDictionary returns a boolean if a field has been set.

### GetSynonyms

`func (o *SearchSettings) GetSynonyms() map[string][]string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *SearchSettings) GetSynonymsOk() (*map[string][]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *SearchSettings) SetSynonyms(v map[string][]string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *SearchSettings) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetDistinctAttribute

`func (o *SearchSettings) GetDistinctAttribute() string`

GetDistinctAttribute returns the DistinctAttribute field if non-nil, zero value otherwise.

### GetDistinctAttributeOk

`func (o *SearchSettings) GetDistinctAttributeOk() (*string, bool)`

GetDistinctAttributeOk returns a tuple with the DistinctAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctAttribute

`func (o *SearchSettings) SetDistinctAttribute(v string)`

SetDistinctAttribute sets DistinctAttribute field to given value.

### HasDistinctAttribute

`func (o *SearchSettings) HasDistinctAttribute() bool`

HasDistinctAttribute returns a boolean if a field has been set.

### GetProximityPrecision

`func (o *SearchSettings) GetProximityPrecision() string`

GetProximityPrecision returns the ProximityPrecision field if non-nil, zero value otherwise.

### GetProximityPrecisionOk

`func (o *SearchSettings) GetProximityPrecisionOk() (*string, bool)`

GetProximityPrecisionOk returns a tuple with the ProximityPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximityPrecision

`func (o *SearchSettings) SetProximityPrecision(v string)`

SetProximityPrecision sets ProximityPrecision field to given value.

### HasProximityPrecision

`func (o *SearchSettings) HasProximityPrecision() bool`

HasProximityPrecision returns a boolean if a field has been set.

### GetTypoTolerance

`func (o *SearchSettings) GetTypoTolerance() SearchSettingsTypoTolerance`

GetTypoTolerance returns the TypoTolerance field if non-nil, zero value otherwise.

### GetTypoToleranceOk

`func (o *SearchSettings) GetTypoToleranceOk() (*SearchSettingsTypoTolerance, bool)`

GetTypoToleranceOk returns a tuple with the TypoTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypoTolerance

`func (o *SearchSettings) SetTypoTolerance(v SearchSettingsTypoTolerance)`

SetTypoTolerance sets TypoTolerance field to given value.

### HasTypoTolerance

`func (o *SearchSettings) HasTypoTolerance() bool`

HasTypoTolerance returns a boolean if a field has been set.

### GetFaceting

`func (o *SearchSettings) GetFaceting() SearchSettingsFaceting`

GetFaceting returns the Faceting field if non-nil, zero value otherwise.

### GetFacetingOk

`func (o *SearchSettings) GetFacetingOk() (*SearchSettingsFaceting, bool)`

GetFacetingOk returns a tuple with the Faceting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceting

`func (o *SearchSettings) SetFaceting(v SearchSettingsFaceting)`

SetFaceting sets Faceting field to given value.

### HasFaceting

`func (o *SearchSettings) HasFaceting() bool`

HasFaceting returns a boolean if a field has been set.

### GetPagination

`func (o *SearchSettings) GetPagination() SearchSettingsPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SearchSettings) GetPaginationOk() (*SearchSettingsPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SearchSettings) SetPagination(v SearchSettingsPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *SearchSettings) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetEmbedders

`func (o *SearchSettings) GetEmbedders() map[string]map[string]interface{}`

GetEmbedders returns the Embedders field if non-nil, zero value otherwise.

### GetEmbeddersOk

`func (o *SearchSettings) GetEmbeddersOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddersOk returns a tuple with the Embedders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedders

`func (o *SearchSettings) SetEmbedders(v map[string]map[string]interface{})`

SetEmbedders sets Embedders field to given value.

### HasEmbedders

`func (o *SearchSettings) HasEmbedders() bool`

HasEmbedders returns a boolean if a field has been set.

### GetSearchCutoffMs

`func (o *SearchSettings) GetSearchCutoffMs() int32`

GetSearchCutoffMs returns the SearchCutoffMs field if non-nil, zero value otherwise.

### GetSearchCutoffMsOk

`func (o *SearchSettings) GetSearchCutoffMsOk() (*int32, bool)`

GetSearchCutoffMsOk returns a tuple with the SearchCutoffMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCutoffMs

`func (o *SearchSettings) SetSearchCutoffMs(v int32)`

SetSearchCutoffMs sets SearchCutoffMs field to given value.

### HasSearchCutoffMs

`func (o *SearchSettings) HasSearchCutoffMs() bool`

HasSearchCutoffMs returns a boolean if a field has been set.

### GetLocalizedAttributes

`func (o *SearchSettings) GetLocalizedAttributes() []SearchSettingsLocalizedAttributesInner`

GetLocalizedAttributes returns the LocalizedAttributes field if non-nil, zero value otherwise.

### GetLocalizedAttributesOk

`func (o *SearchSettings) GetLocalizedAttributesOk() (*[]SearchSettingsLocalizedAttributesInner, bool)`

GetLocalizedAttributesOk returns a tuple with the LocalizedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedAttributes

`func (o *SearchSettings) SetLocalizedAttributes(v []SearchSettingsLocalizedAttributesInner)`

SetLocalizedAttributes sets LocalizedAttributes field to given value.

### HasLocalizedAttributes

`func (o *SearchSettings) HasLocalizedAttributes() bool`

HasLocalizedAttributes returns a boolean if a field has been set.

### GetPrefixSearch

`func (o *SearchSettings) GetPrefixSearch() string`

GetPrefixSearch returns the PrefixSearch field if non-nil, zero value otherwise.

### GetPrefixSearchOk

`func (o *SearchSettings) GetPrefixSearchOk() (*string, bool)`

GetPrefixSearchOk returns a tuple with the PrefixSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSearch

`func (o *SearchSettings) SetPrefixSearch(v string)`

SetPrefixSearch sets PrefixSearch field to given value.

### HasPrefixSearch

`func (o *SearchSettings) HasPrefixSearch() bool`

HasPrefixSearch returns a boolean if a field has been set.

### GetFacetSearch

`func (o *SearchSettings) GetFacetSearch() bool`

GetFacetSearch returns the FacetSearch field if non-nil, zero value otherwise.

### GetFacetSearchOk

`func (o *SearchSettings) GetFacetSearchOk() (*bool, bool)`

GetFacetSearchOk returns a tuple with the FacetSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetSearch

`func (o *SearchSettings) SetFacetSearch(v bool)`

SetFacetSearch sets FacetSearch field to given value.

### HasFacetSearch

`func (o *SearchSettings) HasFacetSearch() bool`

HasFacetSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



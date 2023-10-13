# WAFEvents200ResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCount** | Pointer to **int64** |  | [optional] 
**Top10Countries** | Pointer to [**[]WAFEvents200ResultsInnerTop10CountriesInner**](WAFEvents200ResultsInnerTop10CountriesInner.md) |  | [optional] 
**Top10Ips** | Pointer to [**[]WAFEvents200ResultsInnerTop10CountriesInner**](WAFEvents200ResultsInnerTop10CountriesInner.md) |  | [optional] 
**HitCount** | Pointer to **int64** |  | [optional] 
**RuleId** | Pointer to **int64** |  | [optional] 
**IpCount** | Pointer to **int64** |  | [optional] 
**MatchZone** | Pointer to **string** |  | [optional] 
**PathCount** | Pointer to **int64** |  | [optional] 
**MatchesOn** | Pointer to **string** |  | [optional] 
**RuleDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewWAFEvents200ResultsInner

`func NewWAFEvents200ResultsInner() *WAFEvents200ResultsInner`

NewWAFEvents200ResultsInner instantiates a new WAFEvents200ResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFEvents200ResultsInnerWithDefaults

`func NewWAFEvents200ResultsInnerWithDefaults() *WAFEvents200ResultsInner`

NewWAFEvents200ResultsInnerWithDefaults instantiates a new WAFEvents200ResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCount

`func (o *WAFEvents200ResultsInner) GetCountryCount() int64`

GetCountryCount returns the CountryCount field if non-nil, zero value otherwise.

### GetCountryCountOk

`func (o *WAFEvents200ResultsInner) GetCountryCountOk() (*int64, bool)`

GetCountryCountOk returns a tuple with the CountryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCount

`func (o *WAFEvents200ResultsInner) SetCountryCount(v int64)`

SetCountryCount sets CountryCount field to given value.

### HasCountryCount

`func (o *WAFEvents200ResultsInner) HasCountryCount() bool`

HasCountryCount returns a boolean if a field has been set.

### GetTop10Countries

`func (o *WAFEvents200ResultsInner) GetTop10Countries() []WAFEvents200ResultsInnerTop10CountriesInner`

GetTop10Countries returns the Top10Countries field if non-nil, zero value otherwise.

### GetTop10CountriesOk

`func (o *WAFEvents200ResultsInner) GetTop10CountriesOk() (*[]WAFEvents200ResultsInnerTop10CountriesInner, bool)`

GetTop10CountriesOk returns a tuple with the Top10Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop10Countries

`func (o *WAFEvents200ResultsInner) SetTop10Countries(v []WAFEvents200ResultsInnerTop10CountriesInner)`

SetTop10Countries sets Top10Countries field to given value.

### HasTop10Countries

`func (o *WAFEvents200ResultsInner) HasTop10Countries() bool`

HasTop10Countries returns a boolean if a field has been set.

### GetTop10Ips

`func (o *WAFEvents200ResultsInner) GetTop10Ips() []WAFEvents200ResultsInnerTop10CountriesInner`

GetTop10Ips returns the Top10Ips field if non-nil, zero value otherwise.

### GetTop10IpsOk

`func (o *WAFEvents200ResultsInner) GetTop10IpsOk() (*[]WAFEvents200ResultsInnerTop10CountriesInner, bool)`

GetTop10IpsOk returns a tuple with the Top10Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop10Ips

`func (o *WAFEvents200ResultsInner) SetTop10Ips(v []WAFEvents200ResultsInnerTop10CountriesInner)`

SetTop10Ips sets Top10Ips field to given value.

### HasTop10Ips

`func (o *WAFEvents200ResultsInner) HasTop10Ips() bool`

HasTop10Ips returns a boolean if a field has been set.

### GetHitCount

`func (o *WAFEvents200ResultsInner) GetHitCount() int64`

GetHitCount returns the HitCount field if non-nil, zero value otherwise.

### GetHitCountOk

`func (o *WAFEvents200ResultsInner) GetHitCountOk() (*int64, bool)`

GetHitCountOk returns a tuple with the HitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitCount

`func (o *WAFEvents200ResultsInner) SetHitCount(v int64)`

SetHitCount sets HitCount field to given value.

### HasHitCount

`func (o *WAFEvents200ResultsInner) HasHitCount() bool`

HasHitCount returns a boolean if a field has been set.

### GetRuleId

`func (o *WAFEvents200ResultsInner) GetRuleId() int64`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *WAFEvents200ResultsInner) GetRuleIdOk() (*int64, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *WAFEvents200ResultsInner) SetRuleId(v int64)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *WAFEvents200ResultsInner) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetIpCount

`func (o *WAFEvents200ResultsInner) GetIpCount() int64`

GetIpCount returns the IpCount field if non-nil, zero value otherwise.

### GetIpCountOk

`func (o *WAFEvents200ResultsInner) GetIpCountOk() (*int64, bool)`

GetIpCountOk returns a tuple with the IpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpCount

`func (o *WAFEvents200ResultsInner) SetIpCount(v int64)`

SetIpCount sets IpCount field to given value.

### HasIpCount

`func (o *WAFEvents200ResultsInner) HasIpCount() bool`

HasIpCount returns a boolean if a field has been set.

### GetMatchZone

`func (o *WAFEvents200ResultsInner) GetMatchZone() string`

GetMatchZone returns the MatchZone field if non-nil, zero value otherwise.

### GetMatchZoneOk

`func (o *WAFEvents200ResultsInner) GetMatchZoneOk() (*string, bool)`

GetMatchZoneOk returns a tuple with the MatchZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchZone

`func (o *WAFEvents200ResultsInner) SetMatchZone(v string)`

SetMatchZone sets MatchZone field to given value.

### HasMatchZone

`func (o *WAFEvents200ResultsInner) HasMatchZone() bool`

HasMatchZone returns a boolean if a field has been set.

### GetPathCount

`func (o *WAFEvents200ResultsInner) GetPathCount() int64`

GetPathCount returns the PathCount field if non-nil, zero value otherwise.

### GetPathCountOk

`func (o *WAFEvents200ResultsInner) GetPathCountOk() (*int64, bool)`

GetPathCountOk returns a tuple with the PathCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathCount

`func (o *WAFEvents200ResultsInner) SetPathCount(v int64)`

SetPathCount sets PathCount field to given value.

### HasPathCount

`func (o *WAFEvents200ResultsInner) HasPathCount() bool`

HasPathCount returns a boolean if a field has been set.

### GetMatchesOn

`func (o *WAFEvents200ResultsInner) GetMatchesOn() string`

GetMatchesOn returns the MatchesOn field if non-nil, zero value otherwise.

### GetMatchesOnOk

`func (o *WAFEvents200ResultsInner) GetMatchesOnOk() (*string, bool)`

GetMatchesOnOk returns a tuple with the MatchesOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchesOn

`func (o *WAFEvents200ResultsInner) SetMatchesOn(v string)`

SetMatchesOn sets MatchesOn field to given value.

### HasMatchesOn

`func (o *WAFEvents200ResultsInner) HasMatchesOn() bool`

HasMatchesOn returns a boolean if a field has been set.

### GetRuleDescription

`func (o *WAFEvents200ResultsInner) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *WAFEvents200ResultsInner) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *WAFEvents200ResultsInner) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *WAFEvents200ResultsInner) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



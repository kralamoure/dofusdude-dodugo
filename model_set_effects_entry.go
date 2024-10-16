/*
dofusdude

# A project for you - the developer. The all-in-one toolbelt for your next Ankama related project.  ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [PHP](https://github.com/dofusdude/dofusdude-php) - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 0.9.3
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the SetEffectsEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetEffectsEntry{}

// SetEffectsEntry struct for SetEffectsEntry
type SetEffectsEntry struct {
	// minimum int value, can be a single if ignore_int_max and no ignore_int_min
	IntMinimum *int32 `json:"int_minimum,omitempty"`
	// maximum int value, if not ignore_int_max and not ignore_int_min, the effect has a range value
	IntMaximum *int32 `json:"int_maximum,omitempty"`
	Type *SetEffectsEntryType `json:"type,omitempty"`
	// ignore the int min field because the actual value is a string. For readability the templated field is the only important field in this case. 
	IgnoreIntMin *bool `json:"ignore_int_min,omitempty"`
	// ignore the int max field, if ignore_int_min is true, int min is a single value
	IgnoreIntMax *bool `json:"ignore_int_max,omitempty"`
	// all fields from above encoded in a single string
	Formatted *string `json:"formatted,omitempty"`
	// how many items it needs to trigger this effect with the given set
	ItemCombination *int32 `json:"item_combination,omitempty"`
}

// NewSetEffectsEntry instantiates a new SetEffectsEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetEffectsEntry() *SetEffectsEntry {
	this := SetEffectsEntry{}
	return &this
}

// NewSetEffectsEntryWithDefaults instantiates a new SetEffectsEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetEffectsEntryWithDefaults() *SetEffectsEntry {
	this := SetEffectsEntry{}
	return &this
}

// GetIntMinimum returns the IntMinimum field value if set, zero value otherwise.
func (o *SetEffectsEntry) GetIntMinimum() int32 {
	if o == nil || IsNil(o.IntMinimum) {
		var ret int32
		return ret
	}
	return *o.IntMinimum
}

// GetIntMinimumOk returns a tuple with the IntMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetEffectsEntry) GetIntMinimumOk() (*int32, bool) {
	if o == nil || IsNil(o.IntMinimum) {
		return nil, false
	}
	return o.IntMinimum, true
}

// HasIntMinimum returns a boolean if a field has been set.
func (o *SetEffectsEntry) HasIntMinimum() bool {
	if o != nil && !IsNil(o.IntMinimum) {
		return true
	}

	return false
}

// SetIntMinimum gets a reference to the given int32 and assigns it to the IntMinimum field.
func (o *SetEffectsEntry) SetIntMinimum(v int32) {
	o.IntMinimum = &v
}

// GetIntMaximum returns the IntMaximum field value if set, zero value otherwise.
func (o *SetEffectsEntry) GetIntMaximum() int32 {
	if o == nil || IsNil(o.IntMaximum) {
		var ret int32
		return ret
	}
	return *o.IntMaximum
}

// GetIntMaximumOk returns a tuple with the IntMaximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetEffectsEntry) GetIntMaximumOk() (*int32, bool) {
	if o == nil || IsNil(o.IntMaximum) {
		return nil, false
	}
	return o.IntMaximum, true
}

// HasIntMaximum returns a boolean if a field has been set.
func (o *SetEffectsEntry) HasIntMaximum() bool {
	if o != nil && !IsNil(o.IntMaximum) {
		return true
	}

	return false
}

// SetIntMaximum gets a reference to the given int32 and assigns it to the IntMaximum field.
func (o *SetEffectsEntry) SetIntMaximum(v int32) {
	o.IntMaximum = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SetEffectsEntry) GetType() SetEffectsEntryType {
	if o == nil || IsNil(o.Type) {
		var ret SetEffectsEntryType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetEffectsEntry) GetTypeOk() (*SetEffectsEntryType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SetEffectsEntry) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SetEffectsEntryType and assigns it to the Type field.
func (o *SetEffectsEntry) SetType(v SetEffectsEntryType) {
	o.Type = &v
}

// GetIgnoreIntMin returns the IgnoreIntMin field value if set, zero value otherwise.
func (o *SetEffectsEntry) GetIgnoreIntMin() bool {
	if o == nil || IsNil(o.IgnoreIntMin) {
		var ret bool
		return ret
	}
	return *o.IgnoreIntMin
}

// GetIgnoreIntMinOk returns a tuple with the IgnoreIntMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetEffectsEntry) GetIgnoreIntMinOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreIntMin) {
		return nil, false
	}
	return o.IgnoreIntMin, true
}

// HasIgnoreIntMin returns a boolean if a field has been set.
func (o *SetEffectsEntry) HasIgnoreIntMin() bool {
	if o != nil && !IsNil(o.IgnoreIntMin) {
		return true
	}

	return false
}

// SetIgnoreIntMin gets a reference to the given bool and assigns it to the IgnoreIntMin field.
func (o *SetEffectsEntry) SetIgnoreIntMin(v bool) {
	o.IgnoreIntMin = &v
}

// GetIgnoreIntMax returns the IgnoreIntMax field value if set, zero value otherwise.
func (o *SetEffectsEntry) GetIgnoreIntMax() bool {
	if o == nil || IsNil(o.IgnoreIntMax) {
		var ret bool
		return ret
	}
	return *o.IgnoreIntMax
}

// GetIgnoreIntMaxOk returns a tuple with the IgnoreIntMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetEffectsEntry) GetIgnoreIntMaxOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreIntMax) {
		return nil, false
	}
	return o.IgnoreIntMax, true
}

// HasIgnoreIntMax returns a boolean if a field has been set.
func (o *SetEffectsEntry) HasIgnoreIntMax() bool {
	if o != nil && !IsNil(o.IgnoreIntMax) {
		return true
	}

	return false
}

// SetIgnoreIntMax gets a reference to the given bool and assigns it to the IgnoreIntMax field.
func (o *SetEffectsEntry) SetIgnoreIntMax(v bool) {
	o.IgnoreIntMax = &v
}

// GetFormatted returns the Formatted field value if set, zero value otherwise.
func (o *SetEffectsEntry) GetFormatted() string {
	if o == nil || IsNil(o.Formatted) {
		var ret string
		return ret
	}
	return *o.Formatted
}

// GetFormattedOk returns a tuple with the Formatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetEffectsEntry) GetFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.Formatted) {
		return nil, false
	}
	return o.Formatted, true
}

// HasFormatted returns a boolean if a field has been set.
func (o *SetEffectsEntry) HasFormatted() bool {
	if o != nil && !IsNil(o.Formatted) {
		return true
	}

	return false
}

// SetFormatted gets a reference to the given string and assigns it to the Formatted field.
func (o *SetEffectsEntry) SetFormatted(v string) {
	o.Formatted = &v
}

// GetItemCombination returns the ItemCombination field value if set, zero value otherwise.
func (o *SetEffectsEntry) GetItemCombination() int32 {
	if o == nil || IsNil(o.ItemCombination) {
		var ret int32
		return ret
	}
	return *o.ItemCombination
}

// GetItemCombinationOk returns a tuple with the ItemCombination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetEffectsEntry) GetItemCombinationOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemCombination) {
		return nil, false
	}
	return o.ItemCombination, true
}

// HasItemCombination returns a boolean if a field has been set.
func (o *SetEffectsEntry) HasItemCombination() bool {
	if o != nil && !IsNil(o.ItemCombination) {
		return true
	}

	return false
}

// SetItemCombination gets a reference to the given int32 and assigns it to the ItemCombination field.
func (o *SetEffectsEntry) SetItemCombination(v int32) {
	o.ItemCombination = &v
}

func (o SetEffectsEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetEffectsEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IntMinimum) {
		toSerialize["int_minimum"] = o.IntMinimum
	}
	if !IsNil(o.IntMaximum) {
		toSerialize["int_maximum"] = o.IntMaximum
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IgnoreIntMin) {
		toSerialize["ignore_int_min"] = o.IgnoreIntMin
	}
	if !IsNil(o.IgnoreIntMax) {
		toSerialize["ignore_int_max"] = o.IgnoreIntMax
	}
	if !IsNil(o.Formatted) {
		toSerialize["formatted"] = o.Formatted
	}
	if !IsNil(o.ItemCombination) {
		toSerialize["item_combination"] = o.ItemCombination
	}
	return toSerialize, nil
}

type NullableSetEffectsEntry struct {
	value *SetEffectsEntry
	isSet bool
}

func (v NullableSetEffectsEntry) Get() *SetEffectsEntry {
	return v.value
}

func (v *NullableSetEffectsEntry) Set(val *SetEffectsEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableSetEffectsEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableSetEffectsEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetEffectsEntry(val *SetEffectsEntry) *NullableSetEffectsEntry {
	return &NullableSetEffectsEntry{value: val, isSet: true}
}

func (v NullableSetEffectsEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetEffectsEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



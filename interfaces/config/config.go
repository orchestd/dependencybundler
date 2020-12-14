package config

import (
	"time"
)

// Value is a convenient interface to cast value to different types
type Value interface {
	// IsSet will tell if this key really exists in the configuration
	IsSet() bool
	// Raw returns an interface{}. For a specific type value use a corresponding method.
	Raw() (interface{},error)
	// Bool returns the value associated with the key as a boolean.
	Bool() (bool,error)
	// Int returns the value associated with the key as an integer.
	Int() (int,error)
	// Int32 returns the value associated with the key as an integer.
	Int32() (int32,error)
	// Int64 returns the value associated with the key as an integer.
	Int64() (int64,error)
	// Uint returns the value associated with the key as an unsigned integer.
	Uint() (uint,error)
	// Uint32 returns the value associated with the key as an unsigned integer.
	Uint32() (uint32,error)
	// Uint64 returns the value associated with the key as an unsigned integer.
	Uint64() (uint64,error)
	// Float64 returns the value associated with the key as a float64.
	Float64() (float64,error)
	// Time returns the value associated with the key as time.
	Time() (time.Time,error)
	// Duration returns the value associated with the key as a duration.
	Duration() (time.Duration,error)
	// String returns the value associated with the key as a string.
	String() (string,error)
	// IntSlice returns the value associated with the key as a slice of int values.
	IntSlice() ([]int,error)
	// StringSlice returns the value associated with the key as a slice of strings.
	StringSlice() ([]string,error)
	// StringMap returns the value associated with the key as a map of interfaces.
	StringMap() (map[string]interface{},error)
	// StringMapString returns the value associated with the key as a map of strings.
	StringMapString() (map[string]string,error)
	// StringMapStringSlice returns the value associated with the key as a map to a slice of strings.
	StringMapStringSlice() (map[string][]string,error)
	// Unmarshal tries to unmarshal it to a 'result'. 'result' field must be a pointer.
	//
	// It heavy depends on what library is used to provide Config. For example Viper uses 'mapstructure' for that
	Unmarshal(result interface{}) error
}

// Config defines an interface to obtain configuration values from JSON/YAML/TOML or ENV
type Config interface {
	/*
		Get returns a Value associated with a given key, you can later cast this to a type.
		Examples:
		- Get an Int value
				numberOfPossibilities := config.Get("path.to.key").Int() // if key is absent this will return 0
		- It's possible to check if there is an actual value associated with this key
				initTime := time.Now() // default value
				if value := config.Get("path.to.key"); value.IsSet() {
					initTime = value.Time()
				}
	*/
	Get(key string) Value
	// Implementation returns the actual lib/struct that is responsible for the above logic
	Implementation() interface{}
}

// Builder defines configuration builder options
type Builder interface {
	// SetConfigFile tells builder where to look for file with the configuration map
	SetEnv(env string) Builder
	SetConfStruct(conf interface{}) Builder
	SetServiceName(name string) Builder
	SetRepo(resolver ConfParamsResolver) Builder
	Build() (Config, error)
}

type ConfParamsResolver interface {
	ResolveParams() ConfParams
}
type ConfParams map[string]interface{}

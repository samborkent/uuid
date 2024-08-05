# UUID

This package implements UUID version 4 and 7, as defined by [RFC9562](https://datatracker.ietf.org/doc/html/rfc9562).
It also includes a custom implementation of UUID version 8, which is reserved for vendor specific UUID implementations.
For all three UUID implementations special care was taken to maximize performance and get rid of any allocations.

The SetVersion function can be used to set the default version for newly generated UUIDs. Currently, the following version [4, 7, 8] are supported, with 7 as default.
A new UUID is generated with the **New** function.

```
// Set the version of newly generated UUIDs
uuid.SetVersion(7)

// Generate a new UUID v7
uuidV7 := uuid.New()
```

The version of the UUID can be extrated from the **Version** method.

```
version := uuidV7.Version()
```

For UUIDs with version 7 or 8 the creation time of the UUID can be extracted with the **CreationTime** method.

```
creationTime := uuidV7.CreationTime()
```

The hyphen-delimited string representation of the UUID can be returned with the **String** method:

```
uuidString := uuidV7.String()
```

Another utility function that is provided is the **Short** method, which returns the last 12 characters of the UUID string, which is useful for logging or tracing purposes.

```
uuidShort := uuidV7.Short()
```

There are also two validation methods **IsValid** and **IsValidString** provided to check whether a string or a byte array is a valid UUID, although only version 4, 7 and 8 are supported.

```
isValid := uuid.IsValid(byteArray)
isValidString := uuid.IsValidString(uuidString)
```

In case a byte slice or a string is a valid UUID v4, v7 or v8, it can be converted to a UUID with the **ToUUID** and **StringToUUID** functions. If the provided UUID is invalid, these functions will return an error.

```
uuid1, err := uuid.ToUUID(byteSlice)
uuid2, err := uuid.StringToUUID(uuidString)
```

## Version 4

This version is one of the most commonly used unique identifiers. In Go, [Google's implementation](https://github.com/google/uuid) is most commonly used.
The implementation in this package provides slightly better performance, and gets rid of any allocations.

```
BenchmarkUUIDV4Google-16                 2932860               405.1 ns/op            16 B/op          1 allocs/op
BenchmarkUUIDV4GoogleString-16          18487731               108.4 ns/op            48 B/op          1 allocs/op
BenchmarkUUIDV4-16                       3297364               366.3 ns/op             0 B/op          0 allocs/op
BenchmarkUUIDV4String-16                69654795                17.51 ns/op            0 B/op          0 allocs/op
```

## Version 7

This is the default version in this package. UUID v7 is newly defined in [RFC4122](https://datatracker.ietf.org/doc/html/draft-ietf-uuidrev-rfc4122bis-14) and should replace UUID v4 in most use-cases.
UUID v7 provides several benefits compared to UUID v4:

- It is lexicographically sortable.
- It can function as an sequential identifier, a unique identifier, and a timestamp at once. Which enables you to combine these three fields into one for databases.

```
BenchmarkUUIDV7-16                       2984664               410.6 ns/op             0 B/op          0 allocs/op
BenchmarkUUIDV7CreationTime-16          92355261                13.00 ns/op            0 B/op          0 allocs/op
BenchmarkUUIDV7Short-16                 419974106                2.777 ns/op           0 B/op          0 allocs/op
BenchmarkUUIDV7String-16                69436634                17.47 ns/op            0 B/op          0 allocs/op
BenchmarkUUIDV7Version-16               1000000000               0.2382 ns/op          0 B/op          0 allocs/op
```

## Version 8

UUID v4 and v7 require a crypographically secure pseudo-random number generator. In some cases a crypographically secure random number generator is not required and can be replaced for a faster less secure algorithm.
In this specific implementation of UUID v8, the **xoshiro256++** algorithm is used for random number generation, instead of the **crypto/rand** package in the Go standard library.
The scheme of this UUID v8 implementation is further simplified compared to the UUID v7 scheme by filling the first 64 bits with the Unix nanosecond timestamp instead of a Unix millisecond timestamp together with a sequence number. In this implementation the version bits simply override 4 bits of the Unix nanosecond timestamp. The resolution of the Unix nanosecond timestamp is OS specific, on some OS it is a single nanosecond, on Windows about 500 ns.
Use this implementation if you do not require a cryptographically secure UUID and if you need to generate UUIDs at a very high frequency with very little overhead.

```
BenchmarkUUIDV8-16                      20884886                56.36 ns/op            0 B/op          0 allocs/op
BenchmarkUUIDV8CreationTime-16          299049762                4.142 ns/op           0 B/op          0 allocs/op
BenchmarkUUIDV8Short-16                 429292722                2.775 ns/op           0 B/op          0 allocs/op
BenchmarkUUIDV8String-16                67370560                17.65 ns/op            0 B/op          0 allocs/op
BenchmarkUUIDV8Version-16               1000000000               0.2588 ns/op          0 B/op          0 allocs/op
```

From the benchmarks we can see that this UUID v8 implementation is about 7 times faster than both the UUID v4 and v7 implementations.

## Which version to use

- **UUID v4** - If you need the highest level of crypographical security, as this has the maximum number (122) of CSPRNG bits, for example for API keys.
- **UUID v7** - If you need need a unique identifier, while providing sequentiality, sortability, and being able to easily extract a timestamp, while still providing a high level of cryptographic security with 62 CSPRNG bits.
- **UUID v8** - Only use this version if you are certain you will continue using this specific implementation. Suitable for when unique identifiers need to be generated at very high frequency and cryptographic security is not a priority.

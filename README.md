# UUID [WIP]

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

// Or explicitely specify the version
uuidV4 := uuid.NewV4()
```

The hyphen-delimited string representation of the UUID can be returned with the **String** method:

```
uuidString := uuidV7.String()
```

Another utility function that is provided is the **Short** method, which returns the last 12 characters of the UUID string, which is useful for logging or tracing purposes.

```
uuidShort := uuidV7.Short()
```

The version of the UUID can be extrated from the **Version** method.

```
version := uuidV7.Version()
```

For UUIDs with version 7 or 8 the creation time of the UUID can be extracted with the **CreationTime** method.

```
creationTime := uuidV7.CreationTime()
```

The time ordered UUIDs version 7 and 8 can also use the comparision methods **Before** and **After** to do timewise comparisons:

```
uuid1 := uuid.NewV7()
time.Wait(time.Second)
uuid2 := uuid.NewV7()

if uuid1.Before(uuid2) {
    // This will evaluate true, as uuid2 was generated after uuid1.
}
```

There are also two validation methods **IsValid** and **IsValidString** provided to check whether a string or a byte array is a valid UUID, although only version 4, 7 and 8 are supported.

```
isValid := uuid.IsValid(byteArray)
isValidString := uuid.IsValidString(uuidString)
```

In case a byte slice or a string is a valid UUID v4, v7 or v8, it can be converted to a UUID with the **Parse** and **ParseString** functions. If the provided UUID is invalid, these functions will return an error.

```
uuid1, err := uuid.Parse(byteSlice)
uuid2, err := uuid.ParseString(uuidString)
```

Alternatively, the **FromBytes** and **FromString** functions do not return an error, but will return the a Nil UUID for invalid arguments

```
uuid1 := uuid.FromBytes(byteSlice)
uuid2 := uuid.FromString(uuidString)
```

## Version 4

This version is one of the most commonly used unique identifiers. In Go, [Google's implementation](https://github.com/google/uuid) is most commonly used.
The implementation in this package provides similar performance, but gets rid of any allocations.

```
BenchmarkUUIDV4Google-16                 3415486               349.5 ns/op            16 B/op          1 allocs/op
BenchmarkUUIDV4GoogleString-16          33279388                30.87 ns/op           48 B/op          1 allocs/op
BenchmarkUUIDV4-16                       3494116               348.1 ns/op             0 B/op          0 allocs/op
BenchmarkUUIDV4Short-16                 546104943                2.151 ns/op           0 B/op          0 allocs/op
BenchmarkUUIDV4String-16                78622393                14.04 ns/op            0 B/op          0 allocs/op
```

## Version 7

This is the default version in this package. UUID v7 is newly defined in [RFC9562](https://datatracker.ietf.org/doc/html/rfc9562).
UUID v7 provides several differences compared to UUID v4:

- It is lexicographically sortable.
- It can function as an sequential identifier, a unique identifier, and a timestamp at once. Which enables you to combine these three fields into one for databases.

```
BenchmarkUUIDV7-16                       3127352               383.9 ns/op             0 B/op          0 allocs/op
BenchmarkUUIDV7CreationTime-16          120616062               10.03 ns/op            0 B/op          0 allocs/op
BenchmarkUUIDV7Short-16                 542684994                2.159 ns/op           0 B/op          0 allocs/op
BenchmarkUUIDV7String-16                77577738                14.05 ns/op            0 B/op          0 allocs/op
BenchmarkUUIDV7Version-16               1000000000               0.2098 ns/op          0 B/op          0 allocs/op
```

## Version 8

UUID v4 and v7 require a crypographically secure pseudo-random number generator (CSPRNG). In some cases CSPRNG is not required and can be replaced for a faster, but less secure algorithm.
In this specific implementation of UUID v8, the **math/rand/v2** package is used for random number generation instead of the **crypto/rand** package.
The scheme of this UUID v8 implementation is further simplified compared to the UUID v7 scheme by filling the first 64 bits with the Unix nanosecond timestamp instead of a Unix millisecond timestamp together with a sequence number. In this implementation the version bits simply override 4 bits of the Unix nanosecond timestamp. The resolution of the Unix nanosecond timestamp is OS specific, on some OS it is a single nanosecond, on Windows about 500 ns.

```
BenchmarkUUIDV8-16                      25736830                44.56 ns/op            0 B/op          0 allocs/op
BenchmarkUUIDV8CreationTime-16          195271898                6.068 ns/op           0 B/op          0 allocs/op
BenchmarkUUIDV8Short-16                 554492515                2.145 ns/op           0 B/op          0 allocs/op
BenchmarkUUIDV8String-16                79959684                14.01 ns/op            0 B/op          0 allocs/op
BenchmarkUUIDV8Version-16               1000000000               0.2097 ns/op          0 B/op          0 allocs/op
```

From the benchmarks we can see that this UUID v8 implementation is about 7 times faster than both the UUID v4 and v7 implementations.

## Which version to use?

- **UUID v4** - If you need the highest level of crypographical security, for example for authentication token. It offers the maximum number (122) of CSPRNG bits.
- **UUID v7** - If you need need a unique identifier, while providing sequentiality, sortability, and being able to easily extract a timestamp, while still providing a high level of cryptographic security with 62 CSPRNG bits.
- **UUID v8** - Only use this version if you are certain you will continue using this specific implementation. Suitable for when unique identifiers need to be generated at very high frequency and cryptographic security is not a requirement.

## TODO

- Fix UUID v7 and v8 ordering.
- Tests, tests, tests!

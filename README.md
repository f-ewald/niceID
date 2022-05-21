# niceID
niceID creates numeric non-sequential ids for fast lookups in a url-safe base64 
encoding.

Sequential ids are necessary for efficient RDBMS lookups. Unfortunately, they 
pose a security risk because they are easy to guess. If endpoints are not
properly protected, they allow access to other users data.
They also allow deduction of the number of customers, users, etc. UUIDs solve
this problem and are completely random. Unfortunately they are not very well 
suited for database lookups and create fragmentation.

niceID solves this problem by creating a url-safe base64 representation of an id together
with a configurable key and an optional checksum. This key can be stored together with
the id as the primary key. This combination allows the creation of unguessable,
non-sequential numeric ids which yield a high database performance. Upon request,
the lookup is performed as normal with by primary key together with the validation
of the random key.

This allows the following scenario of a user with the id 1,000:
```
http://example.com/user/GmXkMO01pUFe
Key: GmXkMO01pU, ID: Fe (1,000)

Fast lookup via primary key:
SELECT * FROM user WHERE id = 1000 AND key = "GmXkMO01pU"
```

Example with default options (prefix length of 6, no checksum):
```
gen := NewPrefixGenerator()
gen.Encode(1337)
// "ABCDEFKv"
gen.Decode("ABCDEFKv")
// 1337, "ABCDEF"
```

With custom prefix length:
```
gen := NewPrefixGenerator(WithPrefix(0))
gen.Encode(1337)
// "Kv"
gen.Decode("Kv")
// 1337, ""
```

For further information see the unit tests or have a look at the source code. 
Pull requests are welcome.

## Testing
To test this library, execute the following:
```
go test -cover ./...
ok      github.com/f-ewald/niceID       0.124s  coverage: 82.2% of statements
```

## License
MIT
# `.hdata` serialization format

# hdata-encoder (Take-home: DB Integrations)

Implementation of a compact, deterministic encoder/decoder for the take-home specification.
No built-in encoding/decoding libraries used. Handles heterogeneous values. Round-trip guaranteed.

Spec summary:
- DataInput = array of string | int32 | DataInput
- Max array length: 1000
- Max string length: 1,000,000
- Strings are UTF-8
- Must provide encode and decode

---

## Format

Wire format is type-tagged and length-prefixed:

```<TYPE><LENGTH><PAYLOAD>```

Type tags:
- `0x01` → string
- `0x02` → int32
- `0x03` → nested DataInput (array)

Length rules:
- string → byte length (`u32`)
- int32 → always 4 bytes (`big-endian int32`)
- array → number of items (`u32`)

Decoding becomes deterministic and easy to walk.

---

## Example

```
data := DataInput{
 	spec.Str("foo"),
 	DataInput{
 		spec.Str("bar"),
 		spec.I32(42),
 	},
 }
```

Encode → bytes  
Decode → identical structure

---

## Complexity

### Time

Encoding:  
```O(N + total_string_bytes)```
Single linear traversal.

Decoding:  
```O(N + total_string_bytes)```
Sequential scan. No seeking or backtracking.

### Space
- Output ≈ input size + small per-element overhead  
- Decoder allocates only what’s required  
- No extra buffers beyond what’s needed

### Note on design choices

I considered offsets or an index table to allow partial or concurrent decoding.
But with nested arrays + length-prefixing, the format is inherently sequential.
Offsets require an initial pre-scan → double work → no real benefit for this spec.
Sequential decoding stays optimal and simple.

---

## How to run

Run individual test
```shell
go test ./...
```

Or run the cli command

```shell
❯ go run main.go bench -n 10000 -s 100 -i 10
Generating data (n=10000, maxStr=100)...
Generating data (n=10000, maxStr=100)...
Generating data (n=10000, maxStr=100)...
Generating data (n=10000, maxStr=100)...
Generating data (n=10000, maxStr=100)...
Generating data (n=10000, maxStr=100)...
Generating data (n=10000, maxStr=100)...
Generating data (n=10000, maxStr=100)...
Generating data (n=10000, maxStr=100)...
Generating data (n=10000, maxStr=100)...

[0] encode=165ms size=13.32MB decode=27ms
[1] encode=0ms size=0.00MB decode=0ms
[2] encode=151ms size=8.11MB decode=34ms
[3] encode=868ms size=106.24MB decode=230ms
[4] encode=40ms size=3.03MB decode=5ms
[5] encode=2485ms size=358.34MB decode=593ms
[6] encode=351ms size=46.93MB decode=80ms
[7] encode=141ms size=8.27MB decode=26ms
[8] encode=0ms size=0.00MB decode=0ms
[9] encode=242ms size=28.63MB decode=64ms
```
---

## Extending the format (adding more types)

To add a new type:
1. Assign a new 1-byte tag  
2. Add encode logic  
3. Add decode branch  
4. Define its length rule (fixed/variable)

Protocol stays forward-compatible since unknown tags can be skipped safely.

---

## Notes

- Strings stay raw UTF-8
- int32 uses big-endian for cross-language stability
- Nested DataInput uses the same logic recursively
- Invalid tags/lengths fail fast

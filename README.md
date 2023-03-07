# aes
Simple way to encrypt and decrypt a big file

## Test

`go build aes.go`

### For encrypt:

`./aes test.mp4 123 e`

### For decrypt:

`./aes test.mp4_ 123 d`


# paes
Simple way to encrypt and decrypt a big file in pipe style

## Test

`go build paes.go`

### For encrypt:

`cat test.mp4 | ./paes 123 e > test.mp4.crp`

### For decrypt:

`cat test.mp4.crp | ./paes 123 d > test.mp4`



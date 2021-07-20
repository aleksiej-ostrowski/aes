# aes
Simple way to encrypt and decrypt a big file

## Test

`go build aes.go`

### For encrypt:

`time ./aes test.mp4 123 e`

0m18.167s

### For decrypt:

`time ./aes test.mp4_ 123 d`

0m18.468s

## Verification

size(test.mp4) = 2619M
sha(123) = a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3

`time openssl enc -aes-256-ctr -in test.mp4 -out test.mp4.openssl -K a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3 -iv 00000000000000000000000000000000`

0m10.658s


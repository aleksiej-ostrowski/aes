# aes
Simple way to encrypt and decrypt a file

##Test

time ./aes test.mp4 123 e

0m18.167s

##Verification

sha(123) = a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3

time openssl enc -aes-256-ctr -in test.mp4 -out test.mp4.openssl -K a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3 -iv 00000000000000000000000000000000

0m10.658s


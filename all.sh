/usr/local/go/bin/go build aes.go
time ./aes test.mp4 123 e
time ./aes test.mp4_ 123 d
time openssl enc -aes-256-ctr -in test.mp4 -out test.mp4.openssl -K a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3 -iv 00000000000000000000000000000000
echo "It is comparing the original file and the file after encrypting and decrypting..."
md5sum test.mp4
md5sum test.mp4__
echo "It is verifying the encrypting algorithm..."
md5sum test.mp4_
md5sum test.mp4.openssl

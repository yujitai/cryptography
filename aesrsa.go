package main

import (
    "fmt"
    "time"
    "bytes"
    "crypto"
    "crypto/aes"
    "crypto/rsa"
    "crypto/rand"
    "encoding/hex"
    "crypto/cipher"
    "crypto/sha256"
)

var plaintext = []byte("hello taiyi")
var times = 1000

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

func AesEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    plaintext = PKCS7Padding(plaintext, blockSize)
    blockMode := cipher.NewCBCEncrypter(block, iv)
    crypted := make([]byte, len(plaintext))
    blockMode.CryptBlocks(crypted, plaintext)
    return crypted, nil
}

func AesDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
    origData := make([]byte, len(ciphertext))
    blockMode.CryptBlocks(origData, ciphertext)
    origData = PKCS7UnPadding(origData)
    return origData, nil
}

func doAESSpeedTest() {
    key, _ := hex.DecodeString("6368616e676520746869732070617373")

    c := make([]byte, aes.BlockSize + len(plaintext))
    iv := c[:aes.BlockSize]

    t1 := time.Now().UnixNano() / 1e6
    for i := 0; i <= times; i++ {
        ciphertext, err := AesEncrypt(plaintext, key, iv)
        if err != nil {
            panic(err)
        }

        plaintext, err = AesDecrypt(ciphertext, key, iv)
        if err != nil {
            panic(err)
        }
    }
    t2 := time.Now().UnixNano() / 1e6

    fmt.Printf("AES 加解密, 明文: %v, 次数: %v 次, 耗时: %vms\n", string(plaintext), times, t2 - t1)
}

func doRSASpeedTest() {
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        panic(err)
    }
    publicKey := privateKey.PublicKey

    t1 := time.Now().UnixNano() / 1e6
    for i := 0; i <= times; i++ {
        encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, plaintext, nil)
        if err != nil {
            panic(err)
        }

        _, err = privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
        if err != nil {
            panic(err)
        }
    }
    t2 := time.Now().UnixNano() / 1e6

    fmt.Printf("RSA 加解密, 明文: %v, 次数: %v 次, 耗时: %vms\n", string(plaintext), times, t2 - t1)
}

func main() {
    doAESSpeedTest()
    doRSASpeedTest()
}


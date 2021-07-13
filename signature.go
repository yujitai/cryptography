package main

import (
    "crypto"
    "crypto/rand"
    "crypto/md5"
    "crypto/rsa"
    "fmt"
)

func sign() {

    // 生成RSA密钥对
    priv, _ := rsa.GenerateKey(rand.Reader, 2048)
    pub := &priv.PublicKey
	// fmt.Println("公钥: ", pub);
	// fmt.Println("私钥: ", priv);
	fmt.Println("RSA密钥长度: 2048位")
	fmt.Println("")

	// 明文消息
    plaintxt := []byte("hi, taiyi")
    fmt.Println("明文: ", string(plaintxt[:]))
    fmt.Println("明文长度: ", len(plaintxt))
	fmt.Println("")

    // 计算消息摘要
	// 摘要算法使用md5
    h := md5.New()
    h.Write(plaintxt)
    digest := h.Sum(nil)
    fmt.Println("摘要(MD5): ", digest)
    fmt.Println("摘要长度: ", len(digest))
	fmt.Println("")

    // RSA私钥签名
	// 使用RSA-PSS这种签名方式
    opts := rsa.PSSOptions{rsa.PSSSaltLengthAuto, crypto.MD5}
    signature, _ := rsa.SignPSS(rand.Reader, priv, crypto.MD5, digest, &opts)
    fmt.Println("数字签名: ", signature)
    fmt.Println("数字签名长度: ", len(signature))
	fmt.Println("")

    // RSA公钥验签
    err := rsa.VerifyPSS(pub, crypto.MD5, digest, signature, &opts)
    if err == nil {
        fmt.Println("公钥验签成功")
    }
}

func main() {
    sign()
}

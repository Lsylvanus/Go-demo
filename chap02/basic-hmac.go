package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	c := getSha256Code("test@example.com")
	fmt.Println(c)

	c = getHmacCode("test@example.com")
	fmt.Println(c)

	signature := "POST" + "\n" +
		"https://mws.amazonservices.com" + "\n" +
		"/" + "\n" +
		"AWSAccessKeyId=AKIAFJPPO5KLY6G4XO7Q"  +
		"&Action=ListOrders"  +
		"&MWSAuthToken=amzn.mws.4ea38b7b-f563-7709-4bae-87aeaEXAMPLE"  +
		"&SellerId=A2NEXAMPLETF53"  +
		"&SignatureMethod=HmacSHA256"  +
		"&SignatureVersion=2"  +
		"&Timestamp=2017-08-22T07:53:58.144Z"  + "&Version=2013-09-01"

	fmt.Println(signature)
	fmt.Println(ComputeHmac256(signature, ""))
}

func getHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func getSha256Code(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 如果messageMAC是message的合法HMAC标签，函数返回真
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	msg := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(msg)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func getSignatureKey(key, dateStamp, regionName, serviceName string) string {
	kDate := ComputeHmac256(dateStamp, "AWS4"+key)
	kRegion := ComputeHmac256(regionName, kDate)
	kService := ComputeHmac256(serviceName, kRegion)
	kSigning := ComputeHmac256("aws4_request", kService)
	return kSigning
}

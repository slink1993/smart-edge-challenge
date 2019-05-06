package codechal

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Takes in pointer to rsa.PrivateKey
// Returns a Base64 encoded string of the private rsa key
func EncodePrivateKey(key *rsa.PrivateKey) string {
	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	val := pem.EncodeToMemory(privateKey)

	return string(val)
}

// Takes in rsa.PublicKey type
// Returns a Base64 encoded string of the public rsa key
func EncodePublicKey(key rsa.PublicKey) string {
	asn1Bytes, err := asn1.Marshal(key)
	checkError(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	val := pem.EncodeToMemory(pemkey)

	return string(val)
}

// Takes in io.reader, rsa privatekey and input given up program execution
// Returns a Base64 encoded cryptographic signature of the input,
// Calculated using the private key and the SHA256 digest of the input
func GenerateSignature(reader io.Reader, privatekey *rsa.PrivateKey, input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	digest := hash.Sum(nil)
	signature, err := rsa.SignPKCS1v15(reader, privatekey, crypto.SHA256, digest)
	checkFileError(err)
	return base64.StdEncoding.EncodeToString(signature)
}

// Takes in a filename and the desired contents to be written.
// Writes a given string to a file on the filesystem
func WriteToFile(fileName, contents string) {
	file, err := os.Create(fileName)
	checkFileError(err)
	defer file.Close()

	file.Write([]byte(contents))
}

// Takes in a filename, then reads a file from the filesystem.
// Returns a string of the contents of the file
func ReadFile(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	checkFileError(err)
	return string(fileContents)
}

// Prints the expected output with appropriate newlines and tabs
func PrettyPrintJSON(input, signature, pubkey string) {
	type Values struct {
		Message   string `json:"message"`
		Signature string `json:"signature"`
		Pubkey    string `json:"pubkey"`
	}
	values := Values{
		Message:   input,
		Signature: signature,
		Pubkey:    pubkey,
	}
	output, err := json.MarshalIndent(values, "", "  ")
	checkError(err)
	fmt.Println(string(output))
}

// Checks if the length of the user input is 250 characters or less
// Returns bool if true or not
func IsValidInputLength(input string) bool {
	if len(input) > 250 {
		fmt.Println("Input is greater than 250 characters. Please provide input of 250 characters or less")
		return false
	}
	return true
}

// Takes input string first given from running the app(?)
// Write private rsa key, public rsa key, and cryptographic signature to file
func GenerateNewKeys(input string) (string, string) {
	reader := rand.Reader
	bitSize := 2048

	privateKey, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	publicKey := privateKey.PublicKey

	signature := GenerateSignature(reader, privateKey, input)

	WriteToFile(fmt.Sprintf("%s_signature.txt", input), signature)

	encodedPrivateKey := EncodePrivateKey(privateKey)
	encodedPublicKey := EncodePublicKey(publicKey)
	WriteToFile(fmt.Sprintf("%s_private.pem", input), encodedPrivateKey)
	WriteToFile(fmt.Sprintf("%s_public.pem", input), encodedPublicKey)

	return signature, encodedPublicKey
}

func checkError(err error) {
	if err != nil {
		fmt.Println("An error has occurred", err.Error())
		os.Exit(1)
	}
}

func checkFileError(err error) {
	if err != nil {
		log.Println("File does not exist")
	}
}

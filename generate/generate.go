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
	"path/filepath"
)

// Takes in a pointer to an rsa.PrivateKey.
// Returns a Base64 encoded string of the private rsa key.
func EncodePrivateKey(key *rsa.PrivateKey) string {
	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	val := pem.EncodeToMemory(privateKey)

	return string(val)
}

// Takes in rsa.PublicKey type.
// Returns a Base64 encoded string of the public rsa key.
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

// Takes in io.Reader, rsa.PrivateKey and input given up program execution.
// Returns a Base64 encoded cryptographic signature of the input,
// calculated using the private key and the SHA256 digest of the input.
func GenerateSignature(reader io.Reader, privateKey *rsa.PrivateKey, input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	digest := hash.Sum(nil)
	signature, err := rsa.SignPKCS1v15(reader, privateKey, crypto.SHA256, digest)
	checkError(err)
	return base64.StdEncoding.EncodeToString(signature)
}

// Takes in a filename, the directory of where to write the file, and the desired contents to be written.
// Writes a given string to a file on the filesystem.
func WriteToFile(fileName, contents, dir string) {
	file, err := os.Create(filepath.Join(dir, filepath.Base(fileName)))
	defer file.Close()
	checkError(err)
	file.Write([]byte(contents))
}

// Takes in a filename, the directory of where to write the file, then reads a file from the filesystem.
// Returns a string of the contents of the file.
func ReadFile(fileName, dir string) string {
	fileContents, _ := ioutil.ReadFile(filepath.Join(dir, filepath.Base(fileName)))
	return string(fileContents)
}

// Prints the expected output with appropriate newlines and tabs.
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

// Checks if the length of the user input is 250 characters or less.
// Returns bool if true or not.
func IsValidInputLength(input string) bool {
	if len(input) > 250 {
		log.Println("Input is greater than 250 characters. Please provide input of 250 characters or less")
		return false
	}
	return true
}

// Takes the input given from command line execution, along with the directory to save the files.
// Writes a private rsa key, public rsa key, and cryptographic signature to file.
func GenerateNewKeys(input, dir string) (string, string) {
	reader := rand.Reader
	bitSize := 2048

	privateKey, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	publicKey := privateKey.PublicKey

	signature := GenerateSignature(reader, privateKey, input)

	WriteToFile(fmt.Sprintf("%s_signature.txt", input), signature, dir)

	encodedPrivateKey := EncodePrivateKey(privateKey)
	encodedPublicKey := EncodePublicKey(publicKey)
	WriteToFile(fmt.Sprintf("%s_private.pem", input), encodedPrivateKey, dir)
	WriteToFile(fmt.Sprintf("%s_public.pem", input), encodedPublicKey, dir)

	return signature, encodedPublicKey
}

// Takes in error type, and if error is not nil,
// prints details of error and exits program
func checkError(err error) {
	if err != nil {
		log.Println("An error has occurred: ", err.Error())
		os.Exit(1)
	}
}

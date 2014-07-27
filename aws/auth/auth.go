package auth

import (
	"github.com/twhello/aws-to-go/aws/interfaces"
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	ALGORITHM                      = "AWS4-HMAC-SHA256"
	TERMINATOR                     = "aws4_request"
	MAX_EXPIRATION_TIME_IN_SECONDS = 60 * 60 * 24 * 7
	TIME_FORMAT                    = "20060102T150405Z" // Mon Jan 2 15:04:05 MST 2006
	DATE_FORMAT                    = "20060102"
)

/******************************************************************************
Auth Credentials object
*/
type Credentials struct {
	accessKeyId string
	secretKey   string
}

// Instantiate a new Credentials object.
func NewCredentials(accessKeyId, secretKey string) interfaces.IAWSCredentials {
	return &Credentials{accessKeyId, secretKey}
}

// Instantiate a new Credenials object from environmental properties AWS_ACCESS_KEY_ID and AWS_SECRET_KEY.
func NewEnvCredentials() (interfaces.IAWSCredentials, error) {
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_KEY")
	if accessKeyId == "" || secretKey == "" {
		return nil, errors.New("Environmental properties WS_ACCESS_KEY_ID and AWS_SECRET_KEY could not be found.")
	}
	return &Credentials{accessKeyId, secretKey}, nil
}

// Instantiate a new Credenials object from local PROPERTIES file.
func NewFileCredentials() (interfaces.IAWSCredentials, error) {
	// TODO Implement function
	return &Credentials{"", ""}, nil
}

func (c Credentials) AccessKeyId() string {
	return c.accessKeyId
}

func (c Credentials) SecretKey() string {
	return c.secretKey
}

/******************************************************************************
V4Signer object
*/
type V4Signer struct {
	Credentials interfaces.IAWSCredentials
	AWSService  interfaces.IAWSService
}

func (s V4Signer) Sign(req interfaces.IAWSRequest) {

	now := time.Now().UTC()
	timeFormatTime := now.Format(TIME_FORMAT)
	dateFormatTime := now.Format(DATE_FORMAT)
	req.Header().Add("X-Amz-Date", timeFormatTime)

	request := req.BuildRequest()
	canonicalRequest, signedHeaders, contentSha256, contentMd5, _ := createCanonicalRequest(request, timeFormatTime)
	//fmt.Printf("HTTP REQUEST: %+v \n", request)
	//fmt.Println("CanReq:", canonicalRequest)

	key := s.deriveSigningKey(dateFormatTime)
	scope := dateFormatTime + "/" + s.AWSService.RegionName() + "/" + s.AWSService.ServiceName() + "/" + TERMINATOR
	signature := signature(key, stringToSign(timeFormatTime, scope, canonicalRequest))

	header := req.Header()
	header.Add("Host", request.URL.Host)
	header.Add("Date", now.Format(time.RFC1123))
	//header.Add("Content-Length", fmt.Sprintf("%v", contentSize))
	header.Add("Content-Md5", contentMd5)
	header.Add("Authorization", ALGORITHM+" Credential="+s.Credentials.AccessKeyId()+"/"+scope+", SignedHeaders="+signedHeaders+", Signature="+signature)
	header.Add("X-Amz-Content-Sha256", contentSha256)
}

// HMAC(HMAC(HMAC(HMAC("AWS4" + kSecret,"20110909"),"us-east-1"),"iam"),"aws4_request")
func (s V4Signer) deriveSigningKey(dateFormat string) []byte {
	h := hmacHash([]byte("AWS4"+s.Credentials.SecretKey()), []byte(dateFormat))
	h = hmacHash(h, []byte(s.AWSService.RegionName()))
	h = hmacHash(h, []byte(s.AWSService.ServiceName()))
	h = hmacHash(h, []byte(TERMINATOR))
	return h
}

/******************************************************************************
Helper Functions
*/

// Build a CanonicalRequest.
func createCanonicalRequest(request *http.Request, date string) (string, string, string, string, int) {

	method := request.Method
	path := request.URL.Path
	query := request.URL.Query().Encode()
	payload, _ := ioutil.ReadAll(request.Body)
	request.Body.Close()
	//fmt.Printf("%s \n", payload)
	payloadHash := hash(payload)

	payloadSize := len(payload)
	payloadMd5 := md5.Sum(payload)
	payload64 := base64.StdEncoding.EncodeToString(payloadMd5[0:])

	signedHeaders := "content-length;content-md5;content-type;host"

	var b bytes.Buffer
	b.WriteString(method)
	b.WriteRune('\n')
	b.WriteString(path)
	if path == "" {
		b.WriteRune('/')
	}
	b.WriteRune('\n')
	b.WriteString(query)
	b.WriteRune('\n')

	b.WriteString("content-length:")
	if method[0:1] == "P" { // GET, DELETE requests require empty val: "content-length:"
		b.WriteString(fmt.Sprintf("%v", payloadSize))
	}
	b.WriteRune('\n')
	b.WriteString("content-md5:")
	b.WriteString(payload64)
	b.WriteRune('\n')
	b.WriteString("content-type:")
	b.WriteString(request.Header.Get("content-type"))
	b.WriteRune('\n')
	b.WriteString("host:")
	b.WriteString(request.URL.Host)
	b.WriteRune('\n')

	var keys []string
	for key := range request.Header { // Keys must be in alpha order
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		key = strings.ToLower(key)
		if key[0:6] == "x-amz-" {
			signedHeaders += ";" + key
			b.WriteString(key)
			b.WriteRune(':')
			b.WriteString(request.Header.Get(key))
			b.WriteRune('\n')
		}
	}

	b.WriteRune('\n')
	b.WriteString(signedHeaders)
	b.WriteRune('\n')
	b.WriteString(payloadHash)

	return b.String(), signedHeaders, payloadHash, payload64, payloadSize
}

func signature(derivedKey []byte, sts string) string {
	h := hmacHash(derivedKey, []byte(sts))
	return fmt.Sprintf("%x", h)
}

func stringToSign(timeFormat string, scope, canonicalRequest string) string {
	var b bytes.Buffer
	b.WriteString(ALGORITHM)
	b.WriteRune('\n')
	b.WriteString(timeFormat)
	b.WriteRune('\n')
	b.WriteString(scope)
	b.WriteRune('\n')
	b.WriteString(hash([]byte(canonicalRequest)))
	return b.String()
}

func hash(in []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(in))
}

func hmacHash(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

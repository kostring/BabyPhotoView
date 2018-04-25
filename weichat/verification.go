package weichat

import (
	"fmt"
	"log"
	"net/url"
	"errors"
	"crypto/sha1"
	"sort"
	"bytes"
)

func processVerification(parameters url.Values) ([]byte, error) {
	if parameters["timestamp"] == nil || parameters["nonce"] == nil || parameters["signature"] == nil {
		log.Print("Not a verification request")
		return nil, errors.New("Not a verification request")
	}

	inputArgs := [3]string{token, parameters["timestamp"][0], parameters["nonce"][0]}
	argsSlice := inputArgs[:]

	sort.Strings(argsSlice)
	var buf bytes.Buffer
	for _, s := range argsSlice {
		buf.WriteString(s)
	}

	sha1Result := sha1.Sum(buf.Bytes())

	//Reuse buf to get the correct sha1 result string
	buf.Reset()

	for _, b := range sha1Result {
		buf.WriteString(fmt.Sprintf("%.2x", b))
	}

	sha1String := buf.String()

	if sha1String != parameters["signature"][0] {
		log.Printf("Invalid verification request! Signature %s, expect: %s", parameters["signature"][0], sha1String)
		return nil, errors.New("Invalid verification request!")
	}

	return []byte(parameters["echostr"][0]), nil
}

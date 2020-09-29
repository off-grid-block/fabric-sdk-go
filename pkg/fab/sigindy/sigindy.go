/* Package created to Sign using aca-py agent */

package sigindy

import (
	"bytes"
	b64 "encoding/base64"
	"errors"
	"os"

	"encoding/json"
	"net/http"
)

// IndySign will sign the given object using provided DID
func IndySign(digest []byte, did string) (signature map[string]interface{}, err error) {
	type IndyResponse struct {
		signature string
	}
	var result map[string]interface{}
	if len(digest) == 0 {
		return nil, errors.New("Empty digest bytes received while creating Indy signature")
	}
	if did == "" {
		return nil, errors.New("Empty did string received while creating Indy signature")
	}
	if len(did) != 22 {
		return nil, errors.New("DID size not equal to 22")
	}

	// Signing using the Client Agent by passing the proposal bytes and the signing_did
	url := os.Getenv("CLIENT_AGENT_URL") + "/sign_message"
	encoded := b64.StdEncoding.EncodeToString(digest)
	payload := []byte(`{"message":"` + encoded + `","signing_did":"` + did + `"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Add("content-type", "text/plain")
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		return nil, errors.New("Error connecting to Indy server")
	}
	json.NewDecoder(res.Body).Decode(&result)
	defer res.Body.Close()
	return result, nil
}

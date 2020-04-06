package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

func getSig(secret, data string) string {
	// Balance (POST https://api.bitkub.com/api/market/wallet)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sig := hex.EncodeToString(h.Sum(nil))
	payload := fmt.Sprintf(`%s, "sig":"%s"}`, data[:len(data)-1], sig)
	return payload
}

func getFloat(unk interface{}) (float64, error) {
	var floatType = reflect.TypeOf(float64(0))

	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}

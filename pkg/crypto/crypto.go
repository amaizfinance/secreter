// Copyright 2019 The Secreter Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crypto

import (
	"errors"
)

// Cipher suites used for encryption.
// Will always be the first byte of the encrypted message.
const (
	Curve25519Xchacha20poly1305 byte = iota
	GCPKMSXchacha20poly1305

	HeaderSize = 1
)

// Error definitions
var (
	ErrNoCipherSuites       = errors.New("crypto: no cipher suites found")
	ErrMultipleCipherSuites = errors.New("crypto: more than one provider per section")
)

// Encrypter is the interface that wraps the basic Encrypt method.
//
// Encrypt encrypts the plaintext message.
type Encrypter interface {
	Encrypt(plaintext []byte) ([]byte, error)
}

// Decrypter is the interface that wraps the basic Decrypt method.
//
// Decrypt decrypts the ciphertext.
type Decrypter interface {
	Decrypt(ciphertext []byte) ([]byte, error)
}

// EncryptDecrypter is the interface that groups the basic Encrypt and Decrypt methods.
type EncryptDecrypter interface {
	Encrypter
	Decrypter
}

// ConcatByteSlices concatenates multiple byte slices and returns the resulting slice
func ConcatByteSlices(slices ...[]byte) []byte {
	var capSlice int
	for i := range slices {
		capSlice += len(slices[i])
	}

	res := make([]byte, 0, capSlice)
	for i := range slices {
		res = append(res, slices[i]...)
	}

	return res
}

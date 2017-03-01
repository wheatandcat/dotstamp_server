package utils

import (
	"dotstamp_server/tests"

	. "gopkg.in/check.v1"
)

type TestEncrypt struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestEncrypt{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestEncrypt) TestSrringToEncryption(c *C) {

	r := SrringToEncryption("abc")

	c.Check(r, Equals, "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad")
}

func (t *TestEncrypt) TestIntToEncryption(c *C) {

	r := IntToEncryption(1)

	c.Check(r, Equals, "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b")
}

func (t *TestEncrypt) TestEncrypter(c *C) {

	r, _ := Encrypter([]byte("test@com"))

	c.Check(r, Not(Equals), "test@com")
}

func (t *TestEncrypt) TestDecrypter(c *C) {
	e, _ := Encrypter([]byte("test@com"))
	r, _ := Decrypter([]byte(e))

	c.Check(r, Equals, "test@com")
}

package utils

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"errors"
	"io"
	"io/ioutil"

	base "github.com/multiformats/go-multibase"
)

// CompressBytes func
func CompressBytes(bs []byte) (rs []byte, err error) {
	var bb bytes.Buffer
	w, err := zlib.NewWriterLevel(&bb, zlib.BestCompression)
	if err != nil {
		w = zlib.NewWriter(&bb)
	}
	if w == nil {
		return nil, errors.New("no memeory")
	}
	w.Write(bs)
	w.Close()
	rs = bb.Bytes()
	return rs, nil
}

// UncompressBytes func
func UncompressBytes(bs []byte) (rs []byte, err error) {
	var bb bytes.Buffer
	bb.Write(bs)
	r, err := zlib.NewReader(&bb)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for {
		buf := make([]byte, 4096)
		rc, _ := r.Read(buf)
		if rc > 0 {
			rs = append(rs, buf[:rc]...)
		}
		if rc != 4096 {
			break
		}
	}

	return rs, nil
}

func bytesZip(bs []byte) (z []byte, err error) {
	var b bytes.Buffer
	w, err := flate.NewWriter(&b, flate.BestCompression)
	if err != nil {
		return
	}
	_, err = w.Write(bs)
	w.Close()
	if err != nil {
		return
	}
	z = b.Bytes()
	return z, nil
}

func bytesUnzip(z []byte) (bs []byte, err error) {
	r := flate.NewReader(bytes.NewBuffer(z))
	defer r.Close()
	bs, err = ioutil.ReadAll(r)
	if err != nil && err != io.ErrUnexpectedEOF {
		return
	}
	return bs, nil
}

/*
func bytesZip(bs []byte) (z []byte, err error) {
	fmt.Println("src-len:", len(bs))
	var b bytes.Buffer
	w := lzma.NewWriterLevel(&b, lzma.BestCompression)
	_, err = w.Write(bs)
	w.Close()
	if err != nil {
		return
	}
	z = b.Bytes()
	fmt.Println("zip-len:", len(z))
	return z, nil
}

func bytesUnzip(z []byte) (bs []byte, err error) {
	r := lzma.NewReader(bytes.NewBuffer(z))
	defer r.Close()
	bs, err = ioutil.ReadAll(r)
	if err != nil && err != io.ErrUnexpectedEOF {
		return
	}
	return bs, nil
}
*/

/*
func bytesZip(bs []byte) (z []byte, err error) {
	fmt.Println("src-len:", len(bs))
	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		return
	}
	defer w.Close()
	_, err = w.Write(bs)
	if err != nil {
		return
	}
	w.Flush()
	z = b.Bytes()
	fmt.Println("zip-len:", len(z))
	return z, nil
}

func bytesUnzip(z []byte) (bs []byte, err error) {
	r, err := gzip.NewReader(bytes.NewBuffer(z))
	if err != nil {
		return
	}
	defer r.Close()
	bs, err = ioutil.ReadAll(r)
	if err != nil && err != io.ErrUnexpectedEOF {
		return
	}
	return bs, nil
}
*/

func bytesToBase(bs []byte) (s string, err error) {
	s, err = base.Encode(base.Base58BTC, bs)
	return s, err
}

func baseToBytes(s string) (bs []byte, err error) {
	_, bs, err = base.Decode(s)
	if err != nil {
		return
	}
	return bs, nil
}

package utils

import (
	"bytes"
	"compress/zlib"
	"errors"
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

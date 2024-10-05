package handler

import (
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"sync/atomic"
)

var currSeqNo uint32 = 0

func Handler(w http.ResponseWriter, r *http.Request) {
	seqNo := atomic.AddUint32(&currSeqNo, 1)
	dump, _ := httputil.DumpRequest(r, true)
	if hasReqBody := r.ContentLength > 0; hasReqBody {
		dump = append(dump, '\n')
	} else {
		dump = dump[:len(dump)-1]
	}

	w.Header().Set("Content-Type", "text/plain")

	output(w, seqNo, dump)
	output(os.Stdout, seqNo, dump)
}

func output(w io.Writer, seqNo uint32, dump []byte) {
	title := `
================================
Request ` + strconv.FormatUint(uint64(seqNo), 10) + `
================================

`

	w.Write([]byte(title))
	w.Write(dump)
}

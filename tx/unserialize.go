package tx

import (
	"bytes"
	"encoding/hex"

	"github.com/bcext/gcash/wire"
)

func UnserializeTx(rawtx string) (*wire.MsgTx, error) {
	b, err := hex.DecodeString(rawtx)
	if err != nil {
		return nil, err
	}

	var transaction wire.MsgTx
	buf := bytes.NewBuffer(b)
	err = transaction.Deserialize(buf)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

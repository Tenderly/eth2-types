package types

import (
	"fmt"

	fssz "github.com/ferranbt/fastssz"
)

var _ fssz.HashRoot = (Epoch)(0)
var _ fssz.Marshaler = (*Epoch)(nil)
var _ fssz.Unmarshaler = (*Epoch)(nil)

// Epoch represents a single epoch.
type Epoch uint64

// Mul multiplies epoch by x.
// In case of arithmetic issues (overflow/underflow/div by zero) panic is thrown.
func (e Epoch) Mul(x uint64) Epoch {
	res, err := e.SafeMul(x)
	if err != nil {
		panic(err.Error())
	}
	return res
}

// SafeMul multiplies epoch by x.
// In case of arithmetic issues (overflow/underflow/div by zero) error is returned.
func (e Epoch) SafeMul(x uint64) (Epoch, error) {
	res, err := Mul64(uint64(e), x)
	return Epoch(res), err
}

// Div divides epoch by x.
// In case of arithmetic issues (overflow/underflow/div by zero) panic is thrown.
func (e Epoch) Div(x uint64) Epoch {
	res, err := e.SafeDiv(x)
	if err != nil {
		panic(err.Error())
	}
	return res
}

// SafeDiv divides epoch by x.
// In case of arithmetic issues (overflow/underflow/div by zero) error is returned.
func (e Epoch) SafeDiv(x uint64) (Epoch, error) {
	res, err := Div64(uint64(e), x)
	return Epoch(res), err
}

// Add increases epoch by x.
// In case of arithmetic issues (overflow/underflow/div by zero) panic is thrown.
func (e Epoch) Add(x uint64) Epoch {
	res, err := e.SafeAdd(x)
	if err != nil {
		panic(err.Error())
	}
	return res
}

// SafeAdd increases epoch by x.
// In case of arithmetic issues (overflow/underflow/div by zero) error is returned.
func (e Epoch) SafeAdd(x uint64) (Epoch, error) {
	res, err := Add64(uint64(e), x)
	return Epoch(res), err
}

// AddEpoch increases epoch using another epoch value.
// In case of arithmetic issues (overflow/underflow/div by zero) panic is thrown.
func (e Epoch) AddEpoch(x Epoch) Epoch {
	return e.Add(uint64(x))
}

// SafeAddEpoch increases epoch using another epoch value.
// In case of arithmetic issues (overflow/underflow/div by zero) error is returned.
func (e Epoch) SafeAddEpoch(x Epoch) (Epoch, error) {
	return e.SafeAdd(uint64(x))
}

// Sub subtracts x from the epoch.
// In case of arithmetic issues (overflow/underflow/div by zero) panic is thrown.
func (e Epoch) Sub(x uint64) Epoch {
	res, err := e.SafeSub(x)
	if err != nil {
		panic(err.Error())
	}
	return res
}

// SafeSub subtracts x from the epoch.
// In case of arithmetic issues (overflow/underflow/div by zero) error is returned.
func (e Epoch) SafeSub(x uint64) (Epoch, error) {
	res, err := Sub64(uint64(e), x)
	return Epoch(res), err
}

// Mod returns result of `epoch % x`.
// In case of arithmetic issues (overflow/underflow/div by zero) panic is thrown.
func (e Epoch) Mod(x uint64) Epoch {
	res, err := e.SafeMod(x)
	if err != nil {
		panic(err.Error())
	}
	return res
}

// SafeMod returns result of `epoch % x`.
// In case of arithmetic issues (overflow/underflow/div by zero) error is returned.
func (e Epoch) SafeMod(x uint64) (Epoch, error) {
	res, err := Mod64(uint64(e), x)
	return Epoch(res), err
}

func (s Epoch) GetTree() (*fssz.Node, error) {
	panic("not implemented")
}

// HashTreeRoot returns calculated hash root.
func (e Epoch) HashTreeRoot() ([32]byte, error) {
	return fssz.HashWithDefaultHasher(e)
}

// HashWithDefaultHasher hashes a HashRoot object with a Hasher from the default HasherPool.
func (e Epoch) HashTreeRootWith(hh fssz.HashWalker) error {
	hh.PutUint64(uint64(e))
	return nil
}

// UnmarshalSSZ deserializes the provided bytes buffer into the epoch object.
func (e *Epoch) UnmarshalSSZ(buf []byte) error {
	if len(buf) != e.SizeSSZ() {
		return fmt.Errorf("expected buffer of length %d received %d", e.SizeSSZ(), len(buf))
	}
	*e = Epoch(fssz.UnmarshallUint64(buf))
	return nil
}

// MarshalSSZTo marshals epoch with the provided byte slice.
func (e *Epoch) MarshalSSZTo(dst []byte) ([]byte, error) {
	marshalled, err := e.MarshalSSZ()
	if err != nil {
		return nil, err
	}
	return append(dst, marshalled...), nil
}

// MarshalSSZ marshals epoch into a serialized object.
func (e *Epoch) MarshalSSZ() ([]byte, error) {
	marshalled := fssz.MarshalUint64([]byte{}, uint64(*e))
	return marshalled, nil
}

// SizeSSZ returns the size of the serialized object.
func (e *Epoch) SizeSSZ() int {
	return 8
}

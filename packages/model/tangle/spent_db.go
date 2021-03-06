package tangle

import (
	"encoding/binary"
	"io"

	"github.com/gohornet/hornet/packages/database"
	"github.com/iotaledger/iota.go/trinary"
	"github.com/pkg/errors"
)

var (
	spentAddressesDatabase database.Database
)

func configureSpentAddressesDatabase() {
	if db, err := database.Get(DBPrefixSpentAddresses); err != nil {
		panic(err)
	} else {
		spentAddressesDatabase = db
	}
}

func databaseKeyForAddress(address trinary.Hash) []byte {
	return trinary.MustTrytesToBytes(address)
}

func spentDatabaseContainsAddress(address trinary.Hash) (bool, error) {
	if contains, err := spentAddressesDatabase.Contains(databaseKeyForAddress(address)); err != nil {
		return contains, errors.Wrap(NewDatabaseError(err), "failed to check if the address exists in the spent addresses database")
	} else {
		return contains, nil
	}
}

func storeSpentAddressesInDatabase(spent []trinary.Hash) error {

	var entries []database.Entry

	for _, address := range spent {
		key := databaseKeyForAddress(address)

		entries = append(entries, database.Entry{
			Key:   key,
			Value: []byte{},
		})
	}

	// Now batch insert/delete all entries
	if err := spentAddressesDatabase.Apply(entries, []database.Key{}); err != nil {
		return errors.Wrap(NewDatabaseError(err), "failed to mark addresses as spent")
	}

	return nil
}

func StoreSpentAddressesBytesInDatabase(spentInBytes [][]byte) error {

	var entries []database.Entry

	for _, addressInBytes := range spentInBytes {
		key := addressInBytes

		entries = append(entries, database.Entry{
			Key:   key,
			Value: []byte{},
		})
	}

	// Now batch insert/delete all entries
	if err := spentAddressesDatabase.Apply(entries, []database.Key{}); err != nil {
		return errors.Wrap(NewDatabaseError(err), "failed to mark addresses as spent")
	}

	return nil
}

func StreamSpentAddressesToWriter(buf io.Writer) error {

	err := spentAddressesDatabase.StreamForEachKeyOnly(func(entry database.KeyOnlyEntry) error {
		return binary.Write(buf, binary.BigEndian, entry.Key)
	})

	if err != nil {
		return errors.Wrap(NewDatabaseError(err), "failed to stream spent addresses from database")
	}
	return nil
}

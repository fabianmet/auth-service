//package inmemory contains an inmemory store for our usecase.
// ONLY TO BE USED IN DEVELOPMENT!!
package inmemory

import "errors"

var errInMemoryKeyNotFound error = errors.New("repository: key not found")
var errInMemoryDuplicateKeyFound error = errors.New("repository: duplicate key found")
var errInMemoryKeyAlreadyExists error = errors.New("repository: key already exists")

//InMemoryClient is an in memory client that stores objects in memory.
type InMemoryClient struct {
	contents []inMemoryObject
}

//inMemoryObject contains keys and byteArrays that we use to store our data.
type inMemoryObject struct {
	key       string
	byteArray []byte
}

//NewInMemoryClient returns an empty InMemoryClient used for CRUD on keys.
func NewInMemoryClient() *InMemoryClient {
	return &InMemoryClient{
		contents: []inMemoryObject{},
	}
}

//Read reads from the inmemory client and returns the found key
func (i *InMemoryClient) Read(s string) ([]byte, error) {
	ind, err := i.findInSlice(s)
	if err != nil {
		return nil, err
	}
	return i.contents[ind].byteArray, nil
}

//Delete removes from inmemory store. Its not the nicest but hey its only for development cases! DO NOT USE IN PRODUCTION
func (i *InMemoryClient) Delete(s string) error {

	ind, err := i.findInSlice(s)
	if err != nil {
		return err
	}

	// this deletes an item from a slice, at the index.
	i.contents[ind] = i.contents[len(i.contents)-1]
	i.contents[len(i.contents)-1] = inMemoryObject{}
	i.contents = i.contents[:len(i.contents)-1]

	ind, err = i.findInSlice(s)
	if errors.Is(err, errInMemoryKeyNotFound) {
		return nil
	}
	return errInMemoryDuplicateKeyFound
}

func (i *InMemoryClient) Create(s string, b []byte) error {
	content := inMemoryObject{
		key:       s,
		byteArray: b,
	}
	_, err := i.Read(content.key)
	if errors.Is(err, errInMemoryKeyNotFound) {
		i.contents = append(i.contents, content)
		return nil
	}
	return errInMemoryKeyAlreadyExists
}

func (i *InMemoryClient) UpdateKey(s string, b []byte) error {

	ind, err := i.findInSlice(s)
	if err != nil {
		return err
	}

	content := inMemoryObject{
		key:       s,
		byteArray: b,
	}

	i.contents[ind] = content

	return nil
}

//findInSlice will search through the inmemory store for the object and if it is found return the index, OR error out with an index of -1
func (i *InMemoryClient) findInSlice(s string) (int, error) {
	for ind, v := range i.contents {
		if v.key == s {
			return ind, nil
		}
	}
	return -1, errInMemoryKeyNotFound
}

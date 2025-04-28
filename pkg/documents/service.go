package documents

import (
	"errors"
)

var (
	NotFoundErr = errors.New("not found")
)

type Store struct {
    documents []Document
}

func NewStore() *Store {
    return &Store{
        documents: []Document{},
    }
}

func (s *Store) Create(doc Document) error {
    if doc.Title == "" {
        return errors.New("document title is required")
    }
    s.documents = append(s.documents, doc)
    return nil
}

func (s *Store) GetAll() ([]Document, error){
	return s.documents, nil
}

func (s *Store) Get(id string) (Document, error) {
	for i := 0; i < len(s.documents); i++ {
		if s.documents[i].ID == id {
			return s.documents[i], nil
		}
	}
	return Document{}, NotFoundErr
}

func (s *Store) Update(id string, doc Document) error {
	for i := 0; i < len(s.documents); i++ {
		if s.documents[i].ID == id {
			s.documents[i] = doc
			return nil
		}
	}
	return NotFoundErr
}

func (s *Store) Delete(id string) error {
	for i := 0; i < len(s.documents); i++ {
		if s.documents[i].ID == id {
			// Remove the document at index i
			s.documents = append(s.documents[:i], s.documents[i+1:]...)
			return nil
		}
	}
	return NotFoundErr
}
package protein

import (
	"errors"
)

// ErrStop represents a STOP codon
var ErrStop error = errors.New("STOP codon")

// ErrInvalidBase represents an invalid base that cannot me mapped to an amino acid.
var ErrInvalidBase error = errors.New("Invalid base")

func FromRNA(rna string) ([]string, error) {
	panic("Please implement the FromRNA function")
}

func FromCodon(codon string) (string, error) {
	//panic("Please implement the FromCodon function")
	var res string
	var err error

	switch codon {
	case "AUG":
		res = "Methionine"
	case "UUU", "UUC":
		res = "Phenylalanine"
	case "UUA", "UUG":
		res = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		res = "Serine"
	case "UAU", "UAC":
		res = "Tyrosine"
	case "UGU", "UGC":
		res = "Cysteine"
	case "UGG":
		res = "Tryptophan"
	case "UAA", "UAG", "UGA":
		err = ErrStop
	default:
		err = ErrInvalidBase
	}

	return res, err
}

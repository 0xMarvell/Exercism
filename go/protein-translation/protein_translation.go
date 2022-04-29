// Package protein translates RNA sequences into proteins.
package protein

import (
	"errors"
)

// ErrStop represents a STOP codon
var ErrStop error = errors.New("STOP codon")

// ErrInvalidBase represents an invalid base that cannot me mapped to an amino acid.
var ErrInvalidBase error = errors.New("Invalid base")

// FromRNA translates an RNA-sequence into the proteins. If there is an invalid RNA-sequence
// an error is returned.
func FromRNA(rna string) ([]string, error) {
	var proteins []string

	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if err != nil {
			switch err {
			case ErrStop:
				return proteins, nil
			default:
				return nil, err
			}
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}

// FromCodon translates a 3-letter codon into an amino acid. An error is returned if
// the codon is a stop sequence or an invalid sequence.
func FromCodon(codon string) (string, error) {
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

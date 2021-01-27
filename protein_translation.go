// Package protein translates RNA sequences into proteins.
package protein

import (
	"errors"
)

// ErrStop represents a STOP codon
var ErrStop error = errors.New("STOP codon")

// ErrInvalidBase represents an invalid base that cannot me mapped to an amino acid.
var ErrInvalidBase error = errors.New("Invalid base")

// FromCodon translates a 3-letter codon into an amino acid. An error is returned if
// the codon is a stop sequence or an invalid sequence.
func FromCodon(codon string) (res string, err error) {

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

	return
}

// FromRNA translates an RNA-sequence into the proteins. If there is an invalid RNA-sequence
// an error is returned.
func FromRNA(rna string) (protein []string, err error) {

	if len(rna) == 0 {
		err = ErrInvalidBase
		return
	}

	var aminoAcid string
	for len(rna) > 0 {
		codon := rna[0:3]
		rna = rna[3:]
		aminoAcid, err = FromCodon(codon)
		if err != nil {
			if err == ErrStop {
				err = nil
			}
			break
		}
		protein = append(protein, aminoAcid)
	}
	return
}

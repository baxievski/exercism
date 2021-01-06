package protein

import "errors"

var (
	ErrStop        error = errors.New("Stop")
	ErrInvalidBase error = errors.New("Invalid codon")
)

// FromCodon translates a codon into a protein
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA translates an rna sequence of codons into a sequence of proteins
func FromRNA(rna string) ([]string, error) {
	proteins := []string{}
	for i := 0; i < len(rna); i += 3 {
		p, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			break
		}
		if err == ErrInvalidBase {
			return proteins, err
		}
		proteins = append(proteins, p)
	}
	return proteins, nil
}

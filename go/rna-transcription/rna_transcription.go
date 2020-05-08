// Package strand implements a solution for the exercism rna-transcription chalenge
package strand

// ToRNA transcripts a DNA string to RNA
func ToRNA(dna string) string {
	dnaToRna := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
	rna := []rune{}
	for _, nucleotide := range dna {
		rna = append(rna, dnaToRna[nucleotide])
	}
	return string(rna)
}

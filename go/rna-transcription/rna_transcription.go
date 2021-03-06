// Package starnd converts a DNA strand to an RNA strand
package strand

// ToRNA converts a DNA strand to its RNA equivalent
func ToRNA(dna string) string {
	rna := ""

	for _, nucleotide := range dna {
		switch {
		case nucleotide == 'G':
			rna += "C"
		case nucleotide == 'C':
			rna += "G"
		case nucleotide == 'T':
			rna += "A"
		case nucleotide == 'A':
			rna += "U"
		default:
			return ""
		}
	}
	return rna
}

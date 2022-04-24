package strand

func ToRNA(dna string) string {
	//panic("Please implement the ToRNA function")
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

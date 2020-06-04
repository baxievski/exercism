def to_rna(dna_strand):
    translate = {
        "G": "C",
        "C": "G",
        "T": "A",
        "A": "U",
    }

    result = []
    for nucleotide in dna_strand:
        if nucleotide not in translate:
            raise ValueError(f"'{nucleotide}' is not a nucleotide found in DNA")
        result.append(translate[nucleotide])

    return "".join(result)


if __name__ == "__main__":
    print(f"{to_rna('GCTA')=}")
    print(f"{to_rna('асдадад')=}")

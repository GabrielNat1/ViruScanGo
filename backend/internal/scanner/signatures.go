package scanner

type Signature struct {
	Name        string
	Pattern     []byte
	Offset      int64
	Description string
}

var DefaultSignatures = []Signature{
	{
		Name:        "TEST-MALWARE",
		Pattern:     []byte("X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"),
		Offset:      0,
		Description: "EICAR test signature",
	},
}

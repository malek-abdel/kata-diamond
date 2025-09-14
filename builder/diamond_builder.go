package builder

type DiamondBuilder interface {
	Print(letter string) (string, error)
}

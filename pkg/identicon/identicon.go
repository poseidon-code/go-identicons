package identicon

import (
	"fmt"
	"os"

	h "github.com/poseidon-code/go-identicons/pkg/hash"
	m "github.com/poseidon-code/go-identicons/pkg/matrix"
)

type Identicon struct {
    Options         Configuration
    Text            string
    Hash            string
    Width, Height   int
    Matrix          [][]int
}

func (i *Identicon) New() {
    // handling type (square|wide)
    if i.Options.Square {
        i.Hash, i.Width, i.Height = h.GenerateSHA256(i.Text)
    } else {
        i.Hash, i.Width, i.Height = h.GenerateSHA512(i.Text)
    }

    // handling size (4-8)
    if i.Options.Size<4 || i.Options.Size>8 {
        fmt.Println("Invalid size passed. \nSize must lie between 4 to 8 (inclusive).")
        os.Exit(1)
    }

    // handling vertical dimension (rather than rotating the entire martrix, only the dimensions are switched) (landscape|portrait)
    // handling cell filling (original|invert)
    // handling symmetric filling (asymmetric|symmetric)
    if i.Options.Vertical {
        if i.Options.Symmetric {
            i.Matrix = m.GenerateSymmetric(i.Hash, i.Options.Size, i.Height, i.Width, i.Options.Invert)
        } else {
            i.Matrix = m.Generate(i.Hash, i.Options.Size, i.Height, i.Width, i.Options.Invert)
        }
    } else {
        if i.Options.Symmetric {
            i.Matrix = m.GenerateSymmetric(i.Hash, i.Options.Size, i.Width, i.Height, i.Options.Invert)
        } else {
            i.Matrix = m.Generate(i.Hash, i.Options.Size, i.Width, i.Height, i.Options.Invert)
        }
    }
}
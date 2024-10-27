package histogram

import (
	"fmt"
)

type Histogram struct {
	Distribution   []int64
	Lines          int64
	Files          int64
	ProcessedFiles int64
	Directories    int64
}

func (h Histogram) String() string {
	return fmt.Sprintf("[ distr = %v , %d lines=%d, files=%d, processedFiles=%d, directories=%d]",
		h.Distribution,
		h.Lines,
		h.Lines,
		h.Files,
		h.ProcessedFiles,
		h.Directories)
}

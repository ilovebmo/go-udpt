package udpt

var MAGIC []byte = []byte{0, 0, 4, 23, 39, 16, 25, 128}

var (
	CONNECT  []byte = []byte{0, 0, 0, 0}
	ANNOUNCE []byte = []byte{0, 0, 0, 1}
	SCRAPE   []byte = []byte{0, 0, 0, 2}
	ERROR    []byte = []byte{0, 0, 0, 3}
)

var (
	NONE      []byte = []byte{0, 0, 0, 0}
	COMPLETED []byte = []byte{0, 0, 0, 1}
	STARTED   []byte = []byte{0, 0, 0, 2}
	STOPPED   []byte = []byte{0, 0, 0, 3}
)

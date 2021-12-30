package consts

const (
	LF                  = byte('\n')
	COMMA               = byte(',')
	K                   = 1024
	M                   = 1024 * K
	G                   = 1024 * M
	FileBufferSize      = 64 * K
	FileSortShardSize   = 256 * K
	FileMergeBufferSize = 16 * M
	InsertBatch         = 16 * K
	FileSortLimit       = 2
	SyncLimit           = 8
	PreparedBatch       = 8
)

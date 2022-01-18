package consts

const (
	LF                  = byte('\n')
	COMMA               = byte(',')
	K                   = 1024
	M                   = 1024 * K
	G                   = 1024 * M
	FileBufferSize      = 64 * K
	FileSortShardSize   = 1 * M
	FileMergeBufferSize = 32 * M
	InsertBatch         = 45 * K
	FileSortLimit       = 28
	SyncLimit           = 28
	PreparedBatch       = 3
)

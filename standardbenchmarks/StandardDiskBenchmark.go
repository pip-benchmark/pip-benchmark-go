package standardbenchmarks

import (
	"math"
	"math/rand"
	"os"
	"strconv"

	i64 "github.com/adam-lavrik/go-imath/i64"
	benchmark "github.com/pip-benchmark/pip-benchmark-go/benchmark"
)

type StandardDiskBenchmark struct {
	*benchmark.Benchmark
	nameText        string
	descriptionText string
	bufferSize      int
	chunkSize       int64
	fileSizeMax     int64

	fileName string
	fd       *os.File
	fileSize int64
}

func NewStandardDiskBenchmark() *StandardDiskBenchmark {
	c := StandardDiskBenchmark{
		nameText:        "Disk",
		descriptionText: "Measures disk read and write operations",
		bufferSize:      512,
		chunkSize:       1024000,
		fileSizeMax:     102400000,
		fileSize:        0,
	}
	c.Benchmark = benchmark.NewBenchmark("Disk", "Measures disk read and write operations", "StandardDiskBenchmark")
	c.Benchmark.IExecutable = &c
	return &c
}

func (c *StandardDiskBenchmark) SetUp() error {
	id := int64(math.Ceil(1000000.0 + rand.Float64()*9000000.0))
	c.fileName = "./DiskBenchmark-" + strconv.FormatInt(id, 10) + ".dat"

	file, opnErr := os.OpenFile(c.fileName, os.O_RDWR|os.O_CREATE, 0755)
	if opnErr != nil {
		return opnErr
	}
	c.fd = file
	return nil
}

func (c *StandardDiskBenchmark) Execute() error {
	if c.fd == nil {
		return nil
	}

	if c.fileSize == 0 || rand.Float32() < 0.5 {
		var position int64

		if c.fileSize < c.fileSizeMax {
			position = c.fileSize
		} else {
			position = int64(math.Ceil(rand.Float64() * float64(c.fileSize-c.chunkSize)))
		}

		sizeToWrite := c.chunkSize
		for sizeToWrite > 0 {
			length := i64.Min((int64)(c.bufferSize), sizeToWrite)

			buf := make([]byte, length)
			_, wrErr := c.fd.WriteAt(buf, position)

			if wrErr != nil {
				return wrErr
			}
			position += length
			c.fileSize = i64.Max((int64)(c.fileSize), position)
			sizeToWrite -= length
		}
	} else {
		position := int64(math.Ceil(rand.Float64() * float64(c.fileSize-c.chunkSize)))

		sizeToRead := c.chunkSize
		for sizeToRead > 0 {
			length := i64.Min(int64(c.bufferSize), sizeToRead)

			buf := make([]byte, length)
			_, rdErr := c.fd.ReadAt(buf, position)

			if rdErr != nil {
				return rdErr
			}

			position += length
			c.fileSize = i64.Max(int64(c.fileSize), position)
			sizeToRead -= length
		}
	}
	return nil
}

func (c *StandardDiskBenchmark) TearDown() error {

	if c.fd != nil {
		c.fd.Close()
		c.fd = nil
	}
	return nil
}

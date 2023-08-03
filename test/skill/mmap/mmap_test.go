package mmap

//"golang.org/x/exp/mmap"

// studio.test.mappedByteBuffer.MappedByteBufferTest
func main() {
	/*filePath := "example.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Make sure the file has a size greater than 0 before mapping it.
	file.Truncate(1024) // Set the file size to 1024 bytes (change as needed).

	// Memory-mapping the file.
	mmapFile, err := mmap.Map(file, mmap.RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer mmapFile.Unmap()

	// Now you can use mmapFile, which behaves similarly to MappedByteBuffer in Java.
	// For example, you can write and read data from it as follows:

	// Write data to the mmap.
	data := []byte("Hello, Golang MappedByteBuffer!")
	copy(mmapFile, data)

	// Read data from the mmap.
	readData := make([]byte, len(data))
	copy(readData, mmapFile)

	fmt.Printf("Data read from mmap: %s\n", readData)*/
}

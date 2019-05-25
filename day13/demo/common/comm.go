package common

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

type PkgHeader struct { // 包头包括协议头和数据长度，共六个字节
	HeaderFlag [2]byte
	DataLength uint32
}

const (
	HeaderLength = 6
)

const (
	SocketErrorClientClosed = "Client has been closed"
	SocketErrorServerClosed = "Server has been closed"
	SOCKET_ERROR_TIMEOUT    = "Timeout"
)

type SocketUtil struct { // 创建一个net链接对象
	Conn net.Conn
}

func (fd *SocketUtil) Init(conn net.Conn) {
	fd.Conn = conn
}

func (fd *SocketUtil) WritePkg(data []byte) (int, error) {
	if fd == nil {
		return -1, errors.New(SocketErrorServerClosed)
	}
	if len(data) == 0 {
		return 0, nil
	}
	buff := bytes.NewBuffer([]byte{})
	_ = binary.Write(buff, binary.BigEndian, []byte{0xaa, 0xbb})
	_ = binary.Write(buff, binary.BigEndian, uint32(len(data)))
	binary.Write(buff, binary.BigEndian, data)

	allBytes := buff.Bytes()

	return fd.writeNByte(allBytes)
}

func (fd *SocketUtil) ReadPkg() ([]byte, error) {
	if fd == nil || fd.Conn == nil {
		return nil, errors.New(SocketErrorServerClosed)
	}
	head, err := fd.readHead()
	if err != nil {
		return nil, err
	}
	if head.HeaderFlag != [2]byte{0xaa, 0xbb} {
		return nil, errors.New("Head package error")
	}
	datas, err := fd.readNByte(head.DataLength)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func (fd *SocketUtil) readHead() (*PkgHeader, error) {
	data, err := fd.readNByte(HeaderLength)
	if err != nil {
		return nil, err
	}
	h := PkgHeader{}
	buff := bytes.NewBuffer(data)
	binary.Read(buff, binary.BigEndian, &h.HeaderFlag)
	binary.Read(buff, binary.BigEndian, &h.DataLength)
	return &h, nil
}

func (fd *SocketUtil) readNByte(n uint32) ([]byte, error) {
	data := make([]byte, n)
	for x := 0; x < int(n); {
		length, err := fd.Conn.Read(data[x:])
		if length == 0 {
			return nil, errors.New(SocketErrorClientClosed)
		}
		if err != nil {
			return nil, err
		}
		x += length
	}
	return data, nil
}

func (fd *SocketUtil) writeNByte(data []byte) (int, error) {
	n, err := fd.Conn.Write(data)
	if err != nil {
		return -1, err
	} else {
		return n, nil
	}
}

func (fd *SocketUtil) Close() {
	fd.Conn.Close()
}

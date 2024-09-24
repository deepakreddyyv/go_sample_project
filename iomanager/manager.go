package iomanager 

type IoManager interface {
	Read() ([][]string, error) 
    Write(interface{}) error
}

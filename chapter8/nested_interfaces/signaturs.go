package main
  
import "fmt"
  
type Reader interface{
    read()
}
  
type Writer interface{
    write(string)
}
  
type ReaderWriter interface{
    Reader
    write(string) // неявно встраивается интерфейс Writer
}
  
type File struct{
    text string
}
  
func (f *File) read(){
    fmt.Println(f.text)
}
func (f *File) write(message string){
    f.text = message
    fmt.Println("Запись в файл строки", message)
}
  
func writeToStream(writer Writer, text string){
    writer.write(text)
}
 
func main() {
      
    myFile := &File{}
    writeToStream(myFile, "Hello METANIT.COM")
    myFile.read()
}
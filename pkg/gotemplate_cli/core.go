package gotemplate_cli
import "fmt"
type Engine struct{ready bool}
func NewEngine() *Engine { return &Engine{} }
func (e*Engine) Init(){ e.ready=true; fmt.Println("gotemplate-cli ready") }
func (e*Engine) Process(data string) string{ if !e.ready{ e.Init() }; return "processed: "+data }

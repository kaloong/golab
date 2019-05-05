package obj

type Obj struct {
	Name string
}

func ( o *Obj) Set( n string) {
	o.Name = n
}
func ( o *Obj) Get() string {
	return o.Name
}

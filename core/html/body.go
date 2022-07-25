package core_html

type Body struct{
	Value []byte
}

func (b *Body) ToString() string{
	return string(b.Value)
}
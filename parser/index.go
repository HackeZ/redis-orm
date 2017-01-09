package parser

import (
	"fmt"
	"strings"
)

type Index struct {
	Name       string
	Fields     []*Field
	FieldNames []string
	relation   *Relation
	Obj        *MetaObject
}

func NewIndex(obj *MetaObject) *Index {
	return &Index{Obj: obj}
}

func (idx *Index) LastField() *Field {
	return idx.Fields[len(idx.Fields)-1]
}

func (idx *Index) build(suffix string) error {
	idx.Name = fmt.Sprintf("%sOf%s%s", strings.Join(idx.FieldNames, ""), idx.Obj.Name, suffix)
	for _, name := range idx.FieldNames {
		f := idx.Obj.FieldByName(name)
		if f == nil {
			return fmt.Errorf("%s field not exist", name)
		}
		idx.Fields = append(idx.Fields, f)
	}

	return nil
}

func (idx *Index) GetRelation(storetype, valuetype, modeltype string) *Relation {
	if idx.relation == nil {
		idx.relation = NewRelation(idx.Obj)
	}
	idx.relation.Name = idx.Name + "Relation"
	idx.relation.StoreType = storetype
	idx.relation.ValueType = valuetype
	idx.relation.ModelType = modeltype
	idx.relation.build()
	return idx.relation
}

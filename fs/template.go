package fs

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/ezbuy/redis-orm/parser"
	"github.com/ezbuy/redis-orm/tpl"
)

var RedisOrmTemplate *template.Template

func generate_templates(obj *parser.MetaObject) []string {
	tpls := []string{}
	if obj.Relation != nil {
		tpls = append(tpls, "relation")
	} else {
		tpls = append(tpls, "object")
	}
	return tpls
}

func ExecuteMetaObjectTemplate(output string, obj *parser.MetaObject) error {
	for _, tpl := range generate_templates(obj) {
		filename := filepath.Join(output, strings.Join([]string{"gen", tpl, camel2sep(obj.Name, "."), "go"}, "."))
		fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer fd.Close()
		if err := RedisOrmTemplate.ExecuteTemplate(fd, tpl, obj); err != nil {
			return err
		}

		oscmd := exec.Command("gofmt", "-w", filename)
		oscmd.Run()
	}
	return nil
}

func ExecuteConfigTemplate(output, db string, packageName string) error {
	filename := filepath.Join(output, strings.Join([]string{"gen", "conf", db, "go"}, "."))
	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()
	if err := RedisOrmTemplate.ExecuteTemplate(fd, strings.Join([]string{"conf", db}, "."), map[string]interface{}{
		"GoPackage": packageName,
	}); err != nil {
		return err
	}

	oscmd := exec.Command("gofmt", "-w", filename)
	oscmd.Run()
	return nil
}

func init() {
	funcMap := template.FuncMap{
		"add":           Add,
		"sub":           Sub,
		"divide":        Divide,
		"multiply":      Multiply,
		"minus":         minus,
		"getNullType":   getNullType,
		"join":          strings.Join,
		"preSuffixJoin": preSuffixJoin,
		"repeatJoin":    repeatJoin,
		"camel2list":    camel2list,
		"strDefault":    strDefault,
		"strif":         strif,
		"camel2name":    parser.Camel2Name,
		"toids":         parser.ToIds,
	}

	RedisOrmTemplate = template.New("redis-orm").Funcs(funcMap)
	files := []string{
		"tpl/conf.elastic.gogo",
		"tpl/conf.mongo.gogo",
		"tpl/conf.mssql.gogo",
		"tpl/conf.mysql.gogo",
		"tpl/conf.orm.gogo",
		"tpl/conf.redis.gogo",
		"tpl/object.elastic.gogo",
		"tpl/object.functions.gogo",
		"tpl/object.gogo",
		"tpl/object.index.gogo",
		"tpl/object.mongo.gogo",
		"tpl/object.mssql.gogo",
		"tpl/object.mssql.read.gogo",
		"tpl/object.mssql.write.gogo",
		"tpl/object.mysql.gogo",
		"tpl/object.mysql.read.gogo",
		"tpl/object.mysql.write.gogo",
		"tpl/object.order.by.gogo",
		"tpl/object.range.gogo",
		"tpl/object.redis.gogo",
		"tpl/object.redis.read.gogo",
		"tpl/object.redis.sync.gogo",
		"tpl/object.redis.write.gogo",
		"tpl/object.relation.gogo",
		"tpl/object.unqiue.gogo",
		"tpl/relation.functions.gogo",
		"tpl/relation.geo.gogo",
		"tpl/relation.geo.sync.gogo",
		"tpl/relation.gogo",
		"tpl/relation.list.gogo",
		"tpl/relation.list.sync.gogo",
		"tpl/relation.pair.gogo",
		"tpl/relation.pair.sync.gogo",
		"tpl/relation.set.gogo",
		"tpl/relation.set.sync.gogo",
		"tpl/relation.zset.gogo",
		"tpl/relation.zset.sync.gogo",
	}
	for _, fname := range files {
		data, err := tpl.Asset(fname)
		if err != nil {
			panic(err)
		}
		_, err = RedisOrmTemplate.Parse(string(data))
		if err != nil {
			fmt.Println(fname)
			panic(err)
		}
	}
}

func minus(a, b int) int {
	return a - b
}

var NullTypes = map[string]string{
	"string":    "String",
	"bool":      "Bool",
	"int":       "Int64",
	"int32":     "Int64",
	"int64":     "Int64",
	"bit":       "Bool",
	"time.Time": "String",
	"float":     "Float64",
	"float32":   "Float64",
	"float64":   "Float64",
}

func camel2list(s []string) []string {
	s2 := make([]string, len(s))
	for idx := range s {
		s2[idx] = parser.Camel2Name(s[idx])
	}
	return s2
}

func strif(a bool, b, c string) string {
	if a {
		return b
	}
	return c
}

func strDefault(a, b string) string {
	if a == "" {
		return b
	}
	return a
}

func getNullType(gotype string) string {
	return NullTypes[gotype]
}

func preSuffixJoin(s []string, prefix, suffix, sep string) string {
	sNew := make([]string, 0, len(s))
	for _, each := range s {
		sNew = append(sNew, fmt.Sprintf("%s%s%s", prefix, each, suffix))
	}
	return strings.Join(sNew, sep)
}

func repeatJoin(n int, repeatStr, sep string) string {
	a := make([]string, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, repeatStr)
	}
	return strings.Join(a, sep)
}

func camel2sep(s string, sep string) string {
	nameBuf := bytes.NewBuffer(nil)
	for i := range s {
		n := rune(s[i]) // always ASCII?
		if unicode.IsUpper(n) {
			if i > 0 {
				nameBuf.WriteString(sep)
			}
			n = unicode.ToLower(n)
		}
		nameBuf.WriteRune(n)
	}
	return nameBuf.String()
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Divide(a, b int) int {
	return a / b
}

func Multiply(a, b int) int {
	return a * b
}

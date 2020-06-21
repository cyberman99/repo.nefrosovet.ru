package populator

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"

	randata "github.com/Pallinder/go-randomdata"
	"github.com/fatih/astrewrite"
	"time"
)

type (
	SupportedType string
	StringMeaning string

	ParsedStruct struct {
		Name     string
		Contents []*Content
	}

	Content struct {
		Tag   string
		Type  SupportedType
		Value interface{}
	}
)

const (
	randLimit = 500000

	String   SupportedType = "string"
	Uuid     SupportedType = "UUID"
	Boolean  SupportedType = "bool"
	Int32    SupportedType = "int32"
	Int64    SupportedType = "int64"
	Float64  SupportedType = "float64"
	Float32  SupportedType = "float32"
	NullTime SupportedType = "NullTime"
	Time     SupportedType = "Time"
	NullStr  SupportedType = "NullString"

	FirstName  StringMeaning = "first_name"
	LastName   StringMeaning = "last_name"
	Patronymic StringMeaning = "patronymic"
	StatusCode StringMeaning = "status_code"
)

var (
	now         = time.Now()
	statusCodes = []string{"PLANNED", "ONLINE"}

	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

	supportedTypes = map[SupportedType]func(string) interface{}{
		String:  func(s string) interface{} { return getStringMeaning(s) },
		Uuid:    func(s string) interface{} { return uuid.New() },
		Boolean: func(s string) interface{} { return randata.Boolean() },
		Int32:   func(s string) interface{} { return int32(randata.Number(randLimit)) },
		Int64:   func(s string) interface{} { return int64(randata.Number(randLimit)) },
		Float64: func(s string) interface{} { return float64(randata.Number(randLimit)) },
		Float32: func(s string) interface{} { return float32(randata.Number(randLimit)) },
		NullTime: func(s string) interface{} {
			return pq.NullTime{
				Time:  now.Add(time.Duration(randata.Number(300)) * time.Second),
				Valid: true,
			}
		},
		Time: func(s string) interface{} { return now },
		NullStr: func(s string) interface{} {
			return sql.NullString{
				String: randata.Adjective(),
				Valid:  true,
			}
		},
	}
)

func (ps *ParsedStruct) BuildInsertQuery() (query string, values []interface{}) {
	var (
		insertInto   = `INSERT INTO `
		columns      = "("
		placeHolders = " VALUES ("
	)

	phCount := 1
	commaCount := len(ps.Contents) - 1

	for _, c := range ps.Contents {
		if c.Tag == "end" { // reserved word
			c.Tag = `"end"`
		}
		if commaCount != 0 {
			columns += (c.Tag + ",")
			placeHolders += ("$" + strconv.Itoa(phCount) + ",")
			commaCount--
			phCount++
		} else {
			placeHolders += ("$" + strconv.Itoa(phCount) + ")")
			columns += (c.Tag + ")")
		}
		values = append(values, c.Value)
	}

	return (insertInto + toSnakeCase(ps.Name) + " " + columns + placeHolders + ";"), values
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func parseModels(path string) ([]*ParsedStruct, error) {
	parsed := make([]*ParsedStruct, 0)

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	file, err := parser.ParseFile(token.NewFileSet(), "", f, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	parseFunc := func(n ast.Node) (ast.Node, bool) {
		var ps = new(ParsedStruct)
		ps.Contents = make([]*Content, 0)
		x, ok := n.(*ast.TypeSpec)
		if !ok {
			return n, true
		}

		st, ok := x.Type.(*ast.StructType)
		if !ok {
			return n, true
		}
		if ps.Name != "" {
			return n, true
		}
		ps.Name = x.Name.Name

		for _, field := range st.Fields.List {
			tag := strings.Split(
				strings.SplitAfterN(
					field.Tag.Value, "\"", 2,
				)[1], "\"",
			)[0]

			simpleType, ok := field.Type.(*ast.Ident)
			if ok {
				ps.Contents = append(ps.Contents, &Content{
					Tag:   tag,
					Type:  SupportedType(simpleType.Name),
					Value: nil,
				})
			}
			complexType, ok := field.Type.(*ast.SelectorExpr)
			if ok {
				ps.Contents = append(ps.Contents, &Content{
					Tag:   tag,
					Type:  SupportedType(complexType.Sel.Name),
					Value: nil,
				})
			}
		}
		parsed = append(parsed, ps)
		return x, true
	}
	astrewrite.Walk(file, parseFunc)
	return parsed, nil
}

func buildFixtures(id int32, path string) ([]*ParsedStruct, error) {
	models, err := parseModels(path)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(models); i++ {
		setRandomValues(models[i])
		setIDs(id, models[i])
	}
	return models, nil
}

func setIDs(id int32, m *ParsedStruct) {
	for i := 0; i < len(m.Contents); i++ {
		if m.Contents[i].Tag == "id" {
			m.Contents[i].Value = id
		}
		if strings.Contains(m.Contents[i].Tag, "_id") {
			m.Contents[i].Value = id
		}
	}
}

func setRandomValues(m *ParsedStruct) {
	for i := 0; i < len(m.Contents); i++ {
		if m.Contents[i].Tag == "id" {
			continue
		}

		getTypeFn, ok := supportedTypes[m.Contents[i].Type]
		if !ok {
			log.Println("Unknown data type: ", m.Contents[i].Type)
			continue
		}
		m.Contents[i].Value = getTypeFn(m.Contents[i].Tag)
	}
}

func getStringMeaning(field string) string {
	gender := rand.Intn(2)
	switch StringMeaning(field) {
	case FirstName:
		return randata.FirstName(gender)
	case LastName:
		return randata.LastName()
	case Patronymic:
		return randata.FullName(gender)
	case StatusCode:
		return statusCodes[rand.Intn(2)]
	}
	return "Unknown"
}

func Populate(db *sql.DB, rowsNum int, modelsPath string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for id := 1; id < rowsNum; id++ {
		fixtures, err := buildFixtures(int32(id), dir+"/"+modelsPath)
		if err != nil {
			return err
		}

		fLen := len(fixtures)
		for i := 0; i < fLen; i++ {
			stmt, args := fixtures[i].BuildInsertQuery()
			_, err = db.Exec(stmt, args...)
			if err != nil {
				if pqErr, ok := err.(*pq.Error); ok {
					if pqErr.Code == "23503" { // foreign key violation
						fixtures = append(fixtures, fixtures[i]) // reorder for later processing
						fLen++
						log.Println("REQUEUE: ", stmt, args)
						continue
					}
				}
				return err
			}
			log.Println("PASS: ", stmt, args)
		}
	}
	return nil
}

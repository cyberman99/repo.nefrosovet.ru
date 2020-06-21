package populator

import (
	"database/sql"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"time"

	randata "github.com/Pallinder/go-randomdata"
	"github.com/fatih/astrewrite"
	pg "github.com/lfittl/pg_query_go"
	nodes "github.com/lfittl/pg_query_go/nodes"
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

	ParsedTables map[string][]tableField

	tableField struct {
		fieldName string
		checks    []string
		refTable  *string
		valueUUID uuid.UUID
	}
)

const (
	randLimit = 500000

	String        SupportedType = "string"
	Uuid          SupportedType = "UUID"
	Boolean       SupportedType = "bool"
	Int32         SupportedType = "int32"
	Int64         SupportedType = "int64"
	Float64       SupportedType = "float64"
	Float32       SupportedType = "float32"
	NullTime      SupportedType = "NullTime"
	Time          SupportedType = "Time"
	NullStr       SupportedType = "NullString"
	NullInteger32 SupportedType = "NullInt32"

	FirstName  StringMeaning = "first_name"
	LastName   StringMeaning = "last_name"
	Patronymic StringMeaning = "patronymic"
	Session    StringMeaning = "session"
)

var (
	now = time.Now()

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
		NullInteger32: func(s string) interface{} {
			return sql.NullInt32{
				Int32: int32(randata.Number(randLimit)),
				Valid: true,
			}
		},
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

func getStringMeaning(field string) string {
	gender := rand.Intn(2)
	switch StringMeaning(field) {
	case FirstName:
		return randata.FirstName(gender)
	case LastName:
		return randata.LastName()
	case Patronymic:
		return randata.FullName(gender)
	case Session:
		return randata.IpV6Address()
	}
	return "Unknown"
}

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

func getConstantRandomValue(tField tableField) *string {
	if len(tField.checks) > 0 {
		result := tField.checks[rand.Intn(len(tField.checks))]
		return &result

	}
	return nil
}

func setRandomValues(m *ParsedStruct, parsedTables *ParsedTables) error {
	scTable := toSnakeCase(m.Name)
	tableData, ok := (*parsedTables)[scTable]

	if !ok {
		return errors.New("parsedTable schama has not table: " + scTable)
	}

	for i := 0; i < len(m.Contents); i++ {
		// search this table field in schema
		var tField tableField
		found := false
		for _, mv := range tableData {
			if mv.fieldName == m.Contents[i].Tag {
				tField, found = mv, true
				break
			}
		}
		if !found {
			log.Println("cant find field `", m.Contents[i].Tag, "` in table schema: ", scTable)
			continue
		}
		if m.Contents[i].Tag == "id" {
			m.Contents[i].Value = tField.valueUUID
			continue
		}
		if tField.refTable != nil {
			var idUUID uuid.UUID
			refTableName := *tField.refTable
			refTableData, ok := (*parsedTables)[refTableName]
			if !ok {
				return errors.New("cant find reference table data in table schema: " + refTableName)
			}
			found := false
			for _, v := range refTableData {
				if v.fieldName == "id" {
					idUUID, found = v.valueUUID, true
					break
				}
			}
			if !found {
				return errors.New("cant find `id` field in refTable: " + refTableName)
			}
			m.Contents[i].Value = idUUID
			continue
		}

		rndValue := getConstantRandomValue(tField)
		if rndValue != nil {
			m.Contents[i].Value = *rndValue
			continue
		}
		getTypeFn, ok := supportedTypes[m.Contents[i].Type]
		if !ok {
			log.Println("Unknown data type: ", m.Contents[i].Type)
			continue
		}
		m.Contents[i].Value = getTypeFn(m.Contents[i].Tag)
	}
	return nil
}

// sort fixtures by schema
func sortFixturesBySchema(current []*ParsedStruct, keyMap []string) ([]*ParsedStruct, error) {

	var (
		result []*ParsedStruct
		count  int
	)
	for _, kv := range keyMap {
		for _, v := range current {
			if toSnakeCase(v.Name) == kv {
				result = append(result, v)
				count++
			}
		}
	}
	if count != len(keyMap) {
		return nil, errors.New("The keys in the map do not match the names of the tables in fixtures")
	}
	return result, nil

}

func fillNewUUIDValues(parsedTables *ParsedTables) {
	for i, v := range *parsedTables {
		newFillValues := []tableField{}
		for _, fv := range v {
			newValueUUID := uuid.New()
			fv.valueUUID = newValueUUID
			newFillValues = append(newFillValues, fv)
		}
		(*parsedTables)[i] = newFillValues
	}
}

func buildFixtures(path string, parsedTables *ParsedTables, tableOrder []string) ([]*ParsedStruct, error) {
	fillNewUUIDValues(parsedTables)
	models, err := parseModels(path)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(models); i++ {
		err := setRandomValues(models[i], parsedTables)
		if err != nil {
			return nil, err
		}
	}
	sortedModels, err := sortFixturesBySchema(models, tableOrder)
	if err != nil {
		log.Println(err)
		return models, nil
	}
	return sortedModels, nil
}

func getSchemaTree(filename string) (*pg.ParsetreeList, error) {
	blob, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	tree, err := pg.Parse(string(blob))
	if err != nil {
		return nil, err
	}
	return &tree, nil
}

func getTableOrder(tree pg.ParsetreeList) ([]string, error) {
	var orders []string
	for _, stmt := range tree.Statements {
		raw, ok := stmt.(nodes.RawStmt)
		if !ok {
			return nil, fmt.Errorf("expected RawStmt; got %T", stmt)
		}
		createNode, ok := raw.Stmt.(nodes.CreateStmt)
		if !ok {
			return nil, errors.New("expected createNode")
		}
		tableNameLink := createNode.Relation.Relname
		if tableNameLink == nil {
			return nil, errors.New("tableName null pointer")
		}
		tableName := *tableNameLink
		orders = append(orders, tableName)
	}
	return orders, nil
}

func getCheckValues(constraint nodes.Constraint) ([]string, error) {
	var checkValues []string
	boolExpr, ok := constraint.RawExpr.(nodes.BoolExpr)
	if !ok {
		return nil, errors.New("bad type, it must be BoolExpr")
	}
	boolExprItems := boolExpr.Args.Items
	for _, bv := range boolExprItems {
		boolExprItem, ok := bv.(nodes.A_Expr)
		if !ok {
			return nil, errors.New("bad type, it must be A_Expr")
		}
		itemValue, ok := boolExprItem.Rexpr.(nodes.A_Const)
		if !ok {
			return nil, errors.New("bad type, it must be A_Const")
		}
		itemValueStr, ok := itemValue.Val.(nodes.String)
		if !ok {
			return nil, errors.New("bad type, it must be nodes.String")
		}
		checkValue := itemValueStr.Str
		checkValues = append(checkValues, checkValue)
	}
	return checkValues, nil
}

func getTableField(column nodes.ColumnDef) (*tableField, error) {
	if column.Colname == nil {
		return nil, errors.New("columnName null pointer")
	}
	colName := *column.Colname // column name
	field := tableField{}
	field.fieldName = colName
	for _, iv := range column.Constraints.Items {
		constraint, ok := iv.(nodes.Constraint)
		if !ok {
			continue
		}
		// if it have not references to other tables
		// try to get check values of fields
		if constraint.Pktable == nil {
			// If it have check values
			if constraint.Contype == nodes.CONSTR_CHECK {
				checkVaues, err := getCheckValues(constraint)
				if err != nil {
					return nil, err
				}
				field.checks = checkVaues
			}
			continue
		}
		refTableLink := constraint.Pktable.Relname //
		if refTableLink == nil {
			return nil, errors.New("refTableLink null pointer")
		}
		refTableName := *refTableLink // name of reference table
		field.refTable = &refTableName
	}
	return &field, nil
}

func parseSQLSchema(tree pg.ParsetreeList) (*ParsedTables, error) {
	pTables := make(ParsedTables)
	for _, stmt := range tree.Statements {
		raw, ok := stmt.(nodes.RawStmt)
		if !ok {
			return nil, fmt.Errorf("expected RawStmt; got %T", stmt)
		}
		createNode, ok := raw.Stmt.(nodes.CreateStmt)
		if !ok {
			return nil, errors.New("expected createNode")
		}
		tableNameLink := createNode.Relation.Relname
		if tableNameLink == nil {
			return nil, errors.New("tableName null pointer")
		}
		tableName := *tableNameLink
		for _, elt := range createNode.TableElts.Items {
			// fill  the table creation queue
			columnDef, ok := elt.(nodes.ColumnDef)
			if !ok {
				continue
			}
			field, err := getTableField(columnDef)
			if err != nil {
				return nil, err
			}
			pTables[tableName] = append(pTables[tableName], *field)
		}

	}

	return &pTables, nil
}

func Populate(db *sql.DB, rowsNum int, modelsPath string, sqlSchemaFileName string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	tree, err := getSchemaTree(sqlSchemaFileName)
	if err != nil {
		return err
	}
	parsedTables, err := parseSQLSchema(*tree)
	if err != nil {
		return err
	}
	tableOrder, err := getTableOrder(*tree)
	if err != nil {
		return err
	}
	for id := 1; id < rowsNum; id++ {
		path := dir + "/" + modelsPath
		fixtures, err := buildFixtures(path, parsedTables, tableOrder)
		if err != nil {
			return err
		}
		fLen := len(fixtures)
		for i := 0; i < fLen; i++ {
			stmt, args := fixtures[i].BuildInsertQuery()
			_, err = db.Exec(stmt, args...)
			if err != nil {
				log.Println(err)
				return err
			}
			log.Println("PASS: ", stmt)
		}
	}
	return nil
}

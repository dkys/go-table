package table

import (
	"fmt"
	"github.com/dkys/go-table/border"
	"github.com/dkys/go-table/until"
	"github.com/gookit/color"
	"strings"
)

type Column struct {
	name   string
	width  int
	Length int
	color  string
}

/*type FieldBorder struct {
	Ico string
	Color color.Color
}*/

const (
	AlignContent = iota
	AlignLeft
	AlignRight
)

type Cell struct {
	Content string
	Length  int
}

type Table struct {
	Column      []Column
	Cell        [][]Cell
	Align       int
	Line        string
	Border      border.Border // 整个table的边框颜色
	Color       color.Color   // 表内容颜色
	FieldBorder border.Border // 表内容边框颜色
}

func Create(field ...string) *Table {
	col := make([]Column, 0)
	for _, column := range field {
		l := until.Length(column)
		col = append(col, Column{
			name:   column,
			width:  l,
			Length: l,
		})
	}
	return &Table{
		Column:      col,
		Cell:        make([][]Cell, 0),
		Align:       AlignContent,
		Border:      border.Border{},
		FieldBorder: border.Border{Ico: "|"},
	}
}

// SetAlign 设置对齐方式 默认居中
func (t *Table) SetAlign(val int) {
	t.Align = val
}

// Head 返回拼装好的表头
func (t *Table) Head() string {
	var line strings.Builder
	var dtStr strings.Builder
	line.WriteString("+")
	cell := make([]Cell, 0)
	for _, column := range t.Column {
		line.WriteString(strings.Repeat("-", column.width) + "+")
		cell = append(cell, Cell{Content: column.name, Length: column.Length})
	}
	line.WriteString("\n")
	dtStr.WriteString(t.row(cell))
	t.Line = line.String()
	t.Line = t.Border.Color.Render(t.Line)

	dtStr.WriteString(t.Line)
	return fmt.Sprintf("%s%s", t.Line, dtStr.String())
}

// SetFieldBorderColor 数据边框颜色
func (t *Table) SetFieldBorderColor(color color.Color) {
	t.FieldBorder.Color = color
	t.FieldBorder.Ico = color.Render(t.FieldBorder.Ico)
}

// SetBorderColor 设置边框颜色
func (t *Table) SetBorderColor(color color.Color) {
	t.Border.Color = color
}

// Row 获取拼接后完整的行字符串
func (t *Table) row(row []Cell) string {
	var str strings.Builder
	if t.FieldBorder.Color == 0 {
		t.SetFieldBorderColor(t.Border.Color)
	}
	str.WriteString(t.FieldBorder.Ico)
	for i, field := range row {
		str.WriteString(t.align(t.Column[i].width, field.Length, field.Content))
		str.WriteString(t.FieldBorder.Ico)
	}
	str.WriteString("\n")
	return str.String()
}

// align 返回字段对齐后的字符串
func (t *Table) align(width int, cellLength int, content string) string {
	var dtStr strings.Builder
	var str string
	moreLength := width - cellLength
	switch t.Align {
	case AlignContent:
		str = alignCenter(content, moreLength)
	case AlignLeft:
		str = content + strings.Repeat(" ", moreLength)
	case AlignRight:
		str = strings.Repeat(" ", moreLength) + content
	}
	dtStr.WriteString(str)
	return dtStr.String()
}

// AlignCenter 居中对齐
func alignCenter(content string, moreLength int) string {
	var str strings.Builder
	var IsOdd bool
	if moreLength&1 == 1 {
		IsOdd = true
		moreLength--
	}
	both := moreLength / 2
	left := strings.Repeat(" ", both)
	str.WriteString(left)
	str.WriteString(content)
	if IsOdd {
		left += " "
	}
	str.WriteString(left)

	return str.String()
}

// Data 将数据本地化
func (t *Table) Data(data [][]interface{}) {
	for _, row := range data {
		e := t.AddRow(row)
		if e != nil {
			fmt.Printf("error : %s %+v\n", e.Error(), row)
		}
	}
}

func (t *Table) AddRow(row []interface{}) error {
	if len(row) != len(t.Column) {
		return fmt.Errorf("数据列与表头列不相符！")
	}
	cell := make([]Cell, 0)
	for i, field := range row {
		toStr := until.ToString(field)
		maxWidth := until.Length(toStr)
		cell = append(cell, Cell{Content: toStr, Length: maxWidth})
		if maxWidth > t.Column[i].width {
			t.Column[i].width = maxWidth
		}
	}
	t.Cell = append(t.Cell, cell)
	return nil
}

func (t *Table) String() string {
	head := t.Head()

	var dtStr strings.Builder
	dtStr.WriteString(head)
	for _, row := range t.Cell {
		dtStr.WriteString(t.row(row))
	}

	return fmt.Sprintf("%s%s\n", dtStr.String(), t.Line)
}

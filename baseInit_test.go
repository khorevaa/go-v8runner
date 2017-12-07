package v8runner

import (
	"math/rand"
	"os"
	"path"

	"github.com/khorevaa/go-v8runner/v8tools"
	. "gopkg.in/check.v1"
)

var pwd, _ = os.Getwd()

var _ = Suite(&ТестыНаСозданиеБазыДанных{})

type ТестыНаСозданиеБазыДанных struct {
	conf           *Конфигуратор
	КаталогБазы    string
	ПутьКШаблону   string
	ИмяБазыВСписке string
}

func (s *ТестыНаСозданиеБазыДанных) SetUpTest(c *C) {

	s.conf = НовыйКонфигуратор()
	s.ИмяБазыВСписке = v8tools.NewUID(rand.Intn(16))
	s.КаталогБазы = v8tools.ВременныйКаталогСПрефисом(s.ИмяБазыВСписке)
	s.ПутьКШаблону = path.Join(pwd, "tests", "fixtures", "ТестовыйФайлКонфигурации.cf")

}
func (s *ТестыНаСозданиеБазыДанных) TearDownSuite(c *C) {
	v8tools.ОчиститьВременныйКаталог()
}

func (s *ТестыНаСозданиеБазыДанных) TestКонфигуратор_СоздатьФайловуюБазуПоШаблону(c *C) {

	err := s.conf.СоздатьФайловуюБазуПоШаблону(s.КаталогБазы, s.ПутьКШаблону)

	c.Assert(err, IsNil, Commentf("Ошибка теста: %v", err))

}

func (s *ТестыНаСозданиеБазыДанных) TestКонфигуратор_СоздатьФайловуюБазуПоУмолчанию(c *C) {

	err := s.conf.СоздатьФайловуюБазуПоУмолчанию(s.КаталогБазы)

	c.Assert(err, IsNil, Commentf("Ошибка теста: %v", err))

}

func (s *ТестыНаСозданиеБазыДанных) TestКонфигуратор_СоздатьИменнуюФайловуюБазу(c *C) {

	c.Skip("-live not provided")

	err := s.conf.СоздатьИменнуюФайловуюБазу(s.КаталогБазы, s.ИмяБазыВСписке)

	c.Assert(err, IsNil, Commentf("Ошибка теста: %v", err))

}
func (s *ТестыНаСозданиеБазыДанных) TestКонфигуратор_СоздатьИменнуюФайловуюБазуПоШаблону(c *C) {

	c.Skip("-live not provided")

	err := s.conf.СоздатьИменнуюФайловуюБазуПоШаблону(s.КаталогБазы, s.ПутьКШаблону, s.ИмяБазыВСписке)

	c.Assert(err, IsNil, Commentf("Ошибка теста: %v", err))

}

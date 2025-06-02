package main

import (
    //"fmt"
    "sort"
    "strings"
    "testing"
    "encoding/json"
    "reflect"
)

// Вспомогательная функция для получения отсортированных ключей мапы
func getSortedKeys(m map[string][][]string) []string {
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    return keys
}

// --- Тест: Handle() корректно группирует строки ---
func Test_Handle_ParsesAndGroupsCorrectly(t *testing.T) {
    groups := &StringGrouper{
        Str:         "apple, Apricot, banana, Blueberry",
        MaxPerGroup: 2,
    }
    groups = groups.Handle()

    expectedKeys := []string{"a", "b"}
    keys := getSortedKeys(groups.GroupData)

    if len(keys) != len(expectedKeys) {
        t.Errorf("ожидается %d групп, получено %d", len(expectedKeys), len(keys))
    }

    for i := range expectedKeys {
        if keys[i] != expectedKeys[i] {
            t.Errorf("ожидается группа %q, получена %q", expectedKeys[i], keys[i])
        }
    }
}

// --- Тест: Строки с символами попадают в группу "?" ---
func Test_Handle_GroupsNonLettersToQuestionMark(t *testing.T) {
    groups := &StringGrouper{
        Str:         "$dollar, 123number, apple",
        MaxPerGroup: 2,
    }
    groups = groups.Handle()

    if _, exists := groups.GroupData["?"]; !exists {
        t.Errorf("ожидается группа \"?\", но её нет")
    }

    if len(groups.GroupData["?"][0]) != 2 {
        t.Errorf("ожидается 2 элемента в группе \"?\", получено %d", len(groups.GroupData["?"][0]))
    }
}

// --- Тест: Разбиение на подгруппы работает правильно ---
func Test_Handle_SplitsIntoSubgroups(t *testing.T) {
    groups := &StringGrouper{
        Str:         "apple, apricot, avocado, banana, berry",
        MaxPerGroup: 2,
    }
    groups = groups.Handle()

    aGroup := groups.GroupData["a"]
    if len(aGroup) != 2 || len(aGroup[0]) != 2 || len(aGroup[1]) != 1 {
        t.Errorf("группа 'a' должна быть разделена на 2 подгруппы размером 2 и 1, получено %+v", aGroup)
    }
}

// --- Тест: Формат текстового вывода ---
func Test_AddTextData_FormatsOutputCorrectly(t *testing.T) {
    groups := &StringGrouper{
        Str:         "apple, Apricot, avocado, banana, blueberry",
        MaxPerGroup: 2,
    }
    groups = groups.Handle()
    expected := `"a":
  [1] apple, Apricot
  [2] avocado
"b":
  [1] banana, blueberry`

    if groups.DataText != expected {
        t.Errorf("ожидаемый вывод:\n%s\nполучено:\n%s", expected, groups.DataText)
    }
}

// --- Тест: Пустой ввод должен вызвать ошибку ---
func Test_Parse_EmptyInputReturnsError(t *testing.T) {
    groups := &StringGrouper{
        Str:         "",
        MaxPerGroup: 2,
    }

    err := groups.Parse()
    if err == nil {
        t.Errorf("ожидается ошибка при пустом вводе")
    }
}

// --- Тест: MaxPerGroup <= 0 должно вызывать ошибку ---
func Test_Parse_InvalidMaxPerGroupReturnsError(t *testing.T) {
    groups := &StringGrouper{
        Str:         "apple, banana",
        MaxPerGroup: 0,
    }

    err := groups.Parse()
    if err == nil || !strings.Contains(err.Error(), "maxPerGroup должно быть > 0") {
        t.Errorf("ожидается ошибка при maxPerGroup <= 0")
    }
}

// --- Тест: Все строки пустые → ошибка ---
func Test_Parse_AllEmptyStringsReturnsError(t *testing.T) {
    groups := &StringGrouper{
        Str:         ", ,    ,,",
        MaxPerGroup: 2,
    }

    err := groups.Parse()
    if err == nil || !strings.Contains(err.Error(), "все строки оказались пустыми") {
        t.Errorf("ожидается ошибка, если все строки пустые")
    }
}

// --- Табличные тесты: разные входные данные ---
func Test_Handle_GroupingWithVariousInputs(t *testing.T) {
    tests := []struct {
        name        string
        input       string
        maxPerGroup int
        expected    map[string][][]string
    }{
        {
            name:        "Обычный случай",
            input:       "apple, Apricot, Avocado, banana, Blueberry",
            maxPerGroup: 2,
            expected: map[string][][]string{
                "a": {{"apple", "Apricot"}, {"Avocado"}},
                "b": {{"banana", "Blueberry"}},
            },
        },
        {
            name:        "Смешанные символы",
            input:       "$dollar, 123number, @at, apple",
            maxPerGroup: 3,
            expected: map[string][][]string{
                "?": {{ "$dollar", "123number", "@at"}},
                "a":   {{"apple"}},
            },
        },
        {
            name:        "Группировка по одной подгруппе",
            input:       "apple, banana, cherry",
            maxPerGroup: 5,
            expected: map[string][][]string{
                "a": {{ "apple" }},
                "b": {{ "banana" }},
                "c": {{ "cherry" }},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            groups := &StringGrouper{
                Str:         tt.input,
                MaxPerGroup: tt.maxPerGroup,
            }
            groups = groups.Handle()

            for key, expectedGroup := range tt.expected {
                actualGroup, exists := groups.GroupData[key]
                if !exists {
                    t.Errorf("ключ %q отсутствует в результатах", key)
                    continue
                }
                if len(actualGroup) != len(expectedGroup) {
                    t.Errorf("для ключа %q ожидается %d подгрупп, получено %d", key, len(expectedGroup), len(actualGroup))
                    continue
                }
                for i := range expectedGroup {
                    if len(actualGroup[i]) != len(expectedGroup[i]) {
                        t.Errorf("в подгруппе %d для %q ожидается %d элементов, получено %d", i, key, len(expectedGroup[i]), len(actualGroup[i]))
                        continue
                    }
                    for j := range expectedGroup[i] {
                        if actualGroup[i][j] != expectedGroup[i][j] {
                            t.Errorf("в группе %q, подгруппа %d, элемент %d: ожидается %q, получено %q", key, i, j, expectedGroup[i][j], actualGroup[i][j])
                        }
                    }
                }
            }
        })
    }
}

// --- Тест: JSON вывод совпадает с ожидаемым ---
func Test_AddJSON_FormatsCorrectly(t *testing.T) {
    groups := &StringGrouper{
        Str:         "apple, Apricot, avocado, banana, blueberry",
        MaxPerGroup: 2,
    }
    groups = groups.Handle()

    expectedJSON := `{
        "a": [["apple", "Apricot"], ["avocado"]],
        "b": [["banana", "blueberry"]]
    }`

    var expected, actual map[string][][]string
    if err := json.Unmarshal([]byte(expectedJSON), &expected); err != nil {
        t.Fatal("ошибка парсинга эталонного JSON:", err)
    }

    if err := json.Unmarshal([]byte(groups.DataJson), &actual); err != nil {
        t.Fatal("ошибка парсинга результата:", err)
    }

    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("ожидаемый JSON:\n%+v\nполучено:\n%+v", expected, actual)
    }
}

// --- Тест: Поддержка Unicode-символов ---
func Test_Handle_UnicodeCharacters(t *testing.T) {
    groups := &StringGrouper{
        Str:         "Арбуз, ананас, Банан, баклажан",
        MaxPerGroup: 2,
    }
    groups = groups.Handle()

    keys := getSortedKeys(groups.GroupData)
    expectedKeys := []string{"а", "б"}

    if len(keys) != len(expectedKeys) {
        t.Fatalf("ожидается %d групп, получено %d", len(expectedKeys), len(keys))
    }

    for i := range keys {
        if keys[i] != expectedKeys[i] {
            t.Errorf("ожидается группа %q, получена %q", expectedKeys[i], keys[i])
        }
    }
}

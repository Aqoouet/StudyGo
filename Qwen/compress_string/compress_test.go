package main

import (
    "testing"
    "strings"
)

// -----------------------------
// Тесты для Compress
// -----------------------------

func TestCompress(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        // Основные примеры
        {"aabbaaa", "a2b2a3"},      // "aabbaaa" → 7 символов; "a2b2a3" → 6 → сжато
        {"aaaabbb", "a4b3"},        // "aaaabbb" → 7; "a4b3" → 5 → сжато
        {"abcd", "abcd"},           // длина совпадает → не сжимаем
        {"aaaaaa", "a6"},           // "aaaaaa" → 6; "a6" → 2 → сжато
        {"abba", "abba"},           // "abba" → 4; "a2b2" → 5 → НЕ сжато → вернули оригинал
        {"aaAAbbBB", "aaAAbbBB"},   // "aaAAbbBB" → 8; "a2A2b2B2" → 9 → НЕ сжато
        {"", ""},                   // пустая строка
        {"a", "a"},                 // один символ → не сжимаем
        {"aaaaaaaaaa", "a10"},      // "aaaaaaaaaa" → 10; "a10" → 3 → сжато
        {"zzzzzzzzzzeeeeeeerrrrrrrrtttttttttt", "z10e7r8t10"}, // длинная серия — сжата
        {"abcdefgh", "abcdefgh"},   // нет повторений → не сжато
        {"abcabc", "abcabc"},       // чередование без повторений → не сжато
    }

    for _, test := range tests {
        t.Run(test.input, func(t *testing.T) {
            result := Compress(test.input)
            if result != test.expected {
                t.Errorf("Compress(%q) = %q; want %q", test.input, result, test.expected)
            }
        })
    }
}

// -----------------------------
// Тесты для Decompress
// -----------------------------

func TestDecompress(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        // Базовые случаи
        {"a2b2a3", "aabbaaa"},
        {"a4b3", "aaaabbb"},
        {"abcd", "abcd"},
        {"a", "a"},
        {"a1z2", "azz"},
        {"a10", "aaaaaaaaaa"},
        {"XYZ", "XYZ"},
        {"X2Y3Z", "XXYYYZ"},
        {"", ""},

        // Некорректные участки — остаются как есть
        {"2a", "2a"},              // начинается с числа → оставить как есть
        {"a1bX3c", "abXXXc"},      // X не число → оставить как есть
        {"a-3", "a---"},           // минус перед числом → оставить как есть
        {"a1b2cdef3", "abbcdefff"},// 'd' после числа → продолжить
        {"a1b2c3def4", "abbcccdeffff"}, // частичная обработка
        {"aX", "aX"},              // нет числа после X → оставить как есть
        {"a1bX", "abX"},           // частичная обработка

        // Новые тесты: добавлены по запросу
        {"12a3", "12aaa"},         // только часть — число, но не пара → оставить 12 + обработать a3
        {"a0b1c2", "bcc"},         // нулевое количество → пропустить "a"
        {"a1", "a"},               // одно число после символа → просто символ
        {"😀2😄3", "😀😀😄😄😄"},     // юникодные символы → важно правильно интерпретировать
        {"a1bX3", "abXXX"},        // ошибка в середине → обрабатываем только валидные части
        {"a1b2c", "abbc"},         // c — не число → обрабатываем до него
        {"a99", strings.Repeat("a", 99)}, // большое число → должно обработать
        {"a1b0c5", "accccc"},      // b0 → убираем
        {"5", "5"},                // одиночная цифра → оставить как есть

        // Тесты, где твой код может дать ошибку:
        {"a1b1", "ab"},            // последовательные пары (a1 + b1)
        {"a2b", "aab"},            // число, потом один символ
        {"a1b2c1", "abbc"},       // чередование чисел и символов
        {"a1bXc3", "abXccc"},      // буква между числами
        {"abc1", "abc"},          // число в конце строки без символа перед
        {"aX1b2", "aXbb"},       // число после неверного символа
        {"a1bXc", "abXc"},         // конец строки без числа
    }

    for _, tt := range tests {
        t.Run(tt.input, func(t *testing.T) {
            result := Decompress(tt.input)
            if result != tt.expected {
                t.Errorf("Decompress(%q) = %q; want %q", tt.input, result, tt.expected)
            }
        })
    }
}

// -----------------------------
// Тесты для IsValidCompressedFormat
// -----------------------------

func TestIsValidCompressedFormat(t *testing.T) {
    tests := []struct {
        input    string
        expected bool
    }{
        {"a2b2a3", true},
        {"a4b3", true},
        {"abcd", false},     // символы без чисел считаются допустимыми
        {"a", false},        // одиночный символ
        {"a1z2", true},
        {"a10", true},
        {"2a", false},      // начинается с числа → неверно
        {"a1bX3c", true},  // X не является частью числа → неверно
        {"a-3", true},     // - не часть числа → неверно
        {"a1b2cdef3", true}, // 'd' после числа не имеет числа → неверно
        {"", false},         // пустая строка считается допустимой
    }

    for _, test := range tests {
        t.Run(test.input, func(t *testing.T) {
            result := IsValidCompressedFormat(test.input)
            if result != test.expected {
                t.Errorf("IsValidCompressedFormat(%q) = %v; want %v", test.input, result, test.expected)
            }
        })
    }
}



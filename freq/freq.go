package freq

import(
    "math"
    "unicode"
)

// ScoreEngText scores a text on its similarity to English using the
// Bhattacharyya coefficient: https://en.wikipedia.org/wiki/Bhattacharyya_distance#Bhattacharyya_coefficient
// This method is particularly useful since it heavily penalizes a text for
// each non-alphabetic character that is present
func ScoreEngText(text string) float64 {

    // letter frequency probabilities for the english language
    // https://cs.marlboro.college/2001_2006/fall01/computation/compression/www.data-compression.com/english.html
    freqMap := map[rune]float64 {
        'a': 0.0651738, 'b': 0.0124248, 'c': 0.0217339, 'd': 0.0349835,
        'e': 0.1041442, 'f': 0.0197881, 'g': 0.0158610, 'h': 0.0492888,
        'i': 0.0558094, 'j': 0.0009033, 'k': 0.0050529, 'l': 0.0331490,
        'm': 0.0202124, 'n': 0.0564513, 'o': 0.0596302, 'p': 0.0137645,
        'q': 0.0008606, 'r': 0.0497563, 's': 0.0515760, 't': 0.0729357,
        'u': 0.0225134, 'v': 0.0082903, 'w': 0.0171272, 'x': 0.0013692,
        'y': 0.0145984, 'z': 0.0007836, ' ': 0.1918182 }

    letterMap := make(map[rune]float64, 27)

    var bhatCoefficient float64

    // adds the number of times a symbol shows up in text to letterMap
    for _, char := range(text) {
        letterMap[unicode.ToLower(char)]++
    }

    // normalizes the frequency to a percent of total letters
    for key := range(letterMap){
        letterMap[key] = letterMap[key] / float64( len(text) )
    }

    // BC = sqrt(p * q) where p, q, are the probabilities of a letter showing up in
    // a piece of text. p is from the table above, q is calculated from the given text
    for key := range(letterMap){
        bhatCoefficient += math.Sqrt( freqMap[key] * letterMap[key] )
    }

    return bhatCoefficient

}

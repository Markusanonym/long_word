function LongestWord(sen){
//Remove punctuation
    sen = sen.replace(/[^\w\s]/g, '');
    var words = sen.split(' ');  // Split the string into words
    var longestWord = '';
    for (let i= 0; i < words.length; i++) {
        var word = words[i];
        if (word.length>longestWord.length) {
            longest_word = word;
        }
    }
    return longestWord;
}
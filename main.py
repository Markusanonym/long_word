import re

def LongestWord(sen):
    sen = re.sub(r'[^\w\s]', '', sen)  # Remove punctuation
    words = sen.split()  # Split the string into words
    longest_word = ''
  
    for word in words:
        if len(word) > len(longest_word):
            longest_word = word

    return longest_word
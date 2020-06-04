def find_anagrams(word, candidates):
    result = []
    for candidate_word in candidates:
        if sorted(candidate_word.lower()) != sorted(word.lower()):
            continue
        if candidate_word.lower() == word.lower():
            continue
        result.append(candidate_word)
    return result

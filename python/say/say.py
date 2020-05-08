def say(number):
    output = {
        0: '',
        1: 'one',
        2: 'two',
        3: 'three',
        4: 'four',
        5: 'five',
        6: 'six',
        7: 'seven',
        8: 'eight',
        9: 'nine',
        10: 'ten',
        11: 'eleven',
        12: 'twelve',
        13: 'thirteen',
        14: 'fourteen',
        15: 'fifteen',
        16: 'sixteen',
        17: 'seventeen',
        18: 'eighteen',
        19: 'nineteen',
        20: 'twenty',
        30: 'thirty',
        40: 'forty',
        50: 'fifty',
        60: 'sixty',
        70: 'seventy',
        80: 'eighty',
        90: 'ninety',
        100: 'hundred',
    }
    if number < 0 or number >= 1000000000000:
        raise ValueError("Input must be between 0 and 1000000000000.")
    if number == 0:
        return 'zero'
    if number <= 20:
        return output[number]
    if number < 100:
        return '{}-{}'.format(output[number//10*10], output[number%10])
    if number < 1000 and number % 100 == 0:
        return '{} hundred'.format(output[number//100])
    if number < 1000:
        return '{} {} {}'.format(output[number//100], 'hundred', say(number%100))
    if number < 10000 and number % 1000 == 0:
        return '{} thousand'.format(say(number//1000))
    if number < 1000000000000:
        result = ''
        billions = number // 1000000000
        millions = number % 1000000000 // 1000000
        thousands = number % 1000000 // 1000
        rest = number % 1000
        if billions:
            result += '{} {} '.format(say(billions), 'billion')
        if millions:
            result += '{} {} '.format(say(millions), 'million')
        elif billions and (thousands or rest):
            result += ' '
        if thousands:
            result += '{} {} '.format(say(thousands), 'thousand')
        elif millions and rest:
            result += ' '
        if rest:
            result += '{}'.format(say(rest))
        result = result.replace('  ', ' ')
        return result.strip()
    raise ValueError("Error")


if __name__ == '__main__':
    print('result:', say(987654321123))
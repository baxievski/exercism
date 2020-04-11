def slices(series, length):
    if length > len(series):
        raise ValueError("Slice length is larger than the series length.")
    if not length:
        raise ValueError("Slice length can not be 0.")
    if length < 0:
        raise ValueError("Slice length can not be negative.")
    result = [
        series[i:i+length]
        for i in range(len(series) - length + 1)
    ]
    return result


if __name__ == "__main__":
    print(f'{slices("123456789", 3)=}')

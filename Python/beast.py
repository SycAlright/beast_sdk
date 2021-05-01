# Syc <github.com/SycAlright>
# Beast_SDK Python


beast = ['嗷', '呜', '啊', '~']


def str2hex(text: str):
    ret = ""
    for x in text:
        charHexStr = hex(ord(x))[2:]
        if len(charHexStr) == 3:
            charHexStr = "0" + charHexStr
        elif len(charHexStr) == 2:
            charHexStr = "00" + charHexStr
        ret += charHexStr
    return ret


def hex2str(text: str):
    ret = ""
    for i in range(0, len(text), 4):
        unicodeHexStr = text[i:i + 4]
        charStr = chr(int(unicodeHexStr, 16))
        ret += charStr
    return ret


def encode(str):
    hexArray = list(str2hex(str))
    code = ""
    n = 0
    for x in hexArray:
        k = int(x, 16) + n % 16
        if k >= 16:
            k -= 16
        code += beast[int(k / 4)] + beast[k % 4]
        n += 1
    return code


def decode(str):
    hexArray = list(str)
    code = ""
    for i in range(0, len(hexArray), 2):
        pos1 = beast.index(hexArray[i])
        pos2 = beast.index(hexArray[i + 1])
        k = ((pos1 * 4) + pos2) - (int(i / 2) % 16)
        if k < 0:
            k += 16
        code += hex(k)[2:]
    return hex2str(code)


if __name__ == '__main__':
    print(encode("你好"))
    print(decode("呜嗷嗷嗷啊嗷嗷~啊呜~啊~呜呜嗷"))

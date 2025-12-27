def build_next(pattern: str) -> list:
    """
    构造 KMP 的 next 数组（也叫 lps 数组）
    next[i] 表示 pattern[0:i+1] 的最长相等前后缀长度
    """
    n = len(pattern)
    next_arr = [0] * n
    j = 0

    # 第一个字符没有前后缀，所以从第二个字符开始
    for i in range(1, n):
        # 如果还能回退并且不匹配就继续回退到上一个j-1长度对应的下一个待匹配位
        while j > 0 and pattern[i] != pattern[j]:
            j = next_arr[j-1]
        
        # 匹配成功，前后缀长度加1
        if pattern[i] == pattern[j]:
            j += 1
        
        # 记录前后缀长度
        next_arr[i] = j
    return next_arr

def kmp_search(text: str, pattern: str) -> int:
    """
    在 text 中查找 pattern
    返回 pattern 第一次出现的位置，找不到返回 -1
    """
    if not pattern:
        return 0

    next_arr = build_next(pattern)
    j = 0  # 指向 pattern

    for i in range(len(text)):
        # 不匹配就根据 next 回退
        while j > 0 and text[i] != pattern[j]:
            j = next_arr[j - 1]

        # 匹配成功
        if text[i] == pattern[j]:
            j += 1

        # 完全匹配
        if j == len(pattern):
            return i - j + 1

    return -1

if __name__ == "__main__":
    text = "ababcabcabababd"
    pattern = "ababd"
    index = kmp_search(text, pattern)
    print(f"Pattern found at index: {index}")  # Output: Pattern found at index: 10
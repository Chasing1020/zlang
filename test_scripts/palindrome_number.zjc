// Given an integer x, return true if x is palindrome integer.
// An integer is a palindrome when it reads the same backward as forward.
// For example, 121 is a palindrome while 123 is not.
//
let isPalindrome = function (x) {
    if (x < 0) {
        return false;
    }
    let div = 1;
    for (; x / div >= 10;) {
        div = div * 10;
    }
    for (; x > 0;) {
        let left = x / div;
        let right = x % 10;
        if (left != right) {
            return false;
        }
        x = (x % div) / 10;
        div = div / 100;
    }
    return true;
}


print("Please enter a number:");
let input = int(input());
print(isPalindrome(input));

function isPalindrome(strName, left, right) {
    while (left < right) {
        if (str[left] !== str[right]) {
            return false;
        }
        left++;
        right--;
    }
    return true;
}

function sortBySize(a, b) {
    nums.sort((a, b) => {
        return a - b
    })
}

module.exports = isPalindrome;

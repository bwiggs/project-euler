
function isPalindrome(val) {
    'use strict';
    var str = val + '';
    var len = str.length;
    // this is silly, see below
    var i, j;
    for (i = 0, j = len - 1; i <= j; i++, j--) {
        if (str[i] !== str[j]) {
            return false;
        }
    }
    // just reverse the string and check
    return (str === str.split('').reverse().join(''));
}


var a, b, product, largestPalindrome = 0, passes = 0;
for (a = 999; a >= 100; a--) {
    for (b = a; b >= 100; b--) {
        passes += 1;
        product = a * b;
        if (product < largestPalindrome) {
            break;
        }

        if (isPalindrome(product) && product > largestPalindrome) {
            largestPalindrome = product;
        }
    }
}
console.log('Passes: ', passes);
console.log(largestPalindrome);

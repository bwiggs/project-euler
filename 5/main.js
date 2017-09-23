// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
// https://projecteuler.net/problem=5
//
// 20 - covers 10, 5, 2
// 19
// 18 - covers 3
// 17
// 16
// 15
// 14
// 13
// 12
// 11
// 10
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1

// 20 - 2, 5, 10

var increment = 20;
var multiples = [2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20];
var num = 0;
var solved = false;
while(!solved) {
    num += increment;
    var i;
    for(i = 0; i < multiples.length; i++) {
        if( num % multiples[i] !== 0)
            break;
    }
    if (i == multiples.length) {
        console.log(num);
        solved = true;
    }

}


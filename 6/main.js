// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

range = Array.apply(null, Array(101)).map(function (_, i) {return i;}).slice(1);
sumSquares = range.slice(0).reduce(function(prev, curr) { return prev + Math.pow(curr, 2)});
squareSums = Math.pow(range.slice(0).reduce(function(prev, curr) { return prev + curr; }), 2);

console.log(sumSquares, squareSums, squareSums - sumSquares);

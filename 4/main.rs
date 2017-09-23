fn isPalindrome(palindrome: int) -> bool {
    let pal = palindrome.to_str();
    for c in pal.chars() {
        println!(c);
    }
    return false;
}

fn main() {
    let largestPalindrome = 998899;
    if isPalindrome(largestPalindrome) {
        println!("palindrome!");
    } else {
        println!("newp!");
    }
}


hello team

#### "Random" number generation and array selection

The function `randomNum(input)` generates a seed with the unix time `1706284949383` which is modified by the user input `0-9` and then is transformed to a number `< 100`  by means of the modulos operator `%` and 100. Example:


    1706284949383
    1706284949383 + 5 = 1706284949388
    1706284949388 % 100 = 88  

The selection of the element of the array is done by taking the generated seed and reducing to to a number that is less than the size the array by also using the modulus operator `seed % array_size`. 
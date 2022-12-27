I want to use this post as a reflection for my journey in AdventOfCode 2022. 

<br>
I discovered AoC in 2021 during COVID season. I was bored at home and wanted to find some challenges to work on. And it is at that time that I found AoC in Reddit. My initial goal for AoC was to learn and code the solutions using new languages. 

<br>    
Last year I started with Rust. Frankly, I did not find Rust as quite intimidating as others describe. I like Rust's programing idiom. I think Rust's handling of NULL pointer is great. It is similar to Java's Optional class but more cleaner. I like Rust's pattern matching and I am so happy that Java 18 now support similar syntax. Also, Rust provides lots of out-of-box api that had helped me learn the language. However, one thing I do not like about Rust is its use of **impl for**. I feel that it has became a bit repetitive. I feels weird that I need to write impl for every trait I want to implement. If for me, I would rather write a definition like "impl Action1, Action2, ... for struct" near the struct declaration. And somewhere below, I would then have the detai implementation. These are just my own opnions though. Also, since I did not upload my 2021 code to Github, I cannot remember much about what I wrote before. 

<br>
This year, I tried Golang. Before starting the AoC, I went thought the tutorial from golang website. To be honst, I do not mind Golang's idea of returning error. In Python, you can also return multiple values back to the caller and I will somethimes return a boolean value as logic controller just for simplicity. I even like a lot about its local scope in **if** statements. This allows me to use same auxiliary variable names without worrying it will go out of scope and affect other part of code. Also, **the func (type Type) funcName()** and  **the func (type *Type) funcName()** is really interesting. It provide finer control for immutability and prevent myself from changing values accidentally in my code. I really like this design a lot. Additionally, I realize that Golang looks like a funcational language than a OOP language. "." notation is basically only used after a struct or struct pointer. For example, string operations can only be accomplished using "strings.Split(input, delim)", input.Split(delim) does not work. 

<br>
One thing I got annoyed using Golang is its package level variable definition. Bascially, you cannot define variabels with same naem from different files in the same package. This is terrible due to each problems in AoC has differnt input that requires parser. This makes naming super annoying. But this is just my case, not really comparable.

<br>
Now, AoC. I only finished 17/25 questions this year. There are excuses but pointless here. Questions year involves lots of state machine changes in a 2 dimension space in a period of time. I was not able to complete Question 16 and 19. Both are graph traversing program that find the optimal strategies. I was able to code the base cases but cannot figure out what need to memorize in order for recursions to stop early. I did not check out question 23, 24, 25, that might be a future me task :). 

<br>
Overall, using Golang is not a bad experience. It is strange to use a langauge without set, due to the nature of AoC, that a lot of memorization technique need to be used. Also, compared to Rust, Golang itself only provide some barebone api. Things such as removing a element from slice/array at index is not supported. Also, I did not get to use some of Golang's famous features, such as concurrency and channel. 

<br>
Anyways, thanks for reading my two cents. Have a wonderful days in your life. 

<br>
Just to add here, it is funny that I spend a lot of time thinking and structuring the problem before I started coding. But sometimes it turns out I am just doing too much extra works. LMAO. 
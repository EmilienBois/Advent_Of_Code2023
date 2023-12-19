# Advent_Of_Code2023
Repository containing my works on the advent of code 2023, which I have done during the programming languages course. Most of the code is made in GO language.

DAY 1 :
Language used : GO

Work has been way harder than I thought it would be. However, I manage to do it in 45 min. I'm startled how different it is from during the class. It is quite stressfull.
Let's see how tomorrow goes, I'll start it earlier in the day.

DAY 2 : 
Language used : GO

I started earlier this day, and I enjoyed doing it. 
I found no real difficulties, except dealing with the spaces in the string. I used 40 min for the first part, whereas the second part only took me 5 minutes to complete.
It is a pleasure for now.

DAY 3 :
Language used : GO

I've had some real hard time completing this day. I've spent a lot of time on debugging, showing I need to be more carefull when I code.
First problem I'm confronted : using const to make two dimensional tab. I managed to do the day without using this object, but I still need to know in advance the length of the input to do the exercice. My problem is that I don't know how to make a function return a const.
Second problem was that my number was counted too many times, but I managed it quickly. 
At first I tried to repertory all the symbols in the input, but I noticed later that I forgot some, which made me think of finding not dot or number element to fing symbol, which was way easier.

DAY 4 :
Language used : GO

Quite easy day, it's a relief !

DAY 5 : 
Language used : GO

I've had a really hard time here. Been awoken seen 6:45 am trying to resolve it. However, it is also the first time I've been confronted to the size problem when programming. My first program worked only for part 1 and input test. My second was fast enough for the whole input and test for part 2. I've spent more than 20 min to get the answer for the second part for the whole input. 
Result : 27992443

DAY 6 :
Still on GO language.

I found today's enigma very easy. Though, I tried directly to optimize it by resolving every time the second degree equation, and not iterate at all on the file. It made me loose a lot of time. Although, I'm still trying to understand how to test better.

DAY 7 :
Go Language

At first, I thought today would be easy. However, even if it worked on the inputtest, I don't know why It doesn't work on my input. I abandon for now, but I will come back later.
I finally finished this day on day 12. After all my getpower function were a problem, I didn't take every possibility into account. This was showed to me evenmore on the second part.

DAY 8 :
Language used : GO

Done after day 9, during day 9. First part was quite easy, and I needed more time to do the second part, where I needed to think about getting the smallest path for each begin (each element that ends with). I called all these path steps and put them in the variable liste_steps. I then have found the Smallest common multiple of these steps. FOr this, I first tried to find it starting from 0 up to the wanted number. However, it was still to slow, and I finally decided to use recursive functions : PPCM and PCDG, using mathematical formulas to calculate it fast.
Interesting challenge.

DAY 9 :
Language used : GO

I'm very proud of today. I've done it in half an hour, bu using a very powerful recursive function. So my program stands in less than 100 lines. It feels really good to finally succeed a day.

DAY 11 :
Language used : GO

Interesting day, I have used a lot of functions, and it really helped me going faster. In the end, I realized I could delete one function to optimize my program and make the part 2 doable in less than 7ms.

DAY 13 :
Language used : GO

No big deal for today, I went quite straightforward in the probleme, checking the valid mirror, first vertical, memorizing the pattern that are doone (the mirror has already been found), and next the horizontal mirror. There are a lot of "for" boucle, but breaks and memorizing helped optimize this program. Doing it in 6.856 ms

DAY 14 :
Language used : GO

First part was easy, no big deal. Didn't have time doing the second part, but I'll come back later. I suppose there is, once more, a pattern to fidn.

DAY 15 :
Language used : GO

FIrst part was very easy. I've done it in 10 minutes, so I am proud. For the second, I tried to go to fast and forgot to think carefully. So I have lost some time, eventhough it was finally not that difficult.

DAY 16 :
Language used : GO

Today i wasn't very motivated by the problem. However, it took me about 2 hours to resolve the first part. It was a bit annoying because it was mostly about switch and case. BUT, i thought it was very interesting because it was an infinte boucle, so I needed to implement the number of times my grid hadn't changed, in order to stop the boucle when needed. I don't know if it is the optimize way but it was an interesting approach. For the second part, I am too tired to think about it much more, but I don't see other optimization, so I'm trying to do the same program 400 times. It took me 6min30, but I didn't get the good answer(7952). i'll be coming back later.

DAY 17 : 
Language used : GO

I've been helped a lot by the class we had. Using Djikstra, the problem really is easier. I used an already made program for the priority queue, which helped me but also troubled me because I needed to rewrite some things.

DAY 18 :
Language used : GO

Tried without success today's challenge. I tried to do it withou the formulas on internet, and it just didn't give me a correct result, however it worked well on the inputtest. I tried to first calculate the perimeter : easy ; then calculate the surface, taking line by line. For each line, I used a In boolean countaining the information of whether the actual location was supposed to be in the surface or outside. Then I had a lot of difficulties with the sommet. 

DAY 19 :
Language used : GO

I did first part withouth much difficulty, and I haven't had time to try part 2, though it seems to be interesting and I'll come back later. For the first part, much of my time was occupied by examining the file and making both the rating and workflow map. I think I did well also because I first tried to think about my coding : what maps and slices I would need. 
I first used a map linking a string (the step) to a list of Rules, these being a struct having 4 attributes : sign, category, value and destination. For a default rules I decided that only the attribute destination would be not null. Then I used a list of Rating, rating being a map linking a string to an int (the category to its value)
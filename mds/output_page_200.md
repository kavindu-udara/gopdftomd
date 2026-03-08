194 CHAPTER 5. DIGITAL SIGNAL PROCESSING

Solution to Exercise 5.33 (p. 178)
The delay is not computational delay here the plot shows the rst output value is aligned with the lter's rst
input although in real systems this is an important consideration. Rather, the delay is due to the lter's phase

shift: A phase-shifted sinusoid is equivalent to a time-delayed one: cos (2 fn ) = cos 2 f n
.

All lters have phase shifts. This delay could be removed if the lter introduced no phase shift. Such lters
do not exist in analog form, but digital ones can be programmed, but not in real time. Doing so would
require the output to emerge before the input arrives!
Solution to Exercise 5.34 (p. 179)
We have p + q + 1 multiplications and p + q 1 additions. Thus, the total number of arithmetic operations
equals 2 ( p + q ) .

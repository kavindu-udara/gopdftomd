46 CHAPTER 3. ANALOG SIGNAL PROCESSING

Figure 3.19: All circuits containing sources and resistors can be described by simpler equivalent circuits.
Choosing the one to use depends on the application, not on what is actually inside the circuit. (Solution on p. 94.)
Find the Mayer-Norton equivalent circuit for the circuit below.

Figure 3.20

Equivalent circuits can be used in two basic ways. The rst is to simplify the analysis of a complicated
circuit by realizing the any portion of a circuit can be described by either a Thévenin or Mayer-Norton
equivalent. Which one is used depends on whether what is attached to the terminals is a series con guration
(making the Thévenin equivalent the best) or a parallel one (making Mayer-Norton the best).
Another application is modeling. When we buy a ashlight battery, either equivalent circuit can accu-
rately describe it. These models help us understand the limitations of a battery. Since batteries are labeled
with a voltage speci cation, they should serve as voltage sources and the Thévenin equivalent serves as the
natural choice. If a load resistance R is placed across its terminals, the voltage output can be found using

voltage divider: v =
R + R
v . If we have a load resistance much larger than the battery's equivalent

resistance, then, to a good approximation, the battery does serve as a voltage source. If the load resistance
is much smaller, we certainly don't have a voltage source (the output voltage depends directly on the load
resistance). Consider now the Mayer-Norton equivalent; the current through the load resistance is given by

current divider, and equals i =
R + R
i . For a current that does not vary with the load resistance,

this resistance should be much smaller than the equivalent resistance. If the load resistance is comparable
to the equivalent resistance, the battery serves neither as a voltage source or a current course. Thus, when

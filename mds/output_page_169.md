163

(a)

(b)

Figure 5.12: The initial decomposition of a length-8 DFT into the terms using even- and odd-indexed
inputs marks the rst phase of developing the FFT algorithm. When these half-length transforms are
successively decomposed, we are left with the diagram shown in the bottom panel that depicts the
length-8 FFT computation.

the half-length transforms and combines them through additions and the multiplication by e
, which is

not periodic over
N
, to rewrite the length- N DFT. Figure 5.12 illustrates this decomposition. As it stands,

we now compute two length-
N
transforms (complexity 2 O N = 4 ), multiply one of them by the complex

exponential (complexity O ( N ) ), and add the results (complexity O ( N ) ). At this point, the total complexity
is still dominated by the half-length DFT calculations, but the proportionality coe cient has been reduced.
Now for the fun. Because N = 2 , each of the half-length transforms can be reduced to two quarter-
length transforms, each of these to two eighth-length ones, etc. This decomposition continues until we are
left with length-2 transforms. This transform is quite simple, involving only additions. Thus, the rst stage

of the FFT has
N
length-2 transforms (see the bottom part of Figure 5.12). Pairs of these transforms are

combined by adding one to the other multiplied by a complex exponential. Each pair requires 6 additions

and 4 multiplications, giving a total number of computations equaling 10
N
=
5 N
. This number of

computations does not change from stage to stage. Because the number of stages, the number of times the
length can be divided by two, equals log N , the complexity of the FFT is O ( N log N ) .
Doing an example will make computational savings more obvious. Let's look at the details of a length-8
DFT. As shown on Figure 5.13, we rst decompose the DFT into two length-4 DFTs, with the outputs added
and subtracted together in pairs. Considering Figure 5.13 as the frequency index goes from 0 through 7,
we recycle values from the length-4 DFTs into the nal calculation because of the periodicity of the DFT
output. Examining how pairs of outputs are collected together, we create the basic computational element
known as a butter y (Figure 5.13).

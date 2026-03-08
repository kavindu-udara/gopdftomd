138 CHAPTER 4. FREQUENCY DOMAIN

(a) Assuming the message signal is lowpass and has a bandwidth of W Hz, what values for the square
wave's period T are feasible. In other words, do some combinations of W and T prevent reception?
(b) Assuming reception is possible, can standard radios receive this innovative AM transmission? If so,
show how a coherent receiver could demodulate it; if not, show how the coherent receiver's output
would be corrupted. Assume that the message bandwidth W = 5 kHz.

Problem 4.27: Secret Communication
An amplitude-modulated secret message m ( t ) has the following form.

r ( t ) = A 1 + m ( t ) cos 2 ( f + f ) t

The message signal has a bandwidth of W Hz and a magnitude less than 1 ( j m ( t ) j < 1 ). The idea is to
o set the carrier frequency by f Hz from standard radio carrier frequencies. Thus, o -the-shelf coherent
demodulators would assume the carrier frequency has f Hz. Here, f < W .

(a) Sketch the spectrum of the demodulated signal produced by a coherent demodulator tuned to f Hz.
(b) Will this demodulated signal be a scrambled version of the original? If so, how so; if not, why not?
(c) Can you develop a receiver that can demodulate the message without knowing the o set frequency f ?

Problem 4.28: Signal Scrambling
An excited inventor announces the discovery of a way of using analog technology to render music unlistenable
without knowing the secret recovery method. The idea is to modulate the bandlimited message m ( t ) by a
special periodic signal s ( t ) that is zero during half of its period, which renders the message unlistenable and
super cially, at least, unrecoverable (Figure 4.34).

Figure 4.34

(a) What is the Fourier series for the periodic signal?
(b) What are the restrictions on the period T so that the message signal can be recovered from m ( t ) s ( t ) ?
(c) ELEC 241 students think they have broken the inventor's scheme and are going to announce it to
the world. How would they recover the original message without having detailed knowledge of the
modulating signal?

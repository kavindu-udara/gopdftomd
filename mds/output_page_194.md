188 CHAPTER 5. DIGITAL SIGNAL PROCESSING

(a) Plot the magnitude and phase of this lter's transfer function.

(b) What is this lter's output when x ( n ) = cos
n
+ 2 sin
2 n
?

Problem 5.25: Detective Work
The signal x ( n ) equals ( n ) ( n 1) .

(a) Find the length-8 DFT (discrete Fourier transform) of this signal.
(b) You are told that when x ( n ) served as the input to a linear FIR ( nite impulse response) lter, the
output was y ( n ) = ( n ) ( n 1) + 2 ( n 2) . Is this statement true? If so, indicate why and nd
the system's unit sample response; if not, show why not.

Problem 5.26:
A discrete-time, shift invariant, linear system produces an output y ( n ) = f 1 ; 1 ; 0 ; 0 ; : : : g when its input
x ( n ) equals a unit sample.

(a) Find the di erence equation governing the system.
(b) Find the output when x ( n ) = cos (2 f n ) .
(c) How would you describe this system's function?

Problem 5.27: Time Reversal has Uses
A discrete-time system has transfer function H e . A signal x ( n ) is passed through this system to
yield the signal w ( n ) . The time-reversed signal w ( n ) is then passed through the system to yield the
time-reversed output y ( n ) . What is the transfer function between x ( n ) and y ( n ) ?

Problem 5.28: Removing Hum
The slang word hum represents power line waveforms that creep into signals because of poor circuit con-
struction. Usually, the 60 Hz signal (and its harmonics) are added to the desired signal. What we seek are
lters that can remove hum. In this problem, the signal and the accompanying hum have been sampled; we
want to design a digital lter for hum removal.

(a) Find lter coe cients for the length-3 FIR lter that can remove a sinusoid having digital frequency
f from its input.
(b) Assuming the sampling rate is f to what analog frequency does f correspond?
(c) A more general approach is to design a lter having a frequency response magnitude proportional to
the absolute value of a cosine: j H e j / j cos ( fN ) j . In this way, not only can the fundamental
but also its rst few harmonics be removed. Select the parameter N and the sampling rate so that the
frequencies at which the cosine equals zero correspond to 60 Hz and its odd harmonics through the
fth.
(d) Find the di erence equation that de nes this lter.

Problem 5.29: Digital AM Receiver
Thinking that digital implementations are always better, our clever engineer wants to design a digital AM
receiver. The receiver would bandpass the received signal, pass the result through an A/D converter, perform
all the demodulation with digital signal processing systems, and end with a D/A converter to produce the
analog message signal. Assume in this problem that the carrier frequency is always a large even multiple of
the message signal's bandwidth W .

(a) What is the smallest sampling rate that would be needed?
(b) Show the block diagram of the least complex digital AM receiver.
(c) Assuming the channel adds white noise and that a b -bit A/D converter is used, what is the output's
signal-to-noise ratio?

Problem 5.30: DFTs
A problem on Samantha's homework asks for the 8-point DFT of the discrete-time signal ( n 1)+ ( n 7) .

165

Figure 5.14
of computer memory does the speech consume?

The resulting discrete-time signal, shown in the bottom of Figure 5.14, clearly changes its character with
time. To display these spectral changes, the long signal was sectioned into frames : comparatively short,
contiguous groups of samples. Conceptually, a Fourier transform of each frame is calculated using the FFT.
Each frame is not so long that signi cant signal variations are retained within a frame, but not so short
that we lose the signal's spectral character. Roughly speaking, the speech signal's spectrum is evaluated
over successive time segments and stacked side by side so that the x -axis corresponds to time and the y -axis
frequency, with color indicating the spectral amplitude.
An important detail emerges when we examine each framed signal (Figure 5.15). At the frame's edges,
the signal may change very abruptly, a feature not present in the original signal. A transform of such a
segment reveals a curious oscillation in the spectrum, an artifact directly related to this sharp amplitude
change. A better way to frame signals for spectrograms is to apply a window : Shape the signal values
within a frame so that the signal decays gracefully as it nears the edges. This shaping is accomplished
by multiplying the framed signal by the sequence w ( n ) . In sectioning the signal, we essentially applied a
rectangular window: w ( n ) = 1 , 0 n N 1 . A much more graceful window is the Hanning window ;
it has the cosine shape w ( n ) =
1 cos (2 n=N ) . As shown in Figure 5.15, this shaping greatly reduces
spurious oscillations in each frame's spectrum. Considering the spectrum of the Hanning windowed frame,
we nd that the oscillations resulting from applying the rectangular window obscured a formant (the one
located at a little more than half the Nyquist frequency).

![Im215](../output/images/test_171_Im215.R10.png)

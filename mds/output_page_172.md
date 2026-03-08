166 CHAPTER 5. DIGITAL SIGNAL PROCESSING

Figure 5.15: The top waveform is a segment 1024 samples long taken from the beginning of the Rice
University phrase. Computing Figure 5.14 involved creating frames, here demarcated by the vertical
lines, that were 256 samples long and nding the spectrum of each. If a rectangular window is applied
(corresponding to extracting a frame from the signal), oscillations appear in the spectrum (middle of
bottom row). Applying a Hanning window gracefully tapers the signal toward frame edges, thereby
yielding a more accurate computation of the signal's spectrum at that moment of time.

Figure 5.16: In comparison with the original speech segment shown in the upper plot, the non-
overlapped Hanning windowed version shown below it is very ragged. Clearly, spectral information
extracted from the bottom plot could well miss important features present in the original. (Solution on p. 193.)
What might be the source of these oscillations? To gain some insight, what is the length- 2 N
discrete Fourier transform of a length- N pulse? The pulse emulates the rectangular window, and
certainly has edges. Compare your answer with the length- 2 N transform of a length- N Hanning
window.

If you examine the windowed signal sections in sequence to examine windowing's e ect on signal amplitude,
we see that we have managed to amplitude-modulate the signal with the periodically repeated window
(Figure 5.16). To alleviate this problem, frames are overlapped (typically by half a frame duration). This

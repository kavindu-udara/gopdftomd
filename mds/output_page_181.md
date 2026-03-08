175

Figure 5.21: The blue line shows the Dow Jones Industrial Average from 1997, and the red one the
length-5 boxcar- ltered result that provides a running weekly of this market index. Note the edge
e ects in the ltered output.
h = [1 1 1 1 1]/5;
DJIA = fft(djia, 512);
H = fft(h, 512);
Y = H.*X;
y = ifft(Y);

note: The filter program has the feature that the length of its output equals the length of its
input. To force it to produce a signal having the proper length, the program zero-pads the input
appropriately.

MATLAB's fft function automatically zero-pads its input if the speci ed transform length (its
second argument) exceeds the signal's length. The frequency domain result will have a small
imaginary component largest value is 2 : 2 10 because of the inherent nite precision
nature of computer arithmetic. Because of the unfortunate mis t between signal lengths and
favored FFT lengths, the number of arithmetic operations in the time-domain implementation is
far less than those required by the frequency domain version: 514 versus 62,271. If the input signal
had been one sample shorter, the frequency-domain computations would have been more than a
factor of two less (28,696), but far more than in the time-domain implementation.
An interesting signal processing aspect of this example is demonstrated at the beginning and
end of the output. The ramping up and down that occurs can be traced to assuming the input is
zero before it begins and after it ends. The lter sees these initial and nal values as the di erence
equation passes over the input. These artifacts can be handled in two ways: we can just ignore the
edge e ects or the data from previous and succeeding years' last and rst week, respectively, can
be placed at the ends.
